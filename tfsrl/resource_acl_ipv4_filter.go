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
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	log "github.com/sirupsen/logrus"
)

// resourceAclIpv4FilterString function
func resourceAclIpv4FilterString(d resourceIDStringer) string {
	return resourceIDString(d, "acl_ipv4_filter")
}

// resourceAclIpv4Filter function
func resourceAclIpv4Filter() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceAclIpv4FilterCreate,
		ReadContext:   resourceAclIpv4FilterRead,
		UpdateContext: resourceAclIpv4FilterUpdate,
		DeleteContext: resourceAclIpv4FilterDelete,

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
			"ipv4_filter": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"description": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"entry": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"action": {
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"accept": {
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"log": {
																Type:     schema.TypeBool,
																Optional: true,
																Default:  false,
															},
														},
													},
												},
												"drop": {
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"log": {
																Type:     schema.TypeBool,
																Optional: true,
																Default:  false,
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
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"destination_address": {
													Type:     schema.TypeString,
													Optional: true,
												},
												"destination_port": {
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"operator": {
																Type:     schema.TypeString,
																Optional: true,
															},
															"range": {
																Type:     schema.TypeList,
																Optional: true,
																MaxItems: 1,
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
												"first_fragment": {
													Type:     schema.TypeBool,
													Optional: true,
												},
												"fragment": {
													Type:     schema.TypeBool,
													Optional: true,
												},
												"icmp": {
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
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
												"protocol": {
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
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"operator": {
																Type:     schema.TypeString,
																Optional: true,
															},
															"range": {
																Type:     schema.TypeList,
																Optional: true,
																MaxItems: 1,
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
						"name": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"statistics_per_entry": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"subinterface_specific": {
							Type:     schema.TypeString,
							Optional: true,
							Default:  "disabled",
						},
					},
				},
			},
		},
	}
}

func resourceAclIpv4FilterCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Infof("Beginning Create: %s", resourceAclIpv4FilterString(d))
	target := meta.(*Target)

	rn := "ipv4_filter"
	rk := "name"
	key, err := getResourceListKey(&rn, &rk, d)
	if err != nil {
		return diag.FromErr(err)
	}
	p := "/acl/"
	v := ""

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
	return resourceAclIpv4FilterRead(ctx, d, meta)
}

func resourceAclIpv4FilterRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Infof("Beginning Read: %s", resourceAclIpv4FilterString(d))
	target := meta.(*Target)

	p := fmt.Sprintf("/acl/ipv4-filter[name=%s]", d.Id())

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

func resourceAclIpv4FilterUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Infof("Beginning Update: %s", resourceAclIpv4FilterString(d))
	target := meta.(*Target)

	rn := "ipv4_filter"
	rk := "name"
	key, err := getResourceListKey(&rn, &rk, d)
	if err != nil {
		return diag.FromErr(err)
	}
	p := "/acl/"
	v := ""

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
	return resourceAclIpv4FilterRead(ctx, d, meta)
}

func resourceAclIpv4FilterDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Debugf("Beginning delete: %s", resourceAclIpv4FilterString(d))
	target := meta.(*Target)

	p := fmt.Sprintf("/acl/ipv4-filter[name=%s]", d.Id())

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
