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

// resourceAclCpmFilterIpv6FilterString function
func resourceAclCpmFilterIpv6FilterString(d resourceIDStringer) string {
	return resourceIDString(d, "acl_cpm_filter_ipv6_filter")
}

// resourceAclCpmFilterIpv6Filter function
func resourceAclCpmFilterIpv6Filter() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceAclCpmFilterIpv6FilterCreate,
		ReadContext:   resourceAclCpmFilterIpv6FilterRead,
		UpdateContext: resourceAclCpmFilterIpv6FilterUpdate,
		DeleteContext: resourceAclCpmFilterIpv6FilterDelete,

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
        "ipv6_filter": {
            Type:     schema.TypeList,
            Optional: true,
            MaxItems: 16,
            Elem: &schema.Resource{
            	Schema: map[string]*schema.Schema{
                    "entry": {
                        Type:     schema.TypeList,
                        Optional: true,
                        MaxItems: 16,
                        Elem: &schema.Resource{
                        	Schema: map[string]*schema.Schema{
                                "action": {
                                    Type:     schema.TypeList,
                                    Optional: true,
                                    MaxItems: 16,
                                    Elem: &schema.Resource{
                                    	Schema: map[string]*schema.Schema{
                                            "accept": {
                                                Type:     schema.TypeList,
                                                Optional: true,
                                                MaxItems: 16,
                                                Elem: &schema.Resource{
                                                	Schema: map[string]*schema.Schema{
                                                        "log": {
                                                            Type:     schema.TypeBool,
                                                            Optional: true,
                                                            Default: false,
                                                        },
                                                        "rate_limit": {
                                                            Type:     schema.TypeList,
                                                            Optional: true,
                                                            MaxItems: 16,
                                                            Elem: &schema.Resource{
                                                            	Schema: map[string]*schema.Schema{
                                                                    "distributed_policer": {
                                                                        Type:     schema.TypeString,
                                                                        Optional: true,
                                                                    },
                                                                    "system_cpu_policer": {
                                                                        Type:     schema.TypeString,
                                                                        Optional: true,
                                                                    },
                                                                },
                                                            },
                                                        },
                                                    },
                                                },
                                            },
                                            "drop": {
                                                Type:     schema.TypeList,
                                                Optional: true,
                                                MaxItems: 16,
                                                Elem: &schema.Resource{
                                                	Schema: map[string]*schema.Schema{
                                                        "log": {
                                                            Type:     schema.TypeBool,
                                                            Optional: true,
                                                            Default: false,
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
                                "match": {
                                    Type:     schema.TypeList,
                                    Optional: true,
                                    MaxItems: 16,
                                    Elem: &schema.Resource{
                                    	Schema: map[string]*schema.Schema{
                                            "destination_address": {
                                                Type:     schema.TypeString,
                                                Optional: true,
                                            },
                                            "destination_port": {
                                                Type:     schema.TypeList,
                                                Optional: true,
                                                MaxItems: 16,
                                                Elem: &schema.Resource{
                                                	Schema: map[string]*schema.Schema{
                                                        "operator": {
                                                            Type:     schema.TypeString,
                                                            Optional: true,
                                                        },
                                                        "range": {
                                                            Type:     schema.TypeList,
                                                            Optional: true,
                                                            MaxItems: 16,
                                                            Elem: &schema.Resource{
                                                            	Schema: map[string]*schema.Schema{
                                                                    "end": {
                                                                        Type:     schema.TypeString,
                                                                        Optional: true,
                                                                    },
                                                                    "start": {
                                                                        Type:     schema.TypeString,
                                                                        Optional: true,
                                                                    },
                                                                },
                                                            },
                                                        },
                                                        "value": {
                                                            Type:     schema.TypeString,
                                                            Optional: true,
                                                        },
                                                    },
                                                },
                                            },
                                            "icmp6": {
                                                Type:     schema.TypeList,
                                                Optional: true,
                                                MaxItems: 16,
                                                Elem: &schema.Resource{
                                                	Schema: map[string]*schema.Schema{
                                                        "code": {
                                                            Type:     schema.TypeInt,
                                                            Optional: true,
                                                        },
                                                        "type": {
                                                            Type:     schema.TypeString,
                                                            Optional: true,
                                                        },
                                                    },
                                                },
                                            },
                                            "next_header": {
                                                Type:     schema.TypeString,
                                                Optional: true,
                                            },
                                            "source_address": {
                                                Type:     schema.TypeString,
                                                Optional: true,
                                            },
                                            "source_port": {
                                                Type:     schema.TypeList,
                                                Optional: true,
                                                MaxItems: 16,
                                                Elem: &schema.Resource{
                                                	Schema: map[string]*schema.Schema{
                                                        "operator": {
                                                            Type:     schema.TypeString,
                                                            Optional: true,
                                                        },
                                                        "range": {
                                                            Type:     schema.TypeList,
                                                            Optional: true,
                                                            MaxItems: 16,
                                                            Elem: &schema.Resource{
                                                            	Schema: map[string]*schema.Schema{
                                                                    "end": {
                                                                        Type:     schema.TypeString,
                                                                        Optional: true,
                                                                    },
                                                                    "start": {
                                                                        Type:     schema.TypeString,
                                                                        Optional: true,
                                                                    },
                                                                },
                                                            },
                                                        },
                                                        "value": {
                                                            Type:     schema.TypeString,
                                                            Optional: true,
                                                        },
                                                    },
                                                },
                                            },
                                            "tcp_flags": {
                                                Type:     schema.TypeString,
                                                Optional: true,
                                            },
                                        },
                                    },
                                },
                                "sequence_id": {
                                    Type:     schema.TypeInt,
                                    Required: true,
                                    ForceNew: true,
                                },
                            },
                        },
                    },
                    "statistics_per_entry": {
                        Type:     schema.TypeBool,
                        Optional: true,
                    },
                },
            },
        },

        },
    }
}

func resourceAclCpmFilterIpv6FilterCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Infof("Beginning Create: %s", resourceAclCpmFilterIpv6FilterString(d))
	target := meta.(*Target)
	
	key := "ipv6_filter"

	p := "/acl/cpm-filter/ipv6-filter"
	v := "ipv6_filter"
	
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
	return resourceAclCpmFilterIpv6FilterRead(ctx, d, meta)
}

// func resourceAclCpmFilterIpv6FilterRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
// 	log.Infof("Beginning Read: %s", resourceAclCpmFilterIpv6FilterString(d))
// 	target := meta.(*Target)

// 	
// 	p := "/acl/cpm-filter/ipv6-filter"
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
func resourceAclCpmFilterIpv6FilterRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Infof("Beginning Read: %s", resourceAclCpmFilterIpv6FilterString(d))
	target := meta.(*Target)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	
	p := "/acl/cpm-filter/ipv6-filter"
	

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
			switch x := upd.Values["ipv6-filter"].(type) {
			case map[string]interface{}:
				
				data = append(data, x)
			}
			log.Debugf("get response: index: %d, data: %v", i, data)
			if err := d.Set("ipv6_filter", data); err != nil {
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

func resourceAclCpmFilterIpv6FilterUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Infof("Beginning Update: %s", resourceAclCpmFilterIpv6FilterString(d))
	target := meta.(*Target)
	
	key := "ipv6_filter"

	p := "/acl/cpm-filter/ipv6-filter"
	v := "ipv6_filter"
	

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
	return resourceAclCpmFilterIpv6FilterRead(ctx, d, meta)
}

func resourceAclCpmFilterIpv6FilterDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Debugf("Beginning delete: %s", resourceAclCpmFilterIpv6FilterString(d))
	target := meta.(*Target)

	
	p := "/acl/cpm-filter/ipv6-filter"
	
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
