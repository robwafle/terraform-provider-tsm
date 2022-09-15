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

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	var host, apikey string

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

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

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
