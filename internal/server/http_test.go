package server

import (
	"bytes"
	"encoding/json"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/MrQwenty/fast-context-protocol/internal/protocol"
)

func TestDiscovery(t *testing.T) {
	s := NewHTTPServer(&Catalogue{}, slog.Default())
	req := httptest.NewRequest(http.MethodGet, "/.well-known/fcp", nil)
	rr := httptest.NewRecorder()
	s.Handler().ServeHTTP(rr, req)
	if rr.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rr.Code)
	}
	if got := rr.Header().Get("FCP-Version"); got != protocol.Version {
		t.Fatalf("expected version %s, got %s", protocol.Version, got)
	}
}

func TestResolveEndpoint(t *testing.T) {
	now := time.Now().UTC()
	catalogue := &Catalogue{Nodes: []protocol.ContextNode{{
		ID: "sha256:a", Type: "doc", Content: "hello", ContentType: "text/plain", TokenEstimate: 2,
		ByteLength: 5, Priority: 1, Confidence: 1, Sensitivity: "public", Intents: []string{"general.answer"}, CreatedAt: now,
	}}, byID: map[string]protocol.ContextNode{}}
	s := NewHTTPServer(catalogue, slog.Default())
	body, _ := json.Marshal(protocol.ContextRequest{
		RequestID: "req_1", Intent: protocol.Intent{Type: "general.answer"}, Budget: protocol.Budget{MaxTokens: 10, MaxBytes: 100},
	})
	req := httptest.NewRequest(http.MethodPost, "/fcp/v0.1/context/resolve", bytes.NewReader(body))
	req.Header.Set("FCP-Version", protocol.Version)
	rr := httptest.NewRecorder()
	s.Handler().ServeHTTP(rr, req)
	if rr.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d: %s", rr.Code, rr.Body.String())
	}
}
