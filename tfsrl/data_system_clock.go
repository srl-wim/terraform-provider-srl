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
	"encoding/json"
	"strconv"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	log "github.com/sirupsen/logrus"
)

// daraSystemClockString function
func dataSystemClockString(d resourceIDStringer) string {
	return resourceIDString(d, "system_clock")
}

// dataSystemClock function
func dataSystemClock() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSystemClockRead,

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Timeouts: &schema.ResourceTimeout{
			Read: schema.DefaultTimeout(5 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"clock": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"timezone": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSystemClockRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Infof("Beginning Read: %s", dataSystemClockString(d))
	target := meta.(*Target)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	p := "/system/clock"

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

	notifications := make([]NotificationRspMsg, 0, len(response.GetNotification()))
	for _, notif := range response.GetNotification() {
		msg := NotificationRspMsg{
			Prefix:  gnmiPathToXPath(notif.Prefix),
			Updates: make([]update, 0, len(notif.GetUpdate())),
			Deletes: make([]string, 0, len(notif.GetDelete())),
		}
		msg.Timestamp = notif.Timestamp
		t := time.Unix(0, notif.Timestamp)
		msg.Time = &t

		for i, upd := range notif.GetUpdate() {
			pathElems := make([]string, 0, len(upd.GetPath().GetElem()))
			for _, pElem := range upd.GetPath().GetElem() {
				pathElems = append(pathElems, pElem.GetName())
			}
			log.Infof("pathElems: %v", pathElems)
			var pathElemSplit []string
			if len(pathElems) > 1 {
				pathElemSplit = strings.Split(pathElems[len(pathElems)-1], ":")
			} else {
				pathElemSplit = strings.Split(pathElems[0], ":")
			}
			var pathElem string
			if len(pathElemSplit) > 1 {
				pathElem = pathElemSplit[len(pathElemSplit)-1]
			} else {
				pathElem = pathElemSplit[0]
			}

			log.Infof("pathElem: %s", pathElem)
			value, err := getValue(upd.GetVal())
			if err != nil {
				return diag.FromErr(err)
			}
			value = replaceInKeys(value, "-", "_")
			msg.Updates = append(msg.Updates,
				update{
					Path:   gnmiPathToXPath(upd.GetPath()),
					Values: make(map[string]interface{}),
				})
			msg.Updates[i].Values[pathElem] = value
		}
		for _, del := range notif.GetDelete() {
			msg.Deletes = append(msg.Deletes, gnmiPathToXPath(del))
		}
		notifications = append(notifications, msg)

	}
	x, err := json.Marshal(notifications)
	if err != nil {
		return diag.FromErr(err)
	}
	sb := strings.Builder{}
	sb.Write(x)

	log.Infof("Get response: \n%v\n", sb.String())

	data := make([]map[string]interface{}, 0)
	data = append(data, notifications[0].Updates[0].Values)

	log.Infof("Data: %v", data)

	if err := d.Set("clock", data); err != nil {
		return diag.FromErr(err)
	}
	// always run
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}