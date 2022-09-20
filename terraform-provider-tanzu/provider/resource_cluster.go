package provider

import (
	"context"
	"fmt"
	"os/exec"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	tc "terraform-provider-tanzu/plugin/client"
)

func resourceCluster() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceClusterCreate,
		ReadContext:   resourceClusterRead,
		UpdateContext: resourceClusterUpdate,
		DeleteContext: resourceClusterDelete,
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"display_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"kubernetes_context": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"token": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"auto_install_servicemesh": &schema.Schema{
				Type:     schema.TypeBool,
				Required: true,
			},
			"enable_namespace_exclusions": &schema.Schema{
				Type:     schema.TypeBool,
				Required: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"tags": &schema.Schema{
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
		},
	}
}

func resourceClusterCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	c := m.(*tc.Client)
	displayName := d.Get("display_name").(string)
	kubernetesContext := d.Get("kubernetes_context").(string)
	fmt.Printf("Mapping Cluster From Schema...\n")
	clusterToCreate, mapClusterToCreateError := mapClusterFromSchema(d)

	fmt.Printf("Checking Errors...\n")
	if mapClusterToCreateError != nil {
		fmt.Println(mapClusterToCreateError.Error())
		return diag.FromErr(mapClusterToCreateError)
	}
	fmt.Printf("Done mapping Cluster From Schema...\n")

	onboardUrlResponse, err := c.GetOnboardUrl(nil)
	if err != nil {
		return diag.FromErr(err)
	}
	_ = onboardUrlResponse
	fmt.Printf("-----------------[kubectl-a]----------------------------\n")
	fmt.Printf("onboardUrlResponse.Url:%s\n", onboardUrlResponse.Url)
	fmt.Printf("-----------------[kubectl-b]----------------------------\n")

	onboardCmd := exec.Command("kubectl", "--context", kubernetesContext, "apply", "-f", onboardUrlResponse.Url)
	execOnboardCmdStdout, execOnboardCmdErr := onboardCmd.Output()
	fmt.Printf("\n-----------------[kubectl-c]----------------------------\n")
	fmt.Print(string(execOnboardCmdStdout))

	fmt.Printf("\n-----------------[kubectl-d]----------------------------\n")
	if execOnboardCmdErr != nil {
		fmt.Print(execOnboardCmdErr.Error())
		return diag.FromErr(execOnboardCmdErr)
	}

	cl, err := c.CreateCluster(*clusterToCreate, nil)
	if err != nil {
		return diag.FromErr(err)
	}

	fmt.Printf("\n-----------------[kubectl-1]----------------------------\n")
	fmt.Printf("cl.Token:%s\n", cl.Token)
	fmt.Printf("\n-----------------[kubectl-2]----------------------------\n")
	// kubectl -n vmware-system-tsm create secret generic cluster-token --from-literal=token={token}

	checkSecretExistsCmd := exec.Command("kubectl", "--context", kubernetesContext, "-n", "vmware-system-tsm", "get", "secret", "cluster-token")
	checkSecretExistsCmdStdout, checkSecretExistsCmdErr := checkSecretExistsCmd.Output()
	fmt.Printf("\n-----------------[kubectl-2b]----------------------------\n")
	fmt.Print(string(checkSecretExistsCmdStdout))
	// NOTE: If the cluster has already been registered, token will be null and this command will fail.
	// NOTE: If the secret has already been created, trying to create it again will fail.
	if checkSecretExistsCmdErr != nil && cl.Token != "" {
		fmt.Printf("\nCreating Secret!!!\n")
		connectCmd := exec.Command("kubectl", "--context", kubernetesContext, "-n", "vmware-system-tsm", "create", "secret", "generic", "cluster-token", fmt.Sprintf("--from-literal=token=%s", cl.Token))
		connectCmdStdout, execConnectCmdErr := connectCmd.Output()
		fmt.Printf("\n-----------------[kubectl-3]----------------------------\n")
		fmt.Print(string(connectCmdStdout))
		fmt.Printf("\n-----------------[kubectl-4]----------------------------\n")
		if execConnectCmdErr != nil {
			fmt.Print(execConnectCmdErr.Error())
			return diag.FromErr(err)
		}
	} else {
		fmt.Printf("\nSkipping Secret Creation!!!\n")
	}

	// if err := d.Set("display_name", cl.DisplayName); err != nil {
	// 	return diag.FromErr(err)
	// }

	fmt.Printf("\n Setting ID...\n")
	d.SetId(displayName)

	return diags

}

