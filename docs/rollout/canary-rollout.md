# Canary without Traffic Management

## Purpose

This lab uses Argo Rollouts to replace a standard Kubernetes Deployment with a progressive delivery workflow.

The current scenario demonstrates a canary rollout **without traffic management**.

## Current Strategy

```text
setWeight: 50
pause: 2m
run analysis
promote or abort
```

## How It Works

In this mode, Argo Rollouts approximates the canary weight by scaling the stable and canary ReplicaSets.

It does not use an ingress controller or service mesh to control exact traffic percentages.

## Pause Before Analysis

After shifting part of the workload to the canary version, the rollout pauses for two minutes.

During this period, the request generator continuously sends traffic to the application. This gives Prometheus enough time to scrape application metrics and build a useful data window before the analysis starts.

## Target Behavior

If the canary version is healthy, the rollout completes.

If the canary version generates too many HTTP 5xx responses, the analysis fails and the rollout is aborted.
