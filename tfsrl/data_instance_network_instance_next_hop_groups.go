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

// dataNetworkInstanceInstanceNextHopGroupsString function
func dataNetworkInstanceInstanceNextHopGroupsString(d resourceIDStringer) string {
	return resourceIDString(d, "network_instance_instance_next_hop_groups")
}

// dataNetworkInstanceInstanceNextHopGroups function
func dataNetworkInstanceInstanceNextHopGroups() *schema.Resource {
	return &schema.Resource{
		ReadContext:   dataNetworkInstanceInstanceNextHopGroupsRead,

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
        "next_hop_groups": {
            Type:     schema.TypeList,
            Computed: true,
            Elem: &schema.Resource{
            	Schema: map[string]*schema.Schema{
                    "group": {
                        Type:     schema.TypeList,
                        Computed: true,
                        Elem: &schema.Resource{
                        	Schema: map[string]*schema.Schema{
                                "admin_state": {
                                    Type:     schema.TypeString,
                                    Computed: true,
                                },
                                "blackhole": {
                                    Type:     schema.TypeList,
                                    Computed: true,
                                    Elem: &schema.Resource{
                                    	Schema: map[string]*schema.Schema{
                                            "generate_icmp": {
                                                Type:     schema.TypeBool,
                                                Computed: true,
                                            },
                                        },
                                    },
                                },
                                "name": {
                                    Type:     schema.TypeString,
                                    Required: true,
                                },
                                "nexthop": {
                                    Type:     schema.TypeList,
                                    Computed: true,
                                    Elem: &schema.Resource{
                                    	Schema: map[string]*schema.Schema{
                                            "admin_state": {
                                                Type:     schema.TypeString,
                                                Computed: true,
                                            },
                                            "index": {
                                                Type:     schema.TypeInt,
                                                Required: true,
                                            },
                                            "ip_address": {
                                                Type:     schema.TypeString,
                                                Computed: true,
                                            },
                                            "pushed_mpls_label_stack": {
                                                Type:     schema.TypeString,
                                                Computed: true,
                                            },
                                            "resolve": {
                                                Type:     schema.TypeBool,
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
    }
}

func dataNetworkInstanceInstanceNextHopGroupsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Infof("Beginning Read: %s", dataNetworkInstanceInstanceNextHopGroupsString(d))
	target := meta.(*Target)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	
	
	hkey0 := d.Get("network_instance_id").(string)
    
	//hkey := d.Get("[network_instance_id]").(string)
	

	
	
	//p := fmt.Sprintf("fmt.Sprintf("/network-instance[name=%s]/next-hop-groups",hkey0)", hkey)
	p := fmt.Sprintf("/network-instance[name=%s]/next-hop-groups",hkey0)
	
	

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
			switch x := upd.Values["next-hop-groups"].(type) {
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
			if err := d.Set("next_hop_groups", data); err != nil {
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
