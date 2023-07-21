# Install Linkerd
## Helm chart
### Installing Control Plane
Prerequisite: generate identity certificates for root certificate for mTLS

*requires `step`*

Generate Trust anchor cert
```sh
step certificate create root.linkerd.cluster.local ca.crt ca.key \
--profile root-ca --no-password --insecure --not-after 87600h
```

Issuer certificate and key
```sh
step certificate create identity.linkerd.cluster.local issuer.crt issuer.key \
--profile intermediate-ca --not-after 87600h --no-password --insecure \
--ca ca.crt --ca-key ca.key
```

Add repos
``` sh
# To add the repo for Linkerd stable releases:
helm repo add linkerd https://helm.linkerd.io/stable

# To add the repo for Linkerd edge releases:
helm repo add linkerd-edge https://helm.linkerd.io/edge
```

Install CRDs
```sh
helm install linkerd-crds linkerd/linkerd-crds \
  -n linkerd --create-namespace
```

Install Control Plane
```sh
# in the directory with generated certs
helm install linkerd-control-plane \
  -n linkerd \
  --set-file identityTrustAnchorsPEM=ca.crt \
  --set-file identity.issuer.tls.crtPEM=issuer.crt \
  --set-file identity.issuer.tls.keyPEM=issuer.key \
  linkerd/linkerd-control-plane
```

### Installing linkerd-viz
Install linkerd-viz chart
```sh
helm install linkerd-viz \
  -n linkerd-viz --create-namespace \
  linkerd/linkerd-viz
```

To add existing prometheus as metrics store via federation api
```sh
kubectl apply -f service_monitor.yaml -n [[namespace of existing prometheus]]
```
