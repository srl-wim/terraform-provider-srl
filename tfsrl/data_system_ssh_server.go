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

// dataSystemSshServerString function
func dataSystemSshServerString(d resourceIDStringer) string {
	return resourceIDString(d, "system_ssh_server")
}

// dataSystemSshServer function
func dataSystemSshServer() *schema.Resource {
	return &schema.Resource{
		ReadContext:   dataSystemSshServerRead,

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Timeouts: &schema.ResourceTimeout{
			Read:   schema.DefaultTimeout(5 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
        "ssh_server": {
            Type:     schema.TypeList,
            Computed: true,
            Elem: &schema.Resource{
            	Schema: map[string]*schema.Schema{
                    "network_instance": {
                        Type:     schema.TypeList,
                        Computed: true,
                        Elem: &schema.Resource{
                        	Schema: map[string]*schema.Schema{
                                "admin_state": {
                                    Type:     schema.TypeString,
                                    Computed: true,
                                },
                                "name": {
                                    Type:     schema.TypeString,
                                    Required: true,
                                },
                                "rate_limit": {
                                    Type:     schema.TypeInt,
                                    Computed: true,
                                },
                                "source_address": {
                                    Type:     schema.TypeString,
                                    Computed: true,
                                },
                                "timeout": {
                                    Type:     schema.TypeInt,
                                    Computed: true,
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

func dataSystemSshServerRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Infof("Beginning Read: %s", dataSystemSshServerString(d))
	target := meta.(*Target)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	
	p := "/system/ssh-server"
	

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
			switch x := upd.Values["ssh-server"].(type) {
			case map[string]interface{}:
				
				data = append(data, x)
			}
			log.Debugf("get response: index: %d, data: %v", i, data)
			if err := d.Set("ssh_server", data); err != nil {
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