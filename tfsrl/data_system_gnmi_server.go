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

// dataSystemGnmiServerString function
func dataSystemGnmiServerString(d resourceIDStringer) string {
	return resourceIDString(d, "system_gnmi_server")
}

// dataSystemGnmiServer function
func dataSystemGnmiServer() *schema.Resource {
	return &schema.Resource{
		ReadContext:   dataSystemGnmiServerRead,

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Timeouts: &schema.ResourceTimeout{
			Read:   schema.DefaultTimeout(5 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
        "gnmi_server": {
            Type:     schema.TypeList,
            Computed: true,
            Elem: &schema.Resource{
            	Schema: map[string]*schema.Schema{
                    "admin_state": {
                        Type:     schema.TypeString,
                        Computed: true,
                    },
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
                                "oper_state": {
                                    Type:     schema.TypeString,
                                    Computed: true,
                                },
                                "port": {
                                    Type:     schema.TypeInt,
                                    Computed: true,
                                },
                                "source_address": {
                                    Type:     schema.TypeString,
                                    Computed: true,
                                },
                                "tls_profile": {
                                    Type:     schema.TypeString,
                                    Computed: true,
                                },
                                "use_authentication": {
                                    Type:     schema.TypeBool,
                                    Computed: true,
                                },
                            },
                        },
                    },
                    "rate_limit": {
                        Type:     schema.TypeInt,
                        Computed: true,
                    },
                    "session_limit": {
                        Type:     schema.TypeInt,
                        Computed: true,
                    },
                    "subscription": {
                        Type:     schema.TypeList,
                        Computed: true,
                        Elem: &schema.Resource{
                        	Schema: map[string]*schema.Schema{
                                "id": {
                                    Type:     schema.TypeInt,
                                    Required: true,
                                },
                                "mode": {
                                    Type:     schema.TypeString,
                                    Computed: true,
                                },
                                "paths": {
                                    Type:     schema.TypeString,
                                    Computed: true,
                                },
                                "remote_host": {
                                    Type:     schema.TypeString,
                                    Computed: true,
                                },
                                "remote_port": {
                                    Type:     schema.TypeInt,
                                    Computed: true,
                                },
                                "sample_interval": {
                                    Type:     schema.TypeInt,
                                    Computed: true,
                                },
                                "start_time": {
                                    Type:     schema.TypeString,
                                    Computed: true,
                                },
                                "user": {
                                    Type:     schema.TypeString,
                                    Computed: true,
                                },
                                "user_agent": {
                                    Type:     schema.TypeString,
                                    Computed: true,
                                },
                            },
                        },
                    },
                    "timeout": {
                        Type:     schema.TypeInt,
                        Computed: true,
                    },
                    "trace_options": {
                        Type:     schema.TypeString,
                        Computed: true,
                    },
                    "unix_socket": {
                        Type:     schema.TypeList,
                        Computed: true,
                        Elem: &schema.Resource{
                        	Schema: map[string]*schema.Schema{
                                "admin_state": {
                                    Type:     schema.TypeString,
                                    Computed: true,
                                },
                                "oper_state": {
                                    Type:     schema.TypeString,
                                    Computed: true,
                                },
                                "socket_path": {
                                    Type:     schema.TypeString,
                                    Computed: true,
                                },
                                "tls_profile": {
                                    Type:     schema.TypeString,
                                    Computed: true,
                                },
                                "use_authentication": {
                                    Type:     schema.TypeBool,
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

func dataSystemGnmiServerRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Infof("Beginning Read: %s", dataSystemGnmiServerString(d))
	target := meta.(*Target)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	
	p := "/system/gnmi-server"
	

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
			switch x := upd.Values["gnmi-server"].(type) {
			case map[string]interface{}:
				
				data = append(data, x)
			}
			log.Debugf("get response: index: %d, data: %v", i, data)
			if err := d.Set("gnmi_server", data); err != nil {
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
