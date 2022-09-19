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