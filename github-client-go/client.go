package github

import (
	"fmt"
	"net/http"
	"time"
)

// HostURL - Default Hashicups URL
const HostURL string = "https://api.github.com/users/"

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
		HostURL: fmt.Sprintf("%s/%s", HostURL, *username),
		Auth: AuthStruct{
			Username: *username,
			Token:    *token,
		},
	}
	return &c, nil
}
