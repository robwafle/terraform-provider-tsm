# Examples
Various examples of using the provider are here.  At the moemnt we have:

## global_namespace
Simplest example to use to validate provider setup is correct.  Runs in ~10 seconds to provision a global namespace.

## aks_cluster
This example requires use of the azurerm provider and an Azure Subscription, but it will provision an aks cluster and onboard that aks cluster to your Tanzu Service Mesh instance in about 10 minutes.  It takes 5 minutes to provision the cluster and 5 minutes to onboard it.

## cluster_module
Please ignore this example.

## consume_cluster_module
Please ignore this example.

# Developing with the examples

### How to use the .terraform.rc development overrides (from example folder):
cp ../../.terraform.rc $HOME/.terraform.rc
export TF_CLI_CONFIG_FILE="$HOME/.terraform.rc"