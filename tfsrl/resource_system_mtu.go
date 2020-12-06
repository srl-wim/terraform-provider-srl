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

	
	
	"strconv"
	
	
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	log "github.com/sirupsen/logrus"
)

// resourceSystemMtuString function
func resourceSystemMtuString(d resourceIDStringer) string {
	return resourceIDString(d, "system_mtu")
}

// resourceSystemMtu function
func resourceSystemMtu() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceSystemMtuCreate,
		ReadContext:   resourceSystemMtuRead,
		UpdateContext: resourceSystemMtuUpdate,
		DeleteContext: resourceSystemMtuDelete,

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
        "mtu": {
            Type:     schema.TypeList,
            Optional: true,
            MaxItems: 1,
            Elem: &schema.Resource{
            	Schema: map[string]*schema.Schema{
                    "default_ip_mtu": {
                        Type:     schema.TypeInt,
                        Optional: true,
                        Default: "1500",
                        ValidateFunc: validation.All(
                            validation.IntBetween(1280, 9486),
                        ),
                    },
                    "default_l2_mtu": {
                        Type:     schema.TypeInt,
                        Optional: true,
                        Default: "9232",
                        ValidateFunc: validation.All(
                            validation.IntBetween(1500, 9500),
                        ),
                    },
                    "default_port_mtu": {
                        Type:     schema.TypeInt,
                        Optional: true,
                        Default: "9232",
                        ValidateFunc: validation.All(
                            validation.IntBetween(1500, 9500),
                        ),
                    },
                    "min_path_mtu": {
                        Type:     schema.TypeInt,
                        Optional: true,
                        Default: "552",
                        ValidateFunc: validation.All(
                            validation.IntBetween(552, 9232),
                        ),
                    },
                },
            },
        },

        },
    }
}

func resourceSystemMtuCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Infof("Beginning Create: %s", resourceSystemMtuString(d))
	target := meta.(*Target)
	
	key := "mtu"
	
	p := "/system/mtu"
	
	v := "mtu"
	
	
	hid := make([]string, 0)
	//hid := ""
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
	return resourceSystemMtuRead(ctx, d, meta)
}

func resourceSystemMtuRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Infof("Beginning Read: %s", resourceSystemMtuString(d))
	target := meta.(*Target)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	

	
	
	p := "/system/mtu"
	
	

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
			switch x := upd.Values["mtu"].(type) {
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
				
				data = append(data, x)
			}
			log.Debugf("get response: index: %d, data: %v", i, data)
			if err := d.Set("mtu", data); err != nil {
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

func resourceSystemMtuUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Infof("Beginning Update: %s", resourceSystemMtuString(d))
	target := meta.(*Target)
	
	key := "mtu"
	
	p := "/system/mtu"
	
	v := "mtu"
	
	
	hid := make([]string, 0)
	//hid := ""
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
	return resourceSystemMtuRead(ctx, d, meta)
}

func resourceSystemMtuDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Debugf("Beginning delete: %s", resourceSystemMtuString(d))
	target := meta.(*Target)

	
	
	p := "/system/mtu"
	
	
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
