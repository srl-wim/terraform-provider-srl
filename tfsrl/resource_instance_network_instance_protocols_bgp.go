/*
Package tfsrl is a generated package which contains definitions
of structs which represent a Terraform resource. 

This package was generated by ygocodegen
using the following YANG input files:
	- yang/20_06_2/srl/
Imported modules were sourced from:
	- yang/20_06_2/ietf/
*/
package tfsrl

import (
	"context"
	"strings"
	
	"fmt"
	
	"strconv"
	
	
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	log "github.com/sirupsen/logrus"
)

// resourceNetworkInstanceInstanceProtocolsBgpString function
func resourceNetworkInstanceInstanceProtocolsBgpString(d resourceIDStringer) string {
	return resourceIDString(d, "network_instance_instance_protocols_bgp")
}

// resourceNetworkInstanceInstanceProtocolsBgp function
func resourceNetworkInstanceInstanceProtocolsBgp() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceNetworkInstanceInstanceProtocolsBgpCreate,
		ReadContext:   resourceNetworkInstanceInstanceProtocolsBgpRead,
		UpdateContext: resourceNetworkInstanceInstanceProtocolsBgpUpdate,
		DeleteContext: resourceNetworkInstanceInstanceProtocolsBgpDelete,

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(5 * time.Minute),
			Read:   schema.DefaultTimeout(5 * time.Minute),
			Update: schema.DefaultTimeout(5 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
        "network_instance_id": {
            Type:     schema.TypeString,
            Required: true,
            ForceNew: true,
        },
        "bgp": {
            Type:     schema.TypeList,
            Optional: true,
            MaxItems: 1,
            Elem: &schema.Resource{
            	Schema: map[string]*schema.Schema{
                    "admin_state": {
                        Type:     schema.TypeString,
                        Optional: true,
                        Default: "enable",
                    },
                    "as_path_options": {
                        Type:     schema.TypeList,
                        Optional: true,
                        MaxItems: 1,
                        Elem: &schema.Resource{
                        	Schema: map[string]*schema.Schema{
                                "allow_own_as": {
                                    Type:     schema.TypeInt,
                                    Optional: true,
                                    Default: "0",
                                },
                                "remove_private_as": {
                                    Type:     schema.TypeList,
                                    Optional: true,
                                    MaxItems: 1,
                                    Elem: &schema.Resource{
                                    	Schema: map[string]*schema.Schema{
                                            "ignore_peer_as": {
                                                Type:     schema.TypeBool,
                                                Optional: true,
                                                Default: false,
                                            },
                                            "leading_only": {
                                                Type:     schema.TypeBool,
                                                Optional: true,
                                                Default: false,
                                            },
                                            "mode": {
                                                Type:     schema.TypeString,
                                                Optional: true,
                                                Default: "disabled",
                                            },
                                        },
                                    },
                                },
                            },
                        },
                    },
                    "authentication": {
                        Type:     schema.TypeList,
                        Optional: true,
                        MaxItems: 1,
                        Elem: &schema.Resource{
                        	Schema: map[string]*schema.Schema{
                                "keychain": {
                                    Type:     schema.TypeString,
                                    Optional: true,
                                },
                            },
                        },
                    },
                    "autonomous_system": {
                        Type:     schema.TypeString,
                        Optional: true,
                    },
                    "convergence": {
                        Type:     schema.TypeList,
                        Optional: true,
                        MaxItems: 1,
                        Elem: &schema.Resource{
                        	Schema: map[string]*schema.Schema{
                                "min_wait_to_advertise": {
                                    Type:     schema.TypeInt,
                                    Optional: true,
                                    Default: "0",
                                },
                            },
                        },
                    },
                    "dynamic_neighbors": {
                        Type:     schema.TypeList,
                        Optional: true,
                        MaxItems: 1,
                        Elem: &schema.Resource{
                        	Schema: map[string]*schema.Schema{
                                "accept": {
                                    Type:     schema.TypeList,
                                    Optional: true,
                                    MaxItems: 1,
                                    Elem: &schema.Resource{
                                    	Schema: map[string]*schema.Schema{
                                            "match": {
                                                Type:     schema.TypeList,
                                                Optional: true,
                                                MaxItems: 16,
                                                Elem: &schema.Resource{
                                                	Schema: map[string]*schema.Schema{
                                                        "allowed_peer_as": {
                                                            Type:     schema.TypeString,
                                                            Optional: true,
                                                        },
                                                        "peer_group": {
                                                            Type:     schema.TypeString,
                                                            Optional: true,
                                                        },
                                                        "prefix": {
                                                            Type:     schema.TypeString,
                                                            Required: true,
                                                            ForceNew: true,
                                                        },
                                                    },
                                                },
                                            },
                                            "max_sessions": {
                                                Type:     schema.TypeInt,
                                                Optional: true,
                                                Default: "0",
                                            },
                                        },
                                    },
                                },
                            },
                        },
                    },
                    "ebgp_default_policy": {
                        Type:     schema.TypeList,
                        Optional: true,
                        MaxItems: 1,
                        Elem: &schema.Resource{
                        	Schema: map[string]*schema.Schema{
                                "export_reject_all": {
                                    Type:     schema.TypeBool,
                                    Optional: true,
                                    Default: true,
                                },
                                "import_reject_all": {
                                    Type:     schema.TypeBool,
                                    Optional: true,
                                    Default: true,
                                },
                            },
                        },
                    },
                    "failure_detection": {
                        Type:     schema.TypeList,
                        Optional: true,
                        MaxItems: 1,
                        Elem: &schema.Resource{
                        	Schema: map[string]*schema.Schema{
                                "enable_bfd": {
                                    Type:     schema.TypeBool,
                                    Optional: true,
                                    Default: false,
                                },
                                "fast_failover": {
                                    Type:     schema.TypeBool,
                                    Optional: true,
                                    Default: true,
                                },
                            },
                        },
                    },
                    "graceful_restart": {
                        Type:     schema.TypeList,
                        Optional: true,
                        MaxItems: 1,
                        Elem: &schema.Resource{
                        	Schema: map[string]*schema.Schema{
                                "admin_state": {
                                    Type:     schema.TypeString,
                                    Optional: true,
                                    Default: "disable",
                                },
                                "stale_routes_time": {
                                    Type:     schema.TypeInt,
                                    Optional: true,
                                    Default: "360",
                                },
                            },
                        },
                    },
                    "ipv4_unicast": {
                        Type:     schema.TypeList,
                        Optional: true,
                        MaxItems: 1,
                        Elem: &schema.Resource{
                        	Schema: map[string]*schema.Schema{
                                "admin_state": {
                                    Type:     schema.TypeString,
                                    Optional: true,
                                    Default: "enable",
                                },
                                "advertise_ipv6_next_hops": {
                                    Type:     schema.TypeBool,
                                    Optional: true,
                                    Default: false,
                                },
                                "convergence": {
                                    Type:     schema.TypeList,
                                    Optional: true,
                                    MaxItems: 1,
                                    Elem: &schema.Resource{
                                    	Schema: map[string]*schema.Schema{
                                            "max_wait_to_advertise": {
                                                Type:     schema.TypeInt,
                                                Optional: true,
                                                Default: "0",
                                            },
                                        },
                                    },
                                },
                                "multipath": {
                                    Type:     schema.TypeList,
                                    Optional: true,
                                    MaxItems: 1,
                                    Elem: &schema.Resource{
                                    	Schema: map[string]*schema.Schema{
                                            "allow_multiple_as": {
                                                Type:     schema.TypeBool,
                                                Optional: true,
                                                Default: true,
                                            },
                                            "max_paths_level_1": {
                                                Type:     schema.TypeString,
                                                Optional: true,
                                                Default: "1",
                                            },
                                            "max_paths_level_2": {
                                                Type:     schema.TypeString,
                                                Optional: true,
                                                Default: "1",
                                            },
                                        },
                                    },
                                },
                                "receive_ipv6_next_hops": {
                                    Type:     schema.TypeBool,
                                    Optional: true,
                                    Default: false,
                                },
                            },
                        },
                    },
                    "ipv6_unicast": {
                        Type:     schema.TypeList,
                        Optional: true,
                        MaxItems: 1,
                        Elem: &schema.Resource{
                        	Schema: map[string]*schema.Schema{
                                "admin_state": {
                                    Type:     schema.TypeString,
                                    Optional: true,
                                    Default: "disable",
                                },
                                "convergence": {
                                    Type:     schema.TypeList,
                                    Optional: true,
                                    MaxItems: 1,
                                    Elem: &schema.Resource{
                                    	Schema: map[string]*schema.Schema{
                                            "max_wait_to_advertise": {
                                                Type:     schema.TypeInt,
                                                Optional: true,
                                                Default: "0",
                                            },
                                        },
                                    },
                                },
                                "multipath": {
                                    Type:     schema.TypeList,
                                    Optional: true,
                                    MaxItems: 1,
                                    Elem: &schema.Resource{
                                    	Schema: map[string]*schema.Schema{
                                            "allow_multiple_as": {
                                                Type:     schema.TypeBool,
                                                Optional: true,
                                                Default: true,
                                            },
                                            "max_paths_level_1": {
                                                Type:     schema.TypeString,
                                                Optional: true,
                                                Default: "1",
                                            },
                                            "max_paths_level_2": {
                                                Type:     schema.TypeString,
                                                Optional: true,
                                                Default: "1",
                                            },
                                        },
                                    },
                                },
                            },
                        },
                    },
                    "local_preference": {
                        Type:     schema.TypeInt,
                        Optional: true,
                        Default: "100",
                    },
                    "preference": {
                        Type:     schema.TypeList,
                        Optional: true,
                        MaxItems: 1,
                        Elem: &schema.Resource{
                        	Schema: map[string]*schema.Schema{
                                "ebgp": {
                                    Type:     schema.TypeInt,
                                    Optional: true,
                                    Default: "170",
                                },
                                "ibgp": {
                                    Type:     schema.TypeInt,
                                    Optional: true,
                                    Default: "170",
                                },
                            },
                        },
                    },
                    "route_advertisement": {
                        Type:     schema.TypeList,
                        Optional: true,
                        MaxItems: 1,
                        Elem: &schema.Resource{
                        	Schema: map[string]*schema.Schema{
                                "rapid_withdrawal": {
                                    Type:     schema.TypeBool,
                                    Optional: true,
                                    Default: false,
                                },
                                "wait_for_fib_install": {
                                    Type:     schema.TypeBool,
                                    Optional: true,
                                    Default: true,
                                },
                            },
                        },
                    },
                    "route_reflector": {
                        Type:     schema.TypeList,
                        Optional: true,
                        MaxItems: 1,
                        Elem: &schema.Resource{
                        	Schema: map[string]*schema.Schema{
                                "client": {
                                    Type:     schema.TypeBool,
                                    Optional: true,
                                    Default: false,
                                },
                                "cluster_id": {
                                    Type:     schema.TypeString,
                                    Optional: true,
                                },
                            },
                        },
                    },
                    "router_id": {
                        Type:     schema.TypeString,
                        Optional: true,
                    },
                    "send_community": {
                        Type:     schema.TypeList,
                        Optional: true,
                        MaxItems: 1,
                        Elem: &schema.Resource{
                        	Schema: map[string]*schema.Schema{
                                "large": {
                                    Type:     schema.TypeBool,
                                    Optional: true,
                                    Default: true,
                                },
                                "standard": {
                                    Type:     schema.TypeBool,
                                    Optional: true,
                                    Default: true,
                                },
                            },
                        },
                    },
                    "trace_options": {
                        Type:     schema.TypeList,
                        Optional: true,
                        MaxItems: 1,
                        Elem: &schema.Resource{
                        	Schema: map[string]*schema.Schema{
                                "flag": {
                                    Type:     schema.TypeList,
                                    Optional: true,
                                    MaxItems: 1,
                                    Elem: &schema.Resource{
                                    	Schema: map[string]*schema.Schema{
                                            "modifier": {
                                                Type:     schema.TypeString,
                                                Optional: true,
                                            },
                                            "name": {
                                                Type:     schema.TypeString,
                                                Required: true,
                                                ForceNew: true,
                                            },
                                        },
                                    },
                                },
                            },
                        },
                    },
                    "transport": {
                        Type:     schema.TypeList,
                        Optional: true,
                        MaxItems: 1,
                        Elem: &schema.Resource{
                        	Schema: map[string]*schema.Schema{
                                "tcp_mss": {
                                    Type:     schema.TypeInt,
                                    Optional: true,
                                    Default: "1024",
                                },
                            },
                        },
                    },
                },
            },
        },

        },
    }
}

