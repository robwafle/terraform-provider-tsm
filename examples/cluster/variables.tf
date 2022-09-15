variable "namespace_exclusions" {
  description = "Create IAM users with these names"
  type = list(object({
    match     = string
    type = string
  }))
    default = [
    {
      match = "terraform"
      type = "~/.ssh/google_compute_engine.pub"
    },
    {
      match = "user"
      type = "~/.ssh/id_rsa.pub"
    }
  ]
}