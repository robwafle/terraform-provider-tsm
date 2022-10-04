terraform {
  required_providers {
    tsm = {
      source  = "robwafle/tsm"
      version = "0.0.1"
    }
  }
}

# NOTE: Values are read from the environment variables: TSM_HOST, TSM_APIKEY
provider "tsm" {
  //host = "https://prod-4.nsxservicemesh.vmware.com"
  //apikey = ""
}