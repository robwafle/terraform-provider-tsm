package client

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
	//logger "github.com/hashicorp/terraform-plugin-log"
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
	//transport := logger.NewLoggingHTTPTransport(http.DefaultTransport)

	c := Client{
		HTTPClient: &http.Client{
			Timeout: 10 * time.Second,
			//Transport: transport,
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

	ar, err := c.SignIn()
	if err != nil {
		return nil, err
	}

	c.Token = ar.Token

	return &c, nil
}

func (c *Client) doRequest(req *http.Request, authToken *string) ([]byte, error) {
	token := c.Token
	//logger.Info(c.ctx, "---------------------------------------------")
	//logger.Info(c.ctx, fmt.Sprintf("Url: %s", req.URL))

	if authToken != nil {
		token = *authToken
	}

	req.Header.Set("csp-auth-token", token)

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	fmt.Printf("\n-------------------------------------\n")
	fmt.Printf("request.URL: %s\n", req.URL)
	fmt.Printf("request.METHOD: %s\n", req.Method)
	fmt.Printf("\n-------------------------------------\n")
	//fmt.Printf("%s", body)
	fmt.Printf("\n-------------------------------------\n")

	if err != nil {
		return nil, err
	}

	// check that this was successful (perhaps shoudl check if 500?)
	httpOK := res.StatusCode == http.StatusOK
	httpSubmitted := res.StatusCode == http.StatusAccepted

	if !(httpSubmitted || httpOK) {
		return nil, fmt.Errorf("status: %d, body: %s", res.StatusCode, body)
	}

	return body, err
}
