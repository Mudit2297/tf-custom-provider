package github

import (
	"context"
	"cpf"
	"cpf/providers/client"
	models "cpf/providers/data/github"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var workflowDispatch = cpf.SchemaMap{
	"ref_branch":          cpf.TypeStringRequiredForceNew(),
	"workflow_file_name":  cpf.TypeStringRequiredForceNew(),
	"owner":               cpf.TypeStringRequired(),
	"repo":                cpf.TypeStringRequiredForceNew(),
	"inputs":              cpf.TypeStringOptionalForceNew(),
	"run_id":              cpf.TypeIntComputed(),
	"run_name":            cpf.TypeStringComputed(),
	"run_number":          cpf.TypeIntComputed(),
	"event":               cpf.TypeStringComputed(),
	"status":              cpf.TypeStringComputed(),
	"conclusion":          cpf.TypeStringComputed(),
	"workflow_id":         cpf.TypeIntComputed(),
	"html_url":            cpf.TypeStringComputed(),
	"total_workflow_runs": cpf.TypeIntComputed(),
	"commiter_name":       cpf.TypeStringComputed(),
	"commiter_email":      cpf.TypeStringComputed(),
	"respository_url":     cpf.TypeStringComputed(),
	"repository_fullname": cpf.TypeStringComputed(),
}

var WorkflowDispatchResource = cpf.ResourceMap{
	"cpf_git_workflow_dispatch": resourceGitWorkflow(),
}

var resWorkflowDispatch = cpf.CustomSchema{
	Schemas: []cpf.SchemaMap{workflowDispatch},
}

func resourceGitWorkflow() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceWorkflowCreate,
		ReadContext:   resourceWorkflowRead,
		UpdateContext: resourceWorkflowCreate,
		DeleteContext: resourceWorkflowDelete,
		Schema:        resWorkflowDispatch.Schema(),
	}
}

func resourceWorkflowCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	var payload models.WorkflowDispatchPayload
	orgOwner := d.Get("owner").(string)
	repo := d.Get("repo").(string)
	workflowName := d.Get("workflow_file_name").(string)
	refBranch := d.Get("ref_branch")
	inputs := d.Get("inputs")

	// Declared empty interface
	data := make(map[string]interface{})

	if inputs != "" {
		json.Unmarshal([]byte(inputs.(string)), &data)
		payload = models.WorkflowDispatchPayload{
			Ref:    refBranch.(string),
			Inputs: data,
		}
	} else {
		payload = models.WorkflowDispatchPayload{
			Ref:    refBranch.(string),
			Inputs: data,
		}
	}

	pl, err := json.Marshal(payload)
	if err != nil {
		return diag.FromErr(err)
	}

	c.URL = fmt.Sprintf("https://api.github.com/repos/%v/%v/actions/workflows/%v/dispatches", orgOwner, repo, workflowName)
	c.Method = "POST"
	c.Body = strings.NewReader(string(pl))

	err = c.CreateGitWorkflowDispatch()
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))
	resourceWorkflowRead(ctx, d, m)
	return diags
}

func resourceWorkflowRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.Client)
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	orgOwner := d.Get("owner").(string)
	repo := d.Get("repo").(string)
	workflowName := d.Get("workflow_file_name").(string)

	c.URL = fmt.Sprintf("https://api.github.com/repos/%v/%v/actions/workflows/%v/runs", orgOwner, repo, workflowName)
	c.Method = "GET"

	err := c.GetGitWorkflowRunsByName()
	if err != nil {
		return diag.FromErr(err)
	}

	d.Set("ref_branch", c.WorkflowRuns.WorkflowRuns[0].HeadBranch)
	d.Set("owner", orgOwner)
	d.Set("workflow_file_name", workflowName)
	d.Set("repo", repo)
	d.Set("run_id", c.WorkflowRuns.WorkflowRuns[0].ID)
	d.Set("run_name", c.WorkflowRuns.WorkflowRuns[0].Name)
	d.Set("run_number", c.WorkflowRuns.WorkflowRuns[0].RunNumber)
	d.Set("event", c.WorkflowRuns.WorkflowRuns[0].Event)
	d.Set("status", c.WorkflowRuns.WorkflowRuns[0].Status)
	d.Set("conclusion", c.WorkflowRuns.WorkflowRuns[0].Conclusion)
	d.Set("workflow_id", c.WorkflowRuns.WorkflowRuns[0].WorkflowID)
	d.Set("html_url", c.WorkflowRuns.WorkflowRuns[0].HTMLURL)
	d.Set("total_workflow_runs", c.WorkflowRuns.TotalCount)
	d.Set("commiter_name", c.WorkflowRuns.WorkflowRuns[0].HeadCommit.Committer.Name)
	d.Set("commiter_email", c.WorkflowRuns.WorkflowRuns[0].HeadCommit.Committer.Email)
	d.Set("respository_url", c.WorkflowRuns.WorkflowRuns[0].Repository.HTMLURL)
	d.Set("repository_fullname", c.WorkflowRuns.WorkflowRuns[0].Repository.FullName)

	return diags
}

func resourceWorkflowUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	d.SetId("")
	diags := resourceWorkflowCreate(ctx, d, m)
	return diags
}

func resourceWorkflowDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	d.SetId("")
	return diags
}
