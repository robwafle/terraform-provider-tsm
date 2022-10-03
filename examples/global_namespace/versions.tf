terraform {
  required_providers {

    tsm = {
      source  = "terraform.vmware.com/csc/tsm"
      version = "~> 0.0.1"
    }


  }

  required_version = ">= 0.14"
}

