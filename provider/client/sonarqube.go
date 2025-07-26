package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
	"strings"
)

// Project represents a SonarQube project
type Project struct {
	Key         string   `json:"key"`
	Name        string   `json:"name"`
	Visibility  string   `json:"visibility"`
	MainBranch  string   `json:"mainBranch"`
	Tags        []string `json:"tags"`
	Qualifier   string   `json:"qualifier"`
	Description string   `json:"description,omitempty"`
}

// QualityGate represents a SonarQube quality gate
type QualityGate struct {
	ID         string      `json:"id"`
	Name       string      `json:"name"`
	Conditions []Condition `json:"conditions"`
	IsDefault  bool       `json:"isDefault"`
}

type Condition struct {
	ID       string `json:"id"`
	Metric   string `json:"metric"`
	Op       string `json:"op"`
	Error    string `json:"error"`
	Warning  string `json:"warning,omitempty"`
}

// Project API Methods
func (c *Client) CreateProject(name, key, visibility string, mainBranch string, tags []string) (*Project, error) {
	params := url.Values{}
	params.Set("name", name)
	params.Set("project", key)
	params.Set("visibility", visibility)
	
	if mainBranch != "" {
		params.Set("mainBranch", mainBranch)
	}
	
	if len(tags) > 0 {
		params.Set("tags", strings.Join(tags, ","))
	}

	resp, err := c.doRequest("POST", "projects/create", params)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var project Project
	if err := json.NewDecoder(resp.Body).Decode(&project); err != nil {
		return nil, err
	}

	return &project, nil
}

func (c *Client) ReadProject(key string) (*Project, error) {
	params := url.Values{}
	params.Set("project", key)

	resp, err := c.doRequest("GET", "projects/search", params)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result struct {
		Components []Project `json:"components"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if len(result.Components) == 0 {
		return nil, fmt.Errorf("project not found: %s", key)
	}

	return &result.Components[0], nil
}

func (c *Client) UpdateProject(key string, name string, visibility string, tags []string) (*Project, error) {
	params := url.Values{}
	params.Set("project", key)
	
	if name != "" {
		params.Set("name", name)
	}
	
	if visibility != "" {
		params.Set("visibility", visibility)
	}
	
	if len(tags) > 0 {
		params.Set("tags", strings.Join(tags, ","))
	}

	resp, err := c.doRequest("POST", "projects/update", params)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var project Project
	if err := json.NewDecoder(resp.Body).Decode(&project); err != nil {
		return nil, err
	}

	return &project, nil
}

func (c *Client) DeleteProject(key string) error {
	params := url.Values{}
	params.Set("project", key)

	resp, err := c.doRequest("POST", "projects/delete", params)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}

// Quality Gate API Methods
func (c *Client) CreateQualityGate(name string) (*QualityGate, error) {
	params := url.Values{}
	params.Set("name", name)

	resp, err := c.doRequest("POST", "qualitygates/create", params)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var gate QualityGate
	if err := json.NewDecoder(resp.Body).Decode(&gate); err != nil {
		return nil, err
	}

	return &gate, nil
}

func (c *Client) CreateQualityGateCondition(gateID, metric, op, error string) (*Condition, error) {
	params := url.Values{}
	params.Set("gateId", gateID)
	params.Set("metric", metric)
	params.Set("op", op)
	params.Set("error", error)

	resp, err := c.doRequest("POST", "qualitygates/create_condition", params)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var condition Condition
	if err := json.NewDecoder(resp.Body).Decode(&condition); err != nil {
		return nil, err
	}

	return &condition, nil
}

func (c *Client) ReadQualityGate(id string) (*QualityGate, error) {
	params := url.Values{}
	params.Set("id", id)

	resp, err := c.doRequest("GET", "qualitygates/show", params)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var gate QualityGate
	if err := json.NewDecoder(resp.Body).Decode(&gate); err != nil {
		return nil, err
	}

	return &gate, nil
}

func (c *Client) UpdateQualityGate(id, name string) (*QualityGate, error) {
	params := url.Values{}
	params.Set("id", id)
	params.Set("name", name)

	resp, err := c.doRequest("POST", "qualitygates/update", params)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var gate QualityGate
	if err := json.NewDecoder(resp.Body).Decode(&gate); err != nil {
		return nil, err
	}

	return &gate, nil
}

func (c *Client) DeleteQualityGate(id string) error {
	params := url.Values{}
	params.Set("id", id)

	resp, err := c.doRequest("POST", "qualitygates/delete", params)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}
