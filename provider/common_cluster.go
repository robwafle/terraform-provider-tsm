package provider

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	tc "terraform-provider-tsm/plugin/client"
)

func MapClusterFromSchema(d *schema.ResourceData) (*tc.Cluster, error) {
	ID := d.Get("id").(string)
	displayName := d.Get("display_name").(string)
	description := d.Get("description").(string)
	token := d.Get("token").(string)
	enableNamespaceExclusions := d.Get("enable_namespace_exclusions").(bool)
	autoInstallServiceMesh := d.Get("auto_install_servicemesh").(bool)

	// _tags := d.Get("tags").(*schema.Set).List()
	// tags := make([]string, len(_tags))

	// // workaround for unit test !?!?
	// // reverse list order to pass unit test... this order probably doesn't really matter anyways..
	// // deep compare doesn't have the ability to ignore hte order of the list
	// for i, tag := range _tags {
	// 	//fmt.Print(tag)
	// 	//tags[i] = tag.(string)
	// 	tags[len(_tags)-i-1] = tag.(string)

	// }
	// //copy(tags, _tags)

	// _labels := d.Get("labels").(map[string]interface{})
	// //labels := []tc.Label{}
	// labels := make([]tc.Label, len(_labels))

	// i := 0
	// labelsLen := len(_labels)
	// for key, value := range _labels {
	// 	//fmt.Printf("\nkey:%s\n", key)
	// 	//fmt.Printf("\nvalue:%s\n", value)
	// 	m := make(map[string]interface{})
	// 	m[key] = value
	// 	label := tc.Label{
	// 		Key:   key,
	// 		Value: value.(string),
	// 	}
	// 	//labels[i] = label
	// 	labels[labelsLen-i-1] = label
	// 	i = i + 1
	// 	//labels = append(labels, label)
	// }

	// _namespace_exclusions := d.Get("namespace_exclusion").(*schema.Set).List()
	// //namespace_exclusions := []tc.NamespaceExclusion{}
	// namespace_exclusions := make([]tc.NamespaceExclusion, len(_namespace_exclusions))

	// for i, namespace_exclusion := range _namespace_exclusions {
	// 	ne, lbok := namespace_exclusion.(map[string]any)
	// 	if lbok {
	// 		namespace_exclusion := tc.NamespaceExclusion{
	// 			Match: ne["match"].(string),
	// 			Type:  ne["type"].(string),
	// 		}
	// 		//namespace_exclusions = append(namespace_exclusions, namespace_exclusion)
	// 		namespace_exclusions[len(_namespace_exclusions)-i-1] = namespace_exclusion
	// 	}
	// }

	cluster := tc.Cluster{
		ID:                        ID,
		DisplayName:               displayName,
		Description:               description,
		Token:                     token,
		AutoInstallServiceMesh:    autoInstallServiceMesh,
		EnableNamespaceExclusions: enableNamespaceExclusions,
		//Tags:                      tags,
		//NamespaceExclusions:       namespace_exclusions,
		//Labels:                    labels,
	}

	return &cluster, nil
}

func MapSchemaFromCluster(cl *tc.Cluster, d *schema.ResourceData) diag.Diagnostics {

	var diags diag.Diagnostics

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
	if cl.Status != nil {
		if err := d.Set("state", cl.Status.State); err != nil {
			return diag.FromErr(err)
		}
	}
	if cl.SyncStatus != nil {
		if err := d.Set("sync_state", cl.SyncStatus.State); err != nil {
			return diag.FromErr(err)
		}
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

	d.SetId(cl.ID)

	// TODO: Understand returning diags better, we don't appear to be using this object
	// // https://developer.hashicorp.com/terraform/tutorials/providers/provider-debug
	// The plan is to wrap this in a unit test so most problems with this will be found at build time
	// but it is possible some things do happen at run time and appending to this and aggregating them instead
	// of failing on the first error sounds like a good thing.// https://developer.hashicorp.com/terraform/tutorials/providers/provider-debug
	return diags
}
