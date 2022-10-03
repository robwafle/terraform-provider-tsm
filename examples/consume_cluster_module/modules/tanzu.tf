module "tsm" {
  for_each = var.tsm
  clusters = lookup(each.value, "clusters", null)
}