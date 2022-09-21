# docker-in-docker
```
minikube start --mount-string="type=bind,source=/run/utmp,dst=/host/var/run/utmp" --profile minikube-one --memory=16384 --cpus=8 --nodes 3 --kubernetes-version=v1.22.6
```

# docker-on-docker
```
minikube start --profile minikube-one --memory=4g --cpus=4 --nodes 3 --kubernetes-version=v1.22.6
```


```
minikube delete --profile minikube-one
```


```
for pod in $(kubectl -n istio-system get pod -listio=sidecar-injector -o jsonpath='{.items[*].metadata.name}'); do \
    kubectl -n istio-system logs ${pod} \
done
```

```
kubectl label namespace default istio-injection=enabled --overwrite
```

```
kubectl apply -f dnsutils.yml
```

