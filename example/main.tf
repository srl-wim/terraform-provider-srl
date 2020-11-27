provider "srl" {
  username = var.username
  password = var.password
  target = var.target
  encoding = "JSON_IETF"
  skip_verify = true
  debug = true
}

/*
data "srl_system_clock" "all" {}

output "data_srl_system_clock" {
  value = data.srl_system_clock.all
}
*/

/*
data "srl_system_ntp" "all" {}

output "data_srl_system_ntp" {
  value = data.srl_system_ntp.all
}
*/

/*
data "srl_acl_ipv4_filter" "tf-test" {
  ipv4_filter {
    name = "tf-test"
  }
}

output "data_srl_acl_ipv4_filter" {
  value = data.srl_acl_ipv4_filter.tf-test
}
*/

/*
resource "srl_system_ntp" "this" {
   ntp {
     admin_state = "enable"
     network_instance = "mgmt"
     server {
       address = "193.104.37.238"
       prefer = true
     }
     server {
       address = "45.87.76.3"
     }
     server {
       address = "185.111.204.220"
     }
    #  server {
    #    address = "162.159.200.1"
    #  }
   }
}

output "resource_srl_system_ntp" {
  value = srl_system_ntp.this
}
*/

/*
resource "srl_system_information" "this" {
  information {
    contact = "Wim Henderickx & Mieke Jacobs"
    location = "Copernicuslann 50, 2018 antwerp"
  }
}

output "resource_srl_system_information" {
  value = srl_system_information.this
}
*/

/*
resource "srl_system_clock" "this" {
  clock {
    timezone = "Europe/Stockholm"
  }
}

output "resource_srl_system_clock" {
  value = srl_system_clock.this
}
*/

/*
resource "srl_acl_ipv4_filter" "tf-test" {
  ipv4_filter {
    description = "terraform test description"
    name = "tf-test"
    entry {
      sequence_id = "10"
      description = "tf entry 10 test1"
      #match {
      #  destination_port {
      #    operator = "eq"
      #    value = "80"
      #  }
      #}
      # action {
      #   accept {
      #     log = false
      #   }
      # }
    }
  }
}

output "resource_srl_acl_ipv4_filter" {
  value = srl_acl_ipv4_filter.tf-test
}
*/




resource "srl_interfaces" "this" {
  interface {
    admin_state = "disable"
    name = "ethernet-1/2"
    description = "tf-test"
    mtu = 1500
    vlan_tagging = false
  }
}

output "resource_interface" {
  value = srl_interfaces.this
}

resource "srl_interfaces_subinterface" "this" {
  interface_id = srl_interfaces.this.id
  subinterface {
    index = 1
    admin_state = "enable"
    description = "terraform test subinterface"
    ip_mtu = 1500
    type = "routed"
  }
}

output "resource_sub_interface" {
  value = srl_interfaces_subinterface.this
}

resource "srl_interfaces_subinterface_ipv4" "this" {
  interface_id = srl_interfaces.this.id
  subinterface_id = srl_interfaces_subinterface.this.id
  ipv4 {

  }
}

resource "srl_interfaces_subinterface_ipv4_address" "this" {
  interface_id = srl_interfaces.this.id
  subinterface_id = srl_interfaces_subinterface.this.id
  ipv4_id = srl_interfaces_subinterface_ipv4.this.id
  address {
    ip_prefix = "10.20.20.20/24"
  }
}

resource "srl_network_instance_instance" "this" {
  network_instance {
    name = "terraform"
    interface {
      name = "ethernet-1/2.1"
    }
    type = "ip-vrf"
    admin_state = "enable"
    description = "terraform created network instance"
    router_id = "10.1.1.2"
  }
}

output "resource_network_instance_instance" {
  value = srl_network_instance_instance.this
}

resource "srl_network_instance_instance_next_hop_groups" "this" {
  network_instance_id = srl_network_instance_instance.this.id
  next_hop_groups {
    group {
      name = "tf-nh-group"
      admin_state = "enable"
    }
  }
}

output "resource_network_instance_next_hop_groups" {
  value = srl_network_instance_instance_next_hop_groups.this
}

resource "srl_network_instance_instance_static_routes" "this" {
  network_instance_id = srl_network_instance_instance.this.id
  static_routes {
    route {
      prefix = "172.4.0.0/24"
      admin_state = "enable"
      metric = "30"
      next_hop_group = srl_network_instance_instance_next_hop_groups.this.next_hop_groups[0].group[0].name
    }
  }
}

output "resource_network_instance_static_routes" {
  value = srl_network_instance_instance_static_routes.this
}

resource "srl_network_instance_instance_protocols_bgp" "this" {
  network_instance_id = srl_network_instance_instance.this.id
  bgp {
    autonomous_system = 65000
    router_id = "10.1.1.2"
    #admin_state = "disable"
  }
}

output "output_srl_network_instance_instance_protocols_bgp" {
  value = srl_network_instance_instance_protocols_bgp.this
}

resource "srl_network_instance_instance_protocols_bgp_group" "this" {
  network_instance_id = srl_network_instance_instance.this.id
  bgp_id = srl_network_instance_instance_protocols_bgp.this.id
  group {
    group_name = "tf-test"
    description = "tf-test-description"
    peer_as = 65001
  }
}

output "output_srl_network_instance_instance_protocols_bgp_group" {
  value = srl_network_instance_instance_protocols_bgp_group.this
}

resource "srl_network_instance_instance_protocols_bgp_neighbor" "this" {
  network_instance_id = srl_network_instance_instance.this.id
  bgp_id = srl_network_instance_instance_protocols_bgp.this.id
  neighbor {
    peer_address = "172.1.1.1"
    admin_state = "enable"
    description = "terraform neighbor"
    peer_as = 65002
    peer_group = srl_network_instance_instance_protocols_bgp_group.this.id
  }
}
