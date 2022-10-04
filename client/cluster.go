package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// GetClusters - Returns list of Clusters
func (c *Client) GetClusters(ctx context.Context) (*Clusters, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", fmt.Sprintf("%s/tsm/v1alpha1/clusters", c.HostURL), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, nil)
	if err != nil {
		return nil, err
	}

	Clusters := Clusters{}
	err = json.Unmarshal(body, &Clusters)
	if err != nil {
		return nil, err
	}

	return &Clusters, nil
}

func (c *Client) GetCluster(ctx context.Context, id string) (*Cluster, error) {

	req, err := http.NewRequestWithContext(ctx, "GET", fmt.Sprintf("%s/tsm/v1alpha1/clusters/%s", c.HostURL, id), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, nil)
	if err != nil {
		return nil, err
	}

	cluster := Cluster{}
	err = json.Unmarshal(body, &cluster)
	if err != nil {
		return nil, err
	}

	return &cluster, nil
}

// GetOnboardUrl
func (c *Client) GetOnboardUrl(ctx context.Context, authToken *string) (*OnboardUrlResponse, error) {
	tflog.Trace(ctx, "Getting Onboard Url ...")
	req, err := http.NewRequestWithContext(ctx, "GET", fmt.Sprintf("%s/tsm/v1alpha1/clusters/onboard-url", c.HostURL), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, nil)
	if err != nil {
		return nil, err
	}

	onboardUrlResponse := OnboardUrlResponse{}
	err = json.Unmarshal(body, &onboardUrlResponse)
	if err != nil {
		return nil, err
	}

	return &onboardUrlResponse, nil
}

// CreateCluster - Create new Cluster
func (c *Client) CreateCluster(ctx context.Context, cluster Cluster, authToken *string) (*Cluster, error) {
	tflog.Trace(ctx, "Creating Cluster ...")
	putUrl := fmt.Sprintf("%s/tsm/v1alpha1/clusters/%s?createOnly=true", c.HostURL, cluster.DisplayName)

	// set this to nil, because we're not supposed to send it to the PUT
	clusterJSON, err := json.Marshal(cluster)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, "PUT", putUrl, bytes.NewBuffer(clusterJSON))

	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, authToken)
	if err != nil {
		return nil, err
	}

	newCluster := Cluster{}
	err = json.Unmarshal(body, &newCluster)
	if err != nil {
		return nil, err
	}

	return &newCluster, nil
}

func (c *Client) UpdateCluster(ctx context.Context, cluster Cluster, authToken *string) (*Cluster, error) {
	tflog.Trace(ctx, "Updating Cluster ...")
	putUrl := fmt.Sprintf("%s/tsm/v1alpha1/clusters/%s?createOnly=false", c.HostURL, cluster.DisplayName)

	// set this to nil, because we're not supposed to send it to the PUT	// set this to nil, because we're not supposed to send it to the PUT
	cluster.ID = ""
	clusterJSON, err := json.Marshal(cluster)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, "PUT", putUrl, bytes.NewBuffer(clusterJSON))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, authToken)
	if err != nil {
		return nil, err
	}

	newCluster := Cluster{}
	err = json.Unmarshal(body, &newCluster)
	if err != nil {
		return nil, err
	}

	return &newCluster, nil
}

func (c *Client) DeleteCluster(ctx context.Context, id string, authToken *string) (*Cluster, error) {
	tflog.Trace(ctx, "Deleting Cluster ...")
	deleteUrl := fmt.Sprintf("%s/tsm/v1alpha1/clusters/%s", c.HostURL, id)

	req, err := http.NewRequestWithContext(ctx, "DELETE", deleteUrl, nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, authToken)
	if err != nil {
		return nil, err
	}

	newCluster := Cluster{}
	err = json.Unmarshal(body, &newCluster)
	if err != nil {
		return nil, err
	}

	return &newCluster, nil
}
