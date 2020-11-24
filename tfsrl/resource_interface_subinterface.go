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

// resourceInterfacesSubinterfaceString function
func resourceInterfacesSubinterfaceString(d resourceIDStringer) string {
	return resourceIDString(d, "interfaces_subinterface")
}

// resourceInterfacesSubinterface function
func resourceInterfacesSubinterface() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceInterfacesSubinterfaceCreate,
		ReadContext:   resourceInterfacesSubinterfaceRead,
		UpdateContext: resourceInterfacesSubinterfaceUpdate,
		DeleteContext: resourceInterfacesSubinterfaceDelete,

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
        "interface_id": {
            Type:     schema.TypeString,
            Required: true,
            ForceNew: true,
        },
        "subinterface": {
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
                    "bridge_table": {
                        Type:     schema.TypeList,
                        Optional: true,
                        MaxItems: 1,
                        Elem: &schema.Resource{
                        	Schema: map[string]*schema.Schema{
                                "discard_unknown_src_mac": {
                                    Type:     schema.TypeBool,
                                    Optional: true,
                                    Default: false,
                                },
                                "mac_duplication": {
                                    Type:     schema.TypeList,
                                    Optional: true,
                                    MaxItems: 1,
                                    Elem: &schema.Resource{
                                    	Schema: map[string]*schema.Schema{
                                            "action": {
                                                Type:     schema.TypeString,
                                                Optional: true,
                                                Default: "use-net-instance-action",
                                            },
                                        },
                                    },
                                },
                                "mac_learning": {
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
                                            "aging": {
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
                                                    },
                                                },
                                            },
                                        },
                                    },
                                },
                                "mac_limit": {
                                    Type:     schema.TypeList,
                                    Optional: true,
                                    MaxItems: 1,
                                    Elem: &schema.Resource{
                                    	Schema: map[string]*schema.Schema{
                                            "maximum_entries": {
                                                Type:     schema.TypeInt,
                                                Optional: true,
                                                Default: "250",
                                            },
                                            "warning_threshold_pct": {
                                                Type:     schema.TypeInt,
                                                Optional: true,
                                                Default: "95",
                                            },
                                        },
                                    },
                                },
                            },
                        },
                    },
                    "description": {
                        Type:     schema.TypeString,
                        Optional: true,
                    },
                    "index": {
                        Type:     schema.TypeInt,
                        Required: true,
                        ForceNew: true,
                    },
                    "ip_mtu": {
                        Type:     schema.TypeInt,
                        Optional: true,
                    },
                    "l2_mtu": {
                        Type:     schema.TypeInt,
                        Optional: true,
                    },
                    "qos": {
                        Type:     schema.TypeList,
                        Optional: true,
                        MaxItems: 1,
                        Elem: &schema.Resource{
                        	Schema: map[string]*schema.Schema{
                                "input": {
                                    Type:     schema.TypeList,
                                    Optional: true,
                                    MaxItems: 1,
                                    Elem: &schema.Resource{
                                    	Schema: map[string]*schema.Schema{
                                            "classifiers": {
                                                Type:     schema.TypeList,
                                                Optional: true,
                                                MaxItems: 1,
                                                Elem: &schema.Resource{
                                                	Schema: map[string]*schema.Schema{
                                                        "dscp": {
                                                            Type:     schema.TypeString,
                                                            Optional: true,
                                                        },
                                                        "ipv4_dscp": {
                                                            Type:     schema.TypeString,
                                                            Optional: true,
                                                        },
                                                        "ipv6_dscp": {
                                                            Type:     schema.TypeString,
                                                            Optional: true,
                                                        },
                                                        "mpls_traffic_class": {
                                                            Type:     schema.TypeString,
                                                            Optional: true,
                                                        },
                                                    },
                                                },
                                            },
                                        },
                                    },
                                },
                                "output": {
                                    Type:     schema.TypeList,
                                    Optional: true,
                                    MaxItems: 1,
                                    Elem: &schema.Resource{
                                    	Schema: map[string]*schema.Schema{
                                            "rewrite_rules": {
                                                Type:     schema.TypeList,
                                                Optional: true,
                                                MaxItems: 1,
                                                Elem: &schema.Resource{
                                                	Schema: map[string]*schema.Schema{
                                                        "dscp": {
                                                            Type:     schema.TypeString,
                                                            Optional: true,
                                                        },
                                                        "ipv4_dscp": {
                                                            Type:     schema.TypeString,
                                                            Optional: true,
                                                        },
                                                        "ipv6_dscp": {
                                                            Type:     schema.TypeString,
                                                            Optional: true,
                                                        },
                                                        "mpls_traffic_class": {
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
                        },
                    },
                    "type": {
                        Type:     schema.TypeString,
                        Optional: true,
                    },
                    "vlan": {
                        Type:     schema.TypeList,
                        Optional: true,
                        MaxItems: 1,
                        Elem: &schema.Resource{
                        	Schema: map[string]*schema.Schema{
                                "encap": {
                                    Type:     schema.TypeList,
                                    Optional: true,
                                    MaxItems: 1,
                                    Elem: &schema.Resource{
                                    	Schema: map[string]*schema.Schema{
                                            "single_tagged": {
                                                Type:     schema.TypeList,
                                                Optional: true,
                                                MaxItems: 1,
                                                Elem: &schema.Resource{
                                                	Schema: map[string]*schema.Schema{
                                                        "vlan_id": {
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
                        },
                    },
                },
            },
        },

        },
    }
}

func resourceInterfacesSubinterfaceCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Infof("Beginning Create: %s", resourceInterfacesSubinterfaceString(d))
	target := meta.(*Target)
	 
	rn := "subinterface"
	rk := "index"
	key, err := getResourceListKey(&rn, &rk, d)
	if err != nil {
		return diag.FromErr(err)
	}
	
	
	hkey0 := d.Get("interface_id").(string)
    
	//p := fmt.Sprintf("fmt.Sprintf("/interface[name=%s]/",hkey0)", hkey)
	p := fmt.Sprintf("/interface[name=%s]/",hkey0)
	
	v := ""
	
	
	hid := "interface_id"
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
	return resourceInterfacesSubinterfaceRead(ctx, d, meta)
}

func resourceInterfacesSubinterfaceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Infof("Beginning Read: %s", resourceInterfacesSubinterfaceString(d))
	target := meta.(*Target)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	
	
	hkey0 := d.Get("interface_id").(string)
    
	//hkey := d.Get("[interface_id]").(string)
	

	 
	//rn := "subinterface"
	rk := "index"
	key:= d.Id()

	
	//p := fmt.Sprintf("fmt.Sprintf("/interface[name=%s]/subinterface[index=%s]",hkey0, key)", hkey, key)
	p := fmt.Sprintf("/interface[name=%s]/subinterface[index=%s]",hkey0, key)
	
	

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
			switch x := upd.Values["subinterface"].(type) {
			case map[string]interface{}:
				for k, v := range x {
					log.Debugf("BEFORE KEY: %s, VALUE: %v", k, v)
					sk := strings.Split(k, ":")[len(strings.Split(k, ":"))-1]

					switch sk {
					
					case "acl":
						delete(x, k)
					
					case "ipv4":
						delete(x, k)
					
					case "ipv6":
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
				x[rk], err = strconv.ParseInt(key, 10, 32)
				// append the get resp to data
				
				data = append(data, x)
			}
			log.Debugf("get response: index: %d, data: %v", i, data)
			if err := d.Set("subinterface", data); err != nil {
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

func resourceInterfacesSubinterfaceUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Infof("Beginning Update: %s", resourceInterfacesSubinterfaceString(d))
	target := meta.(*Target)
	 
	rn := "subinterface"
	rk := "index"
	key, err := getResourceListKey(&rn, &rk, d)
	if err != nil {
		return diag.FromErr(err)
	}
	
	
	hkey0 := d.Get("interface_id").(string)
    
	//hkey := d.Get("[interface_id]").(string)
	//p := fmt.Sprintf("fmt.Sprintf("/interface[name=%s]/",hkey0)", hkey)
	p := fmt.Sprintf("/interface[name=%s]/",hkey0)
	
	v := ""
	
	
	hid := "interface_id"
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
	return resourceInterfacesSubinterfaceRead(ctx, d, meta)
}

func resourceInterfacesSubinterfaceDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Debugf("Beginning delete: %s", resourceInterfacesSubinterfaceString(d))
	target := meta.(*Target)

	
	
	hkey0 := d.Get("interface_id").(string)
    
	//hkey := d.Get("[interface_id]").(string)
	 
	//p := fmt.Sprintf("fmt.Sprintf("/interface[name=%s]/subinterface[index=%s]",hkey0, d.Id())", hkey, d.Id())
	p := fmt.Sprintf("/interface[name=%s]/subinterface[index=%s]",hkey0, d.Id())
	
	
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
