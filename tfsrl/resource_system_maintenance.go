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

// resourceSystemMaintenanceString function
func resourceSystemMaintenanceString(d resourceIDStringer) string {
	return resourceIDString(d, "system_maintenance")
}

// resourceSystemMaintenance function
func resourceSystemMaintenance() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceSystemMaintenanceCreate,
		ReadContext:   resourceSystemMaintenanceRead,
		UpdateContext: resourceSystemMaintenanceUpdate,
		DeleteContext: resourceSystemMaintenanceDelete,

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
            "group": {
                Type:     schema.TypeList,
                Optional: true,
                MaxItems: 1,
                Elem: &schema.Resource{
                	Schema: map[string]*schema.Schema{
                        "maintenance_mode": {
                            Type:     schema.TypeList,
                            Optional: true,
                            MaxItems: 1,
                            Elem: &schema.Resource{
                            	Schema: map[string]*schema.Schema{
                                    "admin_state": {
                                        Type:     schema.TypeString,
                                        Optional: true,
                                        Default: "disable",
                                    },
                                },
                            },
                        },
                        "maintenance_profile": {
                            Type:     schema.TypeString,
                            Optional: true,
                        },
                        "members": {
                            Type:     schema.TypeList,
                            Optional: true,
                            MaxItems: 1,
                            Elem: &schema.Resource{
                            	Schema: map[string]*schema.Schema{
                                    "bgp": {
                                        Type:     schema.TypeList,
                                        Optional: true,
                                        MaxItems: 1,
                                        Elem: &schema.Resource{
                                        	Schema: map[string]*schema.Schema{
                                                "network_instance": {
                                                    Type:     schema.TypeList,
                                                    Optional: true,
                                                    MaxItems: 1,
                                                    Elem: &schema.Resource{
                                                    	Schema: map[string]*schema.Schema{
                                                            "name": {
                                                                Type:     schema.TypeString,
                                                                Required: true,
                                                            },
                                                            "neighbor": {
                                                                Type:     schema.TypeString,
                                                                Optional: true,
                                                            },
                                                            "peer_group": {
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
                            },
                        },
                        "name": {
                            Type:     schema.TypeString,
                            Required: true,
                        },
                    },
                },
            },
            "profile": {
                Type:     schema.TypeList,
                Optional: true,
                MaxItems: 1,
                Elem: &schema.Resource{
                	Schema: map[string]*schema.Schema{
                        "bgp": {
                            Type:     schema.TypeList,
                            Optional: true,
                            MaxItems: 1,
                            Elem: &schema.Resource{
                            	Schema: map[string]*schema.Schema{
                                    "export_policy": {
                                        Type:     schema.TypeString,
                                        Optional: true,
                                    },
                                    "import_policy": {
                                        Type:     schema.TypeString,
                                        Optional: true,
                                    },
                                },
                            },
                        },
                        "name": {
                            Type:     schema.TypeString,
                            Required: true,
                        },
                    },
                },
            },

        },
    }
}

func resourceSystemMaintenanceCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] %s: beginning create", resourceSystemMaintenanceString(d))
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

	path := "/system/maintenance"
	gnmiPath, err := ParsePath(strings.TrimSpace(path))
	if err != nil {
		log.Printf("[ERROR] Path parsing failed : %v", err)
		return diagnostics
	}

	log.Printf("[DEBUG] %s: get", d.Get("maintenance"))
	specBytes, _ := json.Marshal(d.Get("maintenance"))
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

	log.Printf("[DEBUG] %s: get", d.Get("maintenance"))

	log.Printf("[DEBUG] : Req: %v", req)
	response, err := client.Set(ctx, req)
	if err != nil {
		log.Printf("[ERROR] : Set failed: %v", err)
	}

	log.Printf("[DEBUG] %v: set response", response)

	d.SetId("maintenance")
	return resourceSystemMaintenanceRead(ctx, d, meta)
}

func resourceSystemMaintenanceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] %s: beginning read", resourceSystemMaintenanceString(d))
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
	paths = append(paths, "/system/maintenance")

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

func resourceSystemMaintenanceUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] %s: beginning update", resourceSystemMaintenanceString(d))
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

	path := "/system/maintenance"
	gnmiPath, err := ParsePath(strings.TrimSpace(path))
	if err != nil {
		log.Printf("[ERROR] Path parsing failed : %v", err)
		return diagnostics
	}

	log.Printf("[DEBUG] %s: get", d.Get("maintenance"))
	specBytes, _ := json.Marshal(d.Get("maintenance"))
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

	log.Printf("[DEBUG] %s: get", d.Get("maintenance"))

	log.Printf("[DEBUG] : Req: %v", req)
	response, err := client.Set(ctx, req)
	if err != nil {
		log.Printf("[ERROR] : Set failed: %v", err)
	}

	log.Printf("[DEBUG] %v: set response", response)

	d.SetId("maintenance")
	return resourceSystemMaintenanceRead(ctx, d, meta)
}

func resourceSystemMaintenanceDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] %s: beginning delete", resourceSystemMaintenanceString(d))
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

	path := "/system/maintenance"

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
