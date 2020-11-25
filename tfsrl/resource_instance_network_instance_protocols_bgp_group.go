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

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	log "github.com/sirupsen/logrus"
)

// resourceNetworkInstanceInstanceProtocolsBgpGroupString function
func resourceNetworkInstanceInstanceProtocolsBgpGroupString(d resourceIDStringer) string {
	return resourceIDString(d, "network_instance_instance_protocols_bgp_group")
}

// resourceNetworkInstanceInstanceProtocolsBgpGroup function
func resourceNetworkInstanceInstanceProtocolsBgpGroup() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceNetworkInstanceInstanceProtocolsBgpGroupCreate,
		ReadContext:   resourceNetworkInstanceInstanceProtocolsBgpGroupRead,
		UpdateContext: resourceNetworkInstanceInstanceProtocolsBgpGroupUpdate,
		DeleteContext: resourceNetworkInstanceInstanceProtocolsBgpGroupDelete,

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
        "bgp_id": {
            Type:     schema.TypeString,
            Required: true,
            ForceNew: true,
        },
        "group": {
            Type:     schema.TypeList,
            Optional: true,
            MaxItems: 1,
            Elem: &schema.Resource{
            	Schema: map[string]*schema.Schema{
                    "admin_state": {
                        Type:     schema.TypeString,
                        Optional: true,
                        Default: "enable",
                        ValidateFunc: validation.All(
                            validation.StringInSlice([]string{
                                "disable",
                                "enable",
                            }, false),
                        ),
                    },
                    "description": {
                        Type:     schema.TypeString,
                        Optional: true,
                        ValidateFunc: validation.All(
                            validation.StringLenBetween(1, 255),
                        ),
                    },
                    "group_name": {
                        Type:     schema.TypeString,
                        Required: true,
                        ForceNew: true,
                        ValidateFunc: validation.All(
                            validation.StringLenBetween(1, 255),
                        ),
                    },
                    "local_as": {
                        Type:     schema.TypeList,
                        Optional: true,
                        MaxItems: 4294967295,
                        Elem: &schema.Resource{
                        	Schema: map[string]*schema.Schema{
                                "as_number": {
                                    Type:     schema.TypeInt,
                                    Required: true,
                                    ForceNew: true,
                                    ValidateFunc: validation.All(
                                        validation.IntBetween(1, 4294967295),
                                    ),
                                },
                                "prepend_global_as": {
                                    Type:     schema.TypeBool,
                                    Optional: true,
                                    Default: true,
                                },
                                "prepend_local_as": {
                                    Type:     schema.TypeBool,
                                    Optional: true,
                                    Default: true,
                                },
                            },
                        },
                    },
                    "local_preference": {
                        Type:     schema.TypeInt,
                        Optional: true,
                        ValidateFunc: validation.All(
                            validation.IntBetween(0, 4294967295),
                        ),
                    },
                    "next_hop_self": {
                        Type:     schema.TypeBool,
                        Optional: true,
                        Default: false,
                    },
                    "peer_as": {
                        Type:     schema.TypeInt,
                        Optional: true,
                        ValidateFunc: validation.All(
                            validation.IntBetween(1, 4294967295),
                        ),
                    },
                },
            },
        },

        },
    }
}

func resourceNetworkInstanceInstanceProtocolsBgpGroupCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Infof("Beginning Create: %s", resourceNetworkInstanceInstanceProtocolsBgpGroupString(d))
	target := meta.(*Target)
	 
	rn := "group"
	rk := "group_name"
	key, err := getResourceListKey(&rn, &rk, d)
	if err != nil {
		return diag.FromErr(err)
	}
	
	
	hkey0 := d.Get("network_instance_id").(string)
    
	//p := fmt.Sprintf("fmt.Sprintf("/network-instance[name=%s]/protocols/bgp/",hkey0)", hkey)
	p := fmt.Sprintf("/network-instance[name=%s]/protocols/bgp/",hkey0)
	
	v := ""
	
	
	hid := make([]string, 0)
	
    hid = append(hid, "network_instance_id")
	
    hid = append(hid, "bgp_id")
	
	//hid = append(hid, "bgp_id")
	//hid := "bgp_id"
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
	return resourceNetworkInstanceInstanceProtocolsBgpGroupRead(ctx, d, meta)
}

