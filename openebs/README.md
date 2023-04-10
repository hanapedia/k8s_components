# Kubernetes persistent volume manifests

## Description
- Dynamic Persistent volume creation via Storage classes provided by OpenEBS
- The following steps will create LocalPV with hostpath set at /var/openebs/local

## steps
- Install iSCI on each working nodes and enable iscid
```
ansible-playbook setup_openebs.yaml
```
- Install OpenEBS operator
```
kubectl apply -f https://openebs.github.io/charts/openebs-operator.yaml
```
