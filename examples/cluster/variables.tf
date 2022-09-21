variable "appId" {
  description = "Azure Kubernetes Service Cluster service principal"
}

variable "password" {
  description = "Azure Kubernetes Service Cluster password"
}

variable "tenantId" {}

variable "clusterPrefix" {
  default = "tanzu-two"
}