package common

import (
	git "cpf/providers/terraform-provider-github"
	//hcp "cpf/providers/terraform-provider-hashicups"

	"cpf"
)

var Data = cpf.CustomResources{
	DataSourcesMap: []cpf.ResourcMap{
		//hcp.CoffeeDataSource,
		git.WorkflowDataSource,
	},
}
