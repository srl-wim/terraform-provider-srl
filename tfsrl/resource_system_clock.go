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
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	log "github.com/sirupsen/logrus"
)

// resourceSystemClockString function
func resourceSystemClockString(d resourceIDStringer) string {
	return resourceIDString(d, "system_clock")
}

// resourceSystemClock function
func resourceSystemClock() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceSystemClockCreate,
		ReadContext:   resourceSystemClockRead,
		UpdateContext: resourceSystemClockUpdate,
		DeleteContext: resourceSystemClockDelete,

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
        "clock": {
            Type:     schema.TypeList,
            Optional: true,
            MaxItems: 1,
            Elem: &schema.Resource{
            	Schema: map[string]*schema.Schema{
                    "timezone": {
                        Type:     schema.TypeString,
                        Optional: true,
                    },
                },
            },
        },

        },
    }
}

func resourceSystemClockCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Infof("Beginning Create: %s", resourceSystemClockString(d))
	target := meta.(*Target)
	
	key := "clock"

	p := "/system/clock"
	v := "clock"
	
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
	return resourceSystemClockRead(ctx, d, meta)
}

// func resourceSystemClockRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
// 	log.Infof("Beginning Read: %s", resourceSystemClockString(d))
// 	target := meta.(*Target)

// 	
// 	p := "/system/clock"
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
func resourceSystemClockRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Infof("Beginning Read: %s", resourceSystemClockString(d))
	target := meta.(*Target)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	
	p := "/system/clock"
	

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
			switch x := upd.Values["clock"].(type) {
			case map[string]interface{}:
				
				data = append(data, x)
			}
			log.Debugf("get response: index: %d, data: %v", i, data)
			if err := d.Set("clock", data); err != nil {
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

func resourceSystemClockUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Infof("Beginning Update: %s", resourceSystemClockString(d))
	target := meta.(*Target)
	
	key := "clock"

	p := "/system/clock"
	v := "clock"
	

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
	return resourceSystemClockRead(ctx, d, meta)
}

func resourceSystemClockDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Debugf("Beginning delete: %s", resourceSystemClockString(d))
	target := meta.(*Target)

	
	p := "/system/clock"
	
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
