package common

import (
	git "cpf/providers/terraform-provider-github"
	//hcp "cpf/providers/terraform-provider-hashicups"

	"cpf"
)

var Data = cpf.CustomResources{
	DataSourcesMap: []cpf.ResourceMap{
		//hcp.CoffeeDataSource,
		git.WorkflowDataSource,
	},
}
