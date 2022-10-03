resource "tsm_cluster" "cluster" {
  for_each                          = var.clusters
  display_name                      = lookup(each.value, "display_name", null)
  description                       = lookup(each.value, "description", null)
  kubernetes_context                = lookup(each.value, "kubernetes_context", null)
  auto_install_servicemesh          = lookup(each.value, "auto_install_servicemesh", null)
  enable_namespace_exclusions       = lookup(each.value, "enable_namespace_exclusions", null)
  tags                              = lookup(each.value, "tags", null)
  labels                            = lookup(each.value, "labels", null)
  dynamic "namespace_exclusion" {
    for_each = lookup(each.value, "namespace_exclusions", null)
    content {
      match = namespace_exclusion.value["match"]
      type = namespace_exclusion.value["type"]
    }
  }
}
