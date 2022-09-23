package provider

import (
	"context"
	"fmt"
	tc "terraform-provider-tanzu/plugin/client"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceGlobalNamespace() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceGlobalNamespaceRead,
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
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"display_name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"domain_name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"use_shared_gateway": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
			},
			"mtls_enforced": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
			},
			"ca_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ca": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"color": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"version": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"match_condition": {
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
			"api_discovery_enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
			},
		},
	}
}

func dataSourceGlobalNamespaceRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*tc.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	id := d.Id()

	GlobalNamespace, err := c.GetGlobalNamespace(ctx, id)
	if err != nil {
		return diag.FromErr(err)
	}

	tflog.Debug(ctx, fmt.Sprintf("Setting Root Level Fields ... "))
	if err := d.Set("id", GlobalNamespace.ID); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("name", GlobalNamespace.Name); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("display_name", GlobalNamespace.DisplayName); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("domain_name", GlobalNamespace.DomainName); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("use_shared_gateway", GlobalNamespace.UseSharedGateway); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("mtls_enforced", GlobalNamespace.MtlsEnforced); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("ca_type", GlobalNamespace.CaType); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("ca", GlobalNamespace.Ca); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("description", GlobalNamespace.Description); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("color", GlobalNamespace.Color); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("version", GlobalNamespace.Version); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("api_discovery_enabled", GlobalNamespace.ApiDiscoveryEnabled); err != nil {
		return diag.FromErr(err)
	}

	tflog.Debug(ctx, fmt.Sprintf("Setting MatchConditions ... "))
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
	d.SetId(id)

	tflog.Debug(ctx, fmt.Sprintf("Done with resourceGlobalNamespaceRead ... "))
	return diags
}

func dataSourceGlobalNamespacesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*tc.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	GlobalNamespaces, err := c.GetGlobalNamespaces(ctx)
	if err != nil {
		return diag.FromErr(err)
	}

	GlobalNamespacesList := make([]string, len(GlobalNamespaces.IDs))
	for i, id := range GlobalNamespaces.IDs {
		GlobalNamespacesList[i] = id
	}

	if err := d.Set("GlobalNamespaces", GlobalNamespacesList); err != nil {
		return diag.FromErr(err)
	}

	//d.SetId(display_name)

	return diags
}
