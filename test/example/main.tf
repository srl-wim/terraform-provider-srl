provider "srl" {
  username = var.username
  password = var.password
  target = var.target
  encoding = "JSON_IETF"
}

#resource "srl_system_ntp" "ntp" {
#    admin_state = "enable"
#}

resource "srl_system_clock" "clock" {
    timezone = "Europe/Amsterdam"
}