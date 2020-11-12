provider "srl" {
  username = var.username
  password = var.password
  target = var.target
  encoding = "JSON_IETF"
  skip_verify = true
  debug = true
}

resource "srl_system_ntp" "ntp" {
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


# resource "srl_system_information" "test" {
#   information {
#     contact = "Wim Henderickx"
#     location = "Copernicuslann 50, 2018 antwerp"
#   }
# }

resource "srl_system_clock" "test" {
  clock {
    timezone = "Europe/Amsterdam"
  }
}

resource "srl_acl_ipv4_filter" "test" {
  ipv4_filter {
    description = "terraform test description new"
    name = "tf-test"
    entry {
      sequence_id = "11"
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

# resource "srl_acl_ipv4_filter" "test2" {
#   ipv4_filter {
#     description = "terraform test2 description"
#     name = "tf-test2"
#     entry {
#       sequence_id = "10"
#       description = "tf entry 10 test2"
#       #match {
#       #  destination_port {
#       #    operator = "eq"
#       #    value = "80"
#       #  }
#       #}
#       # action {
#       #   accept {
#       #     log = false
#       #   }
#       # }
#    }
#  }
# }