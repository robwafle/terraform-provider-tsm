package client

// Order -
type Order struct {
	ID    int         `json:"id,omitempty"`
	Items []OrderItem `json:"items,omitempty"`
}

// OrderItem -
type OrderItem struct {
	Cluster  Cluster `json:"cluster"`
	Quantity int     `json:"quantity"`
}

// type Tag struct {
// 	Name string `json:"name"`
// }

type Label struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type NamespaceExclusion struct {
	Match string `json:"match"`
	Type  string `json:"type"`
}

type OnboardUrlResponse struct {
	Url string `json:"url"`
}

type Clusters struct {
	IDs []string `json:"ids"`
}

// Cluster -
type Cluster struct {
	ID                        string               `json:"id,omitempty"`
	DisplayName               string               `json:"displayName"`
	Description               string               `json:"description"`
	Labels                    []Label              `json:"labels"`
	Tags                      []string             `json:"tags"`
	NamespaceExclusions       []NamespaceExclusion `json:"namespaceExclusions"`
	EnableNamespaceExclusions bool                 `json:"enableNamespaceExclusions"`
	AutoInstallServiceMesh    bool                 `json:"autoInstallServiceMesh"`
	Token                     string               `json:"token,omitempty"`
	Status                    *Status              `json:"status,omitempty"`
	SyncStatus                *SyncStatus          `json:"syncStatus,omitempty"`
}

type Status struct {
	State   string `json:"state,omitempty"`
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

type SyncStatus struct {
	State string `json:"status,omitempty"`
}
