# Install and configure cert manager
## install with kubectl
```
kubectl apply -f https://github.com/cert-manager/cert-manager/releases/download/v1.10.1/cert-manager.yaml
```

## Create issuer/clusterIssuer resource
issuer resource represents certificate signing authority(CA)
### with dns challenge
*retrieve access key for aws iam role*
```
kubectl create secret route53-secret --namespace=cert-manager --from-literal=secret-access-key=<Your ACCESS Key>
```
create ClusterIssuer
```
kubectl apply -f issuer-manifest.yaml
```

## Create certificate resource with issuer
