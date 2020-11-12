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

// resourceSystemSnmpString function
func resourceSystemSnmpString(d resourceIDStringer) string {
	return resourceIDString(d, "system_snmp")
}

// resourceSystemSnmp function
func resourceSystemSnmp() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceSystemSnmpCreate,
		ReadContext:   resourceSystemSnmpRead,
		UpdateContext: resourceSystemSnmpUpdate,
		DeleteContext: resourceSystemSnmpDelete,

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
        "snmp": {
            Type:     schema.TypeList,
            Optional: true,
            MaxItems: 1,
            Elem: &schema.Resource{
            	Schema: map[string]*schema.Schema{
                    "community": {
                        Type:     schema.TypeString,
                        Optional: true,
                    },
                    "network_instance": {
                        Type:     schema.TypeList,
                        Optional: true,
                        MaxItems: 1,
                        Elem: &schema.Resource{
                        	Schema: map[string]*schema.Schema{
                                "admin_state": {
                                    Type:     schema.TypeString,
                                    Optional: true,
                                },
                                "name": {
                                    Type:     schema.TypeString,
                                    Required: true,
                                    ForceNew: true,
                                },
                                "source_address": {
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
    }
}

func resourceSystemSnmpCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Infof("Beginning Create: %s", resourceSystemSnmpString(d))
	target := meta.(*Target)
	
	key := "snmp"

	p := "/system/snmp"
	v := "snmp"
	
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
	return resourceSystemSnmpRead(ctx, d, meta)
}

func resourceSystemSnmpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Infof("Beginning Read: %s", resourceSystemSnmpString(d))
	target := meta.(*Target)

	
	p := "/system/snmp"
	
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

func resourceSystemSnmpUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Infof("Beginning Update: %s", resourceSystemSnmpString(d))
	target := meta.(*Target)
	
	key := "snmp"

	p := "/system/snmp"
	v := "snmp"
	

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
	return resourceSystemSnmpRead(ctx, d, meta)
}

func resourceSystemSnmpDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Debugf("Beginning delete: %s", resourceSystemSnmpString(d))
	target := meta.(*Target)

	
	p := "/system/snmp"
	
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
