package provider

import (
	"context"
	tc "terraform-provider-tsm/plugin/client"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func flattenMatchConditionData(matchConditions *[]tc.MatchCondition) []interface{} {
	if matchConditions != nil {
		match_conditions := make([]interface{}, len(*matchConditions), len(*matchConditions))

		for i, match_condition := range *matchConditions {
			m := make(map[string]interface{})

			m["cluster_type"] = match_condition.ClusterMatchCondition.Type
			m["cluster_match"] = match_condition.ClusterMatchCondition.Match
			m["namespace_type"] = match_condition.NamespaceMatchCondition.Type
			m["namespace_match"] = match_condition.NamespaceMatchCondition.Match

			match_conditions[i] = m
		}

		return match_conditions
	}

	return make([]interface{}, 0)
}

func dataSourceGlobalNamespace() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceGlobalNamespaceRead,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"last_updated": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"domain_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"use_shared_gateway": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"mtls_enforced": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"ca_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"ca": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"color": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"version": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"match_conditions": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cluster_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"cluster_match": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"namespace_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"namespace_match": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"api_discovery_enabled": {
				Type:     schema.TypeBool,
				Computed: true,
			},
		},
	}
}

func dataSourceGlobalNamespaceRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*tc.Client)

	var diags diag.Diagnostics

	id := d.Get("id").(string)

	globalNamespace, err := c.GetGlobalNamespace(ctx, id)
	if err != nil && err.Error() != "404" {
		return diag.FromErr(err)
	}

	if globalNamespace != nil {
		diags = MapSchemaFromGlobalNamespace(globalNamespace, d)
	} else {
		d.SetId("")
	}

	return diags
}

// func dataSourceGlobalNamespacesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
// 	c := m.(*tc.Client)

// 	// Warning or errors can be collected in a slice type
// 	var diags diag.Diagnostics

// 	GlobalNamespaces, err := c.GetGlobalNamespaces(ctx)
// 	if err != nil {
// 		return diag.FromErr(err)
// 	}

// 	GlobalNamespacesList := make([]string, len(GlobalNamespaces.IDs))
// 	// for i, id := range GlobalNamespaces.IDs {
// 	// 	GlobalNamespacesList[i] = id
// 	// }
// 	copy(GlobalNamespacesList, GlobalNamespaces.IDs)

// 	if err := d.Set("GlobalNamespaces", GlobalNamespacesList); err != nil {
// 		return diag.FromErr(err)
// 	}

// 	//d.SetId(display_name)

// 	return diags
// }
