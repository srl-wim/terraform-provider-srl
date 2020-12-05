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

output "resource_interface_wan01e1_1" {
  value = srl_interfaces.wan01e1_1
}

resource "srl_interfaces_subinterface" "wan01e1_1_1" {
  interface_id = srl_interfaces.wan01e1_1.id
  subinterface {
    index = 1
    admin_state = "enable"
    description = "tf-ethernet1/1.1"
    ip_mtu = 1500
    type = "routed"
  }
}

resource "srl_interfaces_subinterface_ipv4" "wan01e1_1_1" {
  interface_id = srl_interfaces.wan01e1_1.id
  subinterface_id = srl_interfaces_subinterface.wan01e1_1_1.id
  ipv4 {
  }
}

resource "srl_interfaces_subinterface_ipv4_address" "wan01e1_1_1" {
  interface_id = srl_interfaces.wan01e1_1.id
  subinterface_id = srl_interfaces_subinterface.wan01e1_1_1.id
  ipv4_id = srl_interfaces_subinterface_ipv4.wan01e1_1_1.id
  address {
    ip_prefix = "10.0.0.1/24"
  }
}

resource "srl_network_instance_instance" "wan01default" {
  network_instance {
    name = "default"
    interface {
      name = format("%s.%s", srl_interfaces.wan01e1_1.id ,srl_interfaces_subinterface.wan01e1_1_1.id)
    }
    type = "default"
    admin_state = "enable"
    description = "terraform created default network instance"
    router_id = "10.1.1.1"
  }
}