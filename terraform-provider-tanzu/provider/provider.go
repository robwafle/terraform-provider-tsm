package provider

import (
	"context"
	"fmt"

	tc "terraform-provider-tanzu/plugin/client"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Provider -
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"host": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("TANZU_HOST", nil),
			},
			"apikey": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("TANZU_APIKEY", nil),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"tanzu_cluster": resourceCluster(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"tanzu_cluster": dataSourceCluster(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}

// func kubectlConfigure(ctx context.Context, d *schema.ResourceData) diag.Diagnostics {
// 	var diags diag.Diagnostics

// 	var cluster_name, resource_group string

// 	fmt.Printf("Configuring Kubectl ...\n")

// 	// Authenticate to cluster using az aks get credentials.
// 	// TODO: learn how to write config file given previous terraform step using client cert and certificate authority cert
// 	// example: az aks get-credentials --name tanzu-two --resource-group tanzu-two-rg

// 	rgHVal, rgOk := d.GetOk("resource_group")
// 	if rgOk {
// 		tmpRg := rgHVal.(string)
// 		resource_group = tmpRg
// 	}
// 	cnHVal, cnOk := d.GetOk("cluster_name")
// 	if cnOk {
// 		tempCn := cnHVal.(string)
// 		cluster_name = tempCn
// 	}

// 	if cluster_name != "" && resource_group != "" {
// 		fmt.Printf("\n-----------------[kubectl-config]----------------------------\n")
// 		fmt.Printf("Configuring .kube/config using az aks get-credentials. Using resource_group: %s, cluster_name: %s ...", resource_group, cluster_name)
// 		kubeConfig := exec.Command("az", "aks", "get-credentials", "--resource-group", resource_group, "--name", cluster_name)
// 		execkubeConfigStdout, execkubeConfigErr := kubeConfig.Output()

// 		fmt.Print(string(execkubeConfigStdout))

// 		fmt.Printf("\n-----------------[kubectl-config]----------------------------\n")
// 		if execkubeConfigErr != nil {
// 			fmt.Print(execkubeConfigErr.Error())
// 			diags = append(diags, diag.Diagnostic{
// 				Severity: diag.Error,
// 				Summary:  "Unable to configure .kube/config for Tanzu Provider",
// 				Detail:   fmt.Sprintf("Unable to configure .kube/config Error: %s, resource_group: %s, cluster_name: %s", execkubeConfigErr.Error(), resource_group, cluster_name),
// 			})

// 			return diags
// 		}
// 	}

// 	return diags
// }

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	var diags diag.Diagnostics

	var host, apikey string

	fmt.Printf("Configuring Provider ...\n")

	// authenticate to tanzu
	apikeyVal, ok := d.GetOk("apikey")
	if ok {
		tempApiKey := apikeyVal.(string)
		apikey = tempApiKey
	}

	hVal, ok := d.GetOk("host")
	if ok {
		tempHost := hVal.(string)
		host = tempHost
	}

	if apikey != "" {
		c, err := tc.NewClient(ctx, &host, &apikey)
		fmt.Printf("===================================================")
		fmt.Printf("host: %s, apikey: %s", host, apikey)
		fmt.Printf("===================================================")
		if err != nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Unable to create Tanzu client",
				Detail:   fmt.Sprintf("Unable to authenticate apikey for authenticated Tanzu client. Error: %s, host: %s, apikey: %s", err.Error(), host, apikey),
			})

			return nil, diags
		}

		return c, diags
	}

	c, err := tc.NewClient(ctx, &host, nil)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to create Tanzu client",
			Detail:   fmt.Sprintf("Unable to create unauthenticated Tanzu client. Error: %s, host: %s", err.Error(), host),
		})
		return nil, diags
	}

	return c, diags
}
