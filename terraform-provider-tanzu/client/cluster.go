package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// GetClusters - Returns list of Clusters
func (c *Client) GetClusters() (*Clusters, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/tsm/v1alpha1/clusters", c.HostURL), nil)
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

// GetCluster - Returns specific Cluster (no auth required)
func (c *Client) GetCluster(id string) (*Cluster, error) {
	//logger.Info(c.ctx, "---------------------------------------------")

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/tsm/v1alpha1/clusters/%s", c.HostURL, id), nil)
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
func (c *Client) GetOnboardUrl(authToken *string) (*OnboardUrlResponse, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/tsm/v1alpha1/clusters/onboard-url", c.HostURL), nil)
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
func (c *Client) CreateCluster(cluster Cluster, authToken *string) (*Cluster, error) {
	putUrl := fmt.Sprintf("%s/tsm/v1alpha1/clusters/%s?createOnly=true", c.HostURL, cluster.DisplayName)

	// set this to nil, because we're not supposed to send it to the PUT
	clusterJSON, err := json.Marshal(cluster)
	if err != nil {
		return nil, err
	}

	fmt.Printf("---------------------------------------------\n")
	fmt.Printf("putUrl: %s\n", putUrl)
	fmt.Printf("%s\n", clusterJSON)
	fmt.Printf("---------------------------------------------\n")
	req, err := http.NewRequest("PUT", putUrl, bytes.NewBuffer(clusterJSON))

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

func (c *Client) UpdateCluster(cluster Cluster, authToken *string) (*Cluster, error) {
	putUrl := fmt.Sprintf("%s/tsm/v1alpha1/clusters/%s?createOnly=false", c.HostURL, cluster.DisplayName)

	// set this to nil, because we're not supposed to send it to the PUT
	clusterJSON, err := json.Marshal(cluster)
	if err != nil {
		return nil, err
	}

	fmt.Printf("---------------------------------------------\n")
	fmt.Printf("putUrl: %s\n", putUrl)
	fmt.Printf("%s\n", clusterJSON)
	fmt.Printf("---------------------------------------------\n")
	req, err := http.NewRequest("PUT", putUrl, bytes.NewBuffer(clusterJSON))
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

func (c *Client) DeleteCluster(id string, authToken *string) (*Cluster, error) {
	deleteUrl := fmt.Sprintf("%s/tsm/v1alpha1/clusters/%s", c.HostURL, id)

	fmt.Printf("---------------------------------------------\n")
	fmt.Printf("deleteUrl: %s\n", deleteUrl)
	fmt.Printf("---------------------------------------------\n")

	req, err := http.NewRequest("DELETE", deleteUrl, nil)
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
