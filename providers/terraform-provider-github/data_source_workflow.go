package github

import (
	"context"
	"cpf"
	"cpf/providers/client"
	"fmt"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var workflow = cpf.SchemaMap{
	"id":                 cpf.TypeIntComputed(),
	"path":               cpf.TypeStringComputed(),
	"state":              cpf.TypeStringComputed(),
	"html_url":           cpf.TypeStringComputed(),
	"workflow_file_name": cpf.TypeStringRequired(),
	"owner":              cpf.TypeStringRequired(),
	"repo":               cpf.TypeStringRequired(),
}

var WorkflowDataSource = cpf.ResourceMap{
	"cpf_git_workflow": dataSourceGitWorkflow(),
}

var dataWorkflow = cpf.CustomSchema{
	Schemas: []cpf.SchemaMap{workflow},
}

func dataSourceGitWorkflow() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceGitWorkflowRead,
		Schema:      dataWorkflow.Schema(),
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

	err := c.GetGitWorkflowByName()
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
