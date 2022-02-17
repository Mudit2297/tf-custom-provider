package common

import (
	git "cpf/providers/terraform-provider-github"
	//hcp "cpf/providers/terraform-provider-hashicups"

	"cpf"
)

var PSchema = cpf.CustomSchema{
	Schemas: []cpf.SchemaMap{
		//hcp.HashicupsProviderSchema,
		git.GitHubProviderSchema,
	},
}
