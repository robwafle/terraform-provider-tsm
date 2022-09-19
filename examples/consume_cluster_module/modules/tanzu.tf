module "tanzu" {
  for_each = var.tanzu
  clusters = lookup(each.value, "clusters", null)
}