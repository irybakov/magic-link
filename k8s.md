# k8s install

create directory to store kafka data and define it in `k3d-cluster.yaml`

```
k3d cluster create --config k3d-cluster.yaml

kubectl apply -f kafka-k8s
```