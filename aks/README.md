
# connected arc as described here

https://learn.microsoft.com/en-us/azure/azure-arc/kubernetes/quickstart-connect-cluster?tabs=azure-cli


az provider register --namespace Microsoft.OperationsManagement
az provider register --namespace Microsoft.OperationalInsights

# create a resource group
az group create --name aks-one-rg --location eastus

# create the cluster
az aks create -g aks-one-rg -n aks-one --enable-managed-identity --node-count 3 --enable-azure-rbac --node-vm-size --generate-ssh-keys


# create the arc cluster
az group create --name aks-one-arc --location EastUS --output table

# get the creds so we can connect the cluster
az aks get-credentials --name tsm-two --resource-group tsm-two-rg

# connect the cluster
az connectedk8s connect --name aks-one --resource-group aks-one-arc

# create a proxy connection to the cluster
az connectedk8s proxy --name aks-one --resource-group aks-one-arc

# grant role binding (first command returns email address)
AAD_ENTITY_OBJECT_ID=$(az ad signed-in-user show --query userPrincipalName -o tsv) 
kubectl create clusterrolebinding rob-admin --clusterrole cluster-admin --user=$AAD_ENTITY_OBJECT_ID
