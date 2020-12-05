provider "srl" {
  alias = "srl2"
  username = var.username
  password = var.password
  target = var.targetwan02
  encoding = "JSON_IETF"
  skip_verify = true
  debug = true
}

resource "srl_interfaces" "wan02e1_1" {
  provider = srl.srl2
  interface {
    admin_state = "enable"
    name = "ethernet-1/1"
    description = "tf-ethernet1/1"
    mtu = 2000
    vlan_tagging = false
  }
}

output "resource_interface_wan02e1_1" {
  value = srl_interfaces.wan02e1_1
}

resource "srl_interfaces_subinterface" "wan02e1_1_1" {
  provider = srl.srl2
  interface_id = srl_interfaces.wan02e1_1.id
  subinterface {
    index = 1
    admin_state = "enable"
    description = "tf-ethernet1/1.1"
    ip_mtu = 1500
    type = "routed"
  }
}

output "resource_subinterface_wan02e1_1" {
  value = srl_interfaces_subinterface.wan02e1_1_1
}

resource "srl_interfaces_subinterface_ipv4" "wan02e1_1_1" {
  provider = srl.srl2
  interface_id = srl_interfaces.wan02e1_1.id
  subinterface_id = srl_interfaces_subinterface.wan02e1_1_1.id
  ipv4 {
  }
}

resource "srl_interfaces_subinterface_ipv4_address" "wan02e1_1_1" {
  provider = srl.srl2
  interface_id = srl_interfaces.wan02e1_1.id
  subinterface_id = srl_interfaces_subinterface.wan02e1_1_1.id
  ipv4_id = srl_interfaces_subinterface_ipv4.wan02e1_1_1.id
  address {
    ip_prefix = "10.0.0.2/24"
  }
}

resource "srl_network_instance_instance" "wan02default" {
  provider = srl.srl2
  network_instance {
    name = "default"
    type = "default"
    interface {
      name = format("%s.%s", srl_interfaces.wan02e1_1.id ,srl_interfaces_subinterface.wan02e1_1_1.id)
    }
    admin_state = "enable"
    description = "terraform created default network instance"
    router_id = "10.1.1.2"
  }
}