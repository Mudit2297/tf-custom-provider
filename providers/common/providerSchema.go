package common

import (
	gh "cpf/providers/terraform-provider-github"
	//hcp "cpf/providers/terraform-provider-hashicups"

	"cpf"
)

var PSchema = cpf.CustomSchema{
	Schemas: []cpf.SchemaMap{
		//hcp.HashicupsProviderSchema,
		gh.GitHubProviderSchema,
	},
}
