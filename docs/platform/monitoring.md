# Monitoring

## Purpose

Monitoring provides visibility into platform and application behavior.

This project uses:

- Prometheus
- Grafana

## Role in This Project

Prometheus collects metrics that will later be used by Argo Rollouts AnalysisTemplates.

Grafana provides dashboards for observing platform and application behavior.

## Why Monitoring Is Installed Before the Demo App

Progressive Delivery requires feedback.

Without metrics, a rollout can only move forward based on static timing. With Prometheus metrics, the rollout can make release decisions based on actual system behavior.

## Target Phase 2 Integration

```mermaid
flowchart LR
    App[Demo Application] --> Metrics[/metrics endpoint]
    Metrics --> Prometheus[Prometheus]
    Prometheus --> AnalysisTemplate[Argo Rollouts AnalysisTemplate]
    AnalysisTemplate --> Decision[Promote or Rollback]
    Prometheus --> Grafana[Grafana Dashboards]
```

## Metrics Planned for Phase 2

- HTTP success rate
- HTTP 5xx error rate
- Request latency
- Pod health
- Rollout status
