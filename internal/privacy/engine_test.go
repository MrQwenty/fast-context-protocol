package privacy

import (
	"strings"
	"testing"
	"time"

	"github.com/MrQwenty/fast-context-protocol/internal/protocol"
)

func strictPolicy(mode protocol.PrivacyMode) protocol.PrivacyPolicy {
	return protocol.PrivacyPolicy{
		Mode: mode, LocalOnly: true, FailClosed: true,
		MinimumConfidence: 0.85, MaxResidualRisk: 0,
	}
}

func TestAnonymizeStructuredAndContextualPII(t *testing.T) {
	engine := NewEngine(nil, nil)
	engine.Now = func() time.Time { return time.Date(2026, 8, 3, 9, 0, 0, 0, time.UTC) }
	input := "Nome: Mario Rossi\nEmail: mario.rossi@example.com\nTelefono: +39 333 123 4567\nIBAN: IT60X0542811101000000123456\nCarta: 4111 1111 1111 1111\nData di nascita: 12/04/1985"
	response, err := engine.Sanitize(protocol.SanitizeRequest{
		DocumentID: "doc-1", ContentType: "text/plain", Content: input,
		Policy: strictPolicy(protocol.PrivacyModeAnonymize),
	})
	if err != nil {
		t.Fatalf("sanitize: %v", err)
	}
	for _, sensitive := range []string{"Mario Rossi", "mario.rossi@example.com", "+39 333 123 4567", "IT60X0542811101000000123456", "4111 1111 1111 1111", "12/04/1985"} {
		if strings.Contains(response.Content, sensitive) {
			t.Fatalf("sensitive value leaked: %q in %q", sensitive, response.Content)
		}
	}
	if !response.Receipt.Passed || response.Receipt.ResidualRisk != 0 {
		t.Fatalf("unexpected receipt: %+v", response.Receipt)
	}
	if response.Receipt.Transformed < 6 {
		t.Fatalf("expected at least 6 transformations, got %d", response.Receipt.Transformed)
	}
}

func TestPseudonymsAreStableWithinScope(t *testing.T) {
	engine := NewEngine([]byte("0123456789abcdef0123456789abcdef"), nil)
	policy := strictPolicy(protocol.PrivacyModePseudonymize)
	policy.ScopeID = "workspace-a"
	request := protocol.SanitizeRequest{DocumentID: "doc-a", ContentType: "text/plain", Content: "a@example.com and a@example.com", Policy: policy}
	first, err := engine.Sanitize(request)
	if err != nil {
		t.Fatal(err)
	}
	request.DocumentID = "doc-b"
	second, err := engine.Sanitize(request)
	if err != nil {
		t.Fatal(err)
	}
	if first.Content != second.Content {
		t.Fatalf("expected stable pseudonyms, got %q and %q", first.Content, second.Content)
	}
	if strings.Contains(first.Content, "a@example.com") {
		t.Fatal("original email leaked")
	}
}

func TestAnonymizationIsUnlinkableAcrossDocuments(t *testing.T) {
	engine := NewEngine(nil, nil)
	policy := strictPolicy(protocol.PrivacyModeAnonymize)
	first, err := engine.Sanitize(protocol.SanitizeRequest{DocumentID: "doc-a", ContentType: "text/plain", Content: "a@example.com", Policy: policy})
	if err != nil {
		t.Fatal(err)
	}
	second, err := engine.Sanitize(protocol.SanitizeRequest{DocumentID: "doc-b", ContentType: "text/plain", Content: "a@example.com", Policy: policy})
	if err != nil {
		t.Fatal(err)
	}
	if first.Content == second.Content {
		t.Fatalf("anonymized identifiers should not be linkable across documents: %q", first.Content)
	}
}

func TestCustomDictionaryAndAllowList(t *testing.T) {
	engine := NewEngine(nil, nil)
	policy := strictPolicy(protocol.PrivacyModeRedact)
	policy.CustomTerms = []string{"Project Aurora", "ACME Internal"}
	policy.AllowList = []string{"ACME Internal"}
	response, err := engine.Sanitize(protocol.SanitizeRequest{DocumentID: "doc", ContentType: "text/plain", Content: "Project Aurora for ACME Internal", Policy: policy})
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(response.Content, "Project Aurora") {
		t.Fatal("custom identifier was not redacted")
	}
	if !strings.Contains(response.Content, "ACME Internal") {
		t.Fatal("allow-listed term was transformed")
	}
}

func TestReversibleVaultDoesNotExposeOriginals(t *testing.T) {
	vaultKey := []byte("0123456789abcdef0123456789abcdef")
	engine := NewEngine([]byte("abcdef0123456789abcdef0123456789"), vaultKey)
	policy := strictPolicy(protocol.PrivacyModePseudonymize)
	policy.Reversible = true
	policy.VaultKeyID = "local-key-1"
	response, err := engine.Sanitize(protocol.SanitizeRequest{DocumentID: "doc", ContentType: "text/plain", Content: "owner@example.com", Policy: policy})
	if err != nil {
		t.Fatal(err)
	}
	if response.Vault == nil {
		t.Fatal("expected sealed vault")
	}
	if strings.Contains(response.Vault.Ciphertext, "owner@example.com") {
		t.Fatal("vault ciphertext exposes original")
	}
	mapping, err := OpenVault(*response.Vault, vaultKey)
	if err != nil {
		t.Fatal(err)
	}
	if len(mapping) != 1 {
		t.Fatalf("unexpected mapping: %#v", mapping)
	}
	for _, original := range mapping {
		if original != "owner@example.com" {
			t.Fatalf("unexpected original %q", original)
		}
	}
}

func TestFailClosedOnUnsupportedBinaryDocument(t *testing.T) {
	engine := NewEngine(nil, nil)
	_, err := engine.Sanitize(protocol.SanitizeRequest{DocumentID: "doc", ContentType: "application/pdf", Content: "%PDF", Policy: strictPolicy(protocol.PrivacyModeAnonymize)})
	if err == nil || !strings.Contains(err.Error(), "extract text locally") {
		t.Fatalf("expected local extraction error, got %v", err)
	}
}
