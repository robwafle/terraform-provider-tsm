package provider

import (
	"context"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	tc "terraform-provider-tsm/plugin/client"
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
		Schema: ResourceGlobalNamespaceSchema(),
	}
}

func ResourceGlobalNamespaceSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"last_updated": {
			Type:     schema.TypeString,
			Computed: true,
			Optional: true,
		},
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"display_name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"domain_name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"use_shared_gateway": {
			Type:     schema.TypeBool,
			Required: true,
		},
		"mtls_enforced": {
			Type:     schema.TypeBool,
			Required: true,
		},
		"ca_type": {
			Type:     schema.TypeString,
			Required: true,
		},
		"ca": {
			Type:     schema.TypeString,
			Required: true,
		},
		"description": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"color": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "test description",
		},
		"version": {
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
		"api_discovery_enabled": {
			Type:     schema.TypeBool,
			Required: true,
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
	resource, maperror := MapGlobalNamespaceFromSchema(d)

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
func resourceGlobalNamespaceRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*tc.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	id := d.Id()

	globalNamespace, err := c.GetGlobalNamespace(ctx, id)
	if err != nil && err.Error() != "404" {
		return diag.FromErr(err)
	}

	if globalNamespace != nil {
		tflog.Debug(ctx, "Setting Root Level Fields ... ")
		// DO NOT SET THE ID
		//d.Set("id", globalNamespace.ID)
		tflog.Debug(ctx, fmt.Sprintf("Setting Name: %s", globalNamespace.Name))
		d.Set("name", globalNamespace.Name)

		tflog.Debug(ctx, fmt.Sprintf("Setting DisplayName: %s", globalNamespace.DisplayName))
		d.Set("display_name", globalNamespace.DisplayName)

		tflog.Debug(ctx, fmt.Sprintf("Setting DomainName: %s", globalNamespace.DomainName))
		d.Set("domain_name", globalNamespace.DomainName)

		d.Set("use_shared_gateway", globalNamespace.UseSharedGateway)
		d.Set("mtls_enforced", globalNamespace.MtlsEnforced)
		d.Set("ca_type", globalNamespace.CaType)
		d.Set("ca", globalNamespace.Ca)
		d.Set("description", globalNamespace.Description)
		d.Set("color", globalNamespace.Color)
		d.Set("version", globalNamespace.Version)
		d.Set("api_discovery_enabled", globalNamespace.ApiDiscoveryEnabled)

		tflog.Debug(ctx, "Setting MatchConditions ... ")
		// Set NamespaceExclusions
		match_conditions := make([]interface{}, len(globalNamespace.MatchConditions))

		for i, mc := range globalNamespace.MatchConditions {
			match_condition := make(map[string]interface{})
			match_condition["cluster_type"] = mc.ClusterMatchCondition.Type
			match_condition["cluster_match"] = mc.ClusterMatchCondition.Match
			match_condition["namespace_type"] = mc.NamespaceMatchCondition.Type
			match_condition["namespace_match"] = mc.NamespaceMatchCondition.Match
			match_conditions[i] = match_condition
		}

		if err := d.Set("match_condition", match_conditions); err != nil {
			return diag.FromErr(err)
		}
		tflog.Debug(ctx, "Setting Id ... ")
		//d.SetId(globalNamespace.ID)
	}
	tflog.Debug(ctx, "Done with resourceGlobalNamespaceRead ...")
	return diags
}

func resourceGlobalNamespaceUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*tc.Client)

	resourceToUpdate, resourceToUpdateError := MapGlobalNamespaceFromSchema(d)
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
