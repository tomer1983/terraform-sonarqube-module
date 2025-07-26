package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"host": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("SONARQUBE_HOST", nil),
			},
			"token": {
				Type:        schema.TypeString,
				Required:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("SONARQUBE_TOKEN", nil),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"sonarqube_project":      resourceSonarqubeProject(),
			"sonarqube_qualitygate":  resourceSonarqubeQualityGate(),
			"sonarqube_user":         resourceSonarqubeUser(),
			"sonarqube_group":        resourceSonarqubeGroup(),
			"sonarqube_portfolio":    resourceSonarqubePortfolio(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"sonarqube_project":      dataSourceSonarqubeProject(),
			"sonarqube_quality_gate": dataSourceSonarqubeQualityGate(),
			"sonarqube_portfolio":    dataSourceSonarqubePortfolio(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}

type Client struct {
	Host  string
	Token string
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	host := d.Get("host").(string)
	token := d.Get("token").(string)

	client := &Client{
		Host:  host,
		Token: token,
	}

	return client, nil
}
