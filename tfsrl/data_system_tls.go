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

// dataSystemTlsString function
func dataSystemTlsString(d resourceIDStringer) string {
	return resourceIDString(d, "system_tls")
}

// dataSystemTls function
func dataSystemTls() *schema.Resource {
	return &schema.Resource{
		ReadContext:   dataSystemTlsRead,

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Timeouts: &schema.ResourceTimeout{
			Read:   schema.DefaultTimeout(5 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
        "tls": {
            Type:     schema.TypeList,
            Computed: true,
            Elem: &schema.Resource{
            	Schema: map[string]*schema.Schema{
                    "server_profile": {
                        Type:     schema.TypeList,
                        Computed: true,
                        Elem: &schema.Resource{
                        	Schema: map[string]*schema.Schema{
                                "authenticate_client": {
                                    Type:     schema.TypeBool,
                                    Computed: true,
                                },
                                "certificate": {
                                    Type:     schema.TypeString,
                                    Computed: true,
                                },
                                "cipher_list": {
                                    Type:     schema.TypeString,
                                    Computed: true,
                                },
                                "key": {
                                    Type:     schema.TypeString,
                                    Computed: true,
                                },
                                "name": {
                                    Type:     schema.TypeString,
                                    Required: true,
                                },
                                "trust_anchor": {
                                    Type:     schema.TypeString,
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

func dataSystemTlsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Infof("Beginning Read: %s", dataSystemTlsString(d))
	target := meta.(*Target)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	
	p := "/system/tls"
	

	req, err := target.CreateGetRequest(&p, d)
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
			switch x := upd.Values["tls"].(type) {
			case map[string]interface{}:
				
				data = append(data, x)
			}
			log.Debugf("get response: index: %d, data: %v", i, data)
			if err := d.Set("tls", data); err != nil {
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
