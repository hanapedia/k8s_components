# Install istio
## With istio api
### Install the istio control plane
```
istioctl install --set profile=default -y
```
### Installing data plane components
1. create services (e.g. sample)
```
kubectl create ns foo
kubectl apply -f (istioctl kube-inject -f sample/tcp-echo.yaml | psub) -n foo
kubectl apply -f (istioctl kube-inject -f sample/sleep.yaml | psub) -n foo
```
Authorization policy if needed
```
kubectl apply -f sample/tcp-policy-authorization.yaml
```
2. create gateway and virtual services and bind them to the gateway
*Metallb must be installed first*
Create gateway resource to redirect traffic
*if you want to add new ports for ingress, edit istio-ingressgateway, then reference the port in gateway resource*
```
kubectl apply -f sample/istiogateway-manifest.yaml
```

## With kubernetes gateway api (TCP route is experimental 11/29)
- this will be the default in the near future
### Install Istio control plane with minimal profile
```
istioctl install --set profile=minimal -y
```
### Install crd for k8s gateway
```
kubectl get crd gateways.gateway.networking.k8s.io || \
  { kubectl kustomize "github.com/kubernetes-sigs/gateway-api/config/crd?ref=v0.5.1" | kubectl apply -f -; }
```
