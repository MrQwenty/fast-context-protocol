package privacy

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/MrQwenty/fast-context-protocol/internal/protocol"
)

const (
	ProcessorName    = "fcp-local-privacy-gateway"
	ProcessorVersion = "0.1.0"
)

var ErrResidualRisk = errors.New("privacy policy rejected document because residual sensitive data remains")

type Engine struct {
	Secret   []byte
	VaultKey []byte
	Now      func() time.Time
}

func NewEngine(secret, vaultKey []byte) *Engine {
	return &Engine{Secret: append([]byte(nil), secret...), VaultKey: append([]byte(nil), vaultKey...), Now: time.Now}
}

func (e *Engine) Sanitize(request protocol.SanitizeRequest) (protocol.SanitizeResponse, error) {
	policy, err := normalizePolicy(request.Policy)
	if err != nil {
		return protocol.SanitizeResponse{}, err
	}
	if strings.TrimSpace(request.DocumentID) == "" {
		return protocol.SanitizeResponse{}, errors.New("documentId is required")
	}
	if request.Content == "" {
		return protocol.SanitizeResponse{}, errors.New("content is required")
	}
	if !isTextContentType(request.ContentType) {
		return protocol.SanitizeResponse{}, fmt.Errorf("unsupported content type %q: extract text locally before sanitization", request.ContentType)
	}

	detectors := []Detector{NewPatternDetector(), NewContextDetector()}
	if len(policy.CustomTerms) > 0 {
		detectors = append(detectors, NewDictionaryDetector(policy.CustomTerms))
	}
	candidates := detectAll(request.Content, detectors, policy)
	selected := resolveOverlaps(candidates)

	documentSalt := make([]byte, 32)
	if _, err := rand.Read(documentSalt); err != nil {
		return protocol.SanitizeResponse{}, fmt.Errorf("create document salt: %w", err)
	}
	transformer, err := newTransformer(policy, e.Secret, documentSalt)
	if err != nil {
		return protocol.SanitizeResponse{}, err
	}
	output, findings, reverseMap := applyTransformations(request.Content, selected, transformer, policy)

	unresolvedCandidates := resolveOverlaps(detectAll(output, detectors, policy))
	residualRisk := calculateResidualRisk(unresolvedCandidates)
	passed := residualRisk <= policy.MaxResidualRisk
	warnings := make([]string, 0, 2)
	if len(unresolvedCandidates) > 0 {
		warnings = append(warnings, fmt.Sprintf("%d sensitive spans remain after post-transform scan", len(unresolvedCandidates)))
	}
	if policy.Mode == protocol.PrivacyModeAnonymize {
		warnings = append(warnings, "automated anonymization reduces risk but cannot establish legal anonymity without contextual re-identification assessment")
	}

	policyDigest, err := digestJSON(policy)
	if err != nil {
		return protocol.SanitizeResponse{}, err
	}
	inputDigest := digestText(request.Content)
	outputDigest := digestText(output)
	createdAt := e.Now().UTC()
	counts := map[string]int{}
	for _, finding := range findings {
		counts[finding.Type]++
	}
	receiptMaterial := strings.Join([]string{request.DocumentID, string(policy.Mode), policyDigest, inputDigest, outputDigest, createdAt.Format(time.RFC3339Nano)}, "\n")
	receipt := protocol.PrivacyReceipt{
		ReceiptID:        "privacy:" + digestBytes([]byte(receiptMaterial))[7:39],
		DocumentID:       request.DocumentID,
		Mode:             policy.Mode,
		PolicyDigest:     policyDigest,
		InputDigest:      inputDigest,
		OutputDigest:     outputDigest,
		Processor:        ProcessorName,
		ProcessorVersion: ProcessorVersion,
		LocalOnly:        policy.LocalOnly,
		FailClosed:       policy.FailClosed,
		Reversible:       policy.Reversible,
		Detected:         len(selected),
		Transformed:      len(findings),
		Unresolved:       len(unresolvedCandidates),
		EntityCounts:     counts,
		ResidualRisk:     residualRisk,
		Passed:           passed,
		Warnings:         warnings,
		CreatedAt:        createdAt,
	}
	response := protocol.SanitizeResponse{Content: output, Receipt: receipt}
	if policy.Reversible {
		vault, err := sealVault(reverseMap, policy.VaultKeyID, e.VaultKey)
		if err != nil {
			return protocol.SanitizeResponse{}, err
		}
		response.Vault = vault
	}
	if !passed && policy.FailClosed {
		return response, ErrResidualRisk
	}
	return response, nil
}

