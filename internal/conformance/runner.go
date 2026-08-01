package conformance

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/MrQwenty/fast-context-protocol/internal/protocol"
)

type Check struct {
	Name       string `json:"name"`
	Passed     bool   `json:"passed"`
	Detail     string `json:"detail,omitempty"`
	DurationMS int64  `json:"durationMs"`
}

type Report struct {
	Endpoint        string    `json:"endpoint"`
	ProtocolVersion string    `json:"protocolVersion"`
	Passed          bool      `json:"passed"`
	StartedAt       time.Time `json:"startedAt"`
	DurationMS      int64     `json:"durationMs"`
	Checks          []Check   `json:"checks"`
}

type Runner struct {
	Endpoint string
	Client   *http.Client
}

func New(endpoint string, timeout time.Duration) *Runner {
	return &Runner{
		Endpoint: strings.TrimRight(endpoint, "/"),
		Client:   &http.Client{Timeout: timeout},
	}
}

func (r *Runner) Run(ctx context.Context) Report {
	started := time.Now().UTC()
	report := Report{
		Endpoint:        r.Endpoint,
		ProtocolVersion: protocol.Version,
		Passed:          true,
		StartedAt:       started,
	}

	checks := []struct {
		name string
		fn   func(context.Context) error
	}{
		{name: "discovery", fn: r.checkDiscovery},
		{name: "resolve", fn: r.checkResolve},
		{name: "unsupported-version", fn: r.checkUnsupportedVersion},
		{name: "malformed-request", fn: r.checkMalformedRequest},
	}

	for _, item := range checks {
		checkStarted := time.Now()
		err := item.fn(ctx)
		check := Check{
			Name:       item.name,
			Passed:     err == nil,
			DurationMS: time.Since(checkStarted).Milliseconds(),
		}
		if err != nil {
			check.Detail = err.Error()
			report.Passed = false
		}
		report.Checks = append(report.Checks, check)
	}
	report.DurationMS = time.Since(started).Milliseconds()
	return report
}

func (r *Runner) checkDiscovery(ctx context.Context) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, r.Endpoint+"/.well-known/fcp", nil)
	if err != nil {
		return err
	}
	resp, err := r.Client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("expected 200, got %d", resp.StatusCode)
	}
	if resp.Header.Get("FCP-Version") != protocol.Version {
		return fmt.Errorf("missing or invalid FCP-Version response header")
	}
	var discovery protocol.Discovery
	if err := decodeJSON(resp.Body, &discovery); err != nil {
		return fmt.Errorf("decode discovery: %w", err)
	}
	if !contains(discovery.Versions, protocol.Version) {
		return fmt.Errorf("discovery does not advertise version %s", protocol.Version)
	}
	if discovery.Endpoints["resolve"] == "" {
		return fmt.Errorf("discovery does not advertise resolve endpoint")
	}
	return nil
}

func (r *Runner) checkResolve(ctx context.Context) error {
	request := protocol.ContextRequest{
		RequestID: "conformance_resolve",
		Intent:    protocol.Intent{Type: "general.answer"},
		Consumer:  protocol.Consumer{ID: "fcpconform", ModelFamily: "generic-transformer", Modalities: []string{"text"}},
		Budget:    protocol.Budget{MaxTokens: 1024, MaxBytes: 64 * 1024, MaxLatencyMS: 1000},
	}
	body, err := json.Marshal(request)
	if err != nil {
		return err
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, r.Endpoint+"/fcp/v0.1/context/resolve", bytes.NewReader(body))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/fcp+json")
	req.Header.Set("FCP-Version", protocol.Version)
	resp, err := r.Client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		data, _ := io.ReadAll(io.LimitReader(resp.Body, 4096))
		return fmt.Errorf("expected 200, got %d: %s", resp.StatusCode, strings.TrimSpace(string(data)))
	}
	var plan protocol.ContextPlan
	if err := decodeJSON(resp.Body, &plan); err != nil {
		return fmt.Errorf("decode plan: %w", err)
	}
	if plan.RequestID != request.RequestID {
		return fmt.Errorf("plan requestId mismatch")
	}
	if plan.ProtocolVersion != protocol.Version {
		return fmt.Errorf("plan protocolVersion mismatch")
	}
	if plan.PlanID == "" || plan.ContextRoot == "" {
		return fmt.Errorf("planId and contextRoot are required")
	}
	if plan.EstimatedTokens > request.Budget.MaxTokens || plan.EstimatedBytes > request.Budget.MaxBytes {
		return fmt.Errorf("provider exceeded declared budget")
	}
	return nil
}

func (r *Runner) checkUnsupportedVersion(ctx context.Context) error {
	body := []byte(`{"requestId":"conformance_version","intent":{"type":"general.answer"},"budget":{"maxTokens":1,"maxBytes":1}}`)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, r.Endpoint+"/fcp/v0.1/context/resolve", bytes.NewReader(body))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/fcp+json")
	req.Header.Set("FCP-Version", "99.0")
	resp, err := r.Client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusUpgradeRequired {
		return fmt.Errorf("expected 426, got %d", resp.StatusCode)
	}
	if !strings.Contains(resp.Header.Get("FCP-Supported-Versions"), protocol.Version) {
		return fmt.Errorf("supported versions header does not include %s", protocol.Version)
	}
	return nil
}

func (r *Runner) checkMalformedRequest(ctx context.Context) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, r.Endpoint+"/fcp/v0.1/context/resolve", strings.NewReader(`{"unknown":true}`))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/fcp+json")
	req.Header.Set("FCP-Version", protocol.Version)
	resp, err := r.Client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusBadRequest {
		return fmt.Errorf("expected 400, got %d", resp.StatusCode)
	}
	if !strings.HasPrefix(resp.Header.Get("Content-Type"), "application/problem+json") {
		return fmt.Errorf("expected application/problem+json error")
	}
	return nil
}

func decodeJSON(reader io.Reader, value any) error {
	decoder := json.NewDecoder(io.LimitReader(reader, 2<<20))
	decoder.DisallowUnknownFields()
	return decoder.Decode(value)
}

func contains(values []string, expected string) bool {
	for _, value := range values {
		if value == expected {
			return true
		}
	}
	return false
}
