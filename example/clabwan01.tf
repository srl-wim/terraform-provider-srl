provider "srl" {
  username = var.username
  password = var.password
  target = var.targetwan01
  encoding = "JSON_IETF"
  skip_verify = true
  debug = true
}

resource "srl_interfaces" "wan01e1_1" {
  interface {
    admin_state = "enable"
    name = "ethernet-1/1"
    description = "tf-ethernet1/1"
    mtu = 2000
    vlan_tagging = false
  }
}

