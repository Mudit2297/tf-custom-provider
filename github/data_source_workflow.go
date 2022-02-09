package github

import (
	"context"
	git "git/github-client-go"
	gitsch "git/github/git_schemas"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceGitWorkflow() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceGitWorkflowRead,
		Schema: map[string]*schema.Schema{
			"id":         gitsch.WorkflowId(),
			"node_id":    gitsch.WorkflowNodeId(),
			"path":       gitsch.WorkflowPath(),
			"state":      gitsch.WorkflowState(),
			"created_at": gitsch.WorkflowCreatedAt(),
			"updated_at": gitsch.WorkflowUpdatedAt(),
			"url":        gitsch.WorkflowURL(),
			"html_url":   gitsch.WorkflowHtmlURL(),
			"badge_url":  gitsch.WorkflowBadgeURL(),
			"workflow_file_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"owner": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"repo": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
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
	d.Set("url", p.URL)
	d.Set("workflow_file_name", p.Name)
	d.Set("id", p.ID)
	d.Set("state", p.State)
	d.SetId(strconv.Itoa(p.ID))
	return diags
}
