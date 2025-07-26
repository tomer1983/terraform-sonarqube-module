// Package client provides a SonarQube API client
package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/hashicorp/go-retryablehttp"
)

type Client struct {
	host      string
	token     string
	client    *retryablehttp.Client
}

func NewClient(host, token string) *Client {
	retryClient := retryablehttp.NewClient()
	retryClient.RetryMax = 3

	return &Client{
		host:      host,
		token:     token,
		client:    retryClient,
	}
}

// Base API calls
func (c *Client) doRequest(method, path string, body interface{}) (*http.Response, error) {
	url := fmt.Sprintf("%s/api/%s", c.host, path)
	
	req, err := retryablehttp.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.token))
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("API request failed with status %d", resp.StatusCode)
	}

	return resp, nil
}

// Project API calls
func (c *Client) CreateProject(name, key string) error {
	path := "projects/create"
	body := map[string]string{
		"name": name,
		"project": key,
	}

	_, err := c.doRequest("POST", path, body)
	return err
}

func (c *Client) DeleteProject(key string) error {
	path := fmt.Sprintf("projects/delete?project=%s", key)
	_, err := c.doRequest("POST", path, nil)
	return err
}

// Add more API methods for other resources (quality gates, users, groups, etc.)
