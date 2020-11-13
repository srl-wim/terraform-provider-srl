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

// resourceSystemSshServerString function
func resourceSystemSshServerString(d resourceIDStringer) string {
	return resourceIDString(d, "system_ssh_server")
}

// resourceSystemSshServer function
func resourceSystemSshServer() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceSystemSshServerCreate,
		ReadContext:   resourceSystemSshServerRead,
		UpdateContext: resourceSystemSshServerUpdate,
		DeleteContext: resourceSystemSshServerDelete,

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
        "ssh_server": {
            Type:     schema.TypeList,
            Optional: true,
            MaxItems: 1,
            Elem: &schema.Resource{
            	Schema: map[string]*schema.Schema{
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
                                "rate_limit": {
                                    Type:     schema.TypeInt,
                                    Optional: true,
                                    Default: "20",
                                },
                                "source_address": {
                                    Type:     schema.TypeString,
                                    Optional: true,
                                },
                                "timeout": {
                                    Type:     schema.TypeInt,
                                    Optional: true,
                                    Default: "0",
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

func resourceSystemSshServerCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Infof("Beginning Create: %s", resourceSystemSshServerString(d))
	target := meta.(*Target)
	
	key := "ssh_server"

	p := "/system/ssh-server"
	v := "ssh_server"
	
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
	return resourceSystemSshServerRead(ctx, d, meta)
}

func resourceSystemSshServerRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Infof("Beginning Read: %s", resourceSystemSshServerString(d))
	target := meta.(*Target)

	
	p := "/system/ssh-server"
	
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

func resourceSystemSshServerUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Infof("Beginning Update: %s", resourceSystemSshServerString(d))
	target := meta.(*Target)
	
	key := "ssh_server"

	p := "/system/ssh-server"
	v := "ssh_server"
	

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
	return resourceSystemSshServerRead(ctx, d, meta)
}

func resourceSystemSshServerDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Debugf("Beginning delete: %s", resourceSystemSshServerString(d))
	target := meta.(*Target)

	
	p := "/system/ssh-server"
	
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