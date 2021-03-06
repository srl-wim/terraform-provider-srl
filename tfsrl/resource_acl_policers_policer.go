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
	"regexp"
	
	
	
	"fmt"
	
	
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	log "github.com/sirupsen/logrus"
)

// resourceAclPolicersPolicerString function
func resourceAclPolicersPolicerString(d resourceIDStringer) string {
	return resourceIDString(d, "acl_policers_policer")
}

// resourceAclPolicersPolicer function
func resourceAclPolicersPolicer() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceAclPolicersPolicerCreate,
		ReadContext:   resourceAclPolicersPolicerRead,
		UpdateContext: resourceAclPolicersPolicerUpdate,
		DeleteContext: resourceAclPolicersPolicerDelete,

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
        "policer": {
            Type:     schema.TypeList,
            Optional: true,
            MaxItems: 1,
            Elem: &schema.Resource{
            	Schema: map[string]*schema.Schema{
                    "entry_specific": {
                        Type:     schema.TypeBool,
                        Optional: true,
                        Default: false,
                    },
                    "max_burst": {
                        Type:     schema.TypeInt,
                        Optional: true,
                        ValidateFunc: validation.All(
                            validation.IntBetween(1, 125000000),
                        ),
                    },
                    "name": {
                        Type:     schema.TypeString,
                        Required: true,
                        ForceNew: true,
                        ValidateFunc: validation.All(
                            validation.StringLenBetween(1, 255),
                            validation.StringMatch(regexp.MustCompile("[A-Za-z0-9 !@#$%!^(MISSING)&()|+=`~.,'/_:;?-]*"), "must match regex"),
                        ),
                    },
                    "peak_rate": {
                        Type:     schema.TypeInt,
                        Optional: true,
                        ValidateFunc: validation.All(
                            validation.IntBetween(1, 1000000),
                        ),
                    },
                },
            },
        },

        },
    }
}

func resourceAclPolicersPolicerCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Infof("Beginning Create: %s", resourceAclPolicersPolicerString(d))
	target := meta.(*Target)
	 
	rn := "policer"
	rk := "name"
	key, err := getResourceListKey(&rn, &rk, d)
	if err != nil {
		return diag.FromErr(err)
	}
	
	p := "/acl/policers/"
	
	v := ""
	
	
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
	return resourceAclPolicersPolicerRead(ctx, d, meta)
}

func resourceAclPolicersPolicerRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Infof("Beginning Read: %s", resourceAclPolicersPolicerString(d))
	target := meta.(*Target)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	

	 
	//rn := "policer"
	rk := "name"
	key:= d.Id()

	
	//p := fmt.Sprintf("fmt.Sprintf("/acl/policers/policer[name=%s]", key)", key)
	p := fmt.Sprintf("/acl/policers/policer[name=%s]", key)
	
	

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
			switch x := upd.Values["policer"].(type) {
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
				 
				// add key to the get resp data since it is not returned in the gnmi data
				x[rk] = key
				// append the get resp to data
				
				data = append(data, x)
			}
			log.Debugf("get response: index: %d, data: %v", i, data)
			if err := d.Set("policer", data); err != nil {
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

func resourceAclPolicersPolicerUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Infof("Beginning Update: %s", resourceAclPolicersPolicerString(d))
	target := meta.(*Target)
	 
	rn := "policer"
	rk := "name"
	key, err := getResourceListKey(&rn, &rk, d)
	if err != nil {
		return diag.FromErr(err)
	}
	
	p := "/acl/policers/"
	
	v := ""
	
	
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
	return resourceAclPolicersPolicerRead(ctx, d, meta)
}

func resourceAclPolicersPolicerDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Debugf("Beginning delete: %s", resourceAclPolicersPolicerString(d))
	target := meta.(*Target)

	
	 
	//p := fmt.Sprintf("fmt.Sprintf("/acl/policers/policer[name=%s]", d.Id())", d.Id())
	p := fmt.Sprintf("/acl/policers/policer[name=%s]", d.Id())
	
	
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
