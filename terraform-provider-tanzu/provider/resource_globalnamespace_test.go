package provider_test

import (
	"testing"

	tc "terraform-provider-tanzu/plugin/client"
	tp "terraform-provider-tanzu/plugin/provider"

	"github.com/go-test/deep"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func TestMapGlobalNamespaceFromSchema(t *testing.T) {

	// MatchConditions := []tc.MatchCondition{
	// 	tc.MatchCondition{
	// 		tc.NamespaceMatchCondition{
	// 			Match: "ns_match1",
	// 			Type:  "ns_type1",
	// 		},
	// 		tc.ClusterMatchCondition{
	// 			Match: "cl_match1",
	// 			Type:  "cl_nstype1",
	// 		},
	// 	},
	// 	tc.MatchCondition{
	// 		tc.NamespaceMatchCondition{
	// 			Match: "ns_match2",
	// 			Type:  "ns_type2",
	// 		},
	// 		tc.ClusterMatchCondition{
	// 			Match: "cl_match2",
	// 			Type:  "cl_nstype2",
	// 		},
	// 	},
	// }

	idAndDisplayNameAreTheSameValue := "idAndDisplayNameAreTheSameValue"
	expected := tc.GlobalNamespace{
		ID:                  idAndDisplayNameAreTheSameValue,
		Name:                idAndDisplayNameAreTheSameValue,
		DisplayName:         idAndDisplayNameAreTheSameValue,
		DomainName:          "vmware.com",
		UseSharedGateway:    true,
		MtlsEnforced:        true,
		CaType:              "PreExistingCA",
		Ca:                  "default",
		Description:         "created in unit test",
		Color:               "#00FF00",
		Version:             "2.0",
	//	MatchConditions:     MatchConditions,
		MatchConditions:     []tc.MatchCondition{},
		ApiDiscoveryEnabled: true,
	}

	//TODO?: perhaps find some automatic way to map struct to map (perhaps json marshal / unmarshal?)
	// import structs above... but we need a different struct with a different json mapping matching the terrform schema
	// for example, DisplayName needs to have json annotation display_name instead of displayName
	//testData := structs.Map(&u1)

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
		//"match_condition":       expected.MatchConditions,
		"api_discovery_enabled": expected.ApiDiscoveryEnabled,
	}

	d := schema.TestResourceDataRaw(t, tp.ResourceGlobalNamespaceSchema(), testData)

	actual, error := tp.MapGlobalNamespaceFromSchema(d)
	if error != nil {
		t.Errorf("Error: %s", error.Error())
	}

	if diff := deep.Equal(actual, &expected); diff != nil {
		t.Error(diff)
	}

}
