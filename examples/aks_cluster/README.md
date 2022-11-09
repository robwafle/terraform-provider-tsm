## Requirements

| Name | Version |
|------|---------|
| <a name="requirement_azurerm"></a> [azurerm](#requirement\_azurerm) | >= 3.27.0 |
| <a name="requirement_tsm"></a> [tsm](#requirement\_tsm) | >= 0.0.81 |

## Providers

| Name | Version |
|------|---------|
| <a name="provider_azurerm"></a> [azurerm](#provider\_azurerm) | 3.27.0 |
| <a name="provider_null"></a> [null](#provider\_null) | 3.1.1 |
| <a name="provider_tsm"></a> [tsm](#provider\_tsm) | 0.0.81 |

## Modules

No modules.

## Resources

| Name | Type |
|------|------|
| [azurerm_kubernetes_cluster.k8s](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/kubernetes_cluster) | resource |
| [azurerm_resource_group.default](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/resource_group) | resource |
| [null_resource.kubectl](https://registry.terraform.io/providers/hashicorp/null/latest/docs/resources/resource) | resource |
| [tsm_cluster.aks](https://registry.terraform.io/providers/robwafle/tsm/latest/docs/resources/cluster) | resource |
| [tsm_globalnamespace.default](https://registry.terraform.io/providers/robwafle/tsm/latest/docs/resources/globalnamespace) | resource |
| [azurerm_kubernetes_cluster.k8s](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/data-sources/kubernetes_cluster) | data source |
| [tsm_cluster.aks](https://registry.terraform.io/providers/robwafle/tsm/latest/docs/data-sources/cluster) | data source |
| [tsm_globalnamespace.default](https://registry.terraform.io/providers/robwafle/tsm/latest/docs/data-sources/globalnamespace) | data source |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_appId"></a> [appId](#input\_appId) | Azure Kubernetes Service Cluster service principal | `any` | n/a | yes |
| <a name="input_clusterPrefix"></a> [clusterPrefix](#input\_clusterPrefix) | n/a | `string` | `"tsm-one"` | no |
| <a name="input_password"></a> [password](#input\_password) | Azure Kubernetes Service Cluster password | `any` | n/a | yes |
| <a name="input_subscriptionId"></a> [subscriptionId](#input\_subscriptionId) | n/a | `any` | n/a | yes |
| <a name="input_tenantId"></a> [tenantId](#input\_tenantId) | n/a | `any` | n/a | yes |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_kube_admin_config"></a> [kube\_admin\_config](#output\_kube\_admin\_config) | kubeconfig for kubectl admin access. |
| <a name="output_kube_config"></a> [kube\_config](#output\_kube\_config) | kubeconfig for kubectl access. |
| <a name="output_kubernetes_cluster_name"></a> [kubernetes\_cluster\_name](#output\_kubernetes\_cluster\_name) | n/a |
| <a name="output_resource_group_name"></a> [resource\_group\_name](#output\_resource\_group\_name) | n/a |
| <a name="output_tsm_cluster"></a> [tsm\_cluster](#output\_tsm\_cluster) | n/a |
| <a name="output_tsm_globalnamespace"></a> [tsm\_globalnamespace](#output\_tsm\_globalnamespace) | n/a |
