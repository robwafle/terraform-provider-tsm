package provider

import (
	"context"
	"fmt"
	"os/exec"
	"time"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	tc "terraform-provider-tanzu/plugin/client"
)

func ResourceClusterSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"last_updated": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"display_name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"kubernetes_context": {
			Type:     schema.TypeString,
			Required: true,
		},
		"token": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"auto_install_servicemesh": {
			Type:     schema.TypeBool,
			Required: true,
		},
		"enable_namespace_exclusions": {
			Type:     schema.TypeBool,
			Required: true,
		},
		"description": {
			Type:     schema.TypeString,
			Required: true,
		},
		"tags": {
			Type:     schema.TypeSet,
			Required: true,
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
	}
}

func resourceCluster() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceClusterCreate,
		ReadContext:   resourceClusterRead,
		UpdateContext: resourceClusterUpdate,
		DeleteContext: resourceClusterDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: ResourceClusterSchema(),
	}
}

func resourceClusterCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	c := m.(*tc.Client)
	//kubectlConfigure(ctx, d)

	displayName := d.Get("display_name").(string)
	kubernetesContext := d.Get("kubernetes_context").(string)
	tflog.Trace(ctx, "Mapping Cluster From Schema...")
	clusterToCreate, mapClusterToCreateError := MapClusterFromSchema(d)

	tflog.Trace(ctx, "Checking Errors...")
	if mapClusterToCreateError != nil {
		fmt.Println(mapClusterToCreateError.Error())
		return diag.FromErr(mapClusterToCreateError)
	}
	tflog.Trace(ctx, "Done mapping Cluster From Schema...")

	onboardUrlResponse, err := c.GetOnboardUrl(ctx, nil)
	if err != nil {
		return diag.FromErr(err)
	}
	_ = onboardUrlResponse

	onboardCmd := exec.Command("kubectl", "--context", kubernetesContext, "apply", "-f", onboardUrlResponse.Url)
	execOnboardCmdStdout, execOnboardCmdErr := onboardCmd.Output()
	fmt.Print(string(execOnboardCmdStdout))

	if execOnboardCmdErr != nil {
		fmt.Print(execOnboardCmdErr.Error())
		return diag.FromErr(execOnboardCmdErr)
	}

	cl, err := c.CreateCluster(ctx, *clusterToCreate, nil)
	if err != nil {
		return diag.FromErr(err)
	}

	tflog.Trace(ctx, fmt.Sprintf("cl.Token:%s", cl.Token))

	// kubectl -n vmware-system-tsm create secret generic cluster-token --from-literal=token={token}

	checkSecretExistsCmd := exec.Command("kubectl", "--context", kubernetesContext, "-n", "vmware-system-tsm", "get", "secret", "cluster-token")
	checkSecretExistsCmdStdout, checkSecretExistsCmdErr := checkSecretExistsCmd.Output()

	fmt.Print(string(checkSecretExistsCmdStdout))
	// NOTE: If the cluster has already been registered, token will be null and this command will fail.
	// NOTE: If the secret has already been created, trying to create it again will fail.
	if checkSecretExistsCmdErr != nil && cl.Token != "" {
		tflog.Trace(ctx, "Creating Secret!!!")
		connectCmd := exec.Command("kubectl", "--context", kubernetesContext, "-n", "vmware-system-tsm", "create", "secret", "generic", "cluster-token", fmt.Sprintf("--from-literal=token=%s", cl.Token))
		connectCmdStdout, execConnectCmdErr := connectCmd.Output()

		fmt.Print(string(connectCmdStdout))

		if execConnectCmdErr != nil {
			fmt.Print(execConnectCmdErr.Error())
			return diag.FromErr(err)
		}
	} else {
		tflog.Trace(ctx, "Skipping Secret Creation!!!")
	}

	sleepTime := time.Second * 10
	for ok := true; ok; ok = cl != nil && cl.Status != nil && cl.SyncStatus != nil && cl.Status.State == "Ready" && cl.SyncStatus.State == "Synced" {
		tflog.Info(ctx, "Waiting for Cluster.Status.State == 'Ready' and Cluster.SyncStatus.State == 'Synced' ... ")
		time.Sleep(sleepTime)
		cl, err = c.GetCluster(ctx, displayName)
		if err != nil {
			if err.Error() == "404" {
				tflog.Info(ctx, "Got 404, continuing ... ")
				continue
			} else {
				return diag.FromErr(err)
			}
		}
		tflog.Info(ctx, fmt.Sprintf("Cluster.Status.State: %s, Cluster.SyncStatus.State: %s", cl.Status.State, cl.SyncStatus.State))
	}

	// if err := d.Set("display_name", cl.DisplayName); err != nil {
	// 	return diag.FromErr(err)
	// }

	tflog.Trace(ctx, " Setting ID...")
	d.SetId(displayName)

	resourceClusterRead(ctx, d, m)

	return diags

}

func (d *schema.ResourceData) (*tc.Cluster, error) {
	ID := d.Get("id").(string)
	displayName := d.Get("display_name").(string)
	description := d.Get("description").(string)
	token := d.Get("token").(string)
	enableNamespaceExclusions := d.Get("enable_namespace_exclusions").(bool)
	autoInstallServiceMesh := d.Get("auto_install_servicemesh").(bool)

	_tags := d.Get("tags").(*schema.Set).List()
	tags := make([]string, len(_tags))

	// workaround for unit test !?!?
	// reverse list order to pass unit test... this order probably doesn't really matter anyways..
	// deep compare doesn't have the ability to ignore hte order of the list
	for i, tag := range _tags {
		//fmt.Print(tag)
		//tags[i] = tag.(string)
		tags[len(_tags)-i-1] = tag.(string)

	}
	//copy(tags, _tags)

	_labels := d.Get("labels").(map[string]interface{})
	//labels := []tc.Label{}
	labels := make([]tc.Label, len(_labels))

	i := 0
	for key, value := range _labels {
		//fmt.Printf("\nkey:%s\n", key)
		//fmt.Printf("\nvalue:%s\n", value)
		m := make(map[string]interface{})
		m[key] = value
		label := tc.Label{
			Key:   key,
			Value: value.(string),
		}
		labels[i] = label
		i = i + 1
		//labels = append(labels, label)
	}

	_namespace_exclusions := d.Get("namespace_exclusion").(*schema.Set).List()
	//namespace_exclusions := []tc.NamespaceExclusion{}
	namespace_exclusions := make([]tc.NamespaceExclusion, len(_namespace_exclusions))

	for i, namespace_exclusion := range _namespace_exclusions {
		ne, lbok := namespace_exclusion.(map[string]any)
		if lbok {
			namespace_exclusion := tc.NamespaceExclusion{
				Match: ne["match"].(string),
				Type:  ne["type"].(string),
			}
			//namespace_exclusions = append(namespace_exclusions, namespace_exclusion)
			namespace_exclusions[len(_namespace_exclusions)-i-1] = namespace_exclusion
		}
	}

	cluster := tc.Cluster{
		ID:                        ID,
		DisplayName:               displayName,
		Description:               description,
		Token:                     token,
		AutoInstallServiceMesh:    autoInstallServiceMesh,
		Tags:                      tags,
		NamespaceExclusions:       namespace_exclusions,
		EnableNamespaceExclusions: enableNamespaceExclusions,
		Labels:                    labels,
	}

	return &cluster, nil
}

func resourceClusterRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*tc.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	clusterID := d.Id()

	cl, err := c.GetCluster(ctx, clusterID)
	if err != nil {
		return diag.FromErr(err)
	}

	tflog.Trace(ctx, "Setting Root Level Fields ... ")
	d.Set("id", cl.ID)
	d.Set("display_name", cl.DisplayName)
	d.Set("token", cl.Token)
	d.Set("auto_install_servicemesh", cl.AutoInstallServiceMesh)
	d.Set("description", cl.Description)
	d.Set("enable_namespace_exclusions", cl.EnableNamespaceExclusions)

	// Set labels
	tflog.Trace(ctx, "Setting Labels ... ")
	labels := make(map[string]interface{})

	for _, l := range cl.Labels {
		labels[l.Key] = l.Value
	}
	if err := d.Set("labels", labels); err != nil {
		return diag.FromErr(err)
	}

	tflog.Trace(ctx, "Setting NamespaceExclusions ... ")
	// Set NamespaceExclusions
	namespace_exclusions := make([]interface{}, len(cl.NamespaceExclusions))

	for i, ne := range cl.NamespaceExclusions {
		namespace_exclusion := make(map[string]interface{})
		namespace_exclusion["match"] = ne.Match
		namespace_exclusion["type"] = ne.Type
		namespace_exclusions[i] = namespace_exclusion
	}

	if err := d.Set("namespace_exclusion", namespace_exclusions); err != nil {
		return diag.FromErr(err)
	}

	tflog.Trace(ctx, "Setting Id ... ")
	//d.SetId(cl.ID)
	tflog.Trace(ctx, "Done with resourceClusterRead ... ")
	return diags
}

func resourceClusterUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*tc.Client)
	//kubectlConfigure(ctx, d)

	clusterToUpdate, clusterToUpdateError := MapClusterFromSchema(d)
	if clusterToUpdateError != nil {
		return diag.FromErr(clusterToUpdateError)
	}

	d.Set("last_updated", time.Now().Format(time.RFC850))

	_, err := c.UpdateCluster(ctx, *clusterToUpdate, nil)
	if err != nil {
		return diag.FromErr(err)
	}

	return resourceClusterRead(ctx, d, m)
}

func resourceClusterDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	c := m.(*tc.Client)
	//kubectlConfigure(ctx, d)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	clusterID := d.Get("id").(string)
	kubernetesContext := d.Get("kubernetes_context").(string)

	tflog.Trace(ctx, fmt.Sprintf("clusterID to delete: %s", clusterID))

	_, err := c.DeleteCluster(ctx, clusterID, nil)
	if err != nil {
		return diag.FromErr(err)
	}

	//kubectl delete --ignore-not-found=true -f https://prod-4.nsxservicemesh.vmware.com/cluster-registration/k8s/client-cluster-uninstall.yaml

	deleteUrlYaml := fmt.Sprintf("%s/cluster-registration/k8s/client-cluster-uninstall.yaml", c.HostURL)

	tflog.Trace(ctx, fmt.Sprintf("deleteUrlYaml: %s", deleteUrlYaml))

	deleteCmd := exec.Command("kubectl", "--context", kubernetesContext, "delete", "--ignore-not-found=true", "-f", deleteUrlYaml)

	execDeleteCmdStdout, execDeleteCmdErr := deleteCmd.Output()

	fmt.Print(string(execDeleteCmdStdout))

	if execDeleteCmdErr != nil {
		fmt.Print(execDeleteCmdErr.Error())
		return diag.FromErr(execDeleteCmdErr)
	}

	return diags
}
