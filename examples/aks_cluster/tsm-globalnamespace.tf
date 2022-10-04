resource "tsm_globalnamespace" "default" {
  depends_on = [tsm_cluster.aks]

  name = "global-default"
  display_name = "global-default"
  domain_name = "global-default.gns"
  use_shared_gateway = true
  mtls_enforced = true
  api_discovery_enabled = true
  ca_type = "PreExistingCA"
  ca = "default"
  description = "created via terraform test update"
  color = "#00FF00"
  version = "2.0"
    
  match_condition {
    cluster_type = "EXACT"
    cluster_match = tsm_cluster.aks.display_name
    namespace_type = "EXACT"
    namespace_match = "default"
  }
}