func resourceNetworkInstanceInstanceProtocolsBgpCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Infof("Beginning Create: %s", resourceNetworkInstanceInstanceProtocolsBgpString(d))
	target := meta.(*Target)
	
	key := "bgp"
	
	
	hkey0 := d.Get("network_instance_id").(string)
    
	//hkey := d.Get("[network_instance_id]").(string)
	//p := fmt.Sprintf("fmt.Sprintf("/network-instance[name=%s]/protocols/bgp",hkey0)", hkey)
	p := fmt.Sprintf("/network-instance[name=%s]/protocols/bgp",hkey0)
	
	v := "bgp"
	
	
	hid := "network_instance_id"
	req, err := target.CreateSetRequest(&p, &v, &hid, d)
	
	if err != nil {
		return diag.FromErr(err)
	}

	response, err := target.Set(ctx, req)
	if err != nil {
		return diag.FromErr(err)
	}

	replaceInKeys(d.Get(v), "-", "_")
	log.Debugf("Set response: %v", response)

	d.SetId(key)
	return resourceNetworkInstanceInstanceProtocolsBgpRead(ctx, d, meta)
}

func resourceNetworkInstanceInstanceProtocolsBgpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Infof("Beginning Read: %s", resourceNetworkInstanceInstanceProtocolsBgpString(d))
	target := meta.(*Target)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	
	
	hkey0 := d.Get("network_instance_id").(string)
    
	//hkey := d.Get("[network_instance_id]").(string)
	

	
	
	//p := fmt.Sprintf("fmt.Sprintf("/network-instance[name=%s]/protocols/bgp",hkey0)", hkey)
	p := fmt.Sprintf("/network-instance[name=%s]/protocols/bgp",hkey0)
	
	

	req, err := target.CreateGetRequest(&p, "CONFIG", d)
	if err != nil {
		return diag.FromErr(err)
	}
	log.Infof("Get Request: %v", req)
	response, err := target.Get(ctx, req)
	if err != nil {
		return diag.FromErr(err)
	}
	log.Debugf("Get Gnmi read response: %v", response)

	u, err := target.HandleGetRespone(response)
	if err != nil {
		return diag.FromErr(err)
	}
	for i, upd := range u {
		// we expect a single response in the get since we target the explicit resource
		log.Debugf("get response: index: %d, update: %v", i, upd)
		if i <= 0 {
			data := make([]map[string]interface{}, 0)
			switch x := upd.Values["bgp"].(type) {
			case map[string]interface{}:
				for k, v := range x {
					log.Debugf("BEFORE KEY: %s, VALUE: %v", k, v)
					sk := strings.Split(k, ":")[len(strings.Split(k, ":"))-1]

					switch sk {
					
					case "import_policy":
						delete(x, k)
					
					case "export_policy":
						delete(x, k)
					
					case "group":
						delete(x, k)
					
					case "neighbor":
						delete(x, k)
					
					default:
						if k != sk {
							delete(x, k)
							x[sk] = v
						}
					}
                }
                for k, v := range x {
                    log.Debugf("AFTER KEY: %s, VALUE: %v", k, v)
				}
				
				data = append(data, x)
			}
			log.Debugf("get response: index: %d, data: %v", i, data)
			if err := d.Set("bgp", data); err != nil {
				return diag.FromErr(err)
			}
			// always run
			
			d.SetId(strconv.FormatInt(time.Now().Unix(), 10))
			
			return diags
		} else {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Unexpected multiple response",
				Detail:   "We only expect a single response from the read/get response",
			})
			return diags
		}
	}
	// when the response is empty no data exists in the system
	log.Debugf("get response: empty set id to nill")
	d.SetId("")
	return diags
}

