package cpf

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type SchemaMap map[string]*schema.Schema

type ResourceMap map[string]*schema.Resource

type CustomSchema struct {
	Schemas []SchemaMap
}

type CustomResources struct {
	ResourcesMaps  []ResourceMap
	DataSourcesMap []ResourceMap
}

type CustomConfigureContextFunc struct {
	ConfigureContextFunc schema.ConfigureContextFunc
}

func (c *CustomSchema) Schema() SchemaMap {
	var schema = make(SchemaMap)
	for _, v := range c.Schemas {
		for x, y := range v {
			if _, ok := schema[x]; !ok {
				schema[x] = y
			}
		}
	}
	return schema
}

func (c *CustomResources) Resources() ResourceMap {
	var resource = make(ResourceMap)
	for _, v := range c.ResourcesMaps {
		for x, y := range v {
			if _, ok := resource[x]; !ok {
				resource[x] = y
			}
		}
	}
	return resource
}

func (c *CustomResources) DataSources() ResourceMap {
	var resource = make(ResourceMap)
	for _, v := range c.DataSourcesMap {
		for x, y := range v {
			if _, ok := resource[x]; !ok {
				resource[x] = y
			}
		}
	}
	return resource
}
