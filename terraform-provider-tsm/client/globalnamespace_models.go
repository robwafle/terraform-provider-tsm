package client

type GlobalNamespaces struct {
	IDs []string `json:"ids"`
}

// Cluster -
type GlobalNamespace struct {
	ID                  string           `json:"id,omitempty"`
	Name                string           `json:"name"`
	DisplayName         string           `json:"display_name"`
	DomainName          string           `json:"domain_name"`
	UseSharedGateway    bool             `json:"use_shared_gateway"`
	MtlsEnforced        bool             `json:"mtls_enforced"`
	CaType              string           `json:"ca_type"`
	Ca                  string           `json:"ca"`
	Description         string           `json:"description,omitempty"`
	Color               string           `json:"color,omitempty"`
	Version             string           `json:"version,omitempty"`
	MatchConditions     []MatchCondition `json:"match_conditions"`
	ApiDiscoveryEnabled bool             `json:"api_discovery_enabled"`
}

type MatchCondition struct {
	NamespaceMatchCondition NamespaceMatchCondition `json:"namespace"`
	ClusterMatchCondition   ClusterMatchCondition   `json:"cluster"`
}

type NamespaceMatchCondition struct {
	Type  string `json:"type"`
	Match string `json:"match"`
}

type ClusterMatchCondition struct {
	Type  string `json:"type"`
	Match string `json:"match"`
}

// {
// 	"name": "string",
// 	"display_name": "string",
// 	"domain_name": "string",
// 	"use_shared_gateway": true,
// 	"mtls_enforced": true,
// 	"ca_type": "PreExistingCA",
// 	"ca": "string",
// 	"description": "string",
// 	"color": "string",
// 	"version": "string",
// 	"match_conditions": [
// 	  {
// 		"namespace": {
// 		  "type": "string",
// 		  "match": "string"
// 		},
// 		"cluster": {
// 		  "type": "string",
// 		  "match": "string"
// 		}
// 	  }
// 	],
// 	"api_discovery_enabled": true
//   }
