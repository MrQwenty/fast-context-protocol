package server

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/MrQwenty/fast-context-protocol/internal/protocol"
)

type Resolver struct {
	catalogue *Catalogue
	now       func() time.Time
}

func NewResolver(c *Catalogue) *Resolver {
	return &Resolver{catalogue: c, now: time.Now}
}

type candidate struct {
	node  protocol.ContextNode
	score float64
}

func (r *Resolver) Resolve(req protocol.ContextRequest) (protocol.ContextPlan, error) {
	if req.RequestID == "" {
		return protocol.ContextPlan{}, fmt.Errorf("requestId is required")
	}
	if req.Intent.Type == "" {
		return protocol.ContextPlan{}, fmt.Errorf("intent.type is required")
	}
	if req.Budget.MaxTokens <= 0 || req.Budget.MaxBytes <= 0 {
		return protocol.ContextPlan{}, fmt.Errorf("positive maxTokens and maxBytes are required")
	}
	known := make(map[string]struct{}, len(req.KnownContext))
	for _, id := range req.KnownContext {
		known[id] = struct{}{}
	}

	now := r.now().UTC()
	var candidates []candidate
	var omissions []protocol.Omission
	for _, n := range r.catalogue.Nodes {
		if !matches(n.Intents, req.Intent.Type) {
			continue
		}
		if len(n.Targets) > 0 && req.Intent.Target != "" && !matches(n.Targets, req.Intent.Target) {
			continue
		}
		if req.Requirements.MinimumConfidence > 0 && n.Confidence < req.Requirements.MinimumConfidence {
			omissions = append(omissions, protocol.Omission{NodeID: n.ID, Reason: "confidence-below-minimum"})
			continue
		}
		if !sensitivityAccepted(n.Sensitivity, req.Requirements.AcceptSensitivity) {
			omissions = append(omissions, protocol.Omission{NodeID: n.ID, Reason: "sensitivity-not-accepted"})
			continue
		}
		freshness := 1.0
		if n.FreshUntil != nil && n.FreshUntil.Before(now) {
			freshness = 0.25
		}
		tokenCost := n.TokenEstimate
		if _, ok := known[n.ID]; ok {
			tokenCost = 1
		}
		score := (n.Priority * n.Confidence * freshness) / float64(max(tokenCost, 1))
		candidates = append(candidates, candidate{node: n, score: score})
	}

	sort.SliceStable(candidates, func(i, j int) bool {
		if candidates[i].score == candidates[j].score {
			return candidates[i].node.Priority > candidates[j].node.Priority
		}
		return candidates[i].score > candidates[j].score
	})

	plan := protocol.ContextPlan{
		RequestID:          req.RequestID,
		PlanID:             planID(req, now),
		ProtocolVersion:    protocol.Version,
		Complete:           true,
		EstimatedLatencyMS: 5,
		CreatedAt:          now,
		Omissions:          omissions,
	}
	for _, c := range candidates {
		n := c.node
		delivery := protocol.DeliveryInline
		incrementTokens := n.TokenEstimate
		incrementBytes := n.ByteLength
		if _, ok := known[n.ID]; ok {
			delivery = protocol.DeliveryReference
			incrementTokens = 0
			incrementBytes = 0
		}
		if plan.EstimatedTokens+incrementTokens > req.Budget.MaxTokens || plan.EstimatedBytes+incrementBytes > req.Budget.MaxBytes {
			plan.Complete = false
			plan.Omissions = append(plan.Omissions, protocol.Omission{NodeID: n.ID, Reason: "budget-exceeded"})
			continue
		}
		chunkNode := n
		fetchPath := ""
		if delivery == protocol.DeliveryReference {
			chunkNode.Content = ""
		} else if n.ByteLength > 32*1024 {
			delivery = protocol.DeliveryFetch
			chunkNode.Content = ""
			fetchPath = "/fcp/v0.1/context/" + n.ID
		}
		plan.EstimatedTokens += incrementTokens
		plan.EstimatedBytes += incrementBytes
		plan.Chunks = append(plan.Chunks, protocol.PlannedChunk{
			Node:      chunkNode,
			Delivery:  delivery,
			Rank:      len(plan.Chunks) + 1,
			Score:     c.score,
			Reason:    "matched intent and maximized utility within budget",
			FetchPath: fetchPath,
		})
	}
	plan.ContextRoot = contextRoot(plan.Chunks)
	return plan, nil
}

func matches(values []string, actual string) bool {
	for _, v := range values {
		if v == "*" || v == actual || (strings.HasSuffix(v, "*") && strings.HasPrefix(actual, strings.TrimSuffix(v, "*"))) {
			return true
		}
	}
	return len(values) == 0
}

func sensitivityAccepted(value string, accepted []string) bool {
	if value == "" || value == "public" || len(accepted) == 0 {
		return true
	}
	return matches(accepted, value)
}

func planID(req protocol.ContextRequest, now time.Time) string {
	sum := sha256.Sum256([]byte(req.RequestID + "|" + req.Intent.Type + "|" + req.Intent.Target + "|" + now.Format(time.RFC3339Nano)))
	return "plan_" + hex.EncodeToString(sum[:8])
}

func contextRoot(chunks []protocol.PlannedChunk) string {
	h := sha256.New()
	for _, c := range chunks {
		_, _ = h.Write([]byte(c.Node.ID))
	}
	return "sha256:" + hex.EncodeToString(h.Sum(nil))
}
