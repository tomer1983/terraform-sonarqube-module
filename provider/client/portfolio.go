package client

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
)

type Portfolio struct {
	Key         string             `json:"key"`
	Name        string             `json:"name"`
	Description string             `json:"description,omitempty"`
	Selection   PortfolioSelection `json:"selection"`
	Filters     PortfolioFilters   `json:"filters,omitempty"`
}

type PortfolioSelection struct {
	Mode           string   `json:"mode"`
	Projects       []string `json:"projects,omitempty"`
	ProjectPattern string   `json:"projectPattern,omitempty"`
	BranchPattern  string   `json:"branchPattern,omitempty"`
}

type PortfolioFilters struct {
	Languages    []string                    `json:"languages,omitempty"`
	Tags         []string                    `json:"tags,omitempty"`
	QualityGates []string                    `json:"qualityGates,omitempty"`
	Compliance   PortfolioCompliance         `json:"compliance,omitempty"`
	Metrics      map[string]PortfolioMetric  `json:"metrics,omitempty"`
}

type PortfolioCompliance struct {
	MinQualityGateStatus string   `json:"minQualityGateStatus,omitempty"`
	MinCoverage         float64  `json:"minCoverage,omitempty"`
	MaxDuplications     float64  `json:"maxDuplications,omitempty"`
	MaxIssues          int      `json:"maxIssues,omitempty"`
	RequiredRules      []string `json:"requiredRules,omitempty"`
}

type PortfolioMetric struct {
	Operator string `json:"operator"`
	Value    string `json:"value"`
}

func (c *Client) CreatePortfolio(portfolio *Portfolio) error {
	params := url.Values{}
	params.Set("key", portfolio.Key)
	params.Set("name", portfolio.Name)
	
	if portfolio.Description != "" {
		params.Set("description", portfolio.Description)
	}

	resp, err := c.doRequest("POST", "portfolios/create", params)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Configure selection mode and filters
	if err := c.configurePortfolioSelection(portfolio.Key, &portfolio.Selection); err != nil {
		return err
	}

	if portfolio.Selection.Mode == "FILTER" {
		if err := c.configurePortfolioFilters(portfolio.Key, &portfolio.Filters); err != nil {
			return err
		}
	}

	return nil
}

func (c *Client) UpdatePortfolio(portfolio *Portfolio) error {
	params := url.Values{}
	params.Set("key", portfolio.Key)
	params.Set("name", portfolio.Name)
	
	if portfolio.Description != "" {
		params.Set("description", portfolio.Description)
	}

	resp, err := c.doRequest("POST", "portfolios/update", params)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Update selection and filters
	if err := c.configurePortfolioSelection(portfolio.Key, &portfolio.Selection); err != nil {
		return err
	}

	if portfolio.Selection.Mode == "FILTER" {
		if err := c.configurePortfolioFilters(portfolio.Key, &portfolio.Filters); err != nil {
			return err
		}
	}

	return nil
}

func (c *Client) DeletePortfolio(key string) error {
	params := url.Values{}
	params.Set("key", key)

	resp, err := c.doRequest("POST", "portfolios/delete", params)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}

func (c *Client) GetPortfolio(key string) (*Portfolio, error) {
	params := url.Values{}
	params.Set("key", key)

	resp, err := c.doRequest("GET", "portfolios/show", params)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var portfolio Portfolio
	if err := json.NewDecoder(resp.Body).Decode(&portfolio); err != nil {
		return nil, err
	}

	return &portfolio, nil
}

func (c *Client) configurePortfolioSelection(key string, selection *PortfolioSelection) error {
	params := url.Values{}
	params.Set("key", key)
	params.Set("mode", selection.Mode)

	switch selection.Mode {
	case "MANUAL":
		if len(selection.Projects) > 0 {
			params.Set("projects", strings.Join(selection.Projects, ","))
		}
	case "REGEXP":
		if selection.ProjectPattern != "" {
			params.Set("projectPattern", selection.ProjectPattern)
		}
		if selection.BranchPattern != "" {
			params.Set("branchPattern", selection.BranchPattern)
		}
	}

	resp, err := c.doRequest("POST", "portfolios/configure_selection", params)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}

func (c *Client) configurePortfolioFilters(key string, filters *PortfolioFilters) error {
	params := url.Values{}
	params.Set("key", key)

	if len(filters.Languages) > 0 {
		params.Set("languages", strings.Join(filters.Languages, ","))
	}
	if len(filters.Tags) > 0 {
		params.Set("tags", strings.Join(filters.Tags, ","))
	}
	if len(filters.QualityGates) > 0 {
		params.Set("qualityGates", strings.Join(filters.QualityGates, ","))
	}

	// Add compliance settings
	if filters.Compliance.MinQualityGateStatus != "" {
		params.Set("minQualityGateStatus", filters.Compliance.MinQualityGateStatus)
	}
	if filters.Compliance.MinCoverage > 0 {
		params.Set("minCoverage", fmt.Sprintf("%.2f", filters.Compliance.MinCoverage))
	}
	if filters.Compliance.MaxDuplications > 0 {
		params.Set("maxDuplications", fmt.Sprintf("%.2f", filters.Compliance.MaxDuplications))
	}
	if filters.Compliance.MaxIssues > 0 {
		params.Set("maxIssues", fmt.Sprintf("%d", filters.Compliance.MaxIssues))
	}
	if len(filters.Compliance.RequiredRules) > 0 {
		params.Set("requiredRules", strings.Join(filters.Compliance.RequiredRules, ","))
	}

	// Add custom metrics
	for metric, value := range filters.Metrics {
		params.Set(fmt.Sprintf("metric_%s_operator", metric), value.Operator)
		params.Set(fmt.Sprintf("metric_%s_value", metric), value.Value)
	}

	resp, err := c.doRequest("POST", "portfolios/configure_filters", params)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}
