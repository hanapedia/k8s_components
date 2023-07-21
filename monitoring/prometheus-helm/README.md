# Install prometheus via helm
[see docs](https://github.com/prometheus-community/helm-charts/tree/main/charts/kube-prometheus-stack)

## Installation
```
helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
helm repo update

helm install [RELEASE_NAME] prometheus-community/kube-prometheus-stack
```

## Configuration
show available configuration values
```
helm show values prometheus-community/kube-prometheus-stack
```
