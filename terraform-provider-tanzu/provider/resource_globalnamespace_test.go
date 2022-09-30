package provider_test

import (
	"encoding/json"
	"testing"

	tc "terraform-provider-tanzu/plugin/client"
	tp "terraform-provider-tanzu/plugin/provider"

	"github.com/go-test/deep"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func TestMapGlobalNamespaceFromSchema(t *testing.T) {

	MatchConditions := []tc.MatchCondition{
		{
			NamespaceMatchCondition: tc.NamespaceMatchCondition{
				Match: "ns_match1",
				Type:  "ns_type1",
			},
			ClusterMatchCondition: tc.ClusterMatchCondition{
				Match: "cl_match1",
				Type:  "cl_nstype1",
			},
		},
		{
			NamespaceMatchCondition: tc.NamespaceMatchCondition{
				Match: "ns_match2",
				Type:  "ns_type2",
			},
			ClusterMatchCondition: tc.ClusterMatchCondition{
				Match: "cl_match2",
				Type:  "cl_nstype2",
			},
		},
	}

	matchConditionsExclusionsAsArrayOfInterface := make([]interface{}, len(MatchConditions))
	for i, match_condition := range MatchConditions {
		m := make(map[string]interface{})
		m["cluster_type"] = match_condition.ClusterMatchCondition.Type
		m["cluster_match"] = match_condition.ClusterMatchCondition.Match
		m["namespace_type"] = match_condition.NamespaceMatchCondition.Type
		m["namespace_match"] = match_condition.NamespaceMatchCondition.Match
		matchConditionsExclusionsAsArrayOfInterface[i] = m
	}

	idAndDisplayNameAreTheSameValue := "idAndDisplayNameAreTheSameValue"
	expected := tc.GlobalNamespace{
		ID:               idAndDisplayNameAreTheSameValue,
		Name:             idAndDisplayNameAreTheSameValue,
		DisplayName:      idAndDisplayNameAreTheSameValue,
		DomainName:       "vmware.com",
		UseSharedGateway: true,
		MtlsEnforced:     true,
		CaType:           "PreExistingCA",
		Ca:               "default",
		Description:      "created in unit test",
		Color:            "#00FF00",
		Version:          "2.0",
		MatchConditions:  MatchConditions,

		ApiDiscoveryEnabled: true,
	}

	testData := map[string]interface{}{
		"id":                    expected.ID,
		"name":                  expected.Name,
		"display_name":          expected.DisplayName,
		"domain_name":           expected.DomainName,
		"use_shared_gateway":    expected.UseSharedGateway,
		"mtls_enforced":         expected.MtlsEnforced,
		"ca_type":               expected.CaType,
		"ca":                    expected.Ca,
		"description":           expected.Description,
		"color":                 expected.Color,
		"version":               expected.Version,
		"match_condition":       matchConditionsExclusionsAsArrayOfInterface,
		"api_discovery_enabled": expected.ApiDiscoveryEnabled,
	}

	d := schema.TestResourceDataRaw(t, tp.ResourceGlobalNamespaceSchema(), testData)

	actual, error := tp.MapGlobalNamespaceFromSchema(d)
	if error != nil {
		t.Errorf("Error: %s", error.Error())
	}

	if diff := deep.Equal(actual, &expected); diff != nil {
		t.Error(diff)
		actualJSON, error := json.Marshal(actual)
		if error != nil {
			t.Errorf("Error: %s", error.Error())
		}
		t.Errorf("actualJSON: %s", actualJSON)
		expectedJSON, error := json.Marshal(expected)
		if error != nil {
			t.Errorf("Error: %s", error.Error())
		}
		t.Errorf("expectedJSON: %s", expectedJSON)
	}

}
