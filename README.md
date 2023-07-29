# Qdrant 

## Objective

The objectives of this task are as follows:
- [ ] Create a highly scalable Qdrant vector database hosted on AWS.
- [ ] Have automatic snapshotting and backup options available.
- [ ] Have a recovery mechanism from backup for the database.
- [ ] Develop an efficient mechanism to ingest around 1 million records in the database.
- [ ] Set up observability and performance monitoring with alerts on the system.
- [ ] Use Terraform to spin up the required resources.


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
