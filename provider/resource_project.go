package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSonarqubeProject() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceProjectCreate,
		ReadContext:   resourceProjectRead,
		UpdateContext: resourceProjectUpdate,
		DeleteContext: resourceProjectDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"project_key": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"visibility": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "private",
			},
			"main_branch": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "main",
			},
			"tags": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func resourceProjectCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*client.Client)
	
	name := d.Get("name").(string)
	key := d.Get("project_key").(string)
	visibility := d.Get("visibility").(string)
	mainBranch := d.Get("main_branch").(string)
	
	rawTags := d.Get("tags").([]interface{})
	tags := make([]string, len(rawTags))
	for i, v := range rawTags {
		tags[i] = v.(string)
	}

	project, err := client.CreateProject(name, key, visibility, mainBranch, tags)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(project.Key)
	return resourceProjectRead(ctx, d, m)
}

func resourceProjectRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*client.Client)
	
	project, err := client.ReadProject(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("name", project.Name); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("project_key", project.Key); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("visibility", project.Visibility); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("main_branch", project.MainBranch); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("tags", project.Tags); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceProjectUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*client.Client)
	
	if d.HasChanges("name", "visibility", "tags") {
		name := d.Get("name").(string)
		visibility := d.Get("visibility").(string)
		
		rawTags := d.Get("tags").([]interface{})
		tags := make([]string, len(rawTags))
		for i, v := range rawTags {
			tags[i] = v.(string)
		}

		_, err := client.UpdateProject(d.Id(), name, visibility, tags)
		if err != nil {
			return diag.FromErr(err)
		}
	}

	return resourceProjectRead(ctx, d, m)
}

func resourceProjectDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*client.Client)
	
	err := client.DeleteProject(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")
	return nil
}
