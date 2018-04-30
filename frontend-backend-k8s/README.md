## Connecting Frontend to Backend Services

In this example, two apps are built in docker and deployed on Minikube.

frontend app is a dummy proxy. It's originated from nginx and proxies all the
traffic to the backend service.

backend app is the same app with [k8s-multipod](../k8s-multipod) It's a simple
go application that prints several information about the pods it's running on.

With go-app v3, redis DB backend has been implemented.

### Building the docker images

If you're not going to push those docker images to docker hub, and just want
Minikube VM to reach those images, run the following command on your terminal
first:
```
$> eval $(minikube docker-env)
```
So that minikube can build those images in its' local registry and reach them
for deployment purposes. 

#### frontend

```
$> cd frontend
$> docker build -t dummyproxy:v1 .
```
#### backend

```
$> cd backend
$> docker build -t go-app:v3 .
```

The names of the docker images are important. If you prefer to name them
differently, do not forget to rename the `template.spec.containers.image` value
in Deployment descriptions.

### Deploying on Kubernetes

```
$> kubectl apply -f redis.yaml
$> kubectl apply -f backend.yaml
$> kubectl apply -f frontend.yaml
```

#### References:
* [Connect a Front End to a Back End Using a Service](https://kubernetes.io/docs/tasks/access-application-cluster/connecting-frontend-backend/)
