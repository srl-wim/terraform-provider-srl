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
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/google/gnxi/utils/xpath"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/openconfig/gnmi/proto/gnmi"
	"google.golang.org/grpc/metadata"
)

// resourceSystemMtuString function
func resourceSystemMtuString(d resourceIDStringer) string {
	return resourceIDString(d, "system_mtu")
}

// resourceSystemMtu function
func resourceSystemMtu() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceSystemMtuCreate,
		ReadContext:   resourceSystemMtuRead,
		UpdateContext: resourceSystemMtuUpdate,
		DeleteContext: resourceSystemMtuDelete,

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
            "default_ip_mtu": {
                Type:     schema.TypeInt,
                Optional: true,
                Default: "1500",
            },
            "default_port_mtu": {
                Type:     schema.TypeInt,
                Optional: true,
                Default: "9232",
            },
            "min_path_mtu": {
                Type:     schema.TypeInt,
                Optional: true,
                Default: "552",
            },

        },
    }
}

func resourceSystemMtuCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] %s: beginning create", resourceSystemMtuString(d))
	diagnostics := make([]diag.Diagnostic, 0)
	config := meta.(BaseConfig)

	conn, err := createGrpcConn(meta)
	if err != nil {
		log.Printf("[ERROR] Dialing to %q failed: %v", config.target, err)
		return diagnostics
	}

	client := gnmi.NewGNMIClient(conn)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	ctx = metadata.AppendToOutgoingContext(ctx, "username", config.username, "password", config.password)

	path := "/system/mtu"
	gnmiPath, err := ParsePath(strings.TrimSpace(path))
	if err != nil {
		log.Printf("[ERROR] Path parsing failed : %v", err)
		return diagnostics
	}

	log.Printf("[DEBUG] %s: get", d.Get("mtu"))
	specBytes, _ := json.Marshal(d.Get("mtu"))
	fmt.Printf("bytes: %s \n", specBytes)
	value := new(gnmi.TypedValue)
	value.Value = &gnmi.TypedValue_JsonIetfVal{
		JsonIetfVal: bytes.Trim(specBytes, " \r\n\t"),
	}

	gnmiPrefix, err := CreatePrefix("", config.target)
	if err != nil {
		log.Printf("[ERROR] Path prefix failed : %v", err)
		return diagnostics
	}

	req := &gnmi.SetRequest{
		Prefix:  gnmiPrefix,
		Delete:  make([]*gnmi.Path, 0, 0),
		Replace: make([]*gnmi.Update, 0),
		Update:  make([]*gnmi.Update, 0),
	}

	req.Update = append(req.Update, &gnmi.Update{
		Path: gnmiPath,
		Val:  value,
	})

	log.Printf("[DEBUG] %s: get", d.Get("mtu"))

	log.Printf("[DEBUG] : Req: %v", req)
	response, err := client.Set(ctx, req)
	if err != nil {
		log.Printf("[ERROR] : Set failed: %v", err)
	}

	log.Printf("[DEBUG] %v: set response", response)

	d.SetId("mtu")
	return resourceSystemMtuRead(ctx, d, meta)
}

func resourceSystemMtuRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] %s: beginning read", resourceSystemMtuString(d))
	diagnostics := make([]diag.Diagnostic, 0)
	config := meta.(BaseConfig)

	conn, err := createGrpcConn(meta)
	if err != nil {
		log.Printf("[ERROR] Dialing to %q failed: %v", config.target, err)
		return diagnostics
	}

	client := gnmi.NewGNMIClient(conn)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	ctx = metadata.AppendToOutgoingContext(ctx, "username", config.username, "password", config.password)

	encodingVal, ok := gnmi.Encoding_value[strings.Replace(strings.ToUpper(config.encoding), "-", "_", -1)]
	if !ok {
		var gnmiEncodingList []string
		for _, name := range gnmi.Encoding_name {
			gnmiEncodingList = append(gnmiEncodingList, name)
		}
		log.Printf("[ERROR] Supported encodings: %s", strings.Join(gnmiEncodingList, ", "))
	}

	req := &gnmi.GetRequest{
		UseModels: make([]*gnmi.ModelData, 0),
		Path:      make([]*gnmi.Path, 0),
		Encoding:  gnmi.Encoding(encodingVal),
	}
	paths := make([]string, 0)
	paths = append(paths, "/system/mtu")

	for _, path := range paths {
		gnmiPath, err := xpath.ToGNMIPath(path)
		if err != nil {
			log.Printf("[ERROR] Error in parsing xpath %q to gnmi path", path)
		}
		req.Path = append(req.Path, gnmiPath)
	}

	response, err := client.Get(ctx, req)
	if err != nil {
		log.Printf("[ERROR] : get failed: %v", err)
	}

	log.Printf("[DEBUG] %v: get response", response)

	return nil
}

func resourceSystemMtuUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] %s: beginning update", resourceSystemMtuString(d))
	diagnostics := make([]diag.Diagnostic, 0)
	config := meta.(BaseConfig)

	conn, err := createGrpcConn(meta)
	if err != nil {
		log.Printf("[ERROR] Dialing to %q failed: %v", config.target, err)
		return diagnostics
	}

	client := gnmi.NewGNMIClient(conn)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	ctx = metadata.AppendToOutgoingContext(ctx, "username", config.username, "password", config.password)

	path := "/system/mtu"
	gnmiPath, err := ParsePath(strings.TrimSpace(path))
	if err != nil {
		log.Printf("[ERROR] Path parsing failed : %v", err)
		return diagnostics
	}

	log.Printf("[DEBUG] %s: get", d.Get("mtu"))
	specBytes, _ := json.Marshal(d.Get("mtu"))
	fmt.Printf("bytes: %s \n", specBytes)
	value := new(gnmi.TypedValue)
	value.Value = &gnmi.TypedValue_JsonIetfVal{
		JsonIetfVal: bytes.Trim(specBytes, " \r\n\t"),
	}

	gnmiPrefix, err := CreatePrefix("", config.target)
	if err != nil {
		log.Printf("[ERROR] Path prefix failed : %v", err)
		return diagnostics
	}

	req := &gnmi.SetRequest{
		Prefix:  gnmiPrefix,
		Delete:  make([]*gnmi.Path, 0, 0),
		Replace: make([]*gnmi.Update, 0),
		Update:  make([]*gnmi.Update, 0),
	}

	req.Update = append(req.Update, &gnmi.Update{
		Path: gnmiPath,
		Val:  value,
	})

	log.Printf("[DEBUG] %s: get", d.Get("mtu"))

	log.Printf("[DEBUG] : Req: %v", req)
	response, err := client.Set(ctx, req)
	if err != nil {
		log.Printf("[ERROR] : Set failed: %v", err)
	}

	log.Printf("[DEBUG] %v: set response", response)

	d.SetId("mtu")
	return resourceSystemMtuRead(ctx, d, meta)
}

func resourceSystemMtuDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] %s: beginning delete", resourceSystemMtuString(d))
	diagnostics := make([]diag.Diagnostic, 0)
	config := meta.(BaseConfig)

	conn, err := createGrpcConn(meta)
	if err != nil {
		log.Printf("[ERROR] Dialing to %q failed: %v", config.target, err)
		return diagnostics
	}

	client := gnmi.NewGNMIClient(conn)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	ctx = metadata.AppendToOutgoingContext(ctx, "username", config.username, "password", config.password)

	var deleteList []*gnmi.Path

	path := "/system/mtu"

	gnmiPath, err := xpath.ToGNMIPath(path)
	if err != nil {
		log.Printf("[ERROR] Error in parsing xpath %q to gnmi path", path)
	}
	deleteList = append(deleteList, gnmiPath)

	req := &gnmi.SetRequest{
		Delete: deleteList,
	}

	log.Printf("[DEBUG] : Delete Req: %v", req)
	response, err := client.Set(ctx, req)
	if err != nil {
		log.Printf("[ERROR] : Delete Set failed: %v", err)
	}

	log.Printf("[DEBUG] %v: delete set response", response)

	d.SetId("")
	return nil
}
