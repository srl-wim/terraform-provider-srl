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

// resourceNetworkInstanceInstanceProtocolsOspfString function
func resourceNetworkInstanceInstanceProtocolsOspfString(d resourceIDStringer) string {
	return resourceIDString(d, "network_instance_instance_protocols_ospf")
}

// resourceNetworkInstanceInstanceProtocolsOspf function
func resourceNetworkInstanceInstanceProtocolsOspf() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceNetworkInstanceInstanceProtocolsOspfCreate,
		ReadContext:   resourceNetworkInstanceInstanceProtocolsOspfRead,
		UpdateContext: resourceNetworkInstanceInstanceProtocolsOspfUpdate,
		DeleteContext: resourceNetworkInstanceInstanceProtocolsOspfDelete,

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
        "ospf": {
            Type:     schema.TypeList,
            Optional: true,
            MaxItems: 1,
            Elem: &schema.Resource{
            	Schema: map[string]*schema.Schema{
                    "instance": {
                        Type:     schema.TypeList,
                        Optional: true,
                        MaxItems: 16,
                        Elem: &schema.Resource{
                        	Schema: map[string]*schema.Schema{
                                "address_family": {
                                    Type:     schema.TypeString,
                                    Optional: true,
                                },
                                "admin_state": {
                                    Type:     schema.TypeString,
                                    Optional: true,
                                    Default: "disable",
                                },
                                "advertise_router_capability": {
                                    Type:     schema.TypeString,
                                    Optional: true,
                                },
                                "area": {
                                    Type:     schema.TypeList,
                                    Optional: true,
                                    MaxItems: 16,
                                    Elem: &schema.Resource{
                                    	Schema: map[string]*schema.Schema{
                                            "advertise_router_capability": {
                                                Type:     schema.TypeBool,
                                                Optional: true,
                                                Default: true,
                                            },
                                            "area_id": {
                                                Type:     schema.TypeString,
                                                Required: true,
                                                ForceNew: true,
                                            },
                                            "area_range": {
                                                Type:     schema.TypeList,
                                                Optional: true,
                                                MaxItems: 16,
                                                Elem: &schema.Resource{
                                                	Schema: map[string]*schema.Schema{
                                                        "advertise": {
                                                            Type:     schema.TypeBool,
                                                            Optional: true,
                                                            Default: true,
                                                        },
                                                        "ip_prefix_mask": {
                                                            Type:     schema.TypeString,
                                                            Required: true,
                                                            ForceNew: true,
                                                        },
                                                    },
                                                },
                                            },
                                            "blackhole_aggregate": {
                                                Type:     schema.TypeBool,
                                                Optional: true,
                                            },
                                            "export_policy": {
                                                Type:     schema.TypeString,
                                                Optional: true,
                                            },
                                            "interface": {
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
                                                        "advertise_router_capability": {
                                                            Type:     schema.TypeBool,
                                                            Optional: true,
                                                            Default: true,
                                                        },
                                                        "advertise_subnet": {
                                                            Type:     schema.TypeBool,
                                                            Optional: true,
                                                            Default: true,
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
                                                        "dead_interval": {
                                                            Type:     schema.TypeInt,
                                                            Optional: true,
                                                            Default: "40",
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
                                                                },
                                                            },
                                                        },
                                                        "hello_interval": {
                                                            Type:     schema.TypeInt,
                                                            Optional: true,
                                                            Default: "10",
                                                        },
                                                        "interface_name": {
                                                            Type:     schema.TypeString,
                                                            Required: true,
                                                            ForceNew: true,
                                                        },
                                                        "interface_type": {
                                                            Type:     schema.TypeString,
                                                            Optional: true,
                                                        },
                                                        "lsa_filter_out": {
                                                            Type:     schema.TypeString,
                                                            Optional: true,
                                                            Default: "none",
                                                        },
                                                        "metric": {
                                                            Type:     schema.TypeInt,
                                                            Optional: true,
                                                        },
                                                        "mtu": {
                                                            Type:     schema.TypeInt,
                                                            Optional: true,
                                                        },
                                                        "passive": {
                                                            Type:     schema.TypeBool,
                                                            Optional: true,
                                                        },
                                                        "priority": {
                                                            Type:     schema.TypeInt,
                                                            Optional: true,
                                                            Default: "1",
                                                        },
                                                        "retransmit_interval": {
                                                            Type:     schema.TypeInt,
                                                            Optional: true,
                                                            Default: "5",
                                                        },
                                                        "transit_delay": {
                                                            Type:     schema.TypeInt,
                                                            Optional: true,
                                                            Default: "1",
                                                        },
                                                    },
                                                },
                                            },
                                            "nssa": {
                                                Type:     schema.TypeList,
                                                Optional: true,
                                                MaxItems: 1,
                                                Elem: &schema.Resource{
                                                	Schema: map[string]*schema.Schema{
                                                        "area_range": {
                                                            Type:     schema.TypeList,
                                                            Optional: true,
                                                            MaxItems: 16,
                                                            Elem: &schema.Resource{
                                                            	Schema: map[string]*schema.Schema{
                                                                    "advertise": {
                                                                        Type:     schema.TypeBool,
                                                                        Optional: true,
                                                                        Default: true,
                                                                    },
                                                                    "ip_prefix_mask": {
                                                                        Type:     schema.TypeString,
                                                                        Required: true,
                                                                        ForceNew: true,
                                                                    },
                                                                },
                                                            },
                                                        },
                                                        "originate_default_route": {
                                                            Type:     schema.TypeList,
                                                            Optional: true,
                                                            MaxItems: 1,
                                                            Elem: &schema.Resource{
                                                            	Schema: map[string]*schema.Schema{
                                                                    "adjacency_check": {
                                                                        Type:     schema.TypeBool,
                                                                        Optional: true,
                                                                        Default: true,
                                                                    },
                                                                    "type_nssa": {
                                                                        Type:     schema.TypeBool,
                                                                        Optional: true,
                                                                        Default: false,
                                                                    },
                                                                },
                                                            },
                                                        },
                                                        "redistribute_external": {
                                                            Type:     schema.TypeBool,
                                                            Optional: true,
                                                        },
                                                        "summaries": {
                                                            Type:     schema.TypeBool,
                                                            Optional: true,
                                                        },
                                                    },
                                                },
                                            },
                                            "stub": {
                                                Type:     schema.TypeList,
                                                Optional: true,
                                                MaxItems: 1,
                                                Elem: &schema.Resource{
                                                	Schema: map[string]*schema.Schema{
                                                        "default_metric": {
                                                            Type:     schema.TypeInt,
                                                            Optional: true,
                                                        },
                                                        "summaries": {
                                                            Type:     schema.TypeBool,
                                                            Optional: true,
                                                        },
                                                    },
                                                },
                                            },
                                        },
                                    },
                                },
                                "asbr": {
                                    Type:     schema.TypeList,
                                    Optional: true,
                                    MaxItems: 1,
                                    Elem: &schema.Resource{
                                    	Schema: map[string]*schema.Schema{
                                            "trace_path": {
                                                Type:     schema.TypeString,
                                                Optional: true,
                                                Default: "none",
                                            },
                                        },
                                    },
                                },
                                "export_limit": {
                                    Type:     schema.TypeList,
                                    Optional: true,
                                    MaxItems: 1,
                                    Elem: &schema.Resource{
                                    	Schema: map[string]*schema.Schema{
                                            "log_percent": {
                                                Type:     schema.TypeInt,
                                                Optional: true,
                                            },
                                            "number": {
                                                Type:     schema.TypeInt,
                                                Optional: true,
                                            },
                                        },
                                    },
                                },
                                "export_policy": {
                                    Type:     schema.TypeString,
                                    Optional: true,
                                },
                                "external_db_overflow": {
                                    Type:     schema.TypeList,
                                    Optional: true,
                                    MaxItems: 1,
                                    Elem: &schema.Resource{
                                    	Schema: map[string]*schema.Schema{
                                            "interval": {
                                                Type:     schema.TypeInt,
                                                Optional: true,
                                                Default: "0",
                                            },
                                            "limit": {
                                                Type:     schema.TypeInt,
                                                Optional: true,
                                                Default: "0",
                                            },
                                        },
                                    },
                                },
                                "external_preference": {
                                    Type:     schema.TypeInt,
                                    Optional: true,
                                    Default: "150",
                                },
                                "graceful_restart": {
                                    Type:     schema.TypeList,
                                    Optional: true,
                                    MaxItems: 1,
                                    Elem: &schema.Resource{
                                    	Schema: map[string]*schema.Schema{
                                            "helper_mode": {
                                                Type:     schema.TypeBool,
                                                Optional: true,
                                                Default: false,
                                            },
                                            "strict_lsa_checking": {
                                                Type:     schema.TypeBool,
                                                Optional: true,
                                                Default: false,
                                            },
                                        },
                                    },
                                },
                                "max_ecmp_paths": {
                                    Type:     schema.TypeString,
                                    Optional: true,
                                    Default: "1",
                                },
                                "name": {
                                    Type:     schema.TypeString,
                                    Required: true,
                                    ForceNew: true,
                                },
                                "overload": {
                                    Type:     schema.TypeList,
                                    Optional: true,
                                    MaxItems: 1,
                                    Elem: &schema.Resource{
                                    	Schema: map[string]*schema.Schema{
                                            "active": {
                                                Type:     schema.TypeBool,
                                                Optional: true,
                                                Default: false,
                                            },
                                            "overload_include_ext_1": {
                                                Type:     schema.TypeBool,
                                                Optional: true,
                                                Default: false,
                                            },
                                            "overload_include_ext_2": {
                                                Type:     schema.TypeBool,
                                                Optional: true,
                                                Default: false,
                                            },
                                            "overload_include_ext_stub": {
                                                Type:     schema.TypeBool,
                                                Optional: true,
                                                Default: false,
                                            },
                                            "overload_on_boot": {
                                                Type:     schema.TypeList,
                                                Optional: true,
                                                MaxItems: 1,
                                                Elem: &schema.Resource{
                                                	Schema: map[string]*schema.Schema{
                                                        "timeout": {
                                                            Type:     schema.TypeInt,
                                                            Optional: true,
                                                            Default: "60",
                                                        },
                                                    },
                                                },
                                            },
                                            "rtr_adv_lsa_limit": {
                                                Type:     schema.TypeList,
                                                Optional: true,
                                                MaxItems: 1,
                                                Elem: &schema.Resource{
                                                	Schema: map[string]*schema.Schema{
                                                        "log_only": {
                                                            Type:     schema.TypeBool,
                                                            Optional: true,
                                                        },
                                                        "max_lsa_count": {
                                                            Type:     schema.TypeInt,
                                                            Optional: true,
                                                        },
                                                        "overload_timeout": {
                                                            Type:     schema.TypeString,
                                                            Optional: true,
                                                        },
                                                        "warning_threshold": {
                                                            Type:     schema.TypeString,
                                                            Optional: true,
                                                            Default: "0",
                                                        },
                                                    },
                                                },
                                            },
                                        },
                                    },
                                },
                                "preference": {
                                    Type:     schema.TypeInt,
                                    Optional: true,
                                    Default: "10",
                                },
                                "reference_bandwidth": {
                                    Type:     schema.TypeInt,
                                    Optional: true,
                                    Default: "100000000",
                                },
                                "router_id": {
                                    Type:     schema.TypeString,
                                    Optional: true,
                                },
                                "timers": {
                                    Type:     schema.TypeList,
                                    Optional: true,
                                    MaxItems: 1,
                                    Elem: &schema.Resource{
                                    	Schema: map[string]*schema.Schema{
                                            "incremental_spf_wait": {
                                                Type:     schema.TypeInt,
                                                Optional: true,
                                                Default: "1000",
                                            },
                                            "lsa_accumulate": {
                                                Type:     schema.TypeInt,
                                                Optional: true,
                                                Default: "1000",
                                            },
                                            "lsa_arrival": {
                                                Type:     schema.TypeInt,
                                                Optional: true,
                                                Default: "1000",
                                            },
                                            "lsa_generate": {
                                                Type:     schema.TypeList,
                                                Optional: true,
                                                MaxItems: 1,
                                                Elem: &schema.Resource{
                                                	Schema: map[string]*schema.Schema{
                                                        "lsa_initial_wait": {
                                                            Type:     schema.TypeInt,
                                                            Optional: true,
                                                            Default: "5000",
                                                        },
                                                        "lsa_second_wait": {
                                                            Type:     schema.TypeInt,
                                                            Optional: true,
                                                            Default: "5000",
                                                        },
                                                        "max_lsa_wait": {
                                                            Type:     schema.TypeInt,
                                                            Optional: true,
                                                            Default: "5000",
                                                        },
                                                    },
                                                },
                                            },
                                            "redistribute_delay": {
                                                Type:     schema.TypeInt,
                                                Optional: true,
                                                Default: "1000",
                                            },
                                            "spf_wait": {
                                                Type:     schema.TypeList,
                                                Optional: true,
                                                MaxItems: 1,
                                                Elem: &schema.Resource{
                                                	Schema: map[string]*schema.Schema{
                                                        "spf_initial_wait": {
                                                            Type:     schema.TypeInt,
                                                            Optional: true,
                                                            Default: "1000",
                                                        },
                                                        "spf_max_wait": {
                                                            Type:     schema.TypeInt,
                                                            Optional: true,
                                                            Default: "10000",
                                                        },
                                                        "spf_second_wait": {
                                                            Type:     schema.TypeInt,
                                                            Optional: true,
                                                            Default: "1000",
                                                        },
                                                    },
                                                },
                                            },
                                        },
                                    },
                                },
                                "version": {
                                    Type:     schema.TypeString,
                                    Optional: true,
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

func resourceNetworkInstanceInstanceProtocolsOspfCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Infof("Beginning Create: %s", resourceNetworkInstanceInstanceProtocolsOspfString(d))
	target := meta.(*Target)
	
	key := "ospf"
	
	
	hkey0 := d.Get("network_instance_id").(string)
    
	//hkey := d.Get("[network_instance_id]").(string)
	//p := fmt.Sprintf("fmt.Sprintf("/network-instance[name=%s]/protocols/ospf",hkey0)", hkey)
	p := fmt.Sprintf("/network-instance[name=%s]/protocols/ospf",hkey0)
	
	v := "ospf"
	
	
	hid := make([]string, 0)
	
    hid = append(hid, "network_instance_id")
	
	//hid = append(hid, "network_instance_id")
	//hid := "network_instance_id"
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
	return resourceNetworkInstanceInstanceProtocolsOspfRead(ctx, d, meta)
}

func resourceNetworkInstanceInstanceProtocolsOspfRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Infof("Beginning Read: %s", resourceNetworkInstanceInstanceProtocolsOspfString(d))
	target := meta.(*Target)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	
	
	hkey0 := d.Get("network_instance_id").(string)
    
	//hkey := d.Get("[network_instance_id]").(string)
	

	
	
	//p := fmt.Sprintf("fmt.Sprintf("/network-instance[name=%s]/protocols/ospf",hkey0)", hkey)
	p := fmt.Sprintf("/network-instance[name=%s]/protocols/ospf",hkey0)
	
	

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
			switch x := upd.Values["ospf"].(type) {
			case map[string]interface{}:
				for k, v := range x {
					log.Debugf("BEFORE KEY: %s, VALUE: %v", k, v)
					sk := strings.Split(k, ":")[len(strings.Split(k, ":"))-1]

					switch sk {
					
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
			if err := d.Set("ospf", data); err != nil {
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

func resourceNetworkInstanceInstanceProtocolsOspfUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Infof("Beginning Update: %s", resourceNetworkInstanceInstanceProtocolsOspfString(d))
	target := meta.(*Target)
	
	key := "ospf"
	
	
	hkey0 := d.Get("network_instance_id").(string)
    
	//hkey := d.Get("[network_instance_id]").(string)
	//p := fmt.Sprintf("fmt.Sprintf("/network-instance[name=%s]/protocols/ospf",hkey0)", hkey)
	p := fmt.Sprintf("/network-instance[name=%s]/protocols/ospf",hkey0)
	
	v := "ospf"
	
	
	hid := make([]string, 0)
	
    hid = append(hid, "network_instance_id")
	
	//hid = append(hid, "network_instance_id")
	//hid := "network_instance_id"
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
	return resourceNetworkInstanceInstanceProtocolsOspfRead(ctx, d, meta)
}

func resourceNetworkInstanceInstanceProtocolsOspfDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Debugf("Beginning delete: %s", resourceNetworkInstanceInstanceProtocolsOspfString(d))
	target := meta.(*Target)

	
	
	hkey0 := d.Get("network_instance_id").(string)
    
	//hkey := d.Get("[network_instance_id]").(string)
	
	//p := fmt.Sprintf("fmt.Sprintf("/network-instance[name=%s]/protocols/ospf",hkey0)", hkey)
	p := fmt.Sprintf("/network-instance[name=%s]/protocols/ospf",hkey0)
	
	
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
