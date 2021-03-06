provider "srl" {
  alias = "srl4"
  username = var.username
  password = var.password
  target = var.targetwan04
  encoding = "JSON_IETF"
  skip_verify = true
  debug = true
}

resource "srl_interfaces" "wan04e1_1" {
  provider = srl.srl4
  interface {
    admin_state = "enable"
    name = "ethernet-1/1"
    description = "tf-ethernet1/1"
    mtu = 2000
    loopback_mode = false
    vlan_tagging = false
  }
}

resource "srl_interfaces" "wan04e1_2" {
  provider = srl.srl4
  interface {
    admin_state = "enable"
    name = "ethernet-1/2"
    description = "tf-ethernet1/2"
    mtu = 2000
    loopback_mode = false
    vlan_tagging = false
  }
}

resource "srl_interfaces" "wan04e1_3" {
  provider = srl.srl4
  interface {
    admin_state = "enable"
    name = "ethernet-1/3"
    description = "tf-ethernet1/3"
    mtu = 2000
    loopback_mode = false
    vlan_tagging = false
  }
}

resource "srl_interfaces" "wan04e1_4" {
  provider = srl.srl4
  interface {
    admin_state = "enable"
    name = "ethernet-1/4"
    description = "tf-ethernet1/4"
    mtu = 2000
    loopback_mode = false
    vlan_tagging = false
  }
}

resource "srl_interfaces" "wan04e1_5" {
  provider = srl.srl4
  interface {
    admin_state = "enable"
    name = "ethernet-1/5"
    description = "tf-ethernet1/5"
    mtu = 2000
    loopback_mode = false
    vlan_tagging = false
  }
}

resource "srl_interfaces" "wan04e1_6" {
  provider = srl.srl4
  interface {
    admin_state = "enable"
    name = "ethernet-1/6"
    description = "tf-ethernet1/6"
    mtu = 2000
    loopback_mode = false
    vlan_tagging = false
  }
}

resource "srl_interfaces_subinterface" "wan04e1_1_1" {
  provider = srl.srl4
  interface_id = srl_interfaces.wan04e1_1.id
  subinterface {
    index = 1
    admin_state = "enable"
    description = "tf-ethernet1/1.1"
    ip_mtu = 1500
    type = "routed"
  }
}

resource "srl_interfaces_subinterface" "wan04e1_2_1" {
  provider = srl.srl4
  interface_id = srl_interfaces.wan04e1_2.id
  subinterface {
    index = 1
    admin_state = "enable"
    description = "tf-ethernet1/2.1"
    ip_mtu = 1500
    type = "routed"
  }
}

resource "srl_interfaces_subinterface" "wan04e1_3_1" {
  provider = srl.srl4
  interface_id = srl_interfaces.wan04e1_3.id
  subinterface {
    index = 1
    admin_state = "enable"
    description = "tf-ethernet1/3.1"
    ip_mtu = 1500
    type = "routed"
  }
}

resource "srl_interfaces_subinterface" "wan04e1_4_1" {
  provider = srl.srl4
  interface_id = srl_interfaces.wan04e1_4.id
  subinterface {
    index = 1
    admin_state = "enable"
    description = "tf-ethernet1/4.1"
    ip_mtu = 1500
    type = "routed"
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

resource "srl_interfaces_subinterface" "wan04e1_6_1" {
  provider = srl.srl4
  interface_id = srl_interfaces.wan04e1_6.id
  subinterface {
    index = 1
    admin_state = "enable"
    description = "tf-ethernet1/6.1"
    ip_mtu = 1500
    type = "routed"
  }
}

resource "srl_interfaces_subinterface_ipv4" "wan04e1_1_1" {
  provider = srl.srl4
  interface_id = srl_interfaces.wan04e1_1.id
  subinterface_id = srl_interfaces_subinterface.wan04e1_1_1.id
  ipv4 {
    allow_directed_broadcast = false
  }
}

resource "srl_interfaces_subinterface_ipv4" "wan04e1_2_1" {
  provider = srl.srl4
  interface_id = srl_interfaces.wan04e1_2.id
  subinterface_id = srl_interfaces_subinterface.wan04e1_2_1.id
  ipv4 {
    allow_directed_broadcast = false
  }
}

resource "srl_interfaces_subinterface_ipv4" "wan04e1_3_1" {
  provider = srl.srl4
  interface_id = srl_interfaces.wan04e1_3.id
  subinterface_id = srl_interfaces_subinterface.wan04e1_3_1.id
  ipv4 {
    allow_directed_broadcast = false
  }
}

resource "srl_interfaces_subinterface_ipv4" "wan04e1_4_1" {
  provider = srl.srl4
  interface_id = srl_interfaces.wan04e1_4.id
  subinterface_id = srl_interfaces_subinterface.wan04e1_4_1.id
  ipv4 {
    allow_directed_broadcast = false
  }
}

resource "srl_interfaces_subinterface_ipv4" "wan04e1_5_1" {
  provider = srl.srl4
  interface_id = srl_interfaces.wan04e1_5.id
  subinterface_id = srl_interfaces_subinterface.wan04e1_5_1.id
  ipv4 {
    allow_directed_broadcast = false
  }
}

resource "srl_interfaces_subinterface_ipv4" "wan04e1_6_1" {
  provider = srl.srl4
  interface_id = srl_interfaces.wan04e1_6.id
  subinterface_id = srl_interfaces_subinterface.wan04e1_6_1.id
  ipv4 {
    allow_directed_broadcast = false
  }
}

resource "srl_interfaces_subinterface_ipv4_address" "wan04e1_1_1" {
  provider = srl.srl4
  interface_id = srl_interfaces.wan04e1_1.id
  subinterface_id = srl_interfaces_subinterface.wan04e1_1_1.id
  ipv4_id = srl_interfaces_subinterface_ipv4.wan04e1_1_1.id
  address {
    ip_prefix = "10.0.7.2/24"
  }
}

resource "srl_interfaces_subinterface_ipv4_address" "wan04e1_2_1" {
  provider = srl.srl4
  interface_id = srl_interfaces.wan04e1_2.id
  subinterface_id = srl_interfaces_subinterface.wan04e1_2_1.id
  ipv4_id = srl_interfaces_subinterface_ipv4.wan04e1_2_1.id
  address {
    ip_prefix = "10.0.8.2/24"
  }
}

resource "srl_interfaces_subinterface_ipv4_address" "wan04e1_3_1" {
  provider = srl.srl4
  interface_id = srl_interfaces.wan04e1_3.id
  subinterface_id = srl_interfaces_subinterface.wan04e1_3_1.id
  ipv4_id = srl_interfaces_subinterface_ipv4.wan04e1_3_1.id
  address {
    ip_prefix = "10.0.11.2/24"
  }
}

resource "srl_interfaces_subinterface_ipv4_address" "wan04e1_4_1" {
  provider = srl.srl4
  interface_id = srl_interfaces.wan04e1_4.id
  subinterface_id = srl_interfaces_subinterface.wan04e1_4_1.id
  ipv4_id = srl_interfaces_subinterface_ipv4.wan04e1_4_1.id
  address {
    ip_prefix = "10.0.12.2/24"
  }
}

resource "srl_interfaces_subinterface_ipv4_address" "wan04e1_5_1" {
  provider = srl.srl4
  interface_id = srl_interfaces.wan04e1_5.id
  subinterface_id = srl_interfaces_subinterface.wan04e1_5_1.id
  ipv4_id = srl_interfaces_subinterface_ipv4.wan04e1_5_1.id
  address {
    ip_prefix = "10.0.13.2/24"
  }
}

resource "srl_interfaces_subinterface_ipv4_address" "wan04e1_6_1" {
  provider = srl.srl4
  interface_id = srl_interfaces.wan04e1_6.id
  subinterface_id = srl_interfaces_subinterface.wan04e1_6_1.id
  ipv4_id = srl_interfaces_subinterface_ipv4.wan04e1_6_1.id
  address {
    ip_prefix = "10.0.14.2/24"
  }
}

resource "srl_network_instance_instance" "wan04default" {
  provider = srl.srl4
  network_instance {
    name = "default"
    interface {
      name = format("%s.%s", srl_interfaces.wan04e1_1.id ,srl_interfaces_subinterface.wan04e1_1_1.id)
    }
    interface {
      name = format("%s.%s", srl_interfaces.wan04e1_2.id ,srl_interfaces_subinterface.wan04e1_2_1.id)
    }
    interface {
      name = format("%s.%s", srl_interfaces.wan04e1_3.id ,srl_interfaces_subinterface.wan04e1_3_1.id)
    }
    interface {
      name = format("%s.%s", srl_interfaces.wan04e1_4.id ,srl_interfaces_subinterface.wan04e1_4_1.id)
    }
    interface {
      name = format("%s.%s", srl_interfaces.wan04e1_5.id ,srl_interfaces_subinterface.wan04e1_5_1.id)
    }
    interface {
      name = format("%s.%s", srl_interfaces.wan04e1_6.id ,srl_interfaces_subinterface.wan04e1_6_1.id)
    }
    type = "default"
    admin_state = "enable"
    description = "terraform created default network instance"
    router_id = "10.1.1.4"
  }
}