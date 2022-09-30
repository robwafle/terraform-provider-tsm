# env vars to set on your local system BEFORE opening visual studio.:
```
TANZU_HOST: https://prod-<number>.nsxservicemesh.vmware.com
TANZU_APIKEY: <its-a-secret>
```

# command to rebuild, install and test
```
pushd ../../terraform-provider-tanzu ; make install && popd; rm .terraform.lock.hcl ; terraform init -upgrade ; terraform plan
```

# clean up if it all goes wrong
```
rm terraform.tfstate.backup; rm terraform.tfstate; rm tfplan; rm tfplan.txt ; export TF_LOG=TRACE; rm .terraform.lock.hcl ; 
```

# plan, if you get a checksum error you didn't update the version in Makefile.
```
rm terraform.tfstate.backup; rm terraform.tfstate; rm tfplan ; export TF_LOG=TRACE; rm .terraform.lock.hcl ; terraform init -upgrade ; terraform plan -out=tfplan; terraform show tfplan | tee tfplan.txt; 
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

# Unit testing
In vscode, you can right click and choose "Go: Test Function At Cursor" or run the tests manually.
```
go test terraform-provider-tanzu/plugin/provider -v
```