provider "srl" {
  username = var.username
  password = var.password
  target = var.target
  encoding = "JSON_IETF"
  skip_verify = true
  debug = true
}

data "srl_system_clock" "all" {}

output "srl_system_clock" {
  value = data.srl_system_clock.all
}