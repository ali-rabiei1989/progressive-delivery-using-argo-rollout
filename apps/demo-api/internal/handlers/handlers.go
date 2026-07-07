package handlers

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"sync/atomic"
	"time"
)

type Handler struct {
	version          string
	errorEnabled    atomic.Bool
	latencyEnabled  atomic.Bool
	healthMode      string
	failureRate     int
}

func New(version string) *Handler {

	failureRate, err := strconv.Atoi(os.Getenv("APP_FAILURE_RATE"))
	if err != nil || failureRate < 0 || failureRate > 100 {
		failureRate = 0
	}
	return &Handler{
		version:      version,
		healthMode:   getenv("APP_HEALTH", "healthy"),
		failureRate:  failureRate,
	}

}

func (h *Handler) Home(w http.ResponseWriter, r *http.Request) {
	if h.shouldFail() {
		writeJSON(w, http.StatusInternalServerError, map[string]string{
			"error":   "simulated failure",
			"version": h.version,
		})
		return
	}

	if h.maybeFailOrDelay(w) {
		return
	}

	writeJSON(w, http.StatusOK, map[string]string{
		"app":     "demo-api",
		"version": h.version,
		"status":  "running",
	})
}

func (h *Handler) Health(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{
		"status": "healthy",
	})
}

func (h *Handler) Ready(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{
		"status": "ready",
	})
}

func (h *Handler) Version(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{
		"version": h.version,
	})
}

func (h *Handler) EnableError(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeJSON(w, http.StatusMethodNotAllowed, map[string]string{
			"error": "method not allowed",
		})
		return
	}

	h.errorEnabled.Store(true)

	writeJSON(w, http.StatusOK, map[string]string{
		"simulate_error": "enabled",
	})
}

func (h *Handler) EnableLatency(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeJSON(w, http.StatusMethodNotAllowed, map[string]string{
			"error": "method not allowed",
		})
		return
	}

	h.latencyEnabled.Store(true)

	writeJSON(w, http.StatusOK, map[string]string{
		"simulate_latency": "enabled",
	})
}

func (h *Handler) Reset(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeJSON(w, http.StatusMethodNotAllowed, map[string]string{
			"error": "method not allowed",
		})
		return
	}

	h.errorEnabled.Store(false)
	h.latencyEnabled.Store(false)

	writeJSON(w, http.StatusOK, map[string]string{
		"simulate_error":   "disabled",
		"simulate_latency": "disabled",
	})
}

func (h *Handler) maybeFailOrDelay(w http.ResponseWriter) bool {
	if h.latencyEnabled.Load() {
		time.Sleep(3 * time.Second)
	}

	if h.errorEnabled.Load() {
		writeJSON(w, http.StatusInternalServerError, map[string]string{
			"error": "simulated internal server error",
		})
		return true
	}

	return false
}

func writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(data)
}

func (h *Handler) shouldFail() bool {
	if h.healthMode != "unhealthy" {
		return false
	}

	if h.failureRate <= 0 {
		return true
	}

	return rand.Intn(100) < h.failureRate
}

func getenv(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}