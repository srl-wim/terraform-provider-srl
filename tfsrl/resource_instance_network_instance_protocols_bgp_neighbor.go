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
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	log "github.com/sirupsen/logrus"
)

// resourceNetworkInstanceInstanceProtocolsBgpNeighborString function
func resourceNetworkInstanceInstanceProtocolsBgpNeighborString(d resourceIDStringer) string {
	return resourceIDString(d, "network_instance_instance_protocols_bgp_neighbor")
}

// resourceNetworkInstanceInstanceProtocolsBgpNeighbor function
func resourceNetworkInstanceInstanceProtocolsBgpNeighbor() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceNetworkInstanceInstanceProtocolsBgpNeighborCreate,
		ReadContext:   resourceNetworkInstanceInstanceProtocolsBgpNeighborRead,
		UpdateContext: resourceNetworkInstanceInstanceProtocolsBgpNeighborUpdate,
		DeleteContext: resourceNetworkInstanceInstanceProtocolsBgpNeighborDelete,

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
			"bgp_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"neighbor": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"admin_state": {
							Type:     schema.TypeString,
							Optional: true,
							Default:  "enable",
						},
						"as_path_options": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"allow_own_as": {
										Type:     schema.TypeInt,
										Optional: true,
									},
									"remove_private_as": {
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"ignore_peer_as": {
													Type:     schema.TypeBool,
													Optional: true,
													Default:  false,
												},
												"leading_only": {
													Type:     schema.TypeBool,
													Optional: true,
													Default:  false,
												},
												"mode": {
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},
									"replace_peer_as": {
										Type:     schema.TypeBool,
										Optional: true,
									},
								},
							},
						},
						"authentication": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"keychain": {
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
						"description": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"export_policy": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"failure_detection": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"enable_bfd": {
										Type:     schema.TypeBool,
										Optional: true,
									},
									"fast_failover": {
										Type:     schema.TypeBool,
										Optional: true,
									},
								},
							},
						},
						"graceful_restart": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"admin_state": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"stale_routes_time": {
										Type:     schema.TypeInt,
										Optional: true,
									},
								},
							},
						},
						"import_policy": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"ipv4_unicast": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"admin_state": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"advertise_ipv6_next_hops": {
										Type:     schema.TypeBool,
										Optional: true,
									},
									"prefix_limit": {
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"max_received_routes": {
													Type:     schema.TypeInt,
													Optional: true,
												},
												"warning_threshold_pct": {
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},
									"receive_ipv6_next_hops": {
										Type:     schema.TypeBool,
										Optional: true,
									},
								},
							},
						},
						"ipv6_unicast": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"admin_state": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"prefix_limit": {
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"max_received_routes": {
													Type:     schema.TypeInt,
													Optional: true,
												},
												"warning_threshold_pct": {
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},
								},
							},
						},
						"local_as": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 4294967295,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"as_number": {
										Type:     schema.TypeString,
										Required: true,
										ForceNew: true,
									},
									"prepend_global_as": {
										Type:     schema.TypeBool,
										Optional: true,
									},
									"prepend_local_as": {
										Type:     schema.TypeBool,
										Optional: true,
									},
								},
							},
						},
						"local_preference": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"next_hop_self": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"peer_address": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"peer_as": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"peer_group": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"route_reflector": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"client": {
										Type:     schema.TypeBool,
										Optional: true,
									},
									"cluster_id": {
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
						"send_community": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"large": {
										Type:     schema.TypeBool,
										Optional: true,
									},
									"standard": {
										Type:     schema.TypeBool,
										Optional: true,
									},
								},
							},
						},
						"send_default_route": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"export_policy": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"ipv4_unicast": {
										Type:     schema.TypeBool,
										Optional: true,
									},
									"ipv6_unicast": {
										Type:     schema.TypeBool,
										Optional: true,
									},
								},
							},
						},
						"timers": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"connect_retry": {
										Type:     schema.TypeInt,
										Optional: true,
									},
									"hold_time": {
										Type:     schema.TypeInt,
										Optional: true,
									},
									"keepalive_interval": {
										Type:     schema.TypeInt,
										Optional: true,
									},
									"minimum_advertisement_interval": {
										Type:     schema.TypeInt,
										Optional: true,
									},
								},
							},
						},
						"trace_options": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"flag": {
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"modifier": {
													Type:     schema.TypeString,
													Optional: true,
												},
												"name": {
													Type:     schema.TypeString,
													Required: true,
													ForceNew: true,
												},
											},
										},
									},
								},
							},
						},
						"transport": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"local_address": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"passive_mode": {
										Type:     schema.TypeBool,
										Optional: true,
									},
									"tcp_mss": {
										Type:     schema.TypeInt,
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

func resourceNetworkInstanceInstanceProtocolsBgpNeighborCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Infof("Beginning Create: %s", resourceNetworkInstanceInstanceProtocolsBgpNeighborString(d))
	target := meta.(*Target)

	rn := "neighbor"
	rk := "peer-address"
	key, err := getResourceListKey(&rn, &rk, d)
	if err != nil {
		return diag.FromErr(err)
	}

	hkey0 := d.Get("network_instance_id").(string)

	//p := fmt.Sprintf("fmt.Sprintf("/network-instance[name=%s]/protocols/bgp/",hkey0)", hkey)
	p := fmt.Sprintf("/network-instance[name=%s]/protocols/bgp/", hkey0)

	v := ""

	hid := "bgp_id"
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
	return resourceNetworkInstanceInstanceProtocolsBgpNeighborRead(ctx, d, meta)
}

