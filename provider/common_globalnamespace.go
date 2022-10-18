package provider

import (
	tc "terraform-provider-tsm/plugin/client"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func MapGlobalNamespaceFromSchema(d *schema.ResourceData) (*tc.GlobalNamespace, error) {
	ID := d.Get("id").(string)
	Name := d.Get("name").(string)
	DisplayName := d.Get("display_name").(string)
	DomainName := d.Get("domain_name").(string)
	UseSharedGateway := d.Get("use_shared_gateway").(bool)
	MtlsEnforced := d.Get("mtls_enforced").(bool)
	CaType := d.Get("ca_type").(string)
	Ca := d.Get("ca").(string)
	Description := d.Get("description").(string)
	Color := d.Get("color").(string)
	Version := d.Get("version").(string)
	ApiDiscoveryEnabled := d.Get("api_discovery_enabled").(bool)

	_match_conditions := d.Get("match_condition").(*schema.Set).List()
	MatchConditions := []tc.MatchCondition{}

	for _, match_condition := range _match_conditions {
		mc, lbok := match_condition.(map[string]any)
		if lbok {
			NamespaceMatchCondition := tc.NamespaceMatchCondition{
				Match: mc["namespace_match"].(string),
				Type:  mc["namespace_type"].(string),
			}

			ClusterMatchCondition := tc.ClusterMatchCondition{
				Type:  mc["cluster_type"].(string),
				Match: mc["cluster_match"].(string),
			}

			MatchCondition := tc.MatchCondition{
				NamespaceMatchCondition: NamespaceMatchCondition,
				ClusterMatchCondition:   ClusterMatchCondition,
			}

			MatchConditions = append(MatchConditions, MatchCondition)
		}
	}

	globalNamespace := tc.GlobalNamespace{
		ID:                  ID,
		Name:                Name,
		DisplayName:         DisplayName,
		DomainName:          DomainName,
		UseSharedGateway:    UseSharedGateway,
		MtlsEnforced:        MtlsEnforced,
		CaType:              CaType,
		Ca:                  Ca,
		Description:         Description,
		Color:               Color,
		Version:             Version,
		MatchConditions:     MatchConditions,
		ApiDiscoveryEnabled: ApiDiscoveryEnabled,
	}

	return &globalNamespace, nil
}

func MapSchemaFromGlobalNamespace(globalNamespace *tc.GlobalNamespace, d *schema.ResourceData) diag.Diagnostics {

	var diags diag.Diagnostics

	if err := d.Set("id", globalNamespace.ID); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("name", globalNamespace.Name); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("display_name", globalNamespace.DisplayName); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("domain_name", globalNamespace.DomainName); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("use_shared_gateway", globalNamespace.UseSharedGateway); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("mtls_enforced", globalNamespace.MtlsEnforced); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("ca_type", globalNamespace.CaType); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("ca", globalNamespace.Ca); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("description", globalNamespace.Description); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("color", globalNamespace.Color); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("version", globalNamespace.Version); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("api_discovery_enabled", globalNamespace.ApiDiscoveryEnabled); err != nil {
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
	d.SetId(globalNamespace.ID)

	// TODO: Understand returning diags better, we don't appear to be using this object
	// // https://developer.hashicorp.com/terraform/tutorials/providers/provider-debug
	// The plan is to wrap this in a unit test so most problems with this will be found at build time
	// but it is possible some things do happen at run time and appending to this and aggregating them instead
	// of failing on the first error sounds like a good thing.// https://developer.hashicorp.com/terraform/tutorials/providers/provider-debug
	return diags
}
