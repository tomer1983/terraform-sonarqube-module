package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSonarqubeQualityGate() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceQualityGateCreate,
		ReadContext:   resourceQualityGateRead,
		UpdateContext: resourceQualityGateUpdate,
		DeleteContext: resourceQualityGateDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"conditions": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"metric": {
							Type:     schema.TypeString,
							Required: true,
						},
						"op": {
							Type:     schema.TypeString,
							Required: true,
						},
						"error": {
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
		},
	}
}

func resourceQualityGateCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*client.Client)
	
	name := d.Get("name").(string)
	
	gate, err := client.CreateQualityGate(name)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(gate.ID)

	// Create conditions
	if v, ok := d.GetOk("conditions"); ok {
		conditions := v.([]interface{})
		for _, c := range conditions {
			condition := c.(map[string]interface{})
			_, err := client.CreateQualityGateCondition(
				gate.ID,
				condition["metric"].(string),
				condition["op"].(string),
				condition["error"].(string),
			)
			if err != nil {
				return diag.FromErr(err)
			}
		}
	}

	return resourceQualityGateRead(ctx, d, m)
}

func resourceQualityGateRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*client.Client)
	
	gate, err := client.ReadQualityGate(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("name", gate.Name); err != nil {
		return diag.FromErr(err)
	}

	conditions := make([]map[string]interface{}, len(gate.Conditions))
	for i, c := range gate.Conditions {
		conditions[i] = map[string]interface{}{
			"metric": c.Metric,
			"op":     c.Op,
			"error":  c.Error,
		}
	}
	
	if err := d.Set("conditions", conditions); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceQualityGateUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*client.Client)
	
	if d.HasChange("name") {
		_, err := client.UpdateQualityGate(d.Id(), d.Get("name").(string))
		if err != nil {
			return diag.FromErr(err)
		}
	}

	// Handle conditions update by recreating them
	if d.HasChange("conditions") {
		// Delete existing conditions
		old, _ := d.GetChange("conditions")
		oldConditions := old.([]interface{})
		for _, c := range oldConditions {
			condition := c.(map[string]interface{})
			if condition["id"] != nil {
				err := client.DeleteQualityGateCondition(condition["id"].(string))
				if err != nil {
					return diag.FromErr(err)
				}
			}
		}

		// Create new conditions
		conditions := d.Get("conditions").([]interface{})
		for _, c := range conditions {
			condition := c.(map[string]interface{})
			_, err := client.CreateQualityGateCondition(
				d.Id(),
				condition["metric"].(string),
				condition["op"].(string),
				condition["error"].(string),
			)
			if err != nil {
				return diag.FromErr(err)
			}
		}
	}

	return resourceQualityGateRead(ctx, d, m)
}

func resourceQualityGateDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*client.Client)
	
	err := client.DeleteQualityGate(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")
	return nil
}
