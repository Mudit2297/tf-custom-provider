package todoist_schema

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Project Schema for Data source
// func Projects() *schema.Schema {
// 	return &schema.Schema{
// 		Type:     schema.TypeList,
// 		Computed: true,
// 		Elem: &schema.Resource{
// 			Schema: map[string]*schema.Schema{
// 				"id":     DataProjectId(),
// 				"name":   DataProjectName(),
// 				"shared": DataProjectShared(),
// 				"url":    DataProjectURL(),
// 			},
// 		},
// 	}
// }

func WorkflowId() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeInt,
		Computed: true,
	}
}

func WorkflowNodeId() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeString,
		Computed: true,
	}
}

// func WorkflowName() *schema.Schema {
// 	return &schema.Schema{
// 		Type:     schema.TypeString,
// 		Computed: true,
// 	}
// }

func WorkflowPath() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeString,
		Computed: true,
	}
}

func WorkflowState() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeString,
		Computed: true,
	}
}

func WorkflowCreatedAt() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeString,
		Computed: true,
	}
}

func WorkflowUpdatedAt() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeString,
		Computed: true,
	}
}

func WorkflowURL() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeString,
		Computed: true,
	}
}

func WorkflowHtmlURL() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeString,
		Computed: true,
	}
}

func WorkflowBadgeURL() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeString,
		Computed: true,
	}
}
