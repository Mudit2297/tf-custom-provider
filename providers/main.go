package main

import (
	"cpf/providers/common"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"

	"cpf"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: func() *schema.Provider {
			return cpf.Provider(common.PSchema, common.Resource, common.Data, common.ConfigContex)
		},
	})
}
