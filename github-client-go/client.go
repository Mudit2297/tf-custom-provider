package github

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

// HostURL - Default Hashicups URL
const HostURL string = "https://api.github.com"

// Client -
type Client struct {
	HTTPClient *http.Client
	Auth       AuthStruct
	HostURL    string
}

// AuthStruct -
type AuthStruct struct {
	Username string
	Token    string
}

// NewClient -
func NewClient(username, token *string) (*Client, error) {
	c := Client{
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
		// Default Hashicups URL
		//HostURL: HostURL + *username,
		HostURL: fmt.Sprintf("%s/%s/%s", HostURL, "users", *username),
		Auth: AuthStruct{
			Username: *username,
			Token:    *token,
		},
	}
	return &c, nil
}

func (c *Client) newRequest(path, method string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, fmt.Sprintf("%s/%s", HostURL, path), body)
	if err != nil {
		return nil, err
	}
	return req, err
}

func (c *Client) doRequest(req *http.Request) ([]byte, error) {
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.Auth.Token))
	req.Header.Set("Accept", "application/vnd.github.v3+json")
	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	if res.StatusCode == http.StatusOK || res.StatusCode == http.StatusNoContent {
		return body, err
	} else {
		return nil, fmt.Errorf("status: %d, body: %s", res.StatusCode, body)
	}
}
