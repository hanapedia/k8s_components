# Install metallb
## Description
Metallb enables the use of loadbalancer type resource on bare metal k8s cluster
## Installation
Install metallb in the cluster
```
kubectl apply -f https://raw.githubusercontent.com/metallb/metallb/v0.13.7/config/manifests/metallb-native.yaml
```
deploy metallb resources
```
kubectl apply -f metallb-manifest.yaml
```
