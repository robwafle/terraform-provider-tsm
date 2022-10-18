terraform {
  required_providers {
    azurerm = {
      source  = "hashicorp/azurerm"
      version = ">= 3.27.0"
    }
    tsm = {
      source  = "robwafle/tsm"
      version = ">= 0.0.81"
    }
    # cloudflare = {
    #   source  = "cloudflare/cloudflare"
    #   version = "=3.9.1"
    # }
    # helm = {
    #   source  = "helm"
    #   version = "=2.6.0"
    # }
    # kubernetes = {
    #   source  = "kubernetes"
    #   version = "=2.8.0"
    # }
  }
}

