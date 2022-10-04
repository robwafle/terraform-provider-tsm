package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// GetGlobalNamespaces - Returns list of GlobalNamespaces
func (c *Client) GetGlobalNamespaces(ctx context.Context) (*GlobalNamespaces, error) {
	tflog.Debug(ctx, "Getting Global Namespaces ...")
	req, err := http.NewRequestWithContext(ctx, "GET", fmt.Sprintf("%s/tsm/v1alpha1/global-namespaces", c.HostURL), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, nil)
	if err != nil {
		return nil, err
	}

	GlobalNamespaces := GlobalNamespaces{}
	err = json.Unmarshal(body, &GlobalNamespaces)
	if err != nil {
		return nil, err
	}

	return &GlobalNamespaces, nil
}

// GetGlobalNamespace - Returns specific GlobalNamespace (no auth required)
func (c *Client) GetGlobalNamespace(ctx context.Context, id string) (*GlobalNamespace, error) {

	tflog.Debug(ctx, fmt.Sprintf("Getting Global Namespace: %s", id))
	req, err := http.NewRequestWithContext(ctx, "GET", fmt.Sprintf("%s/tsm/v1alpha1/global-namespaces/%s", c.HostURL, id), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, nil)
	if err != nil {
		return nil, err
	}

	GlobalNamespace := GlobalNamespace{}
	err = json.Unmarshal(body, &GlobalNamespace)
	if err != nil {
		return nil, err
	}
	strBody := string(body)
	tflog.Debug(ctx, fmt.Sprintf("Global Namespace:%s", strBody))

	return &GlobalNamespace, nil
}

// CreateGlobalNamespace - Create new GlobalNamespace
func (c *Client) CreateUpdateGlobalNamespace(ctx context.Context, globalNamespace GlobalNamespace, authToken *string) (*GlobalNamespace, error) {
	putUrl := fmt.Sprintf("%s/tsm/v1alpha1/global-namespaces/%s", c.HostURL, globalNamespace.ID)

	// set this to nil, because we're not supposed to send it to the PUT
	globalNamespace.ID = ""

	GlobalNamespaceJSON, err := json.Marshal(globalNamespace)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, "PUT", putUrl, bytes.NewBuffer(GlobalNamespaceJSON))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, authToken)
	if err != nil {
		return nil, err
	}

	newGlobalNamespace := GlobalNamespace{}
	err = json.Unmarshal(body, &newGlobalNamespace)
	if err != nil {
		return nil, err
	}

	return &newGlobalNamespace, nil
}

func (c *Client) DeleteGlobalNamespace(ctx context.Context, id string, authToken *string) (*GlobalNamespace, error) {
	deleteUrl := fmt.Sprintf("%s/tsm/v1alpha1/global-namespaces/%s", c.HostURL, id)

	req, err := http.NewRequestWithContext(ctx, "DELETE", deleteUrl, nil)
	if err != nil {
		return nil, err
	}

	_, err = c.doRequest(req, authToken)
	if err != nil {
		return nil, err
	}

	// TODO: expected return 204, should probably check for that

	// newGlobalNamespace := GlobalNamespace{}
	// err = json.Unmarshal(body, &newGlobalNamespace)
	// if err != nil {
	// 	return nil, err
	// }

	return nil, nil
}
