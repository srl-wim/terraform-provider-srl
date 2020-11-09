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

// resourceSystemIpLoadBalancingString function
func resourceSystemIpLoadBalancingString(d resourceIDStringer) string {
	return resourceIDString(d, "system_ip_load_balancing")
}

// resourceSystemIpLoadBalancing function
func resourceSystemIpLoadBalancing() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceSystemIpLoadBalancingCreate,
		ReadContext:   resourceSystemIpLoadBalancingRead,
		UpdateContext: resourceSystemIpLoadBalancingUpdate,
		DeleteContext: resourceSystemIpLoadBalancingDelete,

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
            "hash_keys": {
                Type:     schema.TypeList,
                Optional: true,
                MaxItems: 1,
                Elem: &schema.Resource{
                	Schema: map[string]*schema.Schema{
                        "hash_seed": {
                            Type:     schema.TypeInt,
                            Optional: true,
                            Default: "0",
                        },
                    },
                },
            },

        },
    }
}

func resourceSystemIpLoadBalancingCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] %s: beginning create", resourceSystemIpLoadBalancingString(d))
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

	path := "/system/ip-load-balancing"
	gnmiPath, err := ParsePath(strings.TrimSpace(path))
	if err != nil {
		log.Printf("[ERROR] Path parsing failed : %v", err)
		return diagnostics
	}

	log.Printf("[DEBUG] %s: get", d.Get("ip-load-balancing"))
	specBytes, _ := json.Marshal(d.Get("ip-load-balancing"))
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

	log.Printf("[DEBUG] %s: get", d.Get("ip-load-balancing"))

	log.Printf("[DEBUG] : Req: %v", req)
	response, err := client.Set(ctx, req)
	if err != nil {
		log.Printf("[ERROR] : Set failed: %v", err)
	}

	log.Printf("[DEBUG] %v: set response", response)

	d.SetId("ip-load-balancing")
	return resourceSystemIpLoadBalancingRead(ctx, d, meta)
}

func resourceSystemIpLoadBalancingRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] %s: beginning read", resourceSystemIpLoadBalancingString(d))
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
	paths = append(paths, "/system/ip-load-balancing")

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

func resourceSystemIpLoadBalancingUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] %s: beginning update", resourceSystemIpLoadBalancingString(d))
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

	path := "/system/ip-load-balancing"
	gnmiPath, err := ParsePath(strings.TrimSpace(path))
	if err != nil {
		log.Printf("[ERROR] Path parsing failed : %v", err)
		return diagnostics
	}

	log.Printf("[DEBUG] %s: get", d.Get("ip-load-balancing"))
	specBytes, _ := json.Marshal(d.Get("ip-load-balancing"))
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

	log.Printf("[DEBUG] %s: get", d.Get("ip-load-balancing"))

	log.Printf("[DEBUG] : Req: %v", req)
	response, err := client.Set(ctx, req)
	if err != nil {
		log.Printf("[ERROR] : Set failed: %v", err)
	}

	log.Printf("[DEBUG] %v: set response", response)

	d.SetId("ip-load-balancing")
	return resourceSystemIpLoadBalancingRead(ctx, d, meta)
}

func resourceSystemIpLoadBalancingDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] %s: beginning delete", resourceSystemIpLoadBalancingString(d))
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

	path := "/system/ip-load-balancing"

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
