# Demo Application

## Purpose

The demo application is a small Go-based HTTP API designed to validate progressive delivery workflows.

It is intentionally simple. The goal is not to demonstrate business functionality, but to provide a controllable workload for testing canary rollout behavior, Prometheus-based analysis, and automatic abort scenarios.

## Endpoints

| Endpoint | Purpose |
|---|---|
| `/` | Main application endpoint |
| `/metrics` | Prometheus metrics endpoint |

## Configuration Profiles

The application behavior is controlled through two Kubernetes ConfigMaps.

| ConfigMap | Purpose |
|---|---|
| `healthy-app` | Simulates a healthy application with no intentional failures |
| `unhealthy-app` | Simulates an unhealthy application by enabling configurable request failures |

## Runtime Configuration

| Variable | Example | Description |
|---|---|---|
| `APP_HEALTH` | `healthy` or `unhealthy` | Controls whether the application behaves normally or simulates failures |
| `APP_FAILURE_RATE` | `0`, `30`, `70` | Percentage of requests that should return HTTP 500 when the app is unhealthy |

## Metrics

```text
demo_api_http_requests_total
demo_api_http_request_duration_seconds
```

The rollout analysis uses `demo_api_http_requests_total` to calculate the HTTP 5xx error percentage during the canary phase.
