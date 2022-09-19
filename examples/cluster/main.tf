resource "tanzu_cluster" "cluster" {
  display_name = "rob-local"
  kubernetes_context = "docker-desktop"
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

  namespace_exclusion {
    match = "default"
    type = "EXACT"
  }

  #  namespace_exclusion {
  #    match = "bob"
  #    type = "EXACT"
  #  }

   namespace_exclusion {
     match = "three"
     type = "EXACT"
   }
}




# data "tanzu_cluster" "stage_cluster" {
#   id = "stage"
# }

# output "stage_cluster" {
#   value = data.tanzu_cluster.stage_cluster
# }

# data "tanzu_cluster" "roblocal_cluster" {
#   id = "rob-local"
# }

# output "roblocal_cluster" {
#   value = data.tanzu_cluster.roblocal_cluster
# }