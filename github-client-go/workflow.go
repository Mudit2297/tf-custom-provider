package github

import (
	"encoding/json"
	"fmt"
	"strings"
)

func (c *Client) GetWorkflowByName(orgOwner string, repo string, workflowName string) (*Workflow, error) {
	path := "repos/" + orgOwner + "/" + repo + "/actions/workflows/" + workflowName
	req, err := c.newRequest(path, "GET", nil)
	if err != nil {
		fmt.Println("Error during request:")
		return nil, err
	}
	res, err := c.doRequest(req)
	if err != nil {
		fmt.Println("Error during do request:")
		return nil, err
	}
	var workflow Workflow
	err = json.Unmarshal(res, &workflow)
	if err != nil {
		fmt.Println("Error during do unmarshel:")
		return nil, err
	}
	return &workflow, nil
}

func (c *Client) CreateWorkflowByName(orgOwner string, repo string, workflowName string, payload WorkflowDispatchPayload) (*WorkflowRuns, error) {
	dispatchPath := "repos/" + orgOwner + "/" + repo + "/actions/workflows/" + workflowName + "/dispatches"
	pl, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	dispatchReq, err := c.newRequest(dispatchPath, "POST", strings.NewReader(string(pl)))
	if err != nil {
		fmt.Println("Error during request:")
		return nil, err
	}
	temp, err := c.doRequest(dispatchReq)
	if err != nil {
		fmt.Println("Error during do request:")
		return nil, err
	}
	fmt.Println(temp)

	p, err := c.GetWorkflowRunsByName(orgOwner, repo, workflowName)
	if err != nil {
		fmt.Println("Error during do request:")
	}

	return p, nil
}

func (c *Client) GetWorkflowRunsByName(orgOwner string, repo string, workflowName string) (*WorkflowRuns, error) {
	runPath := "repos/" + orgOwner + "/" + repo + "/actions/workflows/" + workflowName + "/runs"
	getDispatchReq, err := c.newRequest(runPath, "GET", nil)
	if err != nil {
		fmt.Println("Error during request:")
		return nil, err
	}
	getDispatchRes, err := c.doRequest(getDispatchReq)
	if err != nil {
		fmt.Println("Error during do request:")
		return nil, err
	}

	var workflowRuns WorkflowRuns
	err = json.Unmarshal(getDispatchRes, &workflowRuns)
	if err != nil {
		fmt.Println("Error during do unmarshel:")
		return nil, err
	}
	return &workflowRuns, nil
}
