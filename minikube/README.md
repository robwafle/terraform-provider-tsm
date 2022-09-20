```
minikube start --mount-string="type=bind,source=/run/utmp,dst=/host/var/run/utmp" --profile minikube-one --memory=16384 --cpus=8 --nodes 3 --kubernetes-version=v1.22.6
```

```
minikube delete --profile minikube-one
```


