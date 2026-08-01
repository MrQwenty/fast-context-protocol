package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/MrQwenty/fast-context-protocol/internal/protocol"
)

func main() {
	endpoint := flag.String("endpoint", "http://localhost:8080", "FCP provider base URL")
	intent := flag.String("intent", "general.answer", "intent type")
	target := flag.String("target", "", "intent target")
	maxTokens := flag.Int("max-tokens", 4000, "maximum context tokens")
	maxBytes := flag.Int64("max-bytes", 262144, "maximum context bytes")
	maxLatency := flag.Int64("max-latency-ms", 100, "target planning latency")
	flag.Parse()

	req := protocol.ContextRequest{
		RequestID:    fmt.Sprintf("req_%d", time.Now().UnixNano()),
		Intent:       protocol.Intent{Type: *intent, Target: *target},
		Consumer:     protocol.Consumer{ID: "fcpctl", ModelFamily: "generic-transformer", Modalities: []string{"text"}},
		Budget:       protocol.Budget{MaxTokens: *maxTokens, MaxBytes: *maxBytes, MaxLatencyMS: *maxLatency},
		Requirements: protocol.Requirements{MinimumConfidence: 0.5, IncludeProvenance: true},
	}
	body, err := json.Marshal(req)
	fatalIf(err)

	httpReq, err := http.NewRequest(http.MethodPost, *endpoint+"/fcp/v0.1/context/resolve", bytes.NewReader(body))
	fatalIf(err)
	httpReq.Header.Set("Content-Type", "application/fcp+json")
	httpReq.Header.Set("FCP-Version", protocol.Version)

	client := &http.Client{Timeout: 15 * time.Second}
	resp, err := client.Do(httpReq)
	fatalIf(err)
	defer resp.Body.Close()
	responseBody, err := io.ReadAll(resp.Body)
	fatalIf(err)
	if resp.StatusCode >= 300 {
		fmt.Fprintf(os.Stderr, "provider returned %s\n%s\n", resp.Status, responseBody)
		os.Exit(1)
	}
	var pretty bytes.Buffer
	if err := json.Indent(&pretty, responseBody, "", "  "); err != nil {
		_, _ = os.Stdout.Write(responseBody)
		return
	}
	fmt.Println(pretty.String())
}

func fatalIf(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
