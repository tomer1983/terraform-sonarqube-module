package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSonarqubePortfolio() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourcePortfolioCreate,
		ReadContext:   resourcePortfolioRead,
		UpdateContext: resourcePortfolioUpdate,
		DeleteContext: resourcePortfolioDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"key": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"selection_mode": {
				Type:     schema.TypeString,
				Required: true,
				ValidateFunc: func(v interface{}, k string) (ws []string, errors []error) {
					value := v.(string)
					if value != "MANUAL" && value != "REGEXP" && value != "FILTER" {
						errors = append(errors, fmt.Errorf("selection_mode must be one of MANUAL, REGEXP, or FILTER"))
					}
					return
				},
			},
			"projects": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"project_pattern": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"branch_pattern": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"filters": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"languages": {
							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"tags": {
							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"quality_gates": {
							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"compliance": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"min_quality_gate_status": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"min_coverage": {
										Type:     schema.TypeFloat,
										Optional: true,
									},
									"max_duplications": {
										Type:     schema.TypeFloat,
										Optional: true,
									},
									"max_issues": {
										Type:     schema.TypeInt,
										Optional: true,
									},
									"required_rules": {
										Type:     schema.TypeSet,
										Optional: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},
						"custom_metrics": {
							Type:     schema.TypeMap,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"operator": {
										Type:     schema.TypeString,
										Required: true,
									},
									"value": {
										Type:     schema.TypeString,
										Required: true,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func resourcePortfolioCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*client.Client)

	portfolio := &client.Portfolio{
		Key:         d.Get("key").(string),
		Name:        d.Get("name").(string),
		Description: d.Get("description").(string),
		Selection: client.PortfolioSelection{
			Mode: d.Get("selection_mode").(string),
		},
	}

	// Configure selection based on mode
	switch portfolio.Selection.Mode {
	case "MANUAL":
		if v, ok := d.GetOk("projects"); ok {
			projects := v.(*schema.Set).List()
			portfolio.Selection.Projects = make([]string, len(projects))
			for i, project := range projects {
				portfolio.Selection.Projects[i] = project.(string)
			}
		}
	case "REGEXP":
		if v, ok := d.GetOk("project_pattern"); ok {
			portfolio.Selection.ProjectPattern = v.(string)
		}
		if v, ok := d.GetOk("branch_pattern"); ok {
			portfolio.Selection.BranchPattern = v.(string)
		}
	case "FILTER":
		if v, ok := d.GetOk("filters"); ok {
			filters := v.([]interface{})[0].(map[string]interface{})
			portfolio.Filters = expandPortfolioFilters(filters)
		}
	}

	if err := client.CreatePortfolio(portfolio); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(portfolio.Key)
	return resourcePortfolioRead(ctx, d, m)
}

func resourcePortfolioRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*client.Client)

	portfolio, err := client.GetPortfolio(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	d.Set("name", portfolio.Name)
	d.Set("key", portfolio.Key)
	d.Set("description", portfolio.Description)
	d.Set("selection_mode", portfolio.Selection.Mode)

	switch portfolio.Selection.Mode {
	case "MANUAL":
		d.Set("projects", portfolio.Selection.Projects)
	case "REGEXP":
		d.Set("project_pattern", portfolio.Selection.ProjectPattern)
		d.Set("branch_pattern", portfolio.Selection.BranchPattern)
	case "FILTER":
		if err := d.Set("filters", flattenPortfolioFilters(&portfolio.Filters)); err != nil {
			return diag.FromErr(err)
		}
	}

	return nil
}

func resourcePortfolioUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*client.Client)

	portfolio := &client.Portfolio{
		Key:         d.Get("key").(string),
		Name:        d.Get("name").(string),
		Description: d.Get("description").(string),
		Selection: client.PortfolioSelection{
			Mode: d.Get("selection_mode").(string),
		},
	}

	// Configure selection and filters similar to create
	switch portfolio.Selection.Mode {
	case "MANUAL":
		if v, ok := d.GetOk("projects"); ok {
			projects := v.(*schema.Set).List()
			portfolio.Selection.Projects = make([]string, len(projects))
			for i, project := range projects {
				portfolio.Selection.Projects[i] = project.(string)
			}
		}
	case "REGEXP":
		if v, ok := d.GetOk("project_pattern"); ok {
			portfolio.Selection.ProjectPattern = v.(string)
		}
		if v, ok := d.GetOk("branch_pattern"); ok {
			portfolio.Selection.BranchPattern = v.(string)
		}
	case "FILTER":
		if v, ok := d.GetOk("filters"); ok {
			filters := v.([]interface{})[0].(map[string]interface{})
			portfolio.Filters = expandPortfolioFilters(filters)
		}
	}

	if err := client.UpdatePortfolio(portfolio); err != nil {
		return diag.FromErr(err)
	}

	return resourcePortfolioRead(ctx, d, m)
}

func resourcePortfolioDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*client.Client)

	err := client.DeletePortfolio(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")
	return nil
}

func expandPortfolioFilters(data map[string]interface{}) client.PortfolioFilters {
	filters := client.PortfolioFilters{}

	if v, ok := data["languages"]; ok {
		languages := v.(*schema.Set).List()
		filters.Languages = make([]string, len(languages))
		for i, lang := range languages {
			filters.Languages[i] = lang.(string)
		}
	}

	if v, ok := data["tags"]; ok {
		tags := v.(*schema.Set).List()
		filters.Tags = make([]string, len(tags))
		for i, tag := range tags {
			filters.Tags[i] = tag.(string)
		}
	}

	if v, ok := data["quality_gates"]; ok {
		gates := v.(*schema.Set).List()
		filters.QualityGates = make([]string, len(gates))
		for i, gate := range gates {
			filters.QualityGates[i] = gate.(string)
		}
	}

	if v, ok := data["compliance"]; ok {
		compliance := v.([]interface{})[0].(map[string]interface{})
		filters.Compliance = expandPortfolioCompliance(compliance)
	}

	if v, ok := data["custom_metrics"]; ok {
		metrics := v.(map[string]interface{})
		filters.Metrics = make(map[string]client.PortfolioMetric)
		for metric, value := range metrics {
			metricValue := value.(map[string]interface{})
			filters.Metrics[metric] = client.PortfolioMetric{
				Operator: metricValue["operator"].(string),
				Value:    metricValue["value"].(string),
			}
		}
	}

	return filters
}

func expandPortfolioCompliance(data map[string]interface{}) client.PortfolioCompliance {
	compliance := client.PortfolioCompliance{}

	if v, ok := data["min_quality_gate_status"]; ok {
		compliance.MinQualityGateStatus = v.(string)
	}
	if v, ok := data["min_coverage"]; ok {
		compliance.MinCoverage = v.(float64)
	}
	if v, ok := data["max_duplications"]; ok {
		compliance.MaxDuplications = v.(float64)
	}
	if v, ok := data["max_issues"]; ok {
		compliance.MaxIssues = v.(int)
	}
	if v, ok := data["required_rules"]; ok {
		rules := v.(*schema.Set).List()
		compliance.RequiredRules = make([]string, len(rules))
		for i, rule := range rules {
			compliance.RequiredRules[i] = rule.(string)
		}
	}

	return compliance
}

func flattenPortfolioFilters(filters *client.PortfolioFilters) []interface{} {
	if filters == nil {
		return []interface{}{}
	}

	m := map[string]interface{}{
		"languages":     schema.NewSet(schema.HashString, stringSliceToInterfaceSlice(filters.Languages)),
		"tags":         schema.NewSet(schema.HashString, stringSliceToInterfaceSlice(filters.Tags)),
		"quality_gates": schema.NewSet(schema.HashString, stringSliceToInterfaceSlice(filters.QualityGates)),
		"compliance":    flattenPortfolioCompliance(&filters.Compliance),
	}

	if len(filters.Metrics) > 0 {
		metrics := make(map[string]interface{})
		for k, v := range filters.Metrics {
			metrics[k] = map[string]interface{}{
				"operator": v.Operator,
				"value":    v.Value,
			}
		}
		m["custom_metrics"] = metrics
	}

	return []interface{}{m}
}

func flattenPortfolioCompliance(compliance *client.PortfolioCompliance) []interface{} {
	if compliance == nil {
		return []interface{}{}
	}

	m := map[string]interface{}{
		"min_quality_gate_status": compliance.MinQualityGateStatus,
		"min_coverage":           compliance.MinCoverage,
		"max_duplications":       compliance.MaxDuplications,
		"max_issues":            compliance.MaxIssues,
		"required_rules":        schema.NewSet(schema.HashString, stringSliceToInterfaceSlice(compliance.RequiredRules)),
	}

	return []interface{}{m}
}

func stringSliceToInterfaceSlice(s []string) []interface{} {
	i := make([]interface{}, len(s))
	for idx, v := range s {
		i[idx] = v
	}
	return i
}