func resourceNetworkInstanceInstanceProtocolsBgpGroupRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Infof("Beginning Read: %s", resourceNetworkInstanceInstanceProtocolsBgpGroupString(d))
	target := meta.(*Target)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	
	
	hkey0 := d.Get("network_instance_id").(string)
    
	//hkey := d.Get("[network_instance_id]").(string)
	

	 
	//rn := "group"
	rk := "group_name"
	key:= d.Id()

	
	//p := fmt.Sprintf("fmt.Sprintf("/network-instance[name=%s]/protocols/bgp/group[group-name=%s]",hkey0, key)", hkey, key)
	p := fmt.Sprintf("/network-instance[name=%s]/protocols/bgp/group[group-name=%s]",hkey0, key)
	
	

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
			switch x := upd.Values["group"].(type) {
			case map[string]interface{}:
				for k, v := range x {
					log.Debugf("BEFORE KEY: %s, VALUE: %v", k, v)
					sk := strings.Split(k, ":")[len(strings.Split(k, ":"))-1]

					switch sk {
					
					case "import_policy":
						delete(x, k)
					
					case "export_policy":
						delete(x, k)
					
					case "dynamic_neighbors":
						delete(x, k)
					
					case "transport":
						delete(x, k)
					
					case "ipv4_unicast":
						delete(x, k)
					
					case "ipv6_unicast":
						delete(x, k)
					
					case "failure_detection":
						delete(x, k)
					
					case "send_community":
						delete(x, k)
					
					case "route_reflector":
						delete(x, k)
					
					case "as_path_options":
						delete(x, k)
					
					case "ebgp_default_policy":
						delete(x, k)
					
					case "route_advertisement":
						delete(x, k)
					
					case "authentication":
						delete(x, k)
					
					case "convergence":
						delete(x, k)
					
					case "graceful_restart":
						delete(x, k)
					
					case "trace_options":
						delete(x, k)
					
					case "preference":
						delete(x, k)
					
					case "timers":
						delete(x, k)
					
					case "send_default_route":
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
			if err := d.Set("group", data); err != nil {
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

func resourceNetworkInstanceInstanceProtocolsBgpGroupUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Infof("Beginning Update: %s", resourceNetworkInstanceInstanceProtocolsBgpGroupString(d))
	target := meta.(*Target)
	 
	rn := "group"
	rk := "group_name"
	key, err := getResourceListKey(&rn, &rk, d)
	if err != nil {
		return diag.FromErr(err)
	}
	
	
	hkey0 := d.Get("network_instance_id").(string)
    
	//hkey := d.Get("[network_instance_id]").(string)
	//p := fmt.Sprintf("fmt.Sprintf("/network-instance[name=%s]/protocols/bgp/",hkey0)", hkey)
	p := fmt.Sprintf("/network-instance[name=%s]/protocols/bgp/",hkey0)
	
	v := ""
	
	
	hid := make([]string, 0)
	
    hid = append(hid, "network_instance_id")
	
    hid = append(hid, "bgp_id")
	
	//hid = append(hid, "bgp_id")
	//hid := "bgp_id"
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
	return resourceNetworkInstanceInstanceProtocolsBgpGroupRead(ctx, d, meta)
}

func resourceNetworkInstanceInstanceProtocolsBgpGroupDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Debugf("Beginning delete: %s", resourceNetworkInstanceInstanceProtocolsBgpGroupString(d))
	target := meta.(*Target)

	
	
	hkey0 := d.Get("network_instance_id").(string)
    
	//hkey := d.Get("[network_instance_id]").(string)
	 
	//p := fmt.Sprintf("fmt.Sprintf("/network-instance[name=%s]/protocols/bgp/group[group-name=%s]",hkey0, d.Id())", hkey, d.Id())
	p := fmt.Sprintf("/network-instance[name=%s]/protocols/bgp/group[group-name=%s]",hkey0, d.Id())
	
	
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
