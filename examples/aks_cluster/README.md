# env vars to set on your local system BEFORE opening visual studio.:
```
TSM_HOST: https://prod-<number>.nsxservicemesh.vmware.com
TSM_APIKEY: <its-a-secret>
```

# command to rebuild, install and test
```
pushd ../.. ; if make build; then echo "built!"; else popd; fi && popd && rm tflog.json ; terraform plan -out=tfplan
```

# apply
```
terraform apply -input=false tfplan
```

# destroy
```
terraform apply -input=false -destroy
```

# manual deploy
```
kubectl apply -f https://prod-4.nsxservicemesh.vmware.com/cluster-registration/k8s/operator-deployment.yaml
```

# manual delete
```
kubectl delete --ignore-not-found=true -f https://prod-4.nsxservicemesh.vmware.com/cluster-registration/k8s/client-cluster-uninstall.yaml
```

# kubectl run tmp-shell --rm -i --tty --image nicolaka/netshoot
```
kubectl run netshoot --rm -i --tty --image nicolaka/netshoot
```


# if you want to spin up a container on the host's network namespace.
```
 kubectl run netshoot --rm -i --tty --overrides='{"spec": {"hostNetwork": true}}' --image nicolaka/netshoot
```

