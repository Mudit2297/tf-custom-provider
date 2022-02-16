package framework

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func IntComputedSchema() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeInt,
		Computed: true,
	}
}

func IntRequiredSchema() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeInt,
		Required: true,
	}
}

func IntRequiredForceNewSchema() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeInt,
		Required: true,
		ForceNew: true,
	}
}

func IntOptionalSchema() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeInt,
		Optional: true,
	}
}

func StringComputedSchema() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeString,
		Computed: true,
	}
}

func StringRequiredSchema() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
}

func StringRequiredForceNewSchema() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
		ForceNew: true,
	}
}

func StringOptionalSchema() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
	}
}

func BoolComputedSchema() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeBool,
		Computed: true,
	}
}

func BoolRequiredSchema() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeBool,
		Required: true,
	}
}

func BoolRequiredForceNewSchema() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeBool,
		Required: true,
		ForceNew: true,
	}
}

func NullComputedSchema() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeString,
		Computed: true,
	}
}

func EmptyListComputedSchema() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeList,
		Computed: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{},
		},
	}
}

func EmptyListRequiredSchema() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeList,
		Required: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{},
		},
	}
}
