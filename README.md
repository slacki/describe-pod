# describe-pod

This Docker image contains a Golang webserver that prints all environment variables as a response to HTTP request. 

## Motivation

When testing load balancers for Kubernetes I wanted to see how my traffic was routed. I thought this would be the simplest solution.

## Pod configuration

To make use of this image, you need to set environment variables when deploying it. Please refer to this documentation page https://kubernetes.io/docs/tasks/inject-data-application/environment-variable-expose-pod-information/#use-pod-fields-as-values-for-environment-variables

## Example deployment

To create the deployment...

```
apiVersion: apps/v1
kind: Deployment
metadata:
  name: describe-pod-deployment
spec:
  selector:
    matchLabels:
      app: describe-pods
  replicas: 6 
  template:
    metadata:
      labels:
        app: describe-pods
    spec:
      containers:
      - name: describe-pods
        image: shutah/describe-pod
        ports:
        - containerPort: 80
```

...and to expose it...

```
kubectl expose deployment describe-pod-deployment --type=LoadBalancer --name=describe-pod-service
```

## Official image

https://hub.docker.com/r/shutah/describe-pod

## Licence

See LICENCE.md file.
