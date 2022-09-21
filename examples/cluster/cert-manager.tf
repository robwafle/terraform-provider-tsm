provider "helm" {
  kubernetes {
    host = azurerm_kubernetes_cluster.k8s.kube_config.0.host

    client_certificate     = base64decode(azurerm_kubernetes_cluster.k8s.kube_config.0.client_certificate)
    client_key             = base64decode(azurerm_kubernetes_cluster.k8s.kube_config.0.client_key)
    cluster_ca_certificate = base64decode(azurerm_kubernetes_cluster.k8s.kube_config.0.cluster_ca_certificate)
  }
}

provider "kubernetes" {
  host = azurerm_kubernetes_cluster.k8s.kube_config.0.host

  client_certificate     = base64decode(azurerm_kubernetes_cluster.k8s.kube_config.0.client_certificate)
  client_key             = base64decode(azurerm_kubernetes_cluster.k8s.kube_config.0.client_key)
  cluster_ca_certificate = base64decode(azurerm_kubernetes_cluster.k8s.kube_config.0.cluster_ca_certificate)
}

resource "kubernetes_namespace" "cert_manager" {
  depends_on = [azurerm_kubernetes_cluster.k8s]
  metadata {
    name = "cert-manager"
  }
}

resource "helm_release" "cert_manager" {
  depends_on = [kubernetes_namespace.cert_manager]
  name       = "cert-manager"
  repository = "https://charts.jetstack.io"
  chart      = "cert-manager"
  version    = "v1.7.1"
  namespace  = kubernetes_namespace.cert_manager.metadata.0.name

  set {
    name  = "installCRDs"
    value = "true"
  }
}