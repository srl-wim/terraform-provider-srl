
package tfsrl

import (
	"context"
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Provider function
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"target": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "target to connect to <ip>:<port>",
				DefaultFunc: schema.EnvDefaultFunc("SRL_TARGET", nil),
			},
			"proxy": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "proxy to connect to the device",
				DefaultFunc: schema.EnvDefaultFunc("SRL_PROXY", false),
			},
			"username": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "username to connect to SRL device",
				DefaultFunc: schema.EnvDefaultFunc("SRL_USERNAME", nil),
			},
			"password": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "password to connect to SRL device",
				DefaultFunc: schema.EnvDefaultFunc("SRL_PASSWORD", nil),
			},
			"timeout": {
				Type:        schema.TypeFloat,
				Optional:    true,
				Description: "grpc timeout (default: 30 sec)",
				DefaultFunc: schema.EnvDefaultFunc("SRL_TIMEOUT", defaultTimeout.Seconds()),
			},
			"insecure": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "insecure connection",
				DefaultFunc: schema.EnvDefaultFunc("SRL_INSECURE", false),
			},
			"tls_ca": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "tls certificate authority",
				DefaultFunc: schema.EnvDefaultFunc("SRL_TLS_CA", nil),
			},
			"tls_cert": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "tls certificate",
				DefaultFunc: schema.EnvDefaultFunc("SRL_TLS_CERT", nil),
			},
			"tls_key": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "tls key",
				DefaultFunc: schema.EnvDefaultFunc("SRL_TLS_KEY", nil),
			},
			"skip_verify": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "skip verify tls connection",
				DefaultFunc: schema.EnvDefaultFunc("SRL_SKIP_VERIFY", false),
			},
			"encoding": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "grpc encoding",
				DefaultFunc: schema.EnvDefaultFunc("SRL_ENCODING", defaultEncoding),
			},
			"debug": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "initialize the debug level in the srl provider",
				DefaultFunc: schema.EnvDefaultFunc("SRL_DEBUG", false),
			},
		},
		DataSourcesMap: map[string]*schema.Resource{
			
            "srl_acl_capture_filter_ipv4_filter": dataAclCaptureFilterIpv4Filter(),
            
            "srl_acl_capture_filter_ipv6_filter": dataAclCaptureFilterIpv6Filter(),
            
            "srl_acl_cpm_filter_ipv4_filter": dataAclCpmFilterIpv4Filter(),
            
            "srl_acl_cpm_filter_ipv6_filter": dataAclCpmFilterIpv6Filter(),
            
            "srl_acl_ipv4_filter": dataAclIpv4Filter(),
            
            "srl_acl_ipv6_filter": dataAclIpv6Filter(),
            
            "srl_acl_policers_policer": dataAclPolicersPolicer(),
            
            "srl_acl_policers_system_cpu_policer": dataAclPolicersSystemCpuPolicer(),
            
            "srl_interfaces": dataInterfaces(),
            
            "srl_interfaces_subinterface": dataInterfacesSubinterface(),
            
            "srl_interfaces_subinterface_acl": dataInterfacesSubinterfaceAcl(),
            
            "srl_interfaces_subinterface_ipv4": dataInterfacesSubinterfaceIpv4(),
            
            "srl_interfaces_subinterface_ipv4_address": dataInterfacesSubinterfaceIpv4Address(),
            
            "srl_interfaces_subinterface_ipv4_arp": dataInterfacesSubinterfaceIpv4Arp(),
            
            "srl_interfaces_subinterface_ipv4_dhcp_client": dataInterfacesSubinterfaceIpv4DhcpClient(),
            
            "srl_interfaces_subinterface_ipv4_dhcp_relay": dataInterfacesSubinterfaceIpv4DhcpRelay(),
            
            "srl_interfaces_subinterface_ipv4_vrrp": dataInterfacesSubinterfaceIpv4Vrrp(),
            
            "srl_interfaces_subinterface_ipv6": dataInterfacesSubinterfaceIpv6(),
            
            "srl_interfaces_subinterface_ipv6_address": dataInterfacesSubinterfaceIpv6Address(),
            
            "srl_interfaces_subinterface_ipv6_dhcp_client": dataInterfacesSubinterfaceIpv6DhcpClient(),
            
            "srl_interfaces_subinterface_ipv6_dhcp_relay": dataInterfacesSubinterfaceIpv6DhcpRelay(),
            
            "srl_interfaces_subinterface_ipv6_vrrp": dataInterfacesSubinterfaceIpv6Vrrp(),
            
            "srl_interfaces_subinterface_vlan": dataInterfacesSubinterfaceVlan(),
            
            "srl_network_instance_instance": dataNetworkInstanceInstance(),
            
            "srl_network_instance_instance_aggregate_routes": dataNetworkInstanceInstanceAggregateRoutes(),
            
            "srl_network_instance_instance_next_hop_groups": dataNetworkInstanceInstanceNextHopGroups(),
            
            "srl_network_instance_instance_protocols_bgp": dataNetworkInstanceInstanceProtocolsBgp(),
            
            "srl_network_instance_instance_protocols_bgp_group": dataNetworkInstanceInstanceProtocolsBgpGroup(),
            
            "srl_network_instance_instance_protocols_bgp_neighbor": dataNetworkInstanceInstanceProtocolsBgpNeighbor(),
            
            "srl_network_instance_instance_protocols_isis": dataNetworkInstanceInstanceProtocolsIsis(),
            
            "srl_network_instance_instance_protocols_ospf": dataNetworkInstanceInstanceProtocolsOspf(),
            
            "srl_network_instance_instance_static_routes": dataNetworkInstanceInstanceStaticRoutes(),
            
            "srl_system_aaa": dataSystemAaa(),
            
            "srl_system_authentication": dataSystemAuthentication(),
            
            "srl_system_banner": dataSystemBanner(),
            
            "srl_system_boot": dataSystemBoot(),
            
            "srl_system_bridge_table": dataSystemBridgeTable(),
            
            "srl_system_clock": dataSystemClock(),
            
            "srl_system_configuration": dataSystemConfiguration(),
            
            "srl_system_dns": dataSystemDns(),
            
            "srl_system_ftp_server": dataSystemFtpServer(),
            
            "srl_system_gnmi_server": dataSystemGnmiServer(),
            
            "srl_system_information": dataSystemInformation(),
            
            "srl_system_ip_load_balancing": dataSystemIpLoadBalancing(),
            
            "srl_system_json_rpc_server": dataSystemJsonRpcServer(),
            
            "srl_system_lacp": dataSystemLacp(),
            
            "srl_system_lldp": dataSystemLldp(),
            
            "srl_system_maintenance": dataSystemMaintenance(),
            
            "srl_system_mtu": dataSystemMtu(),
            
            "srl_system_name": dataSystemName(),
            
            "srl_system_ntp": dataSystemNtp(),
            
            "srl_system_sflow": dataSystemSflow(),
            
            "srl_system_snmp": dataSystemSnmp(),
            
            "srl_system_ssh_server": dataSystemSshServer(),
            
            "srl_system_tls": dataSystemTls(),
            
		},
		ResourcesMap: map[string]*schema.Resource{
			
            "srl_acl_capture_filter_ipv4_filter": resourceAclCaptureFilterIpv4Filter(),
            
            "srl_acl_capture_filter_ipv6_filter": resourceAclCaptureFilterIpv6Filter(),
            
            "srl_acl_cpm_filter_ipv4_filter": resourceAclCpmFilterIpv4Filter(),
            
            "srl_acl_cpm_filter_ipv6_filter": resourceAclCpmFilterIpv6Filter(),
            
            "srl_acl_ipv4_filter": resourceAclIpv4Filter(),
            
            "srl_acl_ipv6_filter": resourceAclIpv6Filter(),
            
            "srl_acl_policers_policer": resourceAclPolicersPolicer(),
            
            "srl_acl_policers_system_cpu_policer": resourceAclPolicersSystemCpuPolicer(),
            
            "srl_interfaces": resourceInterfaces(),
            
            "srl_interfaces_subinterface": resourceInterfacesSubinterface(),
            
            "srl_interfaces_subinterface_acl": resourceInterfacesSubinterfaceAcl(),
            
            "srl_interfaces_subinterface_ipv4": resourceInterfacesSubinterfaceIpv4(),
            
            "srl_interfaces_subinterface_ipv4_address": resourceInterfacesSubinterfaceIpv4Address(),
            
            "srl_interfaces_subinterface_ipv4_arp": resourceInterfacesSubinterfaceIpv4Arp(),
            
            "srl_interfaces_subinterface_ipv4_dhcp_client": resourceInterfacesSubinterfaceIpv4DhcpClient(),
            
            "srl_interfaces_subinterface_ipv4_dhcp_relay": resourceInterfacesSubinterfaceIpv4DhcpRelay(),
            
            "srl_interfaces_subinterface_ipv4_vrrp": resourceInterfacesSubinterfaceIpv4Vrrp(),
            
            "srl_interfaces_subinterface_ipv6": resourceInterfacesSubinterfaceIpv6(),
            
            "srl_interfaces_subinterface_ipv6_address": resourceInterfacesSubinterfaceIpv6Address(),
            
            "srl_interfaces_subinterface_ipv6_dhcp_client": resourceInterfacesSubinterfaceIpv6DhcpClient(),
            
            "srl_interfaces_subinterface_ipv6_dhcp_relay": resourceInterfacesSubinterfaceIpv6DhcpRelay(),
            
            "srl_interfaces_subinterface_ipv6_vrrp": resourceInterfacesSubinterfaceIpv6Vrrp(),
            
            "srl_interfaces_subinterface_vlan": resourceInterfacesSubinterfaceVlan(),
            
            "srl_network_instance_instance": resourceNetworkInstanceInstance(),
            
            "srl_network_instance_instance_aggregate_routes": resourceNetworkInstanceInstanceAggregateRoutes(),
            
            "srl_network_instance_instance_next_hop_groups": resourceNetworkInstanceInstanceNextHopGroups(),
            
            "srl_network_instance_instance_protocols_bgp": resourceNetworkInstanceInstanceProtocolsBgp(),
            
            "srl_network_instance_instance_protocols_bgp_group": resourceNetworkInstanceInstanceProtocolsBgpGroup(),
            
            "srl_network_instance_instance_protocols_bgp_neighbor": resourceNetworkInstanceInstanceProtocolsBgpNeighbor(),
            
            "srl_network_instance_instance_protocols_isis": resourceNetworkInstanceInstanceProtocolsIsis(),
            
            "srl_network_instance_instance_protocols_ospf": resourceNetworkInstanceInstanceProtocolsOspf(),
            
            "srl_network_instance_instance_static_routes": resourceNetworkInstanceInstanceStaticRoutes(),
            
            "srl_system_aaa": resourceSystemAaa(),
            
            "srl_system_authentication": resourceSystemAuthentication(),
            
            "srl_system_banner": resourceSystemBanner(),
            
            "srl_system_boot": resourceSystemBoot(),
            
            "srl_system_bridge_table": resourceSystemBridgeTable(),
            
            "srl_system_clock": resourceSystemClock(),
            
            "srl_system_configuration": resourceSystemConfiguration(),
            
            "srl_system_dns": resourceSystemDns(),
            
            "srl_system_ftp_server": resourceSystemFtpServer(),
            
            "srl_system_gnmi_server": resourceSystemGnmiServer(),
            
            "srl_system_information": resourceSystemInformation(),
            
            "srl_system_ip_load_balancing": resourceSystemIpLoadBalancing(),
            
            "srl_system_json_rpc_server": resourceSystemJsonRpcServer(),
            
            "srl_system_lacp": resourceSystemLacp(),
            
            "srl_system_lldp": resourceSystemLldp(),
            
            "srl_system_maintenance": resourceSystemMaintenance(),
            
            "srl_system_mtu": resourceSystemMtu(),
            
            "srl_system_name": resourceSystemName(),
            
            "srl_system_ntp": resourceSystemNtp(),
            
            "srl_system_sflow": resourceSystemSflow(),
            
            "srl_system_snmp": resourceSystemSnmp(),
            
            "srl_system_ssh_server": resourceSystemSshServer(),
            
            "srl_system_tls": resourceSystemTls(),
            
		},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	debug := d.Get("debug").(bool)
	if debug {
		log.SetLevel(log.DebugLevel)
	}
	t := d.Get("target").(string)
	u := d.Get("username").(string)
	p := d.Get("password").(string)
	proxy := d.Get("proxy").(bool)
	i := d.Get("insecure").(bool)
	tlsCA := d.Get("tls_ca").(string)
	tlsCert := d.Get("tls_cert").(string)
	tlsKey := d.Get("tls_key").(string)
	s := d.Get("skip_verify").(bool)
	e := d.Get("encoding").(string)

	log.Debugf("target: %s, username: %s and password: %s", t, u, p)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	if t == "" || u == "" || p == "" {
		return nil, diag.Errorf("target (%s), username (%s) and password (%s) must be set", t, u, p)
	}

	tc := &TargetConfig{
		Target:     &t,
		Proxy:      &proxy,
		Username:   &u,
		Password:   &p,
		Timeout:    defaultTimeout,
		Insecure:   &i,
		TLSCA:      &tlsCA,
		TLSCert:    &tlsCert,
		TLSKey:     &tlsKey,
		SkipVerify: &s,
		Encoding:   &e,
		Debug:      &debug,
	}

	target := &Target{
		Config: tc,
	}

	opts := target.createCollectorDialOpts()
	if err := target.CreateGNMIClient(ctx, opts...); err != nil {
		return nil, diag.FromErr(err)
	}

	return target, diags
}

type resourceIDStringer interface {
	Id() string
}

func resourceIDString(d resourceIDStringer, name string) string {
	id := d.Id()
	if id == "" {
		id = "<new resource>"
	}

	return fmt.Sprintf("%s (ID = %s)", name, id)
}	
