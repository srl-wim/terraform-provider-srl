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

	"fmt"

	"strconv"

	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	log "github.com/sirupsen/logrus"
)

// resourceInterfacesSubinterfaceIpv4String function
func resourceInterfacesSubinterfaceIpv4String(d resourceIDStringer) string {
	return resourceIDString(d, "interfaces_subinterface_ipv4")
}

// resourceInterfacesSubinterfaceIpv4 function
func resourceInterfacesSubinterfaceIpv4() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceInterfacesSubinterfaceIpv4Create,
		ReadContext:   resourceInterfacesSubinterfaceIpv4Read,
		UpdateContext: resourceInterfacesSubinterfaceIpv4Update,
		DeleteContext: resourceInterfacesSubinterfaceIpv4Delete,

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
			"interface_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"subinterface_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"ipv4": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"allow_directed_broadcast": {
							Type:     schema.TypeBool,
							Optional: true,
							Default:  false,
						},
					},
				},
			},
		},
	}
}

func resourceInterfacesSubinterfaceIpv4Create(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Infof("Beginning Create: %s", resourceInterfacesSubinterfaceIpv4String(d))
	target := meta.(*Target)

	key := "ipv4"

	hkey0 := d.Get("interface_id").(string)

	hkey1 := d.Get("subinterface_id").(string)

	//hkey := d.Get("[interface_id subinterface_id]").(string)
	//p := fmt.Sprintf("fmt.Sprintf("/interface[name=%s]/subinterface[index=%s]/ipv4",hkey0,hkey1)", hkey)
	p := fmt.Sprintf("/interface[name=%s]/subinterface[index=%s]/ipv4", hkey0, hkey1)

	v := "ipv4"

	hid := make([]string, 0)

	hid = append(hid, "interface_id")

	hid = append(hid, "subinterface_id")

	//hid = append(hid, "subinterface_id")
	//hid := "subinterface_id"
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
	return resourceInterfacesSubinterfaceIpv4Read(ctx, d, meta)
}

func resourceInterfacesSubinterfaceIpv4Read(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Infof("Beginning Read: %s", resourceInterfacesSubinterfaceIpv4String(d))
	target := meta.(*Target)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	hkey0 := d.Get("interface_id").(string)

	hkey1 := d.Get("subinterface_id").(string)

	//hkey := d.Get("[interface_id subinterface_id]").(string)

	//p := fmt.Sprintf("fmt.Sprintf("/interface[name=%s]/subinterface[index=%s]/ipv4",hkey0,hkey1)", hkey)
	p := fmt.Sprintf("/interface[name=%s]/subinterface[index=%s]/ipv4", hkey0, hkey1)

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
			switch x := upd.Values["ipv4"].(type) {
			case map[string]interface{}:
				for k, v := range x {
					log.Debugf("BEFORE KEY: %s, VALUE: %v", k, v)
					sk := strings.Split(k, ":")[len(strings.Split(k, ":"))-1]

					switch sk {

					case "address":
						delete(x, k)

					case "arp":
						delete(x, k)

					case "dhcp_client":
						delete(x, k)

					case "dhcp_relay":
						delete(x, k)

					case "vrrp":
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

				data = append(data, x)
			}
			log.Debugf("get response: index: %d, data: %v", i, data)
			if err := d.Set("ipv4", data); err != nil {
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

func resourceInterfacesSubinterfaceIpv4Update(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Infof("Beginning Update: %s", resourceInterfacesSubinterfaceIpv4String(d))
	target := meta.(*Target)

	key := "ipv4"

	hkey0 := d.Get("interface_id").(string)

	hkey1 := d.Get("subinterface_id").(string)

	//hkey := d.Get("[interface_id subinterface_id]").(string)
	//p := fmt.Sprintf("fmt.Sprintf("/interface[name=%s]/subinterface[index=%s]/ipv4",hkey0,hkey1)", hkey)
	p := fmt.Sprintf("/interface[name=%s]/subinterface[index=%s]/ipv4", hkey0, hkey1)

	v := "ipv4"

	hid := make([]string, 0)

	hid = append(hid, "interface_id")

	hid = append(hid, "subinterface_id")

	//hid = append(hid, "subinterface_id")
	//hid := "subinterface_id"
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
	return resourceInterfacesSubinterfaceIpv4Read(ctx, d, meta)
}

func resourceInterfacesSubinterfaceIpv4Delete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Debugf("Beginning delete: %s", resourceInterfacesSubinterfaceIpv4String(d))
	target := meta.(*Target)

	hkey0 := d.Get("interface_id").(string)

	hkey1 := d.Get("subinterface_id").(string)

	//hkey := d.Get("[interface_id subinterface_id]").(string)

	//p := fmt.Sprintf("fmt.Sprintf("/interface[name=%s]/subinterface[index=%s]/ipv4",hkey0,hkey1)", hkey)
	p := fmt.Sprintf("/interface[name=%s]/subinterface[index=%s]/ipv4", hkey0, hkey1)

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
