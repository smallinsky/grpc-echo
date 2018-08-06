# grpc-echo  

gRPC microservice for testing purpose only

## Setup:
### Create k8 service for echo-server and echo-client
```bash
kubectl apply -f manifests/ehco-server/service.yaml
kubectl apply -f manifests/echo-client/service.yaml
```

### Create service deployment 
```bash
kubectl apply -f manifests/ehco-server/deployment.yaml
kubectl apply -f manifests/echo-client/deployment.yaml
```
