provider "srl" {
  alias = "srl4"
  username = var.username
  password = var.password
  target = var.targetwan04
  encoding = "JSON_IETF"
  skip_verify = true
  debug = true
}

resource "srl_interfaces" "wan04e1_5" {
  provider = srl.srl4
  interface {
    admin_state = "enable"
    name = "ethernet-1/5"
    description = "tf-ethernet1/5"
    mtu = 1500
    vlan_tagging = false
  }
}

resource "srl_interfaces_subinterface" "wan04e1_5_1" {
  provider = srl.srl4
  interface_id = srl_interfaces.wan04e1_5.id
  subinterface {
    index = 1
    admin_state = "enable"
    description = "tf-ethernet1/5.1"
    ip_mtu = 1500
    type = "routed"
  }
}