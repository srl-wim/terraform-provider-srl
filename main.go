package main

import (
	//"terraform-provider-srl/tf_srl"
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	tf_srl "github.com/henderiw/srl-wim/terraform-provider-srl/tf_srl"
	//"tf_srl tf_srl"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: func() terraform.ResourceProvider {
			return tf_srl.Provider()
		},
	})
}
