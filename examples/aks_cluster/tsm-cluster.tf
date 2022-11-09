resource "tsm_cluster" "aks" {
  depends_on   = [null_resource.kubectl]
  display_name = azurerm_kubernetes_cluster.k8s.name
  //cluster_name =  azurerm_kubernetes_cluster.k8s.name
  //resource_group = azurerm_resource_group.default.name
  kubernetes_context          = azurerm_kubernetes_cluster.k8s.name
  description                 = "created via terraform"
  auto_install_servicemesh    = true
  enable_namespace_exclusions = true
  tags                        = ["hello", "world"]

  labels = {
    L1 = "value one"
    L2 = "value two"
    L3 = "value three"
    L4 = "value four"
  }

  namespace_exclusion {
    match = "one"
    type  = "EXACT"
  }

  namespace_exclusion {
    match = "two"
    type  = "EXACT"
  }

  namespace_exclusion {
    match = "three"
    type  = "EXACT"
  }


}


data "tsm_cluster" "aks" {
  depends_on = [tsm_cluster.aks]
  id         = azurerm_kubernetes_cluster.k8s.name
}

output "tsm_cluster" {
  value = data.tsm_cluster.aks
}
