package main

import (
	//"terraform-provider-srl/tf_srl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	"github.com/wim-srl/terraform-provider-srl/tf_srl"
	//"tf_srl tf_srl"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: func() *schema.Provider {
			return tf_srl.Provider()
		},
	})
}
