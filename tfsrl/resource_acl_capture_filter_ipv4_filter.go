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

// resourceAclCaptureFilterIpv4FilterString function
func resourceAclCaptureFilterIpv4FilterString(d resourceIDStringer) string {
	return resourceIDString(d, "acl_capture_filter_ipv4_filter")
}

// resourceAclCaptureFilterIpv4Filter function
func resourceAclCaptureFilterIpv4Filter() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceAclCaptureFilterIpv4FilterCreate,
		ReadContext:   resourceAclCaptureFilterIpv4FilterRead,
		UpdateContext: resourceAclCaptureFilterIpv4FilterUpdate,
		DeleteContext: resourceAclCaptureFilterIpv4FilterDelete,

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
														Schema: map[string]*schema.Schema{},
													},
												},
												"copy": {
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{},
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
					},
				},
			},
		},
	}
}

func resourceAclCaptureFilterIpv4FilterCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Infof("Beginning Create: %s", resourceAclCaptureFilterIpv4FilterString(d))
	target := meta.(*Target)

	key := "ipv4-filter"

	p := "/acl/capture-filter/ipv4-filter"
	v := "ipv4-filter"

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
	return resourceAclCaptureFilterIpv4FilterRead(ctx, d, meta)
}

func resourceAclCaptureFilterIpv4FilterRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Infof("Beginning Read: %s", resourceAclCaptureFilterIpv4FilterString(d))
	target := meta.(*Target)

	p := "/acl/capture-filter/ipv4-filter"

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

func resourceAclCaptureFilterIpv4FilterUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Infof("Beginning Update: %s", resourceAclCaptureFilterIpv4FilterString(d))
	target := meta.(*Target)

	key := "ipv4-filter"

	p := "/acl/capture-filter/ipv4-filter"
	v := "ipv4-filter"

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
	return resourceAclCaptureFilterIpv4FilterRead(ctx, d, meta)
}

func resourceAclCaptureFilterIpv4FilterDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Debugf("Beginning delete: %s", resourceAclCaptureFilterIpv4FilterString(d))
	target := meta.(*Target)

	p := "/acl/capture-filter/ipv4-filter"

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
