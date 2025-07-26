package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceSonarqubeUser() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceUserRead,

		Schema: map[string]*schema.Schema{
			"login": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"email": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"active": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"local": {
				Type:     schema.TypeBool,
				Computed: true,
			},
		},
	}
}

func dataSourceUserRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*client.Client)

	login := d.Get("login").(string)
	user, err := client.GetUser(login)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(user.Login)
	d.Set("name", user.Name)
	d.Set("email", user.Email)
	d.Set("active", user.Active)
	d.Set("local", user.Local)

	return nil
}

func dataSourceSonarqubeGroup() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceGroupRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"members_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"default": {
				Type:     schema.TypeBool,
				Computed: true,
			},
		},
	}
}

func dataSourceGroupRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*client.Client)

	name := d.Get("name").(string)
	group, err := client.GetGroup(name)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(group.Name)
	d.Set("description", group.Description)
	d.Set("members_count", group.MembersCount)
	d.Set("default", group.Default)

	return nil
}

func dataSourceSonarqubeMetric() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceMetricRead,

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
			"domain": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"type": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceMetricRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*client.Client)

	key := d.Get("key").(string)
	metric, err := client.GetMetric(key)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(metric.Key)
	d.Set("name", metric.Name)
	d.Set("description", metric.Description)
	d.Set("domain", metric.Domain)
	d.Set("type", metric.Type)

	return nil
}

func dataSourceSonarqubeLanguage() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceLanguageRead,

		Schema: map[string]*schema.Schema{
			"key": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"file_suffixes": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func dataSourceLanguageRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*client.Client)

	key := d.Get("key").(string)
	language, err := client.GetLanguage(key)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(language.Key)
	d.Set("name", language.Name)
	d.Set("file_suffixes", language.FileSuffixes)

	return nil
}

func dataSourceSonarqubeRule() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceRuleRead,

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
			"severity": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"template": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"language": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"type": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceRuleRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*client.Client)

	key := d.Get("key").(string)
	rule, err := client.GetRule(key)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(rule.Key)
	d.Set("name", rule.Name)
	d.Set("description", rule.Description)
	d.Set("severity", rule.Severity)
	d.Set("status", rule.Status)
	d.Set("template", rule.Template)
	d.Set("language", rule.Language)
	d.Set("type", rule.Type)

	return nil
}
