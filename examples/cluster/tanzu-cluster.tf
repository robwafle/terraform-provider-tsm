provider "tanzu" {
  //host = "https://prod-4.nsxservicemesh.vmware.com"
  //apikey = ""
}

resource "tanzu_cluster" "cluster" {
  depends_on = [null_resource.kubectl]
  display_name = azurerm_kubernetes_cluster.k8s.name
  //cluster_name =  azurerm_kubernetes_cluster.k8s.name
  //resource_group = azurerm_resource_group.default.name
  kubernetes_context = azurerm_kubernetes_cluster.k8s.name
  description = "created via terraform"
  auto_install_servicemesh = true
  enable_namespace_exclusions = true
  tags = ["hello", "world"]

  labels = {
    L1 = "value one"
    L2 = "value two"
    //L3 = "value three"
    L4 = "value four"
  }

  # namespace_exclusion {
  #   match = "default"
  #   type = "EXACT"
  # }

  #  namespace_exclusion {
  #    match = "bob"
  #    type = "EXACT"
  #  }

  #  namespace_exclusion {
  #    match = "three"
  #    type = "EXACT"
  #  }
}




# # data "tanzu_cluster" "stage_cluster" {
# #   id = "stage"
# # }

# # output "stage_cluster" {
# #   value = data.tanzu_cluster.stage_cluster
# # }

# # data "tanzu_cluster" "roblocal_cluster" {
# #   id = "rob-local"
# # }

# # output "roblocal_cluster" {
# #   value = data.tanzu_cluster.roblocal_cluster
# # }