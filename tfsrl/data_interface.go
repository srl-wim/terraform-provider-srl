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
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	log "github.com/sirupsen/logrus"
)

// dataInterfacesString function
func dataInterfacesString(d resourceIDStringer) string {
	return resourceIDString(d, "interfaces")
}

// dataInterfaces function
func dataInterfaces() *schema.Resource {
	return &schema.Resource{
		ReadContext:   dataInterfacesRead,

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Timeouts: &schema.ResourceTimeout{
			Read:   schema.DefaultTimeout(5 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
        "interface": {
            Type:     schema.TypeList,
            Required: true,
            Elem: &schema.Resource{
            	Schema: map[string]*schema.Schema{
                    "admin_state": {
                        Type:     schema.TypeString,
                        Computed: true,
                    },
                    "description": {
                        Type:     schema.TypeString,
                        Computed: true,
                    },
                    "ethernet": {
                        Type:     schema.TypeList,
                        Computed: true,
                        Elem: &schema.Resource{
                        	Schema: map[string]*schema.Schema{
                                "aggregate_id": {
                                    Type:     schema.TypeString,
                                    Computed: true,
                                },
                                "flow_control": {
                                    Type:     schema.TypeList,
                                    Computed: true,
                                    Elem: &schema.Resource{
                                    	Schema: map[string]*schema.Schema{
                                            "receive": {
                                                Type:     schema.TypeBool,
                                                Computed: true,
                                            },
                                        },
                                    },
                                },
                                "lacp_port_priority": {
                                    Type:     schema.TypeInt,
                                    Computed: true,
                                },
                                "port_speed": {
                                    Type:     schema.TypeString,
                                    Computed: true,
                                },
                            },
                        },
                    },
                    "lag": {
                        Type:     schema.TypeList,
                        Computed: true,
                        Elem: &schema.Resource{
                        	Schema: map[string]*schema.Schema{
                                "lacp": {
                                    Type:     schema.TypeList,
                                    Computed: true,
                                    Elem: &schema.Resource{
                                    	Schema: map[string]*schema.Schema{
                                            "admin_key": {
                                                Type:     schema.TypeInt,
                                                Computed: true,
                                            },
                                            "interval": {
                                                Type:     schema.TypeString,
                                                Computed: true,
                                            },
                                            "lacp_mode": {
                                                Type:     schema.TypeString,
                                                Computed: true,
                                            },
                                            "system_id_mac": {
                                                Type:     schema.TypeString,
                                                Computed: true,
                                            },
                                            "system_priority": {
                                                Type:     schema.TypeInt,
                                                Computed: true,
                                            },
                                        },
                                    },
                                },
                                "lag_type": {
                                    Type:     schema.TypeString,
                                    Computed: true,
                                },
                                "member_speed": {
                                    Type:     schema.TypeString,
                                    Computed: true,
                                },
                                "min_links": {
                                    Type:     schema.TypeInt,
                                    Computed: true,
                                },
                            },
                        },
                    },
                    "loopback_mode": {
                        Type:     schema.TypeBool,
                        Computed: true,
                    },
                    "mtu": {
                        Type:     schema.TypeInt,
                        Computed: true,
                    },
                    "name": {
                        Type:     schema.TypeString,
                        Required: true,
                    },
                    "qos": {
                        Type:     schema.TypeList,
                        Computed: true,
                        Elem: &schema.Resource{
                        	Schema: map[string]*schema.Schema{
                                "output": {
                                    Type:     schema.TypeList,
                                    Computed: true,
                                    Elem: &schema.Resource{
                                    	Schema: map[string]*schema.Schema{
                                            "unicast_queue": {
                                                Type:     schema.TypeList,
                                                Computed: true,
                                                Elem: &schema.Resource{
                                                	Schema: map[string]*schema.Schema{
                                                        "queue_id": {
                                                            Type:     schema.TypeString,
                                                            Required: true,
                                                        },
                                                        "queue_parameters": {
                                                            Type:     schema.TypeList,
                                                            Computed: true,
                                                            Elem: &schema.Resource{
                                                            	Schema: map[string]*schema.Schema{
                                                                    "peak_rate_percent": {
                                                                        Type:     schema.TypeInt,
                                                                        Computed: true,
                                                                    },
                                                                    "strict_priority": {
                                                                        Type:     schema.TypeBool,
                                                                        Computed: true,
                                                                    },
                                                                    "weight": {
                                                                        Type:     schema.TypeInt,
                                                                        Computed: true,
                                                                    },
                                                                },
                                                            },
                                                        },
                                                    },
                                                },
                                            },
                                        },
                                    },
                                },
                            },
                        },
                    },
                    "transceiver": {
                        Type:     schema.TypeList,
                        Computed: true,
                        Elem: &schema.Resource{
                        	Schema: map[string]*schema.Schema{
                                "admin_state": {
                                    Type:     schema.TypeString,
                                    Computed: true,
                                },
                                "ddm_events": {
                                    Type:     schema.TypeBool,
                                    Computed: true,
                                },
                                "forward_error_correction": {
                                    Type:     schema.TypeString,
                                    Computed: true,
                                },
                                "tx_laser": {
                                    Type:     schema.TypeBool,
                                    Computed: true,
                                },
                            },
                        },
                    },
                    "vlan_tagging": {
                        Type:     schema.TypeBool,
                        Computed: true,
                    },
                },
            },
        },

        },
    }
}

func dataInterfacesRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Infof("Beginning Read: %s", dataInterfacesString(d))
	target := meta.(*Target)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	

	 
	rn := "interface"
	rk := "name"
	key, err := getResourceListKey(&rn, &rk, d)

	
	//p := fmt.Sprintf("fmt.Sprintf("/interface[name=%s]", key)", key)
	p := fmt.Sprintf("/interface[name=%s]", key)
	
	

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
			switch x := upd.Values["interface"].(type) {
			case map[string]interface{}:
				for k, v := range x {
					log.Debugf("BEFORE KEY: %s, VALUE: %v", k, v)
					sk := strings.Split(k, ":")[len(strings.Split(k, ":"))-1]

					switch sk {
					
					case "sflow":
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
				 
				// add key to the get resp data since it is not returned in the gnmi data
				x[rk] = key
				// append the get resp to data
				
				data = append(data, x)
			}
			log.Debugf("get response: index: %d, data: %v", i, data)
			if err := d.Set("interface", data); err != nil {
				return diag.FromErr(err)
			}
			// always run
			 
			d.SetId(key)
			
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
