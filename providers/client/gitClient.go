package client

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"golang.org/x/oauth2"
)

func (c *Client) GitHubNewClient(pat string) *Client {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: pat},
	)
	c.HttpClient = oauth2.NewClient(ctx, ts)
	// return c.HttpClient
	return c
}

func (c *Client) GetGitWorkflowByName() error {
	c.NewRequest()
	res, err := c.DoRequest()
	if err != nil {
		return fmt.Errorf("Error in executing request: %v", err)
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("Error in reading response body: %v", err)
	}
	if res.StatusCode == http.StatusOK || res.StatusCode == http.StatusNoContent {
		err = json.Unmarshal(body, &c.Workflow)
		if err != nil {
			return fmt.Errorf("Error during do unmarshal: %v", err)
		}
	} else {
		return fmt.Errorf("status: %d, body: %s", res.StatusCode, body)
	}

	return nil
}

func (c *Client) CreateGitWorkflowDispatch() error {
	c.NewRequest()
	res, err := c.DoRequest()
	if err != nil {
		return fmt.Errorf("Error in executing request: %v", err)
	}
	time.Sleep(5 * time.Second)
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("Error in reading response body: %v", err)
	}
	if res.StatusCode == http.StatusOK || res.StatusCode == http.StatusNoContent {
		return nil

	} else {
		return fmt.Errorf("status: %d, body: %s", res.StatusCode, body)
	}
}

func (c *Client) GetGitWorkflowRunsByName() error {
	c.NewRequest()
	res, err := c.DoRequest()
	if err != nil {
		return fmt.Errorf("Error in executing request: %v", err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("Error in reading response body: %v", err)
	}

	if res.StatusCode == http.StatusOK || res.StatusCode == http.StatusNoContent {
		err = json.Unmarshal(body, &c.WorkflowRuns)
		if err != nil {
			return fmt.Errorf("Error during do unmarshal: %v", err)
		}
		if c.WorkflowRuns.WorkflowRuns[0].Conclusion == "" {
			time.Sleep(5 * time.Second)
			c.GetGitWorkflowRunsByName()
		}
	} else {
		return fmt.Errorf("status: %d, body: %s", res.StatusCode, body)
	}

	return nil
}
