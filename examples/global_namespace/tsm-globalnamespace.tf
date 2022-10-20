resource "tsm_globalnamespace" "default" {

  name                  = "global-default"
  display_name          = "global-default"
  domain_name           = "global-default.gns"
  use_shared_gateway    = true
  mtls_enforced         = true
  api_discovery_enabled = true
  ca_type               = "PreExistingCA"
  ca                    = "default"
  description           = "created via terraform -- silly rob update the terraform don't use the UI"
  color                 = "#00FF00"
  version               = "2.0"

  match_condition {
    cluster_type    = "EXACT"
    cluster_match   = "test-cluster"
    namespace_type  = "EXACT"
    namespace_match = "default"
  }

  match_condition {
    cluster_type    = "EXACT"
    cluster_match   = "test2-cluster"
    namespace_type  = "EXACT"
    namespace_match = "other"
  }
}

data "tsm_globalnamespaces" "default" {
}


output "tsm_globalnamespaces" {
  value = data.tsm_globalnamespace.default
}
