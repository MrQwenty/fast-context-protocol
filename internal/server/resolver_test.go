package server

import (
	"testing"
	"time"

	"github.com/MrQwenty/fast-context-protocol/internal/protocol"
)

func TestResolverRespectsBudgetAndKnownContext(t *testing.T) {
	now := time.Date(2026, 8, 1, 12, 0, 0, 0, time.UTC)
	c := &Catalogue{Nodes: []protocol.ContextNode{
		{ID: "sha256:a", Type: "doc", Content: "A", ContentType: "text/plain", TokenEstimate: 80, ByteLength: 80, Priority: 1, Confidence: 1, Sensitivity: "public", Intents: []string{"code.review"}, CreatedAt: now},
		{ID: "sha256:b", Type: "doc", Content: "B", ContentType: "text/plain", TokenEstimate: 60, ByteLength: 60, Priority: .8, Confidence: 1, Sensitivity: "public", Intents: []string{"code.review"}, CreatedAt: now},
	}, byID: map[string]protocol.ContextNode{}}
	r := NewResolver(c)
	r.now = func() time.Time { return now }

	plan, err := r.Resolve(protocol.ContextRequest{
		RequestID:    "req_test",
		Intent:       protocol.Intent{Type: "code.review"},
		Budget:       protocol.Budget{MaxTokens: 80, MaxBytes: 80},
		KnownContext: []string{"sha256:b"},
	})
	if err != nil {
		t.Fatal(err)
	}
	if len(plan.Chunks) != 2 {
		t.Fatalf("expected 2 chunks, got %d", len(plan.Chunks))
	}
	var sawReference bool
	for _, chunk := range plan.Chunks {
		if chunk.Node.ID == "sha256:b" && chunk.Delivery == protocol.DeliveryReference {
			sawReference = true
		}
	}
	if !sawReference {
		t.Fatal("expected known node to use reference delivery")
	}
	if plan.EstimatedTokens != 80 {
		t.Fatalf("expected 80 estimated tokens, got %d", plan.EstimatedTokens)
	}
}

func TestResolverRejectsInvalidRequest(t *testing.T) {
	r := NewResolver(&Catalogue{})
	_, err := r.Resolve(protocol.ContextRequest{})
	if err == nil {
		t.Fatal("expected validation error")
	}
}
