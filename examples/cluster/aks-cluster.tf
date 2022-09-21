provider "azurerm" {
  features {}
}

resource "azurerm_resource_group" "default" {
  name     = "${var.clusterPrefix}-rg"
  location = "East US 2"

  tags = {
    environment = "Demo"
  }
}

resource "azurerm_kubernetes_cluster" "k8s" {
  name                = "${var.clusterPrefix}-aks"
  location            = azurerm_resource_group.default.location
  resource_group_name = azurerm_resource_group.default.name
  dns_prefix          = "${var.clusterPrefix}-k8s"
  #kubernetes_version  = "1.22.6"

  default_node_pool {
    name            = "default"
    node_count      = 3
    vm_size         = "Standard_B4ms"
    os_disk_size_gb = 30
  }

  identity {
    type = "SystemAssigned"
  }

  role_based_access_control {
    enabled = true
  }

  tags = {
    environment = "Demo"
  }
}

data "azurerm_kubernetes_cluster" "k8s" {
  depends_on = [azurerm_kubernetes_cluster.k8s]
  name                = azurerm_kubernetes_cluster.k8s.name
  resource_group_name = azurerm_resource_group.default.name
}
