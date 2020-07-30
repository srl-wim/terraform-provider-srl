package tf_srl

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"google.golang.org/grpc/metadata"

	"github.com/google/gnxi/utils/xpath"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/openconfig/gnmi/proto/gnmi"
)

func resourceSystemNtpString(d resourceIDStringer) string {
	return resourceIDString(d, "srl_system_ntp")
}

func resourceSystemNtp() *schema.Resource {
	s := map[string]*schema.Schema{
		"admin_state": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"network_instance": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
	}

	return &schema.Resource{
		Schema: s,

		Create: resourceSystemNtpCreate,
		Read:   resourceSystemNtpRead,
		Update: resourceSystemNtpUpdate,
		Delete: resourceSystemNtpDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(5 * time.Minute),
			Read:   schema.DefaultTimeout(5 * time.Minute),
			Update: schema.DefaultTimeout(5 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},
	}
}

func resourceSystemNtpCreate(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] %s: beginning create", resourceSystemNtpString(d))
	log.Printf("[DEBUG] %v: meta", meta)

	config := meta.(BaseConfig)

	conn, err := createGrpcConn(meta)
	if err != nil {
		log.Printf("[ERROR] Dialing to %q failed: %v", config.target, err)
		return err
	}

	client := gnmi.NewGNMIClient(conn)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	ctx = metadata.AppendToOutgoingContext(ctx, "username", config.username, "password", config.password)

	p := fmt.Sprintf("/system/ntp/admin-state:\"%s\"", d.Get("admin_state").(string))
	log.Printf("[DEBUG] %s: path", p)
	path := []string{p}

	updateList, err := buildPbUpdateList(path)

	req := &gnmi.SetRequest{
		Update: updateList,
	}

	log.Printf("[DEBUG] : Req: %v", req)
	response, err := client.Set(ctx, req)
	if err != nil {
		log.Printf("[ERROR] : Set failed: %v", err)
	}

	log.Printf("[DEBUG] %v: set response", response)

	adminState := d.Get("admin_state").(string)
	d.SetId(admin_state)
	return resourceSystemNtpRead(d, meta)
}

func resourceSystemNtpRead(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] %s: beginning read", resourceSystemNtpString(d))
	log.Printf("[DEBUG] %v: meta", meta)

	config := meta.(BaseConfig)

	conn, err := createGrpcConn(meta)
	if err != nil {
		log.Printf("[ERROR] Dialing to %q failed: %v", config.target, err)
		return err
	}

	client := gnmi.NewGNMIClient(conn)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	ctx = metadata.AppendToOutgoingContext(ctx, "username", config.username, "password", config.password)

	log.Printf("[DEBUG] %v: encoding", config.encoding)
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

	path := "/system/ntp"

	gnmiPath, err := xpath.ToGNMIPath(path)
	if err != nil {
		log.Printf("[ERROR] Error in parsing xpath %q to gnmi path", path)
	}
	req.Path = append(req.Path, gnmiPath)

	response, err := client.Get(ctx, req)
	if err != nil {
		log.Printf("[ERROR] : get failed: %v", err)
	}

	log.Printf("[DEBUG] %v: get response", response)

	return nil
}

func resourceSystemNtpUpdate(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] %s: beginning update", resourceSystemNtpString(d))

	config := meta.(BaseConfig)

	conn, err := createGrpcConn(meta)
	if err != nil {
		log.Printf("[ERROR] Dialing to %q failed: %v", config.target, err)
		return err
	}

	client := gnmi.NewGNMIClient(conn)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	ctx = metadata.AppendToOutgoingContext(ctx, "username", config.username, "password", config.password)

	p := fmt.Sprintf("/system/ntp/admin-state:\"%s\"", d.Get("admin_state").(string))
	log.Printf("[DEBUG] %s: path", p)
	path := []string{p}

	updateList, err := buildPbUpdateList(path)

	req := &gnmi.SetRequest{
		Update: updateList,
	}

	log.Printf("[DEBUG] : Req: %v", req)
	response, err := client.Set(ctx, req)
	if err != nil {
		log.Printf("[ERROR] : Set failed: %v", err)
	}

	log.Printf("[DEBUG] %v: set response", response)

	adminState := d.Get("admin_state").(string)
	d.SetId(admin_state)
	return resourceSystemNtpRead(d, meta)
}

func resourceSystemNtpDelete(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] %s: beginning delete", resourceSystemNtpString(d))

	config := meta.(BaseConfig)

	conn, err := createGrpcConn(meta)
	if err != nil {
		log.Printf("[ERROR] Dialing to %q failed: %v", config.target, err)
		return err
	}

	client := gnmi.NewGNMIClient(conn)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	ctx = metadata.AppendToOutgoingContext(ctx, "username", config.username, "password", config.password)

	var deleteList []*gnmi.Path

	path := "/system/ntp"

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
