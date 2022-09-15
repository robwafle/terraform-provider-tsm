# command to rebuild, install and test
```
pushd ../../terraform-provider-tanzu ; make install && popd; rm .terraform.lock.hcl ; terraform init -upgrade ; terraform plan
```


```
rm terraform.tfstate.backup; rm terraform.tfstate; rm tfplan ; export TF_LOG=TRACE; rm .terraform.lock.hcl ; terraform init -upgrade ; terraform plan -out=tfplan; terraform show tfplan | tee tfplan.txt;  terraform apply -input=false tfplan
```

```
terraform apply -input=false tfplan
```




```
kubectl --context docker-desktop apply -f https://prod-4.nsxservicemesh.vmware.com/cluster-registration/k8s/operator-deployment.yaml
```