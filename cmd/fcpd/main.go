package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/MrQwenty/fast-context-protocol/internal/server"
)

func main() {
	listen := flag.String("listen", ":8080", "HTTP listen address")
	catalog := flag.String("catalog", "examples/basic-provider/context.json", "path to context catalogue")
	flag.Parse()

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	catalogue, err := server.LoadCatalogue(*catalog)
	if err != nil {
		logger.Error("failed to load catalogue", "error", err)
		os.Exit(1)
	}

	httpServer := &http.Server{
		Addr:              *listen,
		Handler:           server.NewHTTPServer(catalogue, logger).Handler(),
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       15 * time.Second,
		WriteTimeout:      30 * time.Second,
		IdleTimeout:       60 * time.Second,
	}
	logger.Info("FCP provider listening", "address", *listen, "nodes", len(catalogue.Nodes))
	if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logger.Error("server stopped", "error", err)
		os.Exit(1)
	}
}
