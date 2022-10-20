package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	tc "terraform-provider-tsm/plugin/client"
)

func dataSourceCluster() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceClusterRead,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"auto_install_servicemesh": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"enable_namespace_exclusions": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"token": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"sync_state": {
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
			"labels": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			"namespace_exclusions": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"type": {
							Type:     schema.TypeString,
							Computed: true,
						},

						"match": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceClusterRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*tc.Client)

	var diags diag.Diagnostics

	id := d.Get("id").(string)

	cluster, err := c.GetCluster(ctx, id)
	if err != nil && err.Error() != "404" {
		return diag.FromErr(err)
	}

	if cluster != nil {
		diags = MapSchemaFromCluster(cluster, d)
	} else {
		d.SetId("")
	}

	return diags
}

// func dataSourceClustersRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
// 	c := m.(*tc.Client)

// 	// Warning or errors can be collected in a slice type
// 	var diags diag.Diagnostics

// 	clusters, err := c.GetClusters(ctx)
// 	if err != nil {
// 		return diag.FromErr(err)
// 	}

// 	clustersList := make([]string, len(clusters.IDs))
// 	// for i, id := range clusters.IDs {
// 	// 	clustersList[i] = id
// 	// }
// 	copy(clustersList, clusters.IDs)

// 	if err := d.Set("clusters", clustersList); err != nil {
// 		return diag.FromErr(err)
// 	}

// 	//d.SetId(display_name)

// 	return diags
// }
