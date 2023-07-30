# Qdrant 

## Objective

The objectives of this task are as follows:
- [x] Create a highly scalable Qdrant vector database hosted on AWS.
- [ ] Have automatic snapshotting and backup options available. _in progress_
- [x] Have a recovery mechanism from backup for the database.
- [ ] Develop an efficient mechanism to ingest around 1 million records in the database.
- [x] Set up observability and performance monitoring with alerts on the system. _in progress_
- [x] Use Terraform to spin up the required resources.


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

```bash
docker build --tag qdapi .
docker run -p 8000:8000 -e PORT=8000 -e QDRANT_ADDR=qdrant:6334 -d qdapi
```

## To run the Application using docker-compose
```bash
# production
docker-compose -f docker-compose.prod.yaml up -d 
# development
docker-compose up -d 
```

```bash
$ kubectl --namespace monitoring port-forward svc/prometheus-k8s 10000:9090 >/dev/null &
[1] 26130

$ kubectl --namespace monitoring port-forward svc/grafana 20000:3000 >/dev/null &
[2] 26394

$ kubectl --namespace monitoring port-forward svc/alertmanager-main 30000:9093 >/dev/null & 
[1] 26737
```

## To Run Terraform Code

### Setting up AWS Access

1. Create IAM User:
- Log in to the AWS Management Console using an account with administrative privileges.
- Navigate to the IAM service.
- Click on "Users" in the left navigation pane and create a new user.
- Add the user to a group with access to EC2. You can use an existing group with the `AmazonEC2FullAccess` policy attached, or create a custom group with the necessary EC2 permissions.
- Take note of the Access Key ID and Secret Access Key provided during the user creation process. You will need these to configure AWS CLI access.

2. Configure AWS CLI:
  - Open a terminal or command prompt on your local machine.
  - Run the following command and provide the Access Key ID and Secret Access Key when prompted:
     ```bash
     aws configure
     ```

### Running Terraform

1. Clone the Repository:
   - Clone the repository containing the Terraform code to your local machine using Git or download the code as a ZIP archive and extract it.

2. Navigate to the Terraform Configuration Folder:
   - Using the terminal or command prompt, navigate to the folder that contains the Terraform configuration files (e.g., `cd ./.terraform`).

3. Initialize Terraform:
   - Run the following command to initialize Terraform and download the necessary providers:
     ```bash
     terraform init
     ```

4. Plan the Terraform Deployment (Optional):
   - It's recommended to create a Terraform plan to preview the changes before applying them. Run the following command to generate a plan:
     ```
     terraform plan
     ```

5. Apply the Terraform Configuration:
   - If the plan looks good, apply the Terraform configuration to create the AWS EC2 instances. Run the following command and confirm the action:
     ```
     terraform apply
     ```

6. Verify the EC2 Instances:
   - Once the Terraform apply process is complete, log in to your AWS Management Console and navigate to the EC2 service. You should see the newly created EC2 instances.

### Cleaning Up

If you want to remove the resources created by Terraform, you can use the following command:

```
terraform destroy
```
