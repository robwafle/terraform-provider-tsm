terraform {
  required_providers {
    tanzu = {
      source  = "terraform.vmware.com/csc/tanzu"
      version = "0.0.1"
    }
  }
}



# NOTE: Values are read from the environment variables: TANZU_HOST, TANZU_APIKEY

provider "tanzu" {
  //host = "https://prod-4.nsxservicemesh.vmware.com"
  //apikey = ""
}

resource "tanzu_cluster" "cluster" {
  display_name = "rob-local"
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
  
  
  # namespace_exclusion {
  #   for_each = {for ne in var.namespace_exclusions: ne => ne}

  #   match = "${ne.match}"
  #   type = "${ne.type}"
  # }



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