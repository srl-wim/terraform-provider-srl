/*
Package tfsrl is a generated package which contains definitions
of structs which represent a Terraform resource. 

This package was generated by ygocodegen
using the following YANG input files:
	- yang/srl
Imported modules were sourced from:
	- yang/ietf
*/
package tfsrl

import (
	"context"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	log "github.com/sirupsen/logrus"
)

// resourceInterfacesString function
func resourceInterfacesString(d resourceIDStringer) string {
	return resourceIDString(d, "interfaces")
}

// resourceInterfaces function
func resourceInterfaces() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceInterfacesCreate,
		ReadContext:   resourceInterfacesRead,
		UpdateContext: resourceInterfacesUpdate,
		DeleteContext: resourceInterfacesDelete,

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
                    "description": {
                        Type:     schema.TypeString,
                        Optional: true,
                    },
                    "ethernet": {
                        Type:     schema.TypeList,
                        Optional: true,
                        MaxItems: 1,
                        Elem: &schema.Resource{
                        	Schema: map[string]*schema.Schema{
                                "flow_control": {
                                    Type:     schema.TypeList,
                                    Optional: true,
                                    MaxItems: 1,
                                    Elem: &schema.Resource{
                                    	Schema: map[string]*schema.Schema{
                                            "receive": {
                                                Type:     schema.TypeBool,
                                                Optional: true,
                                            },
                                        },
                                    },
                                },
                            },
                        },
                    },
                    "mtu": {
                        Type:     schema.TypeInt,
                        Optional: true,
                    },
                    "name": {
                        Type:     schema.TypeString,
                        Optional: true,
                    },
                    "qos": {
                        Type:     schema.TypeList,
                        Optional: true,
                        MaxItems: 1,
                        Elem: &schema.Resource{
                        	Schema: map[string]*schema.Schema{
                                "output": {
                                    Type:     schema.TypeList,
                                    Optional: true,
                                    MaxItems: 1,
                                    Elem: &schema.Resource{
                                    	Schema: map[string]*schema.Schema{
                                            "queue": {
                                                Type:     schema.TypeList,
                                                Optional: true,
                                                MaxItems: 8,
                                                Elem: &schema.Resource{
                                                	Schema: map[string]*schema.Schema{
                                                        "queue_id": {
                                                            Type:     schema.TypeString,
                                                            Optional: true,
                                                        },
                                                        "queue_scheduler": {
                                                            Type:     schema.TypeList,
                                                            Optional: true,
                                                            MaxItems: 1,
                                                            Elem: &schema.Resource{
                                                            	Schema: map[string]*schema.Schema{
                                                                    "strict_priority": {
                                                                        Type:     schema.TypeBool,
                                                                        Optional: true,
                                                                        Default: true,
                                                                    },
                                                                    "weight": {
                                                                        Type:     schema.TypeInt,
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
                    "transceiver": {
                        Type:     schema.TypeList,
                        Optional: true,
                        MaxItems: 1,
                        Elem: &schema.Resource{
                        	Schema: map[string]*schema.Schema{
                                "admin_state": {
                                    Type:     schema.TypeString,
                                    Optional: true,
                                },
                                "ddm_events": {
                                    Type:     schema.TypeBool,
                                    Optional: true,
                                },
                                "forward_error_correction": {
                                    Type:     schema.TypeString,
                                    Optional: true,
                                },
                            },
                        },
                    },
                    "vlan_tagging": {
                        Type:     schema.TypeBool,
                        Optional: true,
                    },
                },
            },
        },

        },
    }
}

func resourceInterfacesCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Infof("Beginning Create: %s", resourceInterfacesString(d))
	target := meta.(*Target)
	 
	rn := "interface"
	rk := "name"
	key, err := getResourceListKey(&rn, &rk, d)
	if err != nil {
		return diag.FromErr(err)
	}
	p := "/"
	v := ""
	
	req, err := target.CreateSetRequest(&p, &v, d)
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
	return resourceInterfacesRead(ctx, d, meta)
}

// func resourceInterfacesRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
// 	log.Infof("Beginning Read: %s", resourceInterfacesString(d))
// 	target := meta.(*Target)

// 	 
// 	p := fmt.Sprintf("/interface[name=%s]", d.Id())
// 	
// 	req, err := target.CreateGetRequest(&p, "CONFIG", d)
// 	if err != nil {
// 		return diag.FromErr(err)
// 	}
// 	response, err := target.Get(ctx, req)
// 	if err != nil {
// 		return diag.FromErr(err)
// 	}

// 	log.Debugf("Get Gnmi read response: %v", response)

// 	return nil
// }
func resourceInterfacesRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Infof("Beginning Read: %s", resourceInterfacesString(d))
	target := meta.(*Target)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	 
	//rn := "interface"
	rk := "name"
	key:= d.Id()

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
					
					if k == "sflow" {
                        delete(x, k)
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

func resourceInterfacesUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Infof("Beginning Update: %s", resourceInterfacesString(d))
	target := meta.(*Target)
	 
	rn := "interface"
	rk := "name"
	key, err := getResourceListKey(&rn, &rk, d)
	if err != nil {
		return diag.FromErr(err)
	}
	p := "/"
	v := ""
	

	req, err := target.CreateSetRequest(&p, &v, d)
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
	return resourceInterfacesRead(ctx, d, meta)
}

func resourceInterfacesDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Debugf("Beginning delete: %s", resourceInterfacesString(d))
	target := meta.(*Target)

	 
	p := fmt.Sprintf("/interface[name=%s]", d.Id())
	
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
