package tf_srl

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"username": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "username to connect to SRL device",
				DefaultFunc: schema.EnvDefaultFunc("SRL_USERNAME", defaultUsername),
			},
			"password": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "password to connect to SRL device",
				DefaultFunc: schema.EnvDefaultFunc("SRL_PASSWORD", defaultPassword),
			},
			"target": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "target to connect to <ip>:<port>",
				DefaultFunc: schema.EnvDefaultFunc("SRL_TARGET", defaultTarget),
			},
			"proxy": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "proxy to connect to the device",
				DefaultFunc: schema.EnvDefaultFunc("SRL_PROXY", false),
			},
			"tls_ca": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "tls certificate authority",
				DefaultFunc: schema.EnvDefaultFunc("SRL_TLS_CA", ""),
			},
			"tls_cert": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "tls certificate",
				DefaultFunc: schema.EnvDefaultFunc("SRL_TLS_CERT", ""),
			},
			"tls_key": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "tls key",
				DefaultFunc: schema.EnvDefaultFunc("SRL_TLS_KEY", ""),
			},
			"skip_verify": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "skip verify tls connection",
				DefaultFunc: schema.EnvDefaultFunc("SRL_SKIP_VERIFY", true),
			},
			"insecure": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "insecure connection",
				DefaultFunc: schema.EnvDefaultFunc("SRL_INSECURE", true),
			},
			"timeout": {
				Type:        schema.TypeFloat,
				Optional:    true,
				Description: "grpc timeout (default: 30 sec)",
				DefaultFunc: schema.EnvDefaultFunc("SRL_TIMEOUT", defaultTimeout.Seconds()),
			},
			"encoding": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "grpc encoding",
				DefaultFunc: schema.EnvDefaultFunc("SRL_ENCODING", defaultEncoding),
			},
		},
		DataSourcesMap: map[string]*schema.Resource{},
		ResourcesMap: map[string]*schema.Resource{
			"srl_system_ntp":   resourceSystemNtp(),
			"srl_system_clock": resourceSystemClock(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	u, uOK := d.GetOk("username")
	p, pOK := d.GetOk("password")
	t, tOK := d.GetOk("target")

	log.Printf("[DEBUG] %s: username", u)
	log.Printf("[DEBUG] %s: password", p)
	log.Printf("[DEBUG] %s: target", t)

	if !uOK || !pOK || !tOK {
		return nil, fmt.Errorf("username (%#v), password  (%#v) and target (%#v) must be set", u.(string), p.(string), t.(string))
	}

	encoding := d.Get("encoding")
	log.Printf("[DEBUG] %s: encoding", encoding)
	if encoding == "" {
		encoding = defaultEncoding
	}

	baseConfig := BaseConfig{
		username: u.(string),
		password: p.(string),
		target:   t.(string),
		proxy:    false,
		insecure: true,
		notls:    false,
		timeout:  defaultTimeout,
		encoding: defaultEncoding,
		//timeout:  time.Duration(int64(d.Get("timeout").(float64)) * int64(time.Second)),
	}

	return baseConfig, nil
}

type resourceIDStringer interface {
	Id() string
}

func resourceIDString(d resourceIDStringer, name string) string {
	id := d.Id()
	if id == "" {
		id = "<new resource>"
	}

	return fmt.Sprintf("%s (ID = %s)", name, id)
}
