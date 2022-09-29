package provider_test

import (
	"testing"

	tc "terraform-provider-tanzu/plugin/client"
	tp "terraform-provider-tanzu/plugin/provider"

	"github.com/go-test/deep"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func TestMapClusterFromSchema(t *testing.T) {

	// labels := []tc.Label{
	// 	tc.Label{
	// 		Key:   "key1",
	// 		Value: "value1",
	// 	},
	// 	tc.Label{
	// 		Key:   "key2",
	// 		Value: "value2",
	// 	},
	// }
	// tags := []string{"a", "b"}

	// namespace_exclusions := []tc.NamespaceExclusion{
	// 	tc.NamespaceExclusion{
	// 		Match:   "match1",
	// 		Type: "type1",
	// 	},
	// 	tc.NamespaceExclusion{
	// 		Match:   "match2",
	// 		Type: "type2",
	// 	},
	// }

	idAndDisplayNameAreTheSameValue := "idAndDisplayNameAreTheSameValue"
	expected := tc.Cluster{
		ID:                     idAndDisplayNameAreTheSameValue,
		DisplayName:            idAndDisplayNameAreTheSameValue,
		Description:            "description",
		AutoInstallServiceMesh: true,
		Token:                  "testToken",

		//Tags:                      tags,
		//NamespaceExclusions:       namespace_exclusions,
		//Labels:                    labels,
		Tags:                      []string{},
		NamespaceExclusions:       []tc.NamespaceExclusion{},
		Labels:                    []tc.Label{},	

		EnableNamespaceExclusions: true,
		
	}
	//TODO?: perhaps find some automatic way to map struct to map (perhaps json marshal / unmarshal?)
	// import structs above... but we need a different struct with a different json mapping matching the terrform schema
	// for example, DisplayName needs to have json annotation display_name instead of displayName
	//testData := structs.Map(&u1)

	testData := map[string]interface{}{
		"id":                          expected.ID,
		"display_name":                expected.DisplayName,
		"description":                 expected.Description,
		"auto_install_servicemesh":    expected.AutoInstallServiceMesh,
		"token":                       expected.Token,
		"enable_namespace_exclusions": expected.EnableNamespaceExclusions,
	}

	d := schema.TestResourceDataRaw(t, tp.ResourceClusterSchema(), testData)

	actual, error := tp.MapClusterFromSchema(d)
	if error != nil {
		t.Errorf("Error: %s", error.Error())
	}

	if diff := deep.Equal(actual, &expected); diff != nil {
		t.Error(diff)
	}

}
