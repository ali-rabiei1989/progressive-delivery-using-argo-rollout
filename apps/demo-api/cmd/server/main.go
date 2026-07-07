package main

import (
	"github.com/ali-rabiei1989/enterprise-argo-rollouts-platform/apps/demo-api/internal/handlers"
	"github.com/ali-rabiei1989/enterprise-argo-rollouts-platform/apps/demo-api/internal/middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
	"os"
)

func main() {
	appVersion := getenv("APP_VERSION", "v1")
	port := getenv("PORT", "8080")

	mux := http.NewServeMux()

	h := handlers.New(appVersion)

	mux.HandleFunc("/", h.Home)
	mux.HandleFunc("/health", h.Health)
	mux.HandleFunc("/ready", h.Ready)
	mux.HandleFunc("/version", h.Version)
	mux.HandleFunc("/simulate/error", h.EnableError)
	mux.HandleFunc("/simulate/latency", h.EnableLatency)
	mux.HandleFunc("/reset", h.Reset)
	mux.Handle("/metrics", promhttp.Handler())

	log.Printf("demo-api starting on port %s, version=%s", port, appVersion)
	log.Fatal(http.ListenAndServe(":"+port, middleware.Metrics(appVersion, mux)))
}

func getenv(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}