func mapClusterFromSchema(d *schema.ResourceData) (*tc.Cluster, error) {
	displayName := d.Get("display_name").(string)
	description := d.Get("description").(string)
	enableNamespaceExclusions := d.Get("enable_namespace_exclusions").(bool)
	autoInstallServiceMesh := d.Get("auto_install_servicemesh").(bool)

	fmt.Printf("-----------------[TAGS]----------------------------\n")

	_tags := d.Get("tags").(*schema.Set).List()
	tags := []string{}

	for _, _t := range _tags {
		t := _t.(string)
		tags = append(tags, t)
	}

	fmt.Printf("-----------------[LABELS]----------------------------\n")

	_labels := d.Get("labels").(map[string]any)
	labels := []tc.Label{}
	fmt.Printf("-----------------[LABELS]----------------------------\n")
	fmt.Printf("_labels: %d\n", len(_labels))
	fmt.Printf("-----------------[LABELS]----------------------------\n")
	for key, value := range _labels {
		label := tc.Label{
			Key:   key,
			Value: value.(string),
		}
		labels = append(labels, label)
	}

	fmt.Printf("-----------------[namespace_exclusions111]----------------------------\n")

	_namespace_exclusions := d.Get("namespace_exclusion").(*schema.Set).List()
	namespace_exclusions := []tc.NamespaceExclusion{}
	fmt.Printf("-----------------[namespace_exclusions]----------------------------\n")
	fmt.Printf("_namespace_exclusions: %d\n", len(_namespace_exclusions))
	fmt.Printf("-----------------[namespace_exclusions]----------------------------\n")
	for _, namespace_exclusion := range _namespace_exclusions {
		ne, lbok := namespace_exclusion.(map[string]any)
		if lbok {
			namespace_exclusion := tc.NamespaceExclusion{
				Match: ne["match"].(string),
				Type:  ne["type"].(string),
			}
			namespace_exclusions = append(namespace_exclusions, namespace_exclusion)
		}
	}

	cluster := tc.Cluster{
		DisplayName:               displayName,
		Description:               description,
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

	clusterID := d.Get("id").(string)

	cl, err := c.GetCluster(clusterID)
	if err != nil {
		return diag.FromErr(err)
	}

	fmt.Printf("\nSetting Root Level Fields ... \n")
	d.Set("id", cl.ID)
	d.Set("display_name", cl.DisplayName)
	d.Set("token", cl.Token)
	d.Set("auto_install_servicemesh", cl.AutoInstallServiceMesh)
	d.Set("description", cl.Description)
	d.Set("enable_namespace_exclusions", cl.EnableNamespaceExclusions)

	// Set labels
	fmt.Printf("\nSetting Labels ... \n")
	labels := make(map[string]any)

	for _, l := range cl.Labels {
		labels[l.Key] = l.Value
	}
	if err := d.Set("labels", labels); err != nil {
		return diag.FromErr(err)
	}

	fmt.Printf("\nSetting NamespaceExclusions ... \n")
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
	fmt.Printf("\nSetting Id ... \n")
	d.SetId(cl.ID)
	fmt.Printf("\nDone with resourceClusterRead ... \n")
	return diags
}

func resourceClusterUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*tc.Client)

	clusterToUpdate, clusterToUpdateError := mapClusterFromSchema(d)
	if clusterToUpdateError != nil {
		return diag.FromErr(clusterToUpdateError)
	}

	d.Set("last_updated", time.Now().Format(time.RFC850))

	_, err := c.UpdateCluster(*clusterToUpdate, nil)
	if err != nil {
		return diag.FromErr(err)
	}

	return resourceClusterRead(ctx, d, m)
}

func resourceClusterDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	c := m.(*tc.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	clusterID := d.Get("id").(string)
	kubernetesContext := d.Get("kubernetes_context").(string)
	fmt.Printf("---------------------------------------------\n")
	fmt.Printf("clusterID to delete: %s\n", clusterID)
	fmt.Printf("---------------------------------------------\n")

	_, err := c.DeleteCluster(clusterID, nil)
	if err != nil {
		return diag.FromErr(err)
	}

	//kubectl delete --ignore-not-found=true -f https://prod-4.nsxservicemesh.vmware.com/cluster-registration/k8s/client-cluster-uninstall.yaml

	deleteUrlYaml := fmt.Sprintf("%s/cluster-registration/k8s/client-cluster-uninstall.yaml", c.HostURL)
	fmt.Printf("---------------------------------------------\n")
	fmt.Printf("deleteUrlYaml: %s\n", deleteUrlYaml)
	fmt.Printf("---------------------------------------------\n")
	deleteCmd := exec.Command("kubectl", "--context", kubernetesContext, "delete", "--ignore-not-found=true", "-f", deleteUrlYaml)

	execDeleteCmdStdout, execDeleteCmdErr := deleteCmd.Output()
	fmt.Printf("\n-----------------[kubectl-c]----------------------------\n")
	fmt.Print(string(execDeleteCmdStdout))

	if execDeleteCmdErr != nil {
		fmt.Print(execDeleteCmdErr.Error())
		return diag.FromErr(execDeleteCmdErr)
	}

	return diags
}
