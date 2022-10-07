variable "clusters" {
  type = map(object({
    display_name                = string
    description                 = string
    kubernetes_context          = string
    auto_install_servicemesh    = bool
    enable_namespace_exclusions = bool
    tags                        = list(string)
    labels                      = map(string)
    namespace_exclusions = map(object({
      match = string
      type  = string
    }))
  }))
}