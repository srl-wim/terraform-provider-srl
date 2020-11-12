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

// resourceAclCaptureFilterIpv6FilterString function
func resourceAclCaptureFilterIpv6FilterString(d resourceIDStringer) string {
	return resourceIDString(d, "acl_capture_filter_ipv6_filter")
}

// resourceAclCaptureFilterIpv6Filter function
func resourceAclCaptureFilterIpv6Filter() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceAclCaptureFilterIpv6FilterCreate,
		ReadContext:   resourceAclCaptureFilterIpv6FilterRead,
		UpdateContext: resourceAclCaptureFilterIpv6FilterUpdate,
		DeleteContext: resourceAclCaptureFilterIpv6FilterDelete,

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
												"icmp6": {
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

func resourceAclCaptureFilterIpv6FilterCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Infof("Beginning Create: %s", resourceAclCaptureFilterIpv6FilterString(d))
	target := meta.(*Target)

	key := "ipv6-filter"

	p := "/acl/capture-filter/ipv6-filter"
	v := "ipv6-filter"

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
	return resourceAclCaptureFilterIpv6FilterRead(ctx, d, meta)
}

func resourceAclCaptureFilterIpv6FilterRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Infof("Beginning Read: %s", resourceAclCaptureFilterIpv6FilterString(d))
	target := meta.(*Target)

	p := "/acl/capture-filter/ipv6-filter"

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

func resourceAclCaptureFilterIpv6FilterUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Infof("Beginning Update: %s", resourceAclCaptureFilterIpv6FilterString(d))
	target := meta.(*Target)

	key := "ipv6-filter"

	p := "/acl/capture-filter/ipv6-filter"
	v := "ipv6-filter"

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
	return resourceAclCaptureFilterIpv6FilterRead(ctx, d, meta)
}

func resourceAclCaptureFilterIpv6FilterDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Debugf("Beginning delete: %s", resourceAclCaptureFilterIpv6FilterString(d))
	target := meta.(*Target)

	p := "/acl/capture-filter/ipv6-filter"

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
