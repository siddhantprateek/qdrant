# Qdrant 

## Objective

The objectives of this task are as follows:
- [ ] Create a highly scalable Qdrant vector database hosted on AWS. _in progress_
- [ ] Have automatic snapshotting and backup options available. _in progress_
- [x] Have a recovery mechanism from backup for the database.
- [ ] Develop an efficient mechanism to ingest around 1 million records in the database.
- [ ] Set up observability and performance monitoring with alerts on the system. _in progress_
- [ ] Use Terraform to spin up the required resources. _in progress_


## Tech Stack

- `Go` 
- `Go Fiber` - Go Framework
- `Prometheus`
- `Grafana`
- `Qdrant` - Vector Database
- `Terraform` - Infra-structure as Code (IaC)
- `Aws` - Cloud Provider
- `Kube-Prometheus` deploys the Prometheus Operator and already schedules a Prometheus
called prometheus-k8s with alerts and rules by default.


## K8s Monitoring Pods
```bash
$ kubectl get pods -n monitoring  
NAME                                   READY   STATUS              RESTARTS   AGE
blackbox-exporter-7d8c77d7b9-p4txc     0/3     ContainerCreating   0          31s
grafana-79f47474f7-tsrpc               0/1     Running             0          30s
kube-state-metrics-8cc8f7df6-wslgq     0/3     ContainerCreating   0          30s
node-exporter-bd97l                    0/2     ContainerCreating   0          29s
prometheus-adapter-6b88dfd544-4rr57    0/1     ContainerCreating   0          29s
prometheus-adapter-6b88dfd544-vhb98    0/1     ContainerCreating   0          29s
prometheus-operator-557b4f4977-q76cz   0/2     ContainerCreating   0          29s
```


## For Recovery Mechanism in Database

In the StatefulSet configuration, I have used `volumeClaimTemplates` section to define the PVC template that will be used by each replica of the StatefulSet. Each replica will have its own PVC `PersistantVolumeClaim` with its unique identity, backed by the requested storage.

With this configuration, the Qdrant vector database instances will have their data persisted across restarts and rescheduling events, providing data durability and stability for your deployment.

```yaml
apiVersion: apps/v1
kind: StatefulSet
metadata:
  name: qdrant-db
spec:
  selector:
    matchLabels:
      app: qdrant-db
  serviceName: qdrant-db
  replicas: 3
  template:
    metadata:
      labels:
        app: qdrant-db
    spec:
      containers:
      - name: qdrant-db
        image: qdrant/qdrant
        ports:
        - containerPort: 6333
          name: web
        - containerPort: 6334
          name: grpc        
        volumeMounts:
        - name: qdrant-data
          mountPath: /data
  volumeClaimTemplates:
  - metadata:
      name: qdrant-data
    spec:
      accessModes: [ "ReadWriteOnce" ]
        requests:
      resources:
          storage: 10Gi
```