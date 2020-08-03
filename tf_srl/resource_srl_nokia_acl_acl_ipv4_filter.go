package tf_srl

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSrlNokiaAclAclIpv4FilterString(d resourceIDStringer) string {
	return resourceIDString(d, "srl_nokia_acl_acl_ipv4_filter")
}

func resourceSrlNokiaAclAclIpv4Filter() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				//ValidateFunc: validation.IntBetween(1, 65535),
			},
			"description": {
				Type:     schema.TypeString,
				Required: false,
				//ValidateFunc: validation.IntBetween(1, 65535),
			},
			"statisticsPerEntry": {
				Type:     schema.TypeBool,
				Optional: true,
				//Default:  true,
			},
			"subinterfaceSpecific": {
				Type:     schema.TypeString,
				Optional: true,
				//Default:  true,
				//?? HOW TO DO ENUM TYPE in teraform
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
									"destinationAddress": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"destinationPort": {
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"operator": {
													Type:     schema.TypeString,
													Optional: true,
													//Default:  true,
													//?? HOW TO DO ENUM TYPE in teraform
												},
												"range": {
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"end": {
																Type:     schema.TypeInt,
																Optional: true,
															},
															"start": {
																Type:     schema.TypeInt,
																Optional: true,
															},
														},
													},
												},
												"value": {
													Type:     schema.TypeString,
													Optional: true,
													//Default:  true,
													//?? HOW TO DO ENUM TYPE in teraform
												},
											},
										},
									},
									"firstFragment": {
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
												"Code": {
													Type:     schema.TypeInt,
													Optional: true,
												},
												"Type": {
													Type:     schema.TypeString,
													Optional: true,
													//Default:  true,
													//?? HOW TO DO ENUM/UNION TYPE in teraform
												},
											},
										},
									},
									"protocol": {
										Type:     schema.TypeString,
										Optional: true,
										//Default:  true,
										//?? HOW TO DO ENUM TYPE in teraform
									},
									"sourceAddress": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"sourcePort": {
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"operator": {
													Type:     schema.TypeString,
													Optional: true,
													//Default:  true,
													//?? HOW TO DO ENUM TYPE in teraform
												},
												"range": {
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"end": {
																Type:     schema.TypeInt,
																Optional: true,
															},
															"start": {
																Type:     schema.TypeInt,
																Optional: true,
															},
														},
													},
												},
												"value": {
													Type:     schema.TypeString,
													Optional: true,
													//Default:  true,
													//?? HOW TO DO ENUM TYPE in teraform
												},
											},
										},
									},
									"tcpFlags": {
										Type:     schema.TypeString,
										Optional: true,
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
