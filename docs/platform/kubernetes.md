# Kubernetes

## Purpose

Kubernetes is the runtime platform for the Progressive Delivery lab.

It provides the base primitives required by the rest of the platform:

- Pods
- Services
- Deployments
- Custom Resource Definitions
- Controllers
- Namespaces

## Role in This Project

Kubernetes is not the main subject of the project. It is the execution layer where GitOps, ingress routing, observability, and progressive delivery are demonstrated.

## Platform Components Running on Kubernetes

- MetalLB
- Traefik
- Argo CD
- Argo Rollouts
- Prometheus
- Grafana
- Future demo application

## Design Note

The cluster is treated as an existing platform target. The project focuses on release engineering and progressive delivery rather than cluster provisioning.
