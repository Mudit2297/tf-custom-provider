package github

import (
	"context"
	s "git/custom-provider-framework/genSchemas"
	git "git/github-client-go"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGitWorkflow() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceWorkflowCreate,
		ReadContext:   resourceWorkflowRead,
		UpdateContext: resourceWorkflowCreate,
		DeleteContext: resourceWorkflowDelete,
		Schema: map[string]*schema.Schema{
			"ref_branch":          s.StringRequiredForceNewSchema(),
			"ref_tag":             s.StringOptionalSchema(),
			"workflow_file_name":  s.StringRequiredForceNewSchema(),
			"owner":               s.StringRequiredForceNewSchema(),
			"repo":                s.StringRequiredForceNewSchema(),
			"run_id":              s.IntComputedSchema(),
			"run_name":            s.StringComputedSchema(),
			"run_number":          s.IntComputedSchema(),
			"event":               s.StringComputedSchema(),
			"status":              s.StringComputedSchema(),
			"conclusion":          s.StringComputedSchema(),
			"workflow_id":         s.IntComputedSchema(),
			"html_url":            s.StringComputedSchema(),
			"total_workflow_runs": s.IntComputedSchema(),
			"updated_at":          s.StringComputedSchema(),
		},
	}
}

func resourceWorkflowCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*git.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	var payload git.WorkflowDispatchPayload

	orgOwner := d.Get("owner").(string)
	repo := d.Get("repo").(string)
	workflowName := d.Get("workflow_file_name").(string)
	refBranch := d.Get("ref_branch")
	refTag := d.Get("ref_tag")
	if refBranch != "" {
		payload = git.WorkflowDispatchPayload{
			Ref: refBranch.(string),
		}
	} else if refTag != "" {
		payload = git.WorkflowDispatchPayload{
			Ref: refTag.(string),
		}
	} else {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Resource Creation Failed",
			Detail:   "Atleast one of ref_branch or ref_tag is required",
		})
		return diags
	}

	p, err := c.CreateWorkflowByName(orgOwner, repo, workflowName, payload)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(strconv.Itoa(p.WorkflowRuns[0].ID))
	resourceWorkflowRead(ctx, d, m)
	return diags
}

func resourceWorkflowRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*git.Client)
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	orgOwner := d.Get("owner").(string)
	repo := d.Get("repo").(string)
	workflowName := d.Get("workflow_file_name").(string)

	p, err := c.GetWorkflowRunsByName(orgOwner, repo, workflowName)
	if err != nil {
		return diag.FromErr(err)
	}

	refBranch := d.Get("ref_branch")
	refTag := d.Get("ref_tag")
	if refBranch != nil {
		d.Set("ref_branch", p.WorkflowRuns[0].HeadBranch)
		d.Set("ref_tag", nil)
	} else if refTag != nil {
		d.Set("ref_tag", p.WorkflowRuns[0].HeadBranch)
		d.Set("ref_branch", nil)
	}
	d.Set("owner", orgOwner)
	d.Set("workflow_file_name", workflowName)
	d.Set("repo", repo)
	d.Set("run_id", p.WorkflowRuns[0].ID)
	d.Set("run_name", p.WorkflowRuns[0].Name)
	d.Set("run_number", p.WorkflowRuns[0].RunNumber)
	d.Set("event", p.WorkflowRuns[0].Event)
	d.Set("status", p.WorkflowRuns[0].Status)
	d.Set("conclusion", p.WorkflowRuns[0].Conclusion)
	d.Set("workflow_id", p.WorkflowRuns[0].WorkflowID)
	d.Set("html_url", p.WorkflowRuns[0].HTMLURL)
	d.Set("total_workflow_runs", p.TotalCount)

	// runs := flattenWorkflowRuns(p)
	// if err := d.Set("workflow_runs", runs); err != nil {
	// 	return diag.FromErr(err)
	// }

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

// func flattenWorkflowRuns(runs *git.WorkflowRuns) []interface{} {
// 	if runs != nil {
// 		fRuns := make([]interface{}, len(*&runs.WorkflowRuns), len(*&runs.WorkflowRuns))

// 		for i, run := range runs.WorkflowRuns {
// 			c := make(map[string]interface{})
// 			c["id"] = run.ID
// 			c["name"] = run.Name
// 			c["run_number"] = run.RunNumber
// 			c["event"] = run.Event
// 			c["status"] = run.Status
// 			c["conclusion"] = run.Conclusion
// 			c["workflow_id"] = run.WorkflowID
// 			c["html_url"] = run.HTMLURL

// 			fRuns[i] = c
// 		}

// 		return fRuns
// 	}

// 	return make([]interface{}, 0)
// }

// func flattenCoffee(run git.WorkflowRun) []interface{} {
// 	c := make(map[string]interface{})

// 	c["id"] = run.ID
// 	c["name"] = run.Name
// 	c["run_number"] = run.RunNumber
// 	c["event"] = run.Event
// 	c["status"] = run.Status
// 	c["conclusion"] = run.Conclusion
// 	c["workflow_id"] = run.WorkflowID
// 	c["html_url"] = run.HTMLURL
// 	// c["repository"] = run.Repository.Name
// 	// c["conclusion"] = run.Conclusion
// 	// c["conclusion"] = run.Conclusion
// 	// c["conclusion"] = run.Conclusion
// 	// c["conclusion"] = run.Conclusion
// 	// c["conclusion"] = run.Conclusion

// 	return []interface{}{c}
// }
