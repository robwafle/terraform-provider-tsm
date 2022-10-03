# env vars to set on your local system BEFORE opening visual studio.:
```
TSM_HOST: https://prod-<number>.nsxservicemesh.vmware.com
TSM_APIKEY: <its-a-secret>
```

# command to rebuild, install and test
```
pushd ../../terraform-provider-tsm ; make install && popd; rm .terraform.lock.hcl ; terraform init -upgrade ; terraform plan
```


```
rm terraform.tfstate.backup; rm terraform.tfstate; rm tfplan ; export TF_LOG=TRACE; rm .terraform.lock.hcl ; terraform init -upgrade ; terraform plan -var-file=example.tfvars -out=tfplan; terraform show tfplan | tee tfplan.txt;  terraform apply -input=false tfplan
```

```
terraform apply -input=false tfplan
```


```
kubectl --context docker-desktop apply -f https://prod-4.nsxservicemesh.vmware.com/cluster-registration/k8s/operator-deployment.yaml
```