func normalizePolicy(policy protocol.PrivacyPolicy) (protocol.PrivacyPolicy, error) {
	if policy.Mode == "" {
		policy.Mode = protocol.PrivacyModeAnonymize
	}
	switch policy.Mode {
	case protocol.PrivacyModeRedact, protocol.PrivacyModePseudonymize, protocol.PrivacyModeAnonymize:
	default:
		return policy, fmt.Errorf("unsupported privacy mode %q", policy.Mode)
	}
	if !policy.LocalOnly {
		return policy, errors.New("privacy processing must be localOnly")
	}
	if policy.Reversible && policy.Mode != protocol.PrivacyModePseudonymize {
		return policy, errors.New("reversible processing is allowed only in pseudonymize mode")
	}
	if policy.MinimumConfidence <= 0 {
		policy.MinimumConfidence = 0.85
	}
	if policy.MinimumConfidence > 1 {
		return policy, errors.New("minimumConfidence must be between 0 and 1")
	}
	if policy.MaxResidualRisk < 0 || policy.MaxResidualRisk > 1 {
		return policy, errors.New("maxResidualRisk must be between 0 and 1")
	}
	return policy, nil
}

func detectAll(text string, detectors []Detector, policy protocol.PrivacyPolicy) []Candidate {
	allowedEntities := stringSet(policy.EntityTypes)
	allowList := stringSetFold(policy.AllowList)
	var out []Candidate
	for _, detector := range detectors {
		for _, candidate := range detector.Detect(text) {
			if candidate.Confidence < policy.MinimumConfidence {
				continue
			}
			if len(allowedEntities) > 0 {
				if _, ok := allowedEntities[candidate.Type]; !ok {
					continue
				}
			}
			value := strings.TrimSpace(text[candidate.Start:candidate.End])
			if _, ok := allowList[strings.ToLower(value)]; ok {
				continue
			}
			out = append(out, candidate)
		}
	}
	return out
}

func resolveOverlaps(candidates []Candidate) []Candidate {
	sort.Slice(candidates, func(i, j int) bool {
		if candidates[i].Start != candidates[j].Start {
			return candidates[i].Start < candidates[j].Start
		}
		if candidates[i].Confidence != candidates[j].Confidence {
			return candidates[i].Confidence > candidates[j].Confidence
		}
		return candidates[i].End-candidates[i].Start > candidates[j].End-candidates[j].Start
	})
	selected := make([]Candidate, 0, len(candidates))
	for _, candidate := range candidates {
		overlapped := false
		for i := range selected {
			if candidate.Start >= selected[i].End || candidate.End <= selected[i].Start {
				continue
			}
			overlapped = true
			candidateScore := candidate.Confidence * float64(candidate.End-candidate.Start)
			selectedScore := selected[i].Confidence * float64(selected[i].End-selected[i].Start)
			if candidateScore > selectedScore {
				selected[i] = candidate
			}
			break
		}
		if !overlapped {
			selected = append(selected, candidate)
		}
	}
	sort.Slice(selected, func(i, j int) bool { return selected[i].Start < selected[j].Start })
	return selected
}

type transformer struct {
	policy       protocol.PrivacyPolicy
	secret       []byte
	documentSalt []byte
	seen         map[string]string
	counters     map[string]int
}

func newTransformer(policy protocol.PrivacyPolicy, secret, documentSalt []byte) (*transformer, error) {
	if policy.Mode == protocol.PrivacyModePseudonymize && len(secret) < 16 {
		return nil, errors.New("pseudonymize mode requires a secret of at least 16 bytes")
	}
	return &transformer{policy: policy, secret: secret, documentSalt: documentSalt, seen: map[string]string{}, counters: map[string]int{}}, nil
}

