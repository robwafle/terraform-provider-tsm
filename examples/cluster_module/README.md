## Requirements

| Name | Version |
|------|---------|
| <a name="requirement_tsm"></a> [tsm](#requirement\_tsm) | 0.0.1 |

## Providers

| Name | Version |
|------|---------|
| <a name="provider_tsm"></a> [tsm](#provider\_tsm) | 0.0.1 |

## Modules

No modules.

## Resources

| Name | Type |
|------|------|
| [tsm_cluster.cluster](https://registry.terraform.io/providers/robwafle/tsm/0.0.1/docs/resources/cluster) | resource |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_clusters"></a> [clusters](#input\_clusters) | n/a | <pre>map(object({<br>    display_name                = string<br>    description                 = string<br>    kubernetes_context          = string<br>    auto_install_servicemesh    = bool<br>    enable_namespace_exclusions = bool<br>    tags                        = list(string)<br>    labels                      = map(string)<br>    namespace_exclusions = map(object({<br>      match = string<br>      type  = string<br>    }))<br>  }))</pre> | n/a | yes |

## Outputs

No outputs.
