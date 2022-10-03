package client

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

// SignIn - Get a new token for user
func (c *Client) SignIn(ctx context.Context) (*AuthResponse, error) {
	//var body []byte
	var resp *http.Response

	if c.Auth.Apikey == "" {
		return nil, fmt.Errorf("apikey missing from tsm provider config")
	}
	// rb, err := json.Marshal(c.Auth)
	// if err != nil {
	// 	return nil, err
	// }

	data := url.Values{}
	data.Set("refresh_token", c.Auth.Apikey)

	req, err := http.NewRequestWithContext(ctx, "POST", c.AuthURL, strings.NewReader(data.Encode()))
	if err == nil {
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		httputil.DumpRequestOut(req, true)
		resp, err = (&http.Client{}).Do(req)
	}

	if err == nil {
		defer resp.Body.Close()
		httputil.DumpResponse(resp, true)
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		ar := AuthResponse{}
		err = json.Unmarshal(body, &ar)
		if err != nil {
			return nil, err
		}
		return &ar, nil
	}

	return nil, nil
}
