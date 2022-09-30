package provider_test

import (
	"testing"

	tc "terraform-provider-tanzu/plugin/client"
	tp "terraform-provider-tanzu/plugin/provider"

	"encoding/json"

	"github.com/go-test/deep"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func TestMapClusterFromSchema(t *testing.T) {

	// IMPORTANT NOTE: TestResourceDataRaw() method
	// DOES NOT support an array of strings
	// It only suppports a list of interfaces, so it seems we can only use
	// basic types from our models in our tests.
	tags := []string{"tag1", "tag2"}
	tagsAsArrayOfInterface := make([]interface{}, len(tags))
	for i, tag := range tags {
		tagsAsArrayOfInterface[i] = tag
	}

	labels := []tc.Label{
		{
			Key:   "key1",
			Value: "value1",
		},
		{
			Key:   "key2",
			Value: "value2",
		},
	}

	labelsAsArrayOfInterface := make([]interface{}, len(labels))
	for i, label := range labels {
		m := make(map[string]interface{})
		m[label.Key] = label.Value
		labelsAsArrayOfInterface[i] = m
	}

	namespace_exclusions := []tc.NamespaceExclusion{
		{
			Match: "match1",
			Type:  "type1",
		},
		{
			Match: "match2",
			Type:  "type2",
		},
		{
			Match: "match3",
			Type:  "type3",
		},
	}

	namespaceExclusionsAsArrayOfInterface := make([]interface{}, len(namespace_exclusions))
	for i, namespace_exclusion := range namespace_exclusions {
		m := make(map[string]interface{})
		m["match"] = namespace_exclusion.Match
		m["type"] = namespace_exclusion.Type
		namespaceExclusionsAsArrayOfInterface[i] = m
	}

	idAndDisplayNameAreTheSameValue := "idAndDisplayNameAreTheSameValue"
	expected := tc.Cluster{
		ID:                     idAndDisplayNameAreTheSameValue,
		DisplayName:            idAndDisplayNameAreTheSameValue,
		Description:            "description",
		AutoInstallServiceMesh: true,
		Token:                  "testToken",

		Tags:                tags,
		NamespaceExclusions: namespace_exclusions,
		Labels:              labels,

		//Tags:                      []string{},
		//NamespaceExclusions: []tc.NamespaceExclusion{},
		//Labels:              []tc.Label{},

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
		"tags":                        tagsAsArrayOfInterface,
		"labels":                      labelsAsArrayOfInterface,
		"namespace_exclusion":         namespaceExclusionsAsArrayOfInterface,
	}

	// for _, tag := range tagsAsArrayOfInterface {
	// 	t.Errorf("tag: %s", tag.(string))
	// }

	d := schema.TestResourceDataRaw(t, tp.ResourceClusterSchema(), testData)

	actual, error := tp.MapClusterFromSchema(d)
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
