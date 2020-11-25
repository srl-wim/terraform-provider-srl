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

// resourceNetworkInstanceInstanceNextHopGroupsString function
func resourceNetworkInstanceInstanceNextHopGroupsString(d resourceIDStringer) string {
	return resourceIDString(d, "network_instance_instance_next_hop_groups")
}

// resourceNetworkInstanceInstanceNextHopGroups function
func resourceNetworkInstanceInstanceNextHopGroups() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceNetworkInstanceInstanceNextHopGroupsCreate,
		ReadContext:   resourceNetworkInstanceInstanceNextHopGroupsRead,
		UpdateContext: resourceNetworkInstanceInstanceNextHopGroupsUpdate,
		DeleteContext: resourceNetworkInstanceInstanceNextHopGroupsDelete,

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
        "next_hop_groups": {
            Type:     schema.TypeList,
            Optional: true,
            MaxItems: 1,
            Elem: &schema.Resource{
            	Schema: map[string]*schema.Schema{
                    "group": {
                        Type:     schema.TypeList,
                        Optional: true,
                        MaxItems: 16,
                        Elem: &schema.Resource{
                        	Schema: map[string]*schema.Schema{
                                "admin_state": {
                                    Type:     schema.TypeString,
                                    Optional: true,
                                    Default: "enable",
                                },
                                "blackhole": {
                                    Type:     schema.TypeList,
                                    Optional: true,
                                    MaxItems: 1,
                                    Elem: &schema.Resource{
                                    	Schema: map[string]*schema.Schema{
                                            "generate_icmp": {
                                                Type:     schema.TypeBool,
                                                Optional: true,
                                                Default: false,
                                            },
                                        },
                                    },
                                },
                                "name": {
                                    Type:     schema.TypeString,
                                    Required: true,
                                    ForceNew: true,
                                },
                                "nexthop": {
                                    Type:     schema.TypeList,
                                    Optional: true,
                                    MaxItems: 65536,
                                    Elem: &schema.Resource{
                                    	Schema: map[string]*schema.Schema{
                                            "admin_state": {
                                                Type:     schema.TypeString,
                                                Optional: true,
                                                Default: "enable",
                                            },
                                            "index": {
                                                Type:     schema.TypeInt,
                                                Required: true,
                                                ForceNew: true,
                                            },
                                            "ip_address": {
                                                Type:     schema.TypeString,
                                                Optional: true,
                                            },
                                            "pushed_mpls_label_stack": {
                                                Type:     schema.TypeString,
                                                Optional: true,
                                            },
                                            "resolve": {
                                                Type:     schema.TypeBool,
                                                Optional: true,
                                                Default: true,
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

func resourceNetworkInstanceInstanceNextHopGroupsCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Infof("Beginning Create: %s", resourceNetworkInstanceInstanceNextHopGroupsString(d))
	target := meta.(*Target)
	
	key := "next-hop-groups"
	
	
	hkey0 := d.Get("network_instance_id").(string)
    
	//hkey := d.Get("[network_instance_id]").(string)
	//p := fmt.Sprintf("fmt.Sprintf("/network-instance[name=%s]/next-hop-groups",hkey0)", hkey)
	p := fmt.Sprintf("/network-instance[name=%s]/next-hop-groups",hkey0)
	
	v := "next_hop_groups"
	
	
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
	return resourceNetworkInstanceInstanceNextHopGroupsRead(ctx, d, meta)
}

func resourceNetworkInstanceInstanceNextHopGroupsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Infof("Beginning Read: %s", resourceNetworkInstanceInstanceNextHopGroupsString(d))
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

func resourceNetworkInstanceInstanceNextHopGroupsUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Infof("Beginning Update: %s", resourceNetworkInstanceInstanceNextHopGroupsString(d))
	target := meta.(*Target)
	
	key := "next-hop-groups"
	
	
	hkey0 := d.Get("network_instance_id").(string)
    
	//hkey := d.Get("[network_instance_id]").(string)
	//p := fmt.Sprintf("fmt.Sprintf("/network-instance[name=%s]/next-hop-groups",hkey0)", hkey)
	p := fmt.Sprintf("/network-instance[name=%s]/next-hop-groups",hkey0)
	
	v := "next_hop_groups"
	
	
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
	return resourceNetworkInstanceInstanceNextHopGroupsRead(ctx, d, meta)
}

func resourceNetworkInstanceInstanceNextHopGroupsDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Debugf("Beginning delete: %s", resourceNetworkInstanceInstanceNextHopGroupsString(d))
	target := meta.(*Target)

	
	
	hkey0 := d.Get("network_instance_id").(string)
    
	//hkey := d.Get("[network_instance_id]").(string)
	
	//p := fmt.Sprintf("fmt.Sprintf("/network-instance[name=%s]/next-hop-groups",hkey0)", hkey)
	p := fmt.Sprintf("/network-instance[name=%s]/next-hop-groups",hkey0)
	
	
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
