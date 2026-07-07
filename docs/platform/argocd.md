# Argo CD

## Purpose

Argo CD is the GitOps controller for this platform.

It continuously reconciles Kubernetes resources from Git into the cluster.

## Role in This Project

Argo CD will manage application manifests, including future Rollout resources.

```mermaid
flowchart LR
    Git[Git Repository] --> ArgoCD[Argo CD]
    ArgoCD --> K8s[Kubernetes API]
    K8s --> Resources[Applications and Rollouts]
```

## Responsibilities

- Sync manifests from Git
- Detect drift
- Self-heal resources
- Provide application visibility
- Manage Rollout manifests declaratively

## Design Decisions

### GitOps First

Rollout resources should eventually be managed through Git, not by manual `kubectl apply`.

This makes the project closer to how platform teams operate in real environments.

### Exposed Through Traefik

Argo CD is exposed through Traefik using host-based routing.

### Rollouts UI Extension

The Argo Rollouts UI extension is enabled in Argo CD so Rollout resources can be visualized more clearly from the Argo CD interface.

## Important Note

The Rollouts UI extension does not install the Argo Rollouts controller. The controller is installed separately.
