## Connecting Frontend to Backend Services

This is a sample project deploying a frontend and backend applications on
kubernetes with a stateless redis DB backend.

There are two apps that are deployed on Minikube in this project:

frontend app is a dummy proxy. It's originated from nginx and proxies all the
traffic to the backend service.

backend app is the same app with [k8s-multipod](../k8s-multipod) with some small
modifications. It's a simple go application that prints several information about the pods it's running on.

With version 3, redis DB backend has been implemented.

Deployment and Service yaml files are under [kubernetes](kubernetes) directory.

### Prerequisites:
* minikube and kubectl should be installed on your system.
* Make sure that your minikube is up and running, and the kubectl is correctly
  configured to point minikube:
```sh
$> minikube status:
minikube: Running
cluster: Running
kubectl: Correctly Configured: pointing to minikube-vm at 192.168.99.100
```
* Current version is tested on minikube v0.26.1 and kubernetes v1.10.1. There are
  various Workloads API changes in Kubernetes v1.8 and 1.9, this project uses
  the latest API specifications. See [Reference Documentation](https://kubernetes.io/docs/reference/workloads-18-19/) for details of the API change.

* Backend app requires a multi-stage build, which comes with Docker 17.05 and
  higher versions. If you have the latest minikube, it should be alright. But in
  case you see any problem, check out the [docker documentation on multi-stage
  builds](https://docs.docker.com/develop/develop-images/multistage-build/)

### For the impatient:
* `make deploy` command will build the images and deploy the project to your minikube VM.
```sh
$> make deploy
```

* To get detailed info about make targets and flags
```sh
$> make help
```

### Building the docker images

```sh
$> make build
```
This command builds both frontend and backend apps' images.

#### frontend

```sh
$> make build-fe
```
If you want to change the name or the version of the frontend app
```sh
$> FE_APPNAME=<new_name> FE_VERSION=<new_version> make build-fe
```
When you change the name and/or version of the app, do not forget to update `template.spec.containers.image` value in Deployment descriptions.

#### backend

```sh
$> make build-be
```
If you want to change the name or the version of the frontend app
```sh
$> BE_APPNAME=<new_name> BE_VERSION=<new_version> make build-be
```
When you change the name and/or version of the app, do not forget to update `template.spec.containers.image` value in Deployment descriptions.

### Deploying on Kubernetes

Before deploying the project, a namespace and a context for the project is
created (default: NAMESPACE=fe-be-dev, CTX=dev). And kubectl starts using the
new context. 

```sh
$> make
```
or
```sh
$> make deploy
```

### Testing the App

When you deploy all three parts of the system (redis, backend and frontend), you
can curl the frontend service url

```sh
$> curl $(minikube service -n <namespace> --url frontend)/health_check
{"alive": true, "redis_conn": "good"}
```
```sh
$> curl $(minikube service -n <namespace> --url frontend)
Hello from Kubernetes! Running on go-app-548654c765-vzxbb | version: 0.3
Total number of requests to this pod:4
Total number of requests in all system:7
App Uptime: 27m42.059174051s
Log Time: 2018-04-30 04:15:22.600948094 +0000 UTC m=+1662.061687640
```

For more details about the app, check [Backend README.md](backend/README.md).

### Cleanup
```sh
$> make clean
```
It will remove all the resources that have been created on your minikube.

### TODO List
- [ ] Serve the API through SSL (check nginx ingress)
- [ ] Authentication, Authorization
- [ ] Instead of a stateless DB configuration, check StatefulSet for Redis
  high availability and resilience
  - [ ] Separate PV
  - [ ] redis-slave pods and redis-sentinel for master election?
- [ ] Use [Horizontal Pod
  Autoscaler](https://kubernetes.io/docs/tasks/run-application/horizontal-pod-autoscale/)
  for auto scaling the frontend and backend apps
- [x] Automate the cluster setup (Makefile and/or terraform files would be quite
  handy)
- [x] Deploy all the apps to a specific namespace that can be configured to be used
  on different environments.
- [ ] Write smoke tests for the deployment

#### Notes:

frontend Service is of type `LoadBalancer`, however, as there is no external
Load Balancer created in local environment, frontend LoadBalancer service will
be in `<pending>` state for External-IP
```sh
$> kubectl get services
NAME           TYPE           CLUSTER-IP      EXTERNAL-IP   PORT(S)        AGE
frontend       LoadBalancer   10.99.20.153    <pending>     80:32695/TCP   2h
go-app         ClusterIP      10.103.117.39   <none>        80/TCP         2h
kubernetes     ClusterIP      10.96.0.1       <none>        443/TCP        2h
redis-master   ClusterIP      10.110.233.7    <none>        6379/TCP       2h
```
You don't need to wait for it, the `<pending>` state will not change. As
LoadBalancer services get a node port assigned too, the `minikube service --url
frontend` command will return the service URL. Reference [here](https://github.com/kubernetes/minikube/issues/384#issuecomment-234409957)

#### References:
* [Connect a Front End to a Back End Using a Service](https://kubernetes.io/docs/tasks/access-application-cluster/connecting-frontend-backend/)