func (t *transformer) replacement(entity, original string) string {
	key := entity + "\x00" + original
	if existing, ok := t.seen[key]; ok {
		return existing
	}
	var replacement string
	switch t.policy.Mode {
	case protocol.PrivacyModeRedact:
		replacement = "[REDACTED:" + entity + "]"
	case protocol.PrivacyModePseudonymize:
		mac := hmac.New(sha256.New, t.secret)
		_, _ = mac.Write([]byte(t.policy.ScopeID + "\x00" + key))
		replacement = "<" + entity + ":" + strings.ToUpper(hex.EncodeToString(mac.Sum(nil)[:6])) + ">"
	case protocol.PrivacyModeAnonymize:
		t.counters[entity]++
		mac := hmac.New(sha256.New, t.documentSalt)
		_, _ = mac.Write([]byte(key))
		suffix := strings.ToUpper(hex.EncodeToString(mac.Sum(nil)[:3]))
		replacement = fmt.Sprintf("<%s_%03d_%s>", entity, t.counters[entity], suffix)
	}
	if t.policy.PreserveLength {
		replacement = fitLength(replacement, len(original))
	}
	t.seen[key] = replacement
	return replacement
}

func applyTransformations(text string, candidates []Candidate, transformer *transformer, policy protocol.PrivacyPolicy) (string, []protocol.PrivacyFinding, map[string]string) {
	var output strings.Builder
	findings := make([]protocol.PrivacyFinding, 0, len(candidates))
	reverseMap := map[string]string{}
	cursor := 0
	for _, candidate := range candidates {
		if candidate.Start < cursor || candidate.End > len(text) {
			continue
		}
		original := text[candidate.Start:candidate.End]
		replacement := transformer.replacement(candidate.Type, original)
		output.WriteString(text[cursor:candidate.Start])
		output.WriteString(replacement)
		cursor = candidate.End
		findings = append(findings, protocol.PrivacyFinding{
			Type: candidate.Type, Start: candidate.Start, End: candidate.End, Confidence: candidate.Confidence,
			Detector: candidate.Detector, Strategy: string(policy.Mode), Replacement: replacement,
		})
		if policy.Reversible {
			reverseMap[replacement] = original
		}
	}
	output.WriteString(text[cursor:])
	return output.String(), findings, reverseMap
}

func calculateResidualRisk(candidates []Candidate) float64 {
	if len(candidates) == 0 {
		return 0
	}
	maxConfidence := 0.0
	for _, candidate := range candidates {
		if candidate.Confidence > maxConfidence {
			maxConfidence = candidate.Confidence
		}
	}
	densityPenalty := float64(len(candidates)-1) * 0.03
	risk := maxConfidence + densityPenalty
	if risk > 1 {
		return 1
	}
	return risk
}

func isTextContentType(contentType string) bool {
	contentType = strings.ToLower(strings.TrimSpace(strings.Split(contentType, ";")[0]))
	if strings.HasPrefix(contentType, "text/") {
		return true
	}
	switch contentType {
	case "application/json", "application/xml", "application/yaml", "application/x-yaml", "application/csv", "":
		return true
	default:
		return false
	}
}

func digestText(value string) string { return digestBytes([]byte(value)) }

func digestBytes(value []byte) string {
	sum := sha256.Sum256(value)
	return "sha256:" + hex.EncodeToString(sum[:])
}

func digestJSON(value any) (string, error) {
	data, err := json.Marshal(value)
	if err != nil {
		return "", fmt.Errorf("encode policy: %w", err)
	}
	return digestBytes(data), nil
}

func stringSet(values []string) map[string]struct{} {
	out := make(map[string]struct{}, len(values))
	for _, value := range values {
		out[strings.TrimSpace(value)] = struct{}{}
	}
	return out
}

func stringSetFold(values []string) map[string]struct{} {
	out := make(map[string]struct{}, len(values))
	for _, value := range values {
		out[strings.ToLower(strings.TrimSpace(value))] = struct{}{}
	}
	return out
}

func fitLength(value string, length int) string {
	if length <= 0 {
		return ""
	}
	if len(value) == length {
		return value
	}
	if len(value) > length {
		return value[:length]
	}
	return value + strings.Repeat("_", length-len(value))
}
