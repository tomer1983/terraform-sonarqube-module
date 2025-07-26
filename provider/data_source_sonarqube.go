package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceSonarqubeProject() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceProjectRead,

		Schema: map[string]*schema.Schema{
			"key": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"visibility": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"main_branch": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"tags": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func dataSourceProjectRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*client.Client)

	key := d.Get("key").(string)
	project, err := client.GetProject(key)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(project.Key)
	d.Set("name", project.Name)
	d.Set("visibility", project.Visibility)
	d.Set("main_branch", project.MainBranch)
	d.Set("tags", project.Tags)

	return nil
}

func dataSourceSonarqubeQualityGate() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceQualityGateRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"conditions": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"metric": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"op": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"error": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceQualityGateRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*client.Client)

	name := d.Get("name").(string)
	gate, err := client.GetQualityGateByName(name)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(gate.ID)
	
	conditions := make([]map[string]interface{}, len(gate.Conditions))
	for i, c := range gate.Conditions {
		conditions[i] = map[string]interface{}{
			"metric": c.Metric,
			"op":     c.Op,
			"error":  c.Error,
		}
	}
	
	d.Set("conditions", conditions)

	return nil
}

func dataSourceSonarqubePortfolio() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourcePortfolioRead,

		Schema: map[string]*schema.Schema{
			"key": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"selection_mode": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"projects": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"filters": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"languages": {
							Type:     schema.TypeSet,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"tags": {
							Type:     schema.TypeSet,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
		},
	}
}

func dataSourcePortfolioRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*client.Client)

	key := d.Get("key").(string)
	portfolio, err := client.GetPortfolio(key)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(portfolio.Key)
	d.Set("name", portfolio.Name)
	d.Set("description", portfolio.Description)
	d.Set("selection_mode", portfolio.Selection.Mode)
	
	if portfolio.Selection.Mode == "MANUAL" {
		d.Set("projects", portfolio.Selection.Projects)
	}

	if portfolio.Selection.Mode == "FILTER" {
		filters := []map[string]interface{}{
			{
				"languages": portfolio.Filters.Languages,
				"tags":     portfolio.Filters.Tags,
			},
		}
		d.Set("filters", filters)
	}

	return nil
}
