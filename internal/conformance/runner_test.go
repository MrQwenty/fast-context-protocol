package conformance

import (
	"context"
	"log/slog"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/MrQwenty/fast-context-protocol/internal/protocol"
	"github.com/MrQwenty/fast-context-protocol/internal/server"
)

func TestRunnerAgainstReferenceProvider(t *testing.T) {
	now := time.Now().UTC()
	catalogue := server.NewCatalogueForTesting([]protocol.ContextNode{{
		ID:            "sha256:2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043362938b9824",
		Type:          "test.document",
		ContentType:   "text/plain",
		Content:       "hello",
		TokenEstimate: 2,
		ByteLength:    5,
		Priority:      1,
		Confidence:    1,
		Sensitivity:   "public",
		CreatedAt:     now,
		Provenance:    protocol.Provenance{Provider: "test", Resource: "memory://hello"},
		Intents:       []string{"*"},
		Targets:       []string{"*"},
	}})
	provider := httptest.NewServer(server.NewHTTPServer(catalogue, slog.Default()).Handler())
	defer provider.Close()

	report := New(provider.URL, 5*time.Second).Run(context.Background())
	if !report.Passed {
		t.Fatalf("expected conformance success: %+v", report.Checks)
	}
}
