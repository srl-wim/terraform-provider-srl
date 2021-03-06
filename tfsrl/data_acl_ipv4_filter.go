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

// dataAclIpv4FilterString function
func dataAclIpv4FilterString(d resourceIDStringer) string {
	return resourceIDString(d, "acl_ipv4_filter")
}

// dataAclIpv4Filter function
func dataAclIpv4Filter() *schema.Resource {
	return &schema.Resource{
		ReadContext:   dataAclIpv4FilterRead,

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Timeouts: &schema.ResourceTimeout{
			Read:   schema.DefaultTimeout(5 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
        "ipv4_filter": {
            Type:     schema.TypeList,
            Required: true,
            Elem: &schema.Resource{
            	Schema: map[string]*schema.Schema{
                    "description": {
                        Type:     schema.TypeString,
                        Computed: true,
                    },
                    "entry": {
                        Type:     schema.TypeList,
                        Computed: true,
                        Elem: &schema.Resource{
                        	Schema: map[string]*schema.Schema{
                                "action": {
                                    Type:     schema.TypeList,
                                    Computed: true,
                                    Elem: &schema.Resource{
                                    	Schema: map[string]*schema.Schema{
                                            "accept": {
                                                Type:     schema.TypeList,
                                                Computed: true,
                                                Elem: &schema.Resource{
                                                	Schema: map[string]*schema.Schema{
                                                        "log": {
                                                            Type:     schema.TypeBool,
                                                            Computed: true,
                                                        },
                                                    },
                                                },
                                            },
                                            "drop": {
                                                Type:     schema.TypeList,
                                                Computed: true,
                                                Elem: &schema.Resource{
                                                	Schema: map[string]*schema.Schema{
                                                        "log": {
                                                            Type:     schema.TypeBool,
                                                            Computed: true,
                                                        },
                                                    },
                                                },
                                            },
                                        },
                                    },
                                },
                                "description": {
                                    Type:     schema.TypeString,
                                    Computed: true,
                                },
                                "match": {
                                    Type:     schema.TypeList,
                                    Computed: true,
                                    Elem: &schema.Resource{
                                    	Schema: map[string]*schema.Schema{
                                            "destination_address": {
                                                Type:     schema.TypeString,
                                                Computed: true,
                                            },
                                            "destination_port": {
                                                Type:     schema.TypeList,
                                                Computed: true,
                                                Elem: &schema.Resource{
                                                	Schema: map[string]*schema.Schema{
                                                        "operator": {
                                                            Type:     schema.TypeString,
                                                            Computed: true,
                                                        },
                                                        "range": {
                                                            Type:     schema.TypeList,
                                                            Computed: true,
                                                            Elem: &schema.Resource{
                                                            	Schema: map[string]*schema.Schema{
                                                                    "end": {
                                                                        Type:     schema.TypeString,
                                                                        Computed: true,
                                                                    },
                                                                    "start": {
                                                                        Type:     schema.TypeString,
                                                                        Computed: true,
                                                                    },
                                                                },
                                                            },
                                                        },
                                                        "value": {
                                                            Type:     schema.TypeString,
                                                            Computed: true,
                                                        },
                                                    },
                                                },
                                            },
                                            "first_fragment": {
                                                Type:     schema.TypeBool,
                                                Computed: true,
                                            },
                                            "fragment": {
                                                Type:     schema.TypeBool,
                                                Computed: true,
                                            },
                                            "icmp": {
                                                Type:     schema.TypeList,
                                                Computed: true,
                                                Elem: &schema.Resource{
                                                	Schema: map[string]*schema.Schema{
                                                        "code": {
                                                            Type:     schema.TypeInt,
                                                            Computed: true,
                                                        },
                                                        "type": {
                                                            Type:     schema.TypeString,
                                                            Computed: true,
                                                        },
                                                    },
                                                },
                                            },
                                            "protocol": {
                                                Type:     schema.TypeString,
                                                Computed: true,
                                            },
                                            "source_address": {
                                                Type:     schema.TypeString,
                                                Computed: true,
                                            },
                                            "source_port": {
                                                Type:     schema.TypeList,
                                                Computed: true,
                                                Elem: &schema.Resource{
                                                	Schema: map[string]*schema.Schema{
                                                        "operator": {
                                                            Type:     schema.TypeString,
                                                            Computed: true,
                                                        },
                                                        "range": {
                                                            Type:     schema.TypeList,
                                                            Computed: true,
                                                            Elem: &schema.Resource{
                                                            	Schema: map[string]*schema.Schema{
                                                                    "end": {
                                                                        Type:     schema.TypeString,
                                                                        Computed: true,
                                                                    },
                                                                    "start": {
                                                                        Type:     schema.TypeString,
                                                                        Computed: true,
                                                                    },
                                                                },
                                                            },
                                                        },
                                                        "value": {
                                                            Type:     schema.TypeString,
                                                            Computed: true,
                                                        },
                                                    },
                                                },
                                            },
                                            "tcp_flags": {
                                                Type:     schema.TypeString,
                                                Computed: true,
                                            },
                                        },
                                    },
                                },
                                "sequence_id": {
                                    Type:     schema.TypeInt,
                                    Required: true,
                                },
                            },
                        },
                    },
                    "name": {
                        Type:     schema.TypeString,
                        Required: true,
                    },
                    "statistics_per_entry": {
                        Type:     schema.TypeBool,
                        Computed: true,
                    },
                    "subinterface_specific": {
                        Type:     schema.TypeString,
                        Computed: true,
                    },
                },
            },
        },

        },
    }
}

func dataAclIpv4FilterRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Infof("Beginning Read: %s", dataAclIpv4FilterString(d))
	target := meta.(*Target)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	

	 
	rn := "ipv4_filter"
	rk := "name"
	key, err := getResourceListKey(&rn, &rk, d)

	
	//p := fmt.Sprintf("fmt.Sprintf("/acl/ipv4-filter[name=%s]", key)", key)
	p := fmt.Sprintf("/acl/ipv4-filter[name=%s]", key)
	
	

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
			switch x := upd.Values["ipv4-filter"].(type) {
			case map[string]interface{}:
				for k, v := range x {
					log.Debugf("BEFORE KEY: %s, VALUE: %v", k, v)
					sk := strings.Split(k, ":")[len(strings.Split(k, ":"))-1]

					if _, ok := v.(string); ok {
                        log.Debugf("BEFORE VALUE: %s, %s", k, x[k])
                        x[k] = strings.Split(v.(string), ":")[len(strings.Split(v.(string), ":"))-1]
                        log.Debugf("AFTER VALUE: %s, %s", k, x[k])
					}

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
				 
				// add key to the get resp data since it is not returned in the gnmi data
				x[rk] = key
				// append the get resp to data
				
				data = append(data, x)
			}
			log.Debugf("get response: index: %d, data: %v", i, data)
			if err := d.Set("ipv4_filter", data); err != nil {
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
