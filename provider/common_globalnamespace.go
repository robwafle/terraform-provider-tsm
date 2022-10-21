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

	if globalNamespace == nil {
		return diag.Errorf("globalNamespace was nil in MapSchemaFromGlobalNamespace()")
	}

	if err := d.Set("name", globalNamespace.Name); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Warning,
			Summary:  "Unable to update field in cluster see error.",
			Detail:   err.Error(),
		})
	}

	if err := d.Set("display_name", globalNamespace.DisplayName); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Warning,
			Summary:  "Unable to update field in cluster see error.",
			Detail:   err.Error(),
		})
	}

	if err := d.Set("domain_name", globalNamespace.DomainName); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Warning,
			Summary:  "Unable to update field in cluster see error.",
			Detail:   err.Error(),
		})
	}

	if err := d.Set("use_shared_gateway", globalNamespace.UseSharedGateway); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Warning,
			Summary:  "Unable to update field in cluster see error.",
			Detail:   err.Error(),
		})
	}

	if err := d.Set("mtls_enforced", globalNamespace.MtlsEnforced); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Warning,
			Summary:  "Unable to update field in cluster see error.",
			Detail:   err.Error(),
		})
	}

	if err := d.Set("ca_type", globalNamespace.CaType); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Warning,
			Summary:  "Unable to update field in cluster see error.",
			Detail:   err.Error(),
		})
	}

	if err := d.Set("ca", globalNamespace.Ca); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Warning,
			Summary:  "Unable to update field in cluster see error.",
			Detail:   err.Error(),
		})
	}

	if err := d.Set("description", globalNamespace.Description); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Warning,
			Summary:  "Unable to update field in cluster see error.",
			Detail:   err.Error(),
		})
	}

	if err := d.Set("color", globalNamespace.Color); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Warning,
			Summary:  "Unable to update field in cluster see error.",
			Detail:   err.Error(),
		})
	}

	if err := d.Set("version", globalNamespace.Version); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Warning,
			Summary:  "Unable to update field in cluster see error.",
			Detail:   err.Error(),
		})
	}

	if err := d.Set("api_discovery_enabled", globalNamespace.ApiDiscoveryEnabled); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Warning,
			Summary:  "Unable to update field in cluster see error.",
			Detail:   err.Error(),
		})
	}

	matchConditions := flattenMatchConditionData(&globalNamespace.MatchConditions)
	if err := d.Set("match_conditions", matchConditions); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Warning,
			Summary:  "Unable to update field in cluster see error.",
			Detail:   err.Error(),
		})
	}

	// get global namespace does NOT return an ID, but it matches the name
	d.SetId(globalNamespace.Name)

	// TODO: Understand returning diags better, we don't appear to be using this object
	// // https://developer.hashicorp.com/terraform/tutorials/providers/provider-debug
	// The plan is to wrap this in a unit test so most problems with this will be found at build time
	// but it is possible some things do happen at run time and appending to this and aggregating them instead
	// of failing on the first error sounds like a good thing.// https://developer.hashicorp.com/terraform/tutorials/providers/provider-debug
	return diags
}
