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

// dataNetworkInstanceInstanceProtocolsOspfString function
func dataNetworkInstanceInstanceProtocolsOspfString(d resourceIDStringer) string {
	return resourceIDString(d, "network_instance_instance_protocols_ospf")
}

// dataNetworkInstanceInstanceProtocolsOspf function
func dataNetworkInstanceInstanceProtocolsOspf() *schema.Resource {
	return &schema.Resource{
		ReadContext:   dataNetworkInstanceInstanceProtocolsOspfRead,

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Timeouts: &schema.ResourceTimeout{
			Read:   schema.DefaultTimeout(5 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
        "network_instance_id": {
            Type:     schema.TypeString,
            Required: true,
        },
        "ospf": {
            Type:     schema.TypeList,
            Computed: true,
            Elem: &schema.Resource{
            	Schema: map[string]*schema.Schema{
                    "instance": {
                        Type:     schema.TypeList,
                        Computed: true,
                        Elem: &schema.Resource{
                        	Schema: map[string]*schema.Schema{
                                "address_family": {
                                    Type:     schema.TypeString,
                                    Computed: true,
                                },
                                "admin_state": {
                                    Type:     schema.TypeString,
                                    Computed: true,
                                },
                                "advertise_router_capability": {
                                    Type:     schema.TypeString,
                                    Computed: true,
                                },
                                "area": {
                                    Type:     schema.TypeList,
                                    Computed: true,
                                    Elem: &schema.Resource{
                                    	Schema: map[string]*schema.Schema{
                                            "advertise_router_capability": {
                                                Type:     schema.TypeBool,
                                                Computed: true,
                                            },
                                            "area_id": {
                                                Type:     schema.TypeString,
                                                Required: true,
                                            },
                                            "area_range": {
                                                Type:     schema.TypeList,
                                                Computed: true,
                                                Elem: &schema.Resource{
                                                	Schema: map[string]*schema.Schema{
                                                        "advertise": {
                                                            Type:     schema.TypeBool,
                                                            Computed: true,
                                                        },
                                                        "ip_prefix_mask": {
                                                            Type:     schema.TypeString,
                                                            Required: true,
                                                        },
                                                    },
                                                },
                                            },
                                            "blackhole_aggregate": {
                                                Type:     schema.TypeBool,
                                                Computed: true,
                                            },
                                            "export_policy": {
                                                Type:     schema.TypeString,
                                                Computed: true,
                                            },
                                            "interface": {
                                                Type:     schema.TypeList,
                                                Computed: true,
                                                Elem: &schema.Resource{
                                                	Schema: map[string]*schema.Schema{
                                                        "admin_state": {
                                                            Type:     schema.TypeString,
                                                            Computed: true,
                                                        },
                                                        "advertise_router_capability": {
                                                            Type:     schema.TypeBool,
                                                            Computed: true,
                                                        },
                                                        "advertise_subnet": {
                                                            Type:     schema.TypeBool,
                                                            Computed: true,
                                                        },
                                                        "authentication": {
                                                            Type:     schema.TypeList,
                                                            Computed: true,
                                                            Elem: &schema.Resource{
                                                            	Schema: map[string]*schema.Schema{
                                                                    "keychain": {
                                                                        Type:     schema.TypeString,
                                                                        Computed: true,
                                                                    },
                                                                },
                                                            },
                                                        },
                                                        "dead_interval": {
                                                            Type:     schema.TypeInt,
                                                            Computed: true,
                                                        },
                                                        "failure_detection": {
                                                            Type:     schema.TypeList,
                                                            Computed: true,
                                                            Elem: &schema.Resource{
                                                            	Schema: map[string]*schema.Schema{
                                                                    "enable_bfd": {
                                                                        Type:     schema.TypeBool,
                                                                        Computed: true,
                                                                    },
                                                                },
                                                            },
                                                        },
                                                        "hello_interval": {
                                                            Type:     schema.TypeInt,
                                                            Computed: true,
                                                        },
                                                        "interface_name": {
                                                            Type:     schema.TypeString,
                                                            Required: true,
                                                        },
                                                        "interface_type": {
                                                            Type:     schema.TypeString,
                                                            Computed: true,
                                                        },
                                                        "lsa_filter_out": {
                                                            Type:     schema.TypeString,
                                                            Computed: true,
                                                        },
                                                        "metric": {
                                                            Type:     schema.TypeInt,
                                                            Computed: true,
                                                        },
                                                        "mtu": {
                                                            Type:     schema.TypeInt,
                                                            Computed: true,
                                                        },
                                                        "passive": {
                                                            Type:     schema.TypeBool,
                                                            Computed: true,
                                                        },
                                                        "priority": {
                                                            Type:     schema.TypeInt,
                                                            Computed: true,
                                                        },
                                                        "retransmit_interval": {
                                                            Type:     schema.TypeInt,
                                                            Computed: true,
                                                        },
                                                        "transit_delay": {
                                                            Type:     schema.TypeInt,
                                                            Computed: true,
                                                        },
                                                    },
                                                },
                                            },
                                            "nssa": {
                                                Type:     schema.TypeList,
                                                Computed: true,
                                                Elem: &schema.Resource{
                                                	Schema: map[string]*schema.Schema{
                                                        "area_range": {
                                                            Type:     schema.TypeList,
                                                            Computed: true,
                                                            Elem: &schema.Resource{
                                                            	Schema: map[string]*schema.Schema{
                                                                    "advertise": {
                                                                        Type:     schema.TypeBool,
                                                                        Computed: true,
                                                                    },
                                                                    "ip_prefix_mask": {
                                                                        Type:     schema.TypeString,
                                                                        Required: true,
                                                                    },
                                                                },
                                                            },
                                                        },
                                                        "originate_default_route": {
                                                            Type:     schema.TypeList,
                                                            Computed: true,
                                                            Elem: &schema.Resource{
                                                            	Schema: map[string]*schema.Schema{
                                                                    "adjacency_check": {
                                                                        Type:     schema.TypeBool,
                                                                        Computed: true,
                                                                    },
                                                                    "type_nssa": {
                                                                        Type:     schema.TypeBool,
                                                                        Computed: true,
                                                                    },
                                                                },
                                                            },
                                                        },
                                                        "redistribute_external": {
                                                            Type:     schema.TypeBool,
                                                            Computed: true,
                                                        },
                                                        "summaries": {
                                                            Type:     schema.TypeBool,
                                                            Computed: true,
                                                        },
                                                    },
                                                },
                                            },
                                            "stub": {
                                                Type:     schema.TypeList,
                                                Computed: true,
                                                Elem: &schema.Resource{
                                                	Schema: map[string]*schema.Schema{
                                                        "default_metric": {
                                                            Type:     schema.TypeInt,
                                                            Computed: true,
                                                        },
                                                        "summaries": {
                                                            Type:     schema.TypeBool,
                                                            Computed: true,
                                                        },
                                                    },
                                                },
                                            },
                                        },
                                    },
                                },
                                "asbr": {
                                    Type:     schema.TypeList,
                                    Computed: true,
                                    Elem: &schema.Resource{
                                    	Schema: map[string]*schema.Schema{
                                            "trace_path": {
                                                Type:     schema.TypeString,
                                                Computed: true,
                                            },
                                        },
                                    },
                                },
                                "export_limit": {
                                    Type:     schema.TypeList,
                                    Computed: true,
                                    Elem: &schema.Resource{
                                    	Schema: map[string]*schema.Schema{
                                            "log_percent": {
                                                Type:     schema.TypeInt,
                                                Computed: true,
                                            },
                                            "number": {
                                                Type:     schema.TypeInt,
                                                Computed: true,
                                            },
                                        },
                                    },
                                },
                                "export_policy": {
                                    Type:     schema.TypeString,
                                    Computed: true,
                                },
                                "external_db_overflow": {
                                    Type:     schema.TypeList,
                                    Computed: true,
                                    Elem: &schema.Resource{
                                    	Schema: map[string]*schema.Schema{
                                            "interval": {
                                                Type:     schema.TypeInt,
                                                Computed: true,
                                            },
                                            "limit": {
                                                Type:     schema.TypeInt,
                                                Computed: true,
                                            },
                                        },
                                    },
                                },
                                "external_preference": {
                                    Type:     schema.TypeInt,
                                    Computed: true,
                                },
                                "graceful_restart": {
                                    Type:     schema.TypeList,
                                    Computed: true,
                                    Elem: &schema.Resource{
                                    	Schema: map[string]*schema.Schema{
                                            "helper_mode": {
                                                Type:     schema.TypeBool,
                                                Computed: true,
                                            },
                                            "strict_lsa_checking": {
                                                Type:     schema.TypeBool,
                                                Computed: true,
                                            },
                                        },
                                    },
                                },
                                "max_ecmp_paths": {
                                    Type:     schema.TypeString,
                                    Computed: true,
                                },
                                "name": {
                                    Type:     schema.TypeString,
                                    Required: true,
                                },
                                "overload": {
                                    Type:     schema.TypeList,
                                    Computed: true,
                                    Elem: &schema.Resource{
                                    	Schema: map[string]*schema.Schema{
                                            "active": {
                                                Type:     schema.TypeBool,
                                                Computed: true,
                                            },
                                            "overload_include_ext_1": {
                                                Type:     schema.TypeBool,
                                                Computed: true,
                                            },
                                            "overload_include_ext_2": {
                                                Type:     schema.TypeBool,
                                                Computed: true,
                                            },
                                            "overload_include_ext_stub": {
                                                Type:     schema.TypeBool,
                                                Computed: true,
                                            },
                                            "overload_on_boot": {
                                                Type:     schema.TypeList,
                                                Computed: true,
                                                Elem: &schema.Resource{
                                                	Schema: map[string]*schema.Schema{
                                                        "timeout": {
                                                            Type:     schema.TypeInt,
                                                            Computed: true,
                                                        },
                                                    },
                                                },
                                            },
                                            "rtr_adv_lsa_limit": {
                                                Type:     schema.TypeList,
                                                Computed: true,
                                                Elem: &schema.Resource{
                                                	Schema: map[string]*schema.Schema{
                                                        "log_only": {
                                                            Type:     schema.TypeBool,
                                                            Computed: true,
                                                        },
                                                        "max_lsa_count": {
                                                            Type:     schema.TypeInt,
                                                            Computed: true,
                                                        },
                                                        "overload_timeout": {
                                                            Type:     schema.TypeString,
                                                            Computed: true,
                                                        },
                                                        "warning_threshold": {
                                                            Type:     schema.TypeString,
                                                            Computed: true,
                                                        },
                                                    },
                                                },
                                            },
                                        },
                                    },
                                },
                                "preference": {
                                    Type:     schema.TypeInt,
                                    Computed: true,
                                },
                                "reference_bandwidth": {
                                    Type:     schema.TypeInt,
                                    Computed: true,
                                },
                                "router_id": {
                                    Type:     schema.TypeString,
                                    Computed: true,
                                },
                                "timers": {
                                    Type:     schema.TypeList,
                                    Computed: true,
                                    Elem: &schema.Resource{
                                    	Schema: map[string]*schema.Schema{
                                            "incremental_spf_wait": {
                                                Type:     schema.TypeInt,
                                                Computed: true,
                                            },
                                            "lsa_accumulate": {
                                                Type:     schema.TypeInt,
                                                Computed: true,
                                            },
                                            "lsa_arrival": {
                                                Type:     schema.TypeInt,
                                                Computed: true,
                                            },
                                            "lsa_generate": {
                                                Type:     schema.TypeList,
                                                Computed: true,
                                                Elem: &schema.Resource{
                                                	Schema: map[string]*schema.Schema{
                                                        "lsa_initial_wait": {
                                                            Type:     schema.TypeInt,
                                                            Computed: true,
                                                        },
                                                        "lsa_second_wait": {
                                                            Type:     schema.TypeInt,
                                                            Computed: true,
                                                        },
                                                        "max_lsa_wait": {
                                                            Type:     schema.TypeInt,
                                                            Computed: true,
                                                        },
                                                    },
                                                },
                                            },
                                            "redistribute_delay": {
                                                Type:     schema.TypeInt,
                                                Computed: true,
                                            },
                                            "spf_wait": {
                                                Type:     schema.TypeList,
                                                Computed: true,
                                                Elem: &schema.Resource{
                                                	Schema: map[string]*schema.Schema{
                                                        "spf_initial_wait": {
                                                            Type:     schema.TypeInt,
                                                            Computed: true,
                                                        },
                                                        "spf_max_wait": {
                                                            Type:     schema.TypeInt,
                                                            Computed: true,
                                                        },
                                                        "spf_second_wait": {
                                                            Type:     schema.TypeInt,
                                                            Computed: true,
                                                        },
                                                    },
                                                },
                                            },
                                        },
                                    },
                                },
                                "version": {
                                    Type:     schema.TypeString,
                                    Computed: true,
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

func dataNetworkInstanceInstanceProtocolsOspfRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Infof("Beginning Read: %s", dataNetworkInstanceInstanceProtocolsOspfString(d))
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