func resourceNetworkInstanceInstanceProtocolsBgpUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Infof("Beginning Update: %s", resourceNetworkInstanceInstanceProtocolsBgpString(d))
	target := meta.(*Target)
	
	key := "bgp"
	
	
	hkey0 := d.Get("network_instance_id").(string)
    
	//hkey := d.Get("[network_instance_id]").(string)
	//p := fmt.Sprintf("fmt.Sprintf("/network-instance[name=%s]/protocols/bgp",hkey0)", hkey)
	p := fmt.Sprintf("/network-instance[name=%s]/protocols/bgp",hkey0)
	
	v := "bgp"
	
	
	hid := "network_instance_id"
	req, err := target.CreateSetRequest(&p, &v, &hid, d)
	
	if err != nil {
		return diag.FromErr(err)
	}

	response, err := target.Set(ctx, req)
	if err != nil {
		return diag.FromErr(err)
	}

	replaceInKeys(d.Get(v), "-", "_")
	log.Debugf("Set response: %v", response)

	d.SetId(key)
	return resourceNetworkInstanceInstanceProtocolsBgpRead(ctx, d, meta)
}

func resourceNetworkInstanceInstanceProtocolsBgpDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Debugf("Beginning delete: %s", resourceNetworkInstanceInstanceProtocolsBgpString(d))
	target := meta.(*Target)

	
	
	hkey0 := d.Get("network_instance_id").(string)
    
	//hkey := d.Get("[network_instance_id]").(string)
	
	//p := fmt.Sprintf("fmt.Sprintf("/network-instance[name=%s]/protocols/bgp",hkey0)", hkey)
	p := fmt.Sprintf("/network-instance[name=%s]/protocols/bgp",hkey0)
	
	
	req, err := target.CreateDeleteRequest(&p, d)
	if err != nil {
		return diag.FromErr(err)
	}

	response, err := target.Set(ctx, req)
	if err != nil {
		return diag.FromErr(err)
	}

	log.Debugf("Gnmi Delete Response: %v", response)

	d.SetId("")
	return nil
}
