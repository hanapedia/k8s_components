# Install Chaos mesh
helm is recommended for production environment
1. add repo or update
```fish
helm repo add chaos-mesh https://charts.chaos-mesh.org
helm repo update
```
2. view available charts 
```fish
helm search repo chaos-mesh

```
3. create namespace to install chaos mesh
```fish
kubectl create ns chaos-mesh

```
4. install  
```fish
# containerd
helm install chaos-mesh chaos-mesh/chaos-mesh -n=chaos-mesh --set chaosDaemon.runtime=containerd --set chaosDaemon.socketPath=/run/containerd/containerd.sock --version 2.5.1
# cri-o
helm install chaos-mesh chaos-mesh/chaos-mesh -n=chaos-mesh --set chaosDaemon.runtime=crio --set chaosDaemon.socketPath=/var/run/crio/crio.sock --version 2.5.1
```
## Upgrade, Uninstall
```fish
# upgrade
helm upgrade chaos-mesh chaos-mesh/chaos-mesh
# uninstall
helm uninstall chaos-mesh -n chaos-mesh
```

## Accessing Dashboard
0. Port forward
```fish
kubectl port-forward service/chaos-dashboard 2333 -n chaos-mesh
```
1. Create cluster role and binding
```fish
kubectl apply -f rbac.yaml
```
2. create token
```fish
kubectl create token account-cluster-manager-qwers 
```
3. view token
```fish
kubectl describe secrets account-cluster-manager-qwers
```
