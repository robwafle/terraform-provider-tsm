package provider

import (
	"context"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	tc "terraform-provider-tanzu/plugin/client"
)

func resourceGlobalNamespace() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceGlobalNamespaceCreate,
		ReadContext:   resourceGlobalNamespaceRead,
		UpdateContext: resourceGlobalNamespaceUpdate,
		DeleteContext: resourceGlobalNamespaceDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"last_updated": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"domain_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"use_shared_gateway": &schema.Schema{
				Type:     schema.TypeBool,
				Required: true,
			},
			"mtls_enforced": &schema.Schema{
				Type:     schema.TypeBool,
				Required: true,
			},
			"ca_type": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"ca": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"color": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"version": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"match_condition": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cluster_type": {
							Type:     schema.TypeString,
							Required: true,
						},
						"cluster_match": {
							Type:     schema.TypeString,
							Required: true,
						},
						"namespace_type": {
							Type:     schema.TypeString,
							Required: true,
						},
						"namespace_match": {
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
			"api_discovery_enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Required: true,
			},
		},
	}
}

// {
// 	"displayName": "string",
// 	"description": "string",
// 	"tags": [
// 	  "string"
// 	],
// 	"labels": [
// 	  {
// 		"key": "string",
// 		"value": "string"
// 	  }
// 	],
// 	"autoInstallServiceMesh": true,
// 	"enableNamespaceExclusions": true,
// 	"namespaceExclusions": [
// 	  {
// 		"type": "string",
// 		"match": "string"
// 	  }
// 	],
// 	"proxyConfig": {
// 	  "proxy": "Explicit",
// 	  "protocol": "HTTP",
// 	  "host": "string",
// 	  "port": 0,
// 	  "username": "string",
// 	  "password": "string",
// 	  "certificate": "string"
// 	},
// 	"autoInstallServiceMeshConfig": {
// 	  "restrictDefaultExternalAccess": true
// 	},
// 	"registryAccount": "string",
// 	"caLabels": [
// 	  {
// 		"key": "string",
// 		"value": "string"
// 	  }
// 	]
//   }

func resourceGlobalNamespaceCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	c := m.(*tc.Client)

	tflog.Debug(ctx, "Mapping GlobalNamespace From Schema...")
	resource, maperror := mapGlobalNamespaceFromSchema(ctx, d)

	tflog.Debug(ctx, "Checking Errors...")
	if maperror != nil {
		fmt.Println(maperror.Error())
		return diag.FromErr(maperror)
	}
	tflog.Debug(ctx, "Done mapping GlobalNamespace From Schema...")

	_, createerr := c.CreateUpdateGlobalNamespace(ctx, *resource, nil)
	if createerr != nil {
		return diag.FromErr(createerr)
	}

	tflog.Debug(ctx, "Setting ID...")
	d.SetId(resource.Name)

	resourceGlobalNamespaceRead(ctx, d, m)
	tflog.Debug(ctx, "created a resource")

	return diags

}

func mapGlobalNamespaceFromSchema(ctx context.Context, d *schema.ResourceData) (*tc.GlobalNamespace, error) {
	id := d.Get("id").(string)
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

	tflog.Debug(ctx, "-----------------[match_conditions]----------------------------")

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
		ID:                  id,
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

func resourceGlobalNamespaceRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*tc.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	id := d.Id()

	GlobalNamespace, err := c.GetGlobalNamespace(ctx, id)
	if err != nil {
		return diag.FromErr(err)
	}

	tflog.Debug(ctx, "Setting Root Level Fields ... ")
	d.Set("id", GlobalNamespace.ID)
	tflog.Debug(ctx, fmt.Sprintf("Setting Name: %s", GlobalNamespace.Name))
	d.Set("name", GlobalNamespace.Name)

	tflog.Debug(ctx, fmt.Sprintf("Setting DisplayName: %s", GlobalNamespace.DisplayName))
	d.Set("display_name", GlobalNamespace.DisplayName)

	tflog.Debug(ctx, fmt.Sprintf("Setting DomainName: %s", GlobalNamespace.DomainName))
	d.Set("domain_name", GlobalNamespace.DomainName)

	d.Set("use_shared_gateway", GlobalNamespace.UseSharedGateway)
	d.Set("mtls_enforced", GlobalNamespace.MtlsEnforced)
	d.Set("ca_type", GlobalNamespace.CaType)
	d.Set("ca", GlobalNamespace.Ca)
	d.Set("description", GlobalNamespace.Description)
	d.Set("color", GlobalNamespace.Color)
	d.Set("version", GlobalNamespace.Version)
	d.Set("api_discovery_enabled", GlobalNamespace.ApiDiscoveryEnabled)

	// tflog.Debug(ctx, "Setting MatchConditions ... ")
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
	tflog.Debug(ctx, "Setting Id ... ")
	d.SetId(GlobalNamespace.ID)
	tflog.Debug(ctx, "Done with resourceGlobalNamespaceRead ...")
	return diags
}

func resourceGlobalNamespaceUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*tc.Client)

	resourceToUpdate, resourceToUpdateError := mapGlobalNamespaceFromSchema(ctx, d)
	if resourceToUpdateError != nil {
		return diag.FromErr(resourceToUpdateError)
	}

	d.Set("last_updated", time.Now().Format(time.RFC850))

	_, err := c.CreateUpdateGlobalNamespace(ctx, *resourceToUpdate, nil)
	if err != nil {
		return diag.FromErr(err)
	}

	return resourceGlobalNamespaceRead(ctx, d, m)
}

func resourceGlobalNamespaceDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	c := m.(*tc.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	id := d.Id()

	_, err := c.DeleteGlobalNamespace(ctx, id, nil)
	if err != nil {
		return diag.FromErr(err)
	}

	return diags
}
