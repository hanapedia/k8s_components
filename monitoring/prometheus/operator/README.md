# Prometheus Operator
Installation with kube-prometheus, which uses jsonnet to generate manifests for related tools such as:
- The Prometheus Operator
- HA Prometheus
- HA Alertmanager
- Grafana
- Prometheus Adapater for Kubernetes metrics API
  - Formats metrics collected by Prometheus into metrics API of Kubernetes 
  - Kubernetes metrics servers are used for autoscaling
- Prometheus Node Exporters
  - Export hardware and OS level metrics
- Kube-state-metrics
  - Export metrics of various K8s objects such as pods, depoloyments, and nodes
- Blackbox Exporters
  - Export metrics for "blackbox" probing
  - monitors different endpoints for availability

## Installation
1. Configure components to include in the example.jsonnet by referring to docs
2. generate manifests
```
$ bash build.sh
```
3. apply the manifests
```
$ kubectl apply --server-side -f manifests/setup
$ kubectl apply --server-side -f manifests/
```

## Sevice Monitor, PodMonitor, exporters CRD
ServiceMonitor and PodMonitor custum resources describe the target to be monitored by Prometheus

To configure Prometheus deployment with ServiceMonitor, they must be specified in the servicemonitor selector on Prometheus custom resource
