# Running multi pods on Kubernetes

This is a sample project deploying a simple go application on kubernetes with n
number of replicas.

<i>TODO: Check [Forge](https://forge.sh/) for deployments for Kubernetes</i>

## Go App

A small go app that prints the followings: 
* hostname of the machine and the version of the app
* Total number of requests
* App Uptime
* Log Time

In order to test on your local docker, you can run the app with the following
command:

```bash
$> docker run -d -p 8080:8080 serdard/kubernetes-multipod:v1
$> curl localhost:8080
Hello from Kubernetes! Running on bd76b008ef7d | version: 0.1
Total # of requests to this pod:1
App Uptime: 7.56014836s
Log Time: 2018-02-14 18:26:17.137948958 +0000 UTC

```
The output of the curl command should be similar to this.

## Docker building, tagging and pushing to the docker registry

Summary of steps done to update the docker images.

Docker hub is used for the docker registry: https://hub.docker.com/r/serdard/kubernetes-multipod/

You need to have `DOCKER_ID_USER` environment variable set in order to run the
following commands.

Change the version `v1` to your up-to-date version for version updates.

### Building the image
```bash
$> docker build -t k8s-multipod:v1 .
```

### Tagging and pushing the image
```bash
$> docker tag k8s-multipod:v1 $DOCKER_ID_USER/kubernetes-multipod:v1
$> docker push $DOCKER_ID_USER/kubernetes-multipod
```
## Deploying on minikube

Make sure minikube is up and running
```bash
$>minikube status
minikube: Running
cluster: Running
kubectl: Correctly Configured: pointing to minikube-vm at 192.168.99.100

```
Create the deployment with the following command
```bash
$> kubectl create -f kubernetes/deployment.yaml
```
Check the deployment is created and applied
```bash
$> kubectl get deployment
NAME                  DESIRED   CURRENT   UP-TO-DATE   AVAILABLE   AGE
kubernetes-multipod   3         3         3            3           2h
```
Also deploy the service for exposing a NodePort for the application
```bash
$> kubectl create -f kubernetes/service.yaml
```
Check the service is deployed correctly
```bash
$> kubectl get svc
NAME                  TYPE        CLUSTER-IP       EXTERNAL-IP   PORT(S)      AGE
kubernetes            ClusterIP   10.96.0.1        <none>        443/TCP       2d
kubernetes-multipod   NodePort    10.107.205.127   <none>        80:31026/TCP  2h

```

Now, the app is deployed on 3 PODs and the Port `31026` is exposed to outside.
Run the following curl command to see the output of the app
```bash
$> curl 192.168.99.100:31026
```
Note that the IP comes from minikube (see the command output of `minikube status`) and the port is assigned automatically by Kubernetes (see the `kubectl get svc` command output)

You can also run the curl command in the following way
```bash
$> curl $(minikube service kubernetes-multipod --url)
```
It will give the same output.

## Scaling up/down the app

Change the parameter `replicas` to any reasonable number you want.

The following command will set the number of running pods to 4.
```bash
$> kubectl scale deployments/kubernetes-multipod --replicas=4
```

## Performing a rolling update

Simple run the following command to do the rolling update
```bash
$> kubectl set image deployments/kubernetes-multipod kubernetes-multipod=serdard/kubernetes-multipod:v2
```
The pods will be created and terminated one by one, because
`.spec.strategy.rollingUpdate.maxSurge` is set to 1 on [deployment.yaml](kubernetes/deployment.yaml) file

You can checkout rollout status with the following command:
```bash
$> kubectl rollout status deployment kubernetes-multipod
deployment "kubernetes-multipod" successfully rolled out
```

If, for any reason, you want to rollback the update, you can run the following
command
```bash
$> kubectl rollout undo deployment kubernetes-multipod
deployment "kubernetes-multipod" rolled back
```

## Cleanup

Delete the service and the deployment from your minikube. 
```bash
$> kubectl delete svc kubernetes-multipod
$> kubectl delete deploy kubernetes-multipod
```
