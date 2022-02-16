package github

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"cpf/providers/client"

	"cpf"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var workflow = cpf.SchemaMap{
	// "total_count": cpf.TypeIntComputed(),
	// "workflows": &schema.Schema{
	// 	Type:     schema.TypeList,
	// 	Computed: true,
	// 	Elem: &schema.Resource{
	// 		Schema: cpf.SchemaMap{
	// 			"id":    cpf.TypeIntComputed(),
	// 			"name":  cpf.TypeStringComputed(),
	// 			"path":  cpf.TypeStringComputed(),
	// 			"state": cpf.TypeStringComputed(),
	// 			"url":   cpf.TypeStringComputed(),
	// 			"owner": cpf.TypeStringRequired(),
	// 			"repo":  cpf.TypeStringRequired(),
	// 		},
	// 	},
	// },
	"id":                 cpf.TypeIntComputed(),
	"path":               cpf.TypeStringComputed(),
	"state":              cpf.TypeStringComputed(),
	"html_url":           cpf.TypeStringComputed(),
	"workflow_file_name": cpf.TypeStringRequired(),
	"owner":              cpf.TypeStringRequired(),
	"repo":               cpf.TypeStringRequired(),
}

var WorkflowDataSource = cpf.ResourcMap{
	"cpf_git_workflow": dataSourceGitWorkflow(),
}

var sch = cpf.CustomSchema{
	Schemas: []cpf.SchemaMap{workflow},
}

func dataSourceGitWorkflow() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceGitWorkflowRead,
		Schema:      sch.Schema(),
	}
}

func dataSourceGitWorkflowRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.Client)

	var diags diag.Diagnostics

	orgOwner := d.Get("owner").(string)
	repo := d.Get("repo").(string)
	workflowName := d.Get("workflow_file_name").(string)

	c.URL = fmt.Sprintf("https://api.github.com/repos/%v/%v/actions/workflows/%v", orgOwner, repo, workflowName)
	c.Method = "GET"

	err := c.GetGitWorkflows()
	if err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("html_url", c.Workflow.HTMLURL); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("state", c.Workflow.State); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("path", c.Workflow.Path); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("owner", orgOwner); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("repo", repo); err != nil {
		return diag.FromErr(err)
	}

	// always run
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}
