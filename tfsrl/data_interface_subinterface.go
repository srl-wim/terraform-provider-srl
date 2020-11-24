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
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	log "github.com/sirupsen/logrus"
)

// dataInterfacesSubinterfaceString function
func dataInterfacesSubinterfaceString(d resourceIDStringer) string {
	return resourceIDString(d, "interfaces_subinterface")
}

// dataInterfacesSubinterface function
func dataInterfacesSubinterface() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataInterfacesSubinterfaceRead,

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Timeouts: &schema.ResourceTimeout{
			Read: schema.DefaultTimeout(5 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"interface_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"subinterface": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"admin_state": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"bridge_table": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"discard_unknown_src_mac": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"mac_duplication": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"action": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"mac_learning": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"admin_state": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"aging": {
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"admin_state": {
																Type:     schema.TypeString,
																Computed: true,
															},
														},
													},
												},
											},
										},
									},
									"mac_limit": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"maximum_entries": {
													Type:     schema.TypeInt,
													Computed: true,
												},
												"warning_threshold_pct": {
													Type:     schema.TypeInt,
													Computed: true,
												},
											},
										},
									},
								},
							},
						},
						"description": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"index": {
							Type:     schema.TypeInt,
							Required: true,
						},
						"ip_mtu": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"l2_mtu": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"qos": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"input": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"classifiers": {
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"dscp": {
																Type:     schema.TypeString,
																Computed: true,
															},
															"ipv4_dscp": {
																Type:     schema.TypeString,
																Computed: true,
															},
															"ipv6_dscp": {
																Type:     schema.TypeString,
																Computed: true,
															},
															"mpls_traffic_class": {
																Type:     schema.TypeString,
																Computed: true,
															},
														},
													},
												},
											},
										},
									},
									"output": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"rewrite_rules": {
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"dscp": {
																Type:     schema.TypeString,
																Computed: true,
															},
															"ipv4_dscp": {
																Type:     schema.TypeString,
																Computed: true,
															},
															"ipv6_dscp": {
																Type:     schema.TypeString,
																Computed: true,
															},
															"mpls_traffic_class": {
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
							},
						},
						"type": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataInterfacesSubinterfaceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Infof("Beginning Read: %s", dataInterfacesSubinterfaceString(d))
	target := meta.(*Target)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	hkey0 := d.Get("interface_id").(string)

	//hkey := d.Get("[interface_id]").(string)

	rn := "subinterface"
	rk := "index"
	key, err := getResourceListKey(&rn, &rk, d)

	//p := fmt.Sprintf("fmt.Sprintf("/interface[name=%s]/subinterface[index=%s]",hkey0, key)", hkey, key)
	p := fmt.Sprintf("/interface[name=%s]/subinterface[index=%s]", hkey0, key)

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
			switch x := upd.Values["subinterface"].(type) {
			case map[string]interface{}:
				for k, v := range x {
					log.Debugf("BEFORE KEY: %s, VALUE: %v", k, v)
					sk := strings.Split(k, ":")[len(strings.Split(k, ":"))-1]

					switch sk {

					case "acl":
						delete(x, k)

					case "ipv4":
						delete(x, k)

					case "ipv6":
						delete(x, k)

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
				x[rk], err = strconv.ParseInt(key, 10, 32)
				// append the get resp to data

				data = append(data, x)
			}
			log.Debugf("get response: index: %d, data: %v", i, data)
			if err := d.Set("subinterface", data); err != nil {
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
