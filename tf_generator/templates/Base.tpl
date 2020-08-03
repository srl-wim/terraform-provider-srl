{{define "base"}}
package {{ .PkgName }}

func resource{{ .ResourceName }}String(d resourceIDStringer) string {
	return resourceIDString(d, "srl_{{ ToSnakeCase .ResourceName }}")
}

func resource{{ .ResourceName }}() *schema.Resource {
	s := map[string]*schema.Schema{
	{{range .Resources}}
	    "{{ $.Name }}": {
			Type:     {{ YangToTFType .Type }},
			Optional: {{ .Optional }},
		},
	{{end}}
	}

	return &schema.Resource{
		Schema: s,

		CreateContext: resource{{ .Resource }}Create,
		ReadContext:   resource{{ .Resource }}Read,
		UpdateContext: resource{{ .Resource }}Update,
		DeleteContext: resource{{ .Resource }}Delete,

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(5 * time.Minute),
			Read:   schema.DefaultTimeout(5 * time.Minute),
			Update: schema.DefaultTimeout(5 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},
	}
}

{{template "createfunc" .}}

{{template "readfunc" .}}

{{template "updatefunc" .}}

{{template "deletefunc" .}}
{end}