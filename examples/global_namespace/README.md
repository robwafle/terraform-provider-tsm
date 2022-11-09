## Requirements

| Name | Version |
|------|---------|
| <a name="requirement_terraform"></a> [terraform](#requirement\_terraform) | >= 0.14 |
| <a name="requirement_tsm"></a> [tsm](#requirement\_tsm) | ~> 0.0.1 |

## Providers

| Name | Version |
|------|---------|
| <a name="provider_tsm"></a> [tsm](#provider\_tsm) | 0.0.81 |

## Modules

No modules.

## Resources

| Name | Type |
|------|------|
| [tsm_globalnamespace.default](https://registry.terraform.io/providers/robwafle/tsm/latest/docs/resources/globalnamespace) | resource |
| [tsm_globalnamespaces.default](https://registry.terraform.io/providers/robwafle/tsm/latest/docs/data-sources/globalnamespaces) | data source |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_appId"></a> [appId](#input\_appId) | Azure Kubernetes Service Cluster service principal | `any` | n/a | yes |
| <a name="input_clusterPrefix"></a> [clusterPrefix](#input\_clusterPrefix) | n/a | `string` | `"tsm-one"` | no |
| <a name="input_password"></a> [password](#input\_password) | Azure Kubernetes Service Cluster password | `any` | n/a | yes |
| <a name="input_tenantId"></a> [tenantId](#input\_tenantId) | n/a | `any` | n/a | yes |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_tsm_globalnamespaces"></a> [tsm\_globalnamespaces](#output\_tsm\_globalnamespaces) | n/a |
