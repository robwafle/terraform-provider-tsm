locals {
  clusters = {
    rob-local = {
      display_name                = "rob-local"
      kubernetes_context          = "docker-desktop"
      description                 = "created via terraform"
      auto_install_servicemesh    = true
      enable_namespace_exclusions = true
      tags                        = ["hello", "world"]

      labels = {
        L1 = "value one"
        L2 = "value two"
        //L3 = "value three"
        L4 = "value four"
      }

      namespace_exclusions = {
        one = {
          match = "default"
          type  = "EXACT"
        }
        two = {
          match = "bob"
          type  = "EXACT"
        }
      }
    }
  }
}

module "tsm" {
  providers = {

  }
  for_each = local.tsm
}