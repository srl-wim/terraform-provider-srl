{{define "createfunc"}}
func resource{{ .ResourceName }}Create(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] %s: beginning create", resource{{ .ResourceName }}String(d))
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

	p := fmt.Sprintf("/system/clock/timezone:\"%s\"", d.Get("timezone").(string))
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

	timezone := d.Get("timezone").(string)
	d.SetId(timezone)
	return resource{{ .ResourceName }}Read(ctx, d, meta)
}
{{end}}