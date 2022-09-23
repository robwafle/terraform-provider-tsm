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
				Computed: true,
			},
			"last_updated": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"display_name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"auto_install_servicemesh": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
			},
			"enable_namespace_exclusions": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
			},
			"token": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"tags": &schema.Schema{
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

			"namespace_exclusion": {
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

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	id := d.Id()

	cl, err := c.GetCluster(id)
	if err != nil {
		return diag.FromErr(err)
	}

	// Set top level values
	if err := d.Set("id", cl.ID); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("display_name", cl.DisplayName); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("token", cl.Token); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("auto_install_servicemesh", cl.AutoInstallServiceMesh); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("description", cl.Description); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("enable_namespace_exclusions", cl.EnableNamespaceExclusions); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("state", cl.Status.State); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("sync_state", cl.SyncStatus.State); err != nil {
		return diag.FromErr(err)
	}

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

	d.SetId(id)

	return diags
}

func dataSourceClustersRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*tc.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	clusters, err := c.GetClusters()
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
