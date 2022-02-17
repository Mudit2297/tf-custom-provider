package cpf

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func TypeIntComputed() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeInt,
		Computed: true,
	}
}

func TypeIntRequired() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeInt,
		Required: true,
	}
}

func TypeIntRequiredForceNew() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeInt,
		Required: true,
		ForceNew: true,
	}
}

func TypeIntOptional() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeInt,
		Optional: true,
	}
}

func TypeStringComputed() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeString,
		Computed: true,
	}
}

func TypeStringRequired() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
}

func TypeStringRequiredForceNew() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
		ForceNew: true,
	}
}

func TypeStringOptional() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
	}
}

func TypeBoolComputed() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeBool,
		Computed: true,
	}
}

func TypeBoolRequired() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeBool,
		Required: true,
	}
}

func TypeBoolRequiredForceNew() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeBool,
		Required: true,
		ForceNew: true,
	}
}

func TypeNullComputed() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeString,
		Computed: true,
	}
}

func TypeEmptyListComputed() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeList,
		Computed: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{},
		},
	}
}

func TypeEmptyListRequired() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeList,
		Required: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{},
		},
	}
}
