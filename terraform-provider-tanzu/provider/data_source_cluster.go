package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	tc "terraform-provider-tanzu/plugin/client"
)

func dataSourceCluster() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceClusterRead,
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"cluster_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"resource_group": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"auto_install_servicemesh": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"enable_namespace_exclusions": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"token": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"tags": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"labels": {
				Type:     schema.TypeMap,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			"namespace_exclusion": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"type": {
							Type:     schema.TypeString,
							Required: true,
						},

						"match": {
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceClusterRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	tc := m.(*tc.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	clusterID := d.Get("id").(string)

	cl, err := tc.GetCluster(clusterID)
	if err != nil {
		return diag.FromErr(err)
	}

	// Set top level values
	d.Set("id", cl.ID)
	d.Set("display_name", cl.DisplayName)
	d.Set("token", cl.Token)
	d.Set("auto_install_servicemesh", cl.AutoInstallServiceMesh)
	d.Set("description", cl.Description)
	d.Set("enable_namespace_exclusions", cl.EnableNamespaceExclusions)
	d.Set("state", cl.Status.State)
	d.Set("sync_state", cl.SyncStatus.State)

	// Set labels
	labels := make(map[string]any)

	for _, l := range cl.Labels {
		labels[l.Key] = l.Value
	}
	if err := d.Set("labels", labels); err != nil {
		return diag.FromErr(err)
	}

	// Set NamespaceExclusions
	// namespace_exclusions := make([]map[string]any, 0)

	// for _, ne := range cl.NamespaceExclusions {
	// 	namespace_exclusion := make(map[string]any)
	// 	namespace_exclusion["match"] = ne.Match
	// 	namespace_exclusion["type"] = ne.Type
	// 	namespace_exclusions = append(namespace_exclusions, namespace_exclusion)
	// }

	// if err := d.Set("namespace_exclusions", namespace_exclusions); err != nil {
	// 	return diag.FromErr(err)
	// }

	// if err := d.Set("cluster", clustermap); err != nil {
	// 	return diag.FromErr(err)
	// }

	d.SetId(clusterID)

	return diags
}

func dataSourceCoffees() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceClustersRead,
		Schema: map[string]*schema.Schema{
			"clusters": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceClustersRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	tc := m.(*tc.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	clusters, err := tc.GetClusters()
	if err != nil {
		return diag.FromErr(err)
	}

	clustersList := make([]string, len(clusters.IDs))
	for i, id := range clusters.IDs {
		clustersList[i] = id
	}

	if err := d.Set("clusters", clustersList); err != nil {
		return diag.FromErr(err)
	}

	//d.SetId(display_name)

	return diags
}
