package server

import (
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"strings"
	"sync"

	"github.com/MrQwenty/fast-context-protocol/internal/protocol"
)

type HTTPServer struct {
	catalogue *Catalogue
	resolver  *Resolver
	logger    *slog.Logger
	mu        sync.Mutex
	receipts  []protocol.Receipt
}

func NewHTTPServer(c *Catalogue, logger *slog.Logger) *HTTPServer {
	return &HTTPServer{catalogue: c, resolver: NewResolver(c), logger: logger}
}

func (s *HTTPServer) Handler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /.well-known/fcp", s.discovery)
	mux.HandleFunc("POST /fcp/v0.1/context/resolve", s.resolve)
	mux.HandleFunc("GET /fcp/v0.1/context/{id}", s.fetch)
	mux.HandleFunc("POST /fcp/v0.1/receipts", s.receipt)
	mux.HandleFunc("GET /healthz", func(w http.ResponseWriter, _ *http.Request) { w.WriteHeader(http.StatusNoContent) })
	return s.versionMiddleware(mux)
}

func (s *HTTPServer) versionMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/healthz" && r.URL.Path != "/.well-known/fcp" {
			requested := r.Header.Get("FCP-Version")
			if requested != "" && requested != protocol.Version {
				w.Header().Set("FCP-Supported-Versions", protocol.Version)
				problem(w, http.StatusUpgradeRequired, "unsupported-version", "Unsupported FCP version", "Requested FCP version is not supported")
				return
			}
		}
		w.Header().Set("FCP-Version", protocol.Version)
		next.ServeHTTP(w, r)
	})
}

func (s *HTTPServer) discovery(w http.ResponseWriter, _ *http.Request) {
	writeJSON(w, http.StatusOK, protocol.Discovery{
		Name:            "FCP reference provider",
		Description:     "Minimal Fast Context Protocol 0.1 provider",
		Versions:        []string{protocol.Version},
		Endpoints:       map[string]string{"resolve": "/fcp/v0.1/context/resolve", "fetch": "/fcp/v0.1/context/{digest}", "receipts": "/fcp/v0.1/receipts"},
		DeliveryModes:   []protocol.DeliveryMode{protocol.DeliveryInline, protocol.DeliveryReference, protocol.DeliveryFetch},
		Features:        protocol.Features{KnownContext: true, Fetch: true, Receipts: true},
		MaxRequestBytes: 1 << 20,
	})
}

func (s *HTTPServer) resolve(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	limited := io.LimitReader(r.Body, 1<<20)
	var req protocol.ContextRequest
	dec := json.NewDecoder(limited)
	dec.DisallowUnknownFields()
	if err := dec.Decode(&req); err != nil {
		problem(w, http.StatusBadRequest, "invalid-request", "Invalid context request", err.Error())
		return
	}
	plan, err := s.resolver.Resolve(req)
	if err != nil {
		problem(w, http.StatusUnprocessableEntity, "unresolvable-request", "Context request cannot be resolved", err.Error())
		return
	}
	writeJSON(w, http.StatusOK, plan)
}

func (s *HTTPServer) fetch(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if !strings.HasPrefix(id, "sha256:") {
		problem(w, http.StatusBadRequest, "invalid-digest", "Invalid context digest", "Expected sha256:<hex>")
		return
	}
	node, ok := s.catalogue.Get(id)
	if !ok {
		problem(w, http.StatusNotFound, "unknown-context", "Context not found", "No authorized context exists for this digest")
		return
	}
	writeJSON(w, http.StatusOK, node)
}

func (s *HTTPServer) receipt(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var receipt protocol.Receipt
	if err := json.NewDecoder(io.LimitReader(r.Body, 1<<20)).Decode(&receipt); err != nil {
		problem(w, http.StatusBadRequest, "invalid-receipt", "Invalid receipt", err.Error())
		return
	}
	if receipt.RequestID == "" || receipt.PlanID == "" {
		problem(w, http.StatusUnprocessableEntity, "invalid-receipt", "Invalid receipt", "requestId and planId are required")
		return
	}
	s.mu.Lock()
	s.receipts = append(s.receipts, receipt)
	s.mu.Unlock()
	w.WriteHeader(http.StatusAccepted)
}

func writeJSON(w http.ResponseWriter, status int, value any) {
	w.Header().Set("Content-Type", "application/fcp+json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(value)
}

func problem(w http.ResponseWriter, status int, code, title, detail string) {
	w.Header().Set("Content-Type", "application/problem+json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(map[string]any{
		"type":   fmt.Sprintf("https://fcp.dev/problems/%s", code),
		"title":  title,
		"status": status,
		"detail": detail,
	})
}
