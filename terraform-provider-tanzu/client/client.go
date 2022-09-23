package client

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	logging "github.com/hashicorp/terraform-plugin-sdk/v2/helper/logging"
)

// HostURL - Default Hashicups URL
// TODO: Load from provider config !!
const HostURL string = "https://prod-4.nsxservicemesh.vmware.com"
const AuthURL string = "https://console.cloud.vmware.com/csp/gateway/am/api/auth/api-tokens/authorize?grant_type=refresh_token"

// Client -
type Client struct {
	HostURL    string
	AuthURL    string
	HTTPClient *http.Client
	Token      string
	Auth       AuthStruct
	ctx        context.Context
}

// AuthStruct -
type AuthStruct struct {
	apikey string `json:"apikey"`
}

// AuthResponse -
type AuthResponse struct {
	Token string `json:"access_token"`
}

// NewClient -
func NewClient(ctx context.Context, host *string, apikey *string) (*Client, error) {
	transport := logging.NewLoggingHTTPTransport(http.DefaultTransport)

	c := Client{
		HTTPClient: &http.Client{
			Timeout:   10 * time.Second,
			Transport: transport,
		},
		// Default Hashicups URL
		HostURL: HostURL,
		AuthURL: AuthURL,
	}

	if host != nil {
		c.HostURL = *host
	}

	if ctx != nil {
		c.ctx = ctx
	}

	// If username or password not provided, return empty client
	if apikey == nil {
		return &c, nil
	}

	c.Auth = AuthStruct{
		apikey: *apikey,
	}

	ar, err := c.SignIn(ctx)
	if err != nil {
		return nil, err
	}

	c.Token = ar.Token

	return &c, nil
}

func (c *Client) doRequest(req *http.Request, authToken *string) ([]byte, error) {
	token := c.Token
	tflog.Debug(c.ctx, fmt.Sprintf("Making %s Request to Url: %s", req.Method, req.URL))

	if authToken != nil {
		token = *authToken
	}

	req.Header.Set("csp-auth-token", token)
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	req.Header.Set("accept", "application/json")
	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}

	// check that this was successful (perhaps shoudl check if 500?)
	httpOK := res.StatusCode == http.StatusOK
	httpSubmitted := res.StatusCode == http.StatusAccepted
	httpDeleted := res.StatusCode == http.StatusNoContent
	httpCreated := res.StatusCode == http.StatusCreated
	httpNotFound := res.StatusCode == http.StatusNotFound

	if !(httpSubmitted || httpOK || httpCreated || httpDeleted || httpNotFound) {
		return nil, fmt.Errorf("status: %d, body: %s", res.StatusCode, body)
	}

	return body, err
}