func resourceNetworkInstanceInstanceProtocolsBgpNeighborRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Infof("Beginning Read: %s", resourceNetworkInstanceInstanceProtocolsBgpNeighborString(d))
	target := meta.(*Target)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	hkey0 := d.Get("network_instance_id").(string)

	//hkey := d.Get("[network_instance_id]").(string)

	//rn := "neighbor"
	rk := "peer-address"
	key := d.Id()

	//p := fmt.Sprintf("fmt.Sprintf("/network-instance[name=%s]/protocols/bgp/neighbor[peer-address=%s]",hkey0, key)", hkey, key)
	p := fmt.Sprintf("/network-instance[name=%s]/protocols/bgp/neighbor[peer-address=%s]", hkey0, key)

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
			switch x := upd.Values["neighbor"].(type) {
			case map[string]interface{}:
				for k, v := range x {
					log.Debugf("BEFORE KEY: %s, VALUE: %v", k, v)
					sk := strings.Split(k, ":")[len(strings.Split(k, ":"))-1]

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
			if err := d.Set("neighbor", data); err != nil {
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

func resourceNetworkInstanceInstanceProtocolsBgpNeighborUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Infof("Beginning Update: %s", resourceNetworkInstanceInstanceProtocolsBgpNeighborString(d))
	target := meta.(*Target)

	rn := "neighbor"
	rk := "peer-address"
	key, err := getResourceListKey(&rn, &rk, d)
	if err != nil {
		return diag.FromErr(err)
	}

	hkey0 := d.Get("network_instance_id").(string)

	//hkey := d.Get("[network_instance_id]").(string)
	//p := fmt.Sprintf("fmt.Sprintf("/network-instance[name=%s]/protocols/bgp/",hkey0)", hkey)
	p := fmt.Sprintf("/network-instance[name=%s]/protocols/bgp/", hkey0)

	v := ""

	hid := "bgp_id"
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
	return resourceNetworkInstanceInstanceProtocolsBgpNeighborRead(ctx, d, meta)
}

func resourceNetworkInstanceInstanceProtocolsBgpNeighborDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Debugf("Beginning delete: %s", resourceNetworkInstanceInstanceProtocolsBgpNeighborString(d))
	target := meta.(*Target)

	hkey0 := d.Get("network_instance_id").(string)

	//hkey := d.Get("[network_instance_id]").(string)

	//p := fmt.Sprintf("fmt.Sprintf("/network-instance[name=%s]/protocols/bgp/neighbor[peer-address=%s]",hkey0, d.Id())", hkey, d.Id())
	p := fmt.Sprintf("/network-instance[name=%s]/protocols/bgp/neighbor[peer-address=%s]", hkey0, d.Id())

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
