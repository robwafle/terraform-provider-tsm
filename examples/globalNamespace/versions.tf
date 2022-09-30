terraform {
  required_providers {

    tanzu = {
      source  = "terraform.vmware.com/csc/tanzu"
      version = "~> 0.0.1"
    }


  }

  required_version = ">= 0.14"
}

