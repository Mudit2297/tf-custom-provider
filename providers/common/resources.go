package common

import (
	"cpf"
	git "cpf/providers/terraform-provider-github"
)

var Resource = cpf.CustomResources{
	ResourcesMaps: []cpf.ResourceMap{
		git.WorkflowDispatchResource,
	},
}
