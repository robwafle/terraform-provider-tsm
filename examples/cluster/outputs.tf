output "resource_group_name" {
  value = azurerm_resource_group.default.name
}

output "kubernetes_cluster_name" {
  value = azurerm_kubernetes_cluster.k8s.name
}

# terraform output -raw kube_config > ~/.kube/config
output "kube_config" {
  value       = azurerm_kubernetes_cluster.k8s.kube_config_raw
  description = "kubeconfig for kubectl access."
  sensitive   = true
}

# terraform output -raw kube_admin_config > ~/.kube/config
output "kube_admin_config" {
  value       = azurerm_kubernetes_cluster.k8s.kube_admin_config_raw
  description = "kubeconfig for kubectl admin access."
  sensitive   = true
}


# output "host" {
#   value = azurerm_kubernetes_cluster.k8s.kube_config.0.host
# }

# output "client_key" {
#   value = azurerm_kubernetes_cluster.k8s.kube_config.0.client_key
# }

# output "client_certificate" {
#   value = azurerm_kubernetes_cluster.k8s.kube_config.0.client_certificate
# }

# output "kube_config" {
#   value = azurerm_kubernetes_cluster.k8s.kube_config_raw
# }

# output "cluster_username" {
#   value = azurerm_kubernetes_cluster.k8s.kube_config.0.username
# }

# output "cluster_password" {
#   value = azurerm_kubernetes_cluster.k8s.kube_config.0.password
# }
