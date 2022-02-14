package github

import (
	"context"
	s "git/custom-provider-framework/genSchemas"
	git "git/github-client-go"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceGitWorkflow() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceGitWorkflowRead,
		Schema: map[string]*schema.Schema{
			"id":                 s.IntComputedSchema(),
			"path":               s.StringComputedSchema(),
			"state":              s.StringComputedSchema(),
			"html_url":           s.StringComputedSchema(),
			"workflow_file_name": s.StringRequiredSchema(),
			"owner":              s.StringRequiredSchema(),
			"repo":               s.StringRequiredSchema(),
		},
	}
}

func dataSourceGitWorkflowRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*git.Client)
	var diags diag.Diagnostics
	orgOwner := d.Get("owner").(string)
	repo := d.Get("repo").(string)
	workflowName := d.Get("workflow_file_name").(string)
	p, err := c.GetWorkflowByName(orgOwner, repo, workflowName)
	if err != nil {
		return diag.FromErr(err)
	}
	d.Set("html_url", p.HTMLURL)
	d.Set("workflow_file_name", p.Name)
	d.Set("id", p.ID)
	d.Set("path", p.Path)
	d.Set("state", p.State)
	d.SetId(strconv.Itoa(p.ID))
	return diags
}
