# Rook and Ceph
- Rook is an open source cloud-native storage orchestrator.
- Ceph is a distributed storage system that provides file, block and object storage.

## Installation using Helm
1. Install Rook Operator Chart
- add configuration in `rook-operator-values.yaml`
```
helm repo add rook-release https://charts.rook.io/release
helm install --create-namespace --namespace rook-ceph rook-ceph rook-release/rook-ceph -f rook-operator-values.yaml
```

2. Install Ceph Cluster Helm Chart
- add configuration in `rook-ceph-cluster-values.yaml`
```
helm repo add rook-release https://charts.rook.io/release
helm install --create-namespace --namespace rook-ceph rook-ceph-cluster \
   --set operatorNamespace=rook-ceph rook-release/rook-ceph-cluster -f rook-ceph-cluster-values.yaml
```

## Default configurations
*check https://github.com/rook/rook/tree/master/deploy/charts for latest configuration values*
- `defaults/rook-operator-values.yaml`
- `defaults/rook-ceph-cluster.yaml`
