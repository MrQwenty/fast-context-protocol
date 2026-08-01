package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/MrQwenty/fast-context-protocol/internal/conformance"
)

func main() {
	endpoint := flag.String("endpoint", "http://localhost:8080", "FCP provider base URL")
	timeout := flag.Duration("timeout", 15*time.Second, "timeout for each HTTP request")
	compact := flag.Bool("compact", false, "emit compact JSON")
	flag.Parse()

	report := conformance.New(*endpoint, *timeout).Run(context.Background())
	encoder := json.NewEncoder(os.Stdout)
	if !*compact {
		encoder.SetIndent("", "  ")
	}
	if err := encoder.Encode(report); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
	}
	if !report.Passed {
		os.Exit(1)
	}
}
