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
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	log "github.com/sirupsen/logrus"
)

// resourceSystemConfigurationString function
func resourceSystemConfigurationString(d resourceIDStringer) string {
	return resourceIDString(d, "system_configuration")
}

// resourceSystemConfiguration function
func resourceSystemConfiguration() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceSystemConfigurationCreate,
		ReadContext:   resourceSystemConfigurationRead,
		UpdateContext: resourceSystemConfigurationUpdate,
		DeleteContext: resourceSystemConfigurationDelete,

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
        "configuration": {
            Type:     schema.TypeList,
            Optional: true,
            MaxItems: 1,
            Elem: &schema.Resource{
            	Schema: map[string]*schema.Schema{
                    "max_checkpoints": {
                        Type:     schema.TypeInt,
                        Optional: true,
                        Default: "10",
                    },
                },
            },
        },

        },
    }
}

func resourceSystemConfigurationCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Infof("Beginning Create: %s", resourceSystemConfigurationString(d))
	target := meta.(*Target)
	
	key := "configuration"

	p := "/system/configuration"
	v := "configuration"
	
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
	return resourceSystemConfigurationRead(ctx, d, meta)
}

func resourceSystemConfigurationRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Infof("Beginning Read: %s", resourceSystemConfigurationString(d))
	target := meta.(*Target)

	
	p := "/system/configuration"
	
	req, err := target.CreateGetRequest(&p, d)
	if err != nil {
		return diag.FromErr(err)
	}
	response, err := target.Get(ctx, req)
	if err != nil {
		return diag.FromErr(err)
	}

	log.Debugf("Get Gnmi read response: %v", response)

	return nil
}

func resourceSystemConfigurationUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Infof("Beginning Update: %s", resourceSystemConfigurationString(d))
	target := meta.(*Target)
	
	key := "configuration"

	p := "/system/configuration"
	v := "configuration"
	

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
	return resourceSystemConfigurationRead(ctx, d, meta)
}

func resourceSystemConfigurationDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Debugf("Beginning delete: %s", resourceSystemConfigurationString(d))
	target := meta.(*Target)

	
	p := "/system/configuration"
	
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
