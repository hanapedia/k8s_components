# Elasticsearch installation
## Elastic cloud on kubernetes
1. install crd
```
kubectl create -f https://download.elastic.co/downloads/eck/2.5.0/crds.yaml
```
2. Install the operator with its RBAC rules
```
kubectl apply -f https://download.elastic.co/downloads/eck/2.5.0/operator.yaml
```
3. Monitor operator logs
```
kubectl -n elastic-system logs -f statefulset.apps/elastic-operator
```

## Deploy elasticsearch cluster
```
kubectl apply -f elasticsearch-manifest.yaml -n observability
```
monitor elasticsearch cluster
```
kubectl get elasticsearch
kubectl get pods --selector='elasticsearch.k8s.elastic.co/cluster-name={es cluster name}'
kubectl logs -f podname
```
request elasticsearch access
```
kubectl get service {es cluster name}-es-http
PASSWORD=$(kubectl get secret {es cluster name}-es-elastic-user -o go-template='{{.data.elastic | base64decode}}')
kubectl port-forward service/{es cluster name}-es-http 9200
curl -u "elastic:$PASSWORD" -k "https://localhost:9200"
```

## Deploy Kibana instance
```
kubectl apply -f kibana-manifest.yaml
```
monitor kibana instance
```
kubectl get kibana
kubectl get pod --selector='kibana.k8s.elastic.co/name={kibana name}'
```
access kibana
```
kubectl get service {kibana name}-kb-http
kubectl port-forward service/{kibana name}-kb-http 5601
kubectl get secret {es cluster name}-es-elastic-user -o=jsonpath='{.data.elastic}' | base64 --decode; echo
```
