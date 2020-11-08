provider "srl" {
  username = var.username
  password = var.password
  target = var.target
  encoding = "JSON_IETF"
  skip_verify = true
}

#resource "srl_system_ntp" "ntp" {
#    admin_state = "enable"
#}

resource "srl_system_information" "info" {
    information {
      contact = "Wim henderickx"
      location = "Copernicuslann 50, 2018 antwerp"
    }
}