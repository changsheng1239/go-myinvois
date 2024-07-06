package myinvois

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"strings"
)

//	type PlatformAPI interface {
//		LoginAsTaxpayer()
//		LoginAsIntermediary()
//		GetAllDocumentTypes()
//		GetDocumentType()
//		GetDocumentTypeVersion()
//		GetNotifications()
//	}

const (
	LoginEndpoint = "/connect/token"

	DefaultGrantType = "client_credentials"
	DefaultScope     = "InvoicingAPI"
)

var (
	ErrHttpRequestFailed = errors.New("http request failed")
	ErrReadBodyFailed    = errors.New("failed to read response body")
)

type platformEndpoints struct {
	loginAsTaxpayer *url.URL
}

var (
	PlatformEndpoints = &platformEndpoints{
		loginAsTaxpayer: &url.URL{Path: LoginEndpoint},
	}
)

type PlatformAPI struct {
	c *Client
}

func NewPlatformClient(c *Client) PlatformAPI {
	return PlatformAPI{c: c}
}

type OAuth2Token struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	TokenType   string `json:"token_type"`
	Scope       string `json:"scope"`
}

func (p *PlatformAPI) LoginAsTaxpayer(clientId, clientSecret string) (*OAuth2Token, error) {
	endpoint := p.c.Endpoint.ResolveReference(PlatformEndpoints.loginAsTaxpayer)
	form := url.Values{}
	form.Add("client_id", clientId)
	form.Add("client_secret", clientSecret)
	form.Add("grant_type", DefaultGrantType)
	form.Add("scope", DefaultScope)

	res, err := p.c.httpClient.Post(endpoint.String(), "application/x-www-form-urlencoded", strings.NewReader(form.Encode()))
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrHttpRequestFailed, err)
	}
	defer res.Body.Close()

	var token OAuth2Token
	err = json.NewDecoder(res.Body).Decode(&token)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrReadBodyFailed, err)
	}

	return &token, nil
}

func (p *PlatformAPI) LoginAsIntermediaries(clientId, clientSecret, onbehalfof string) (*OAuth2Token, error) {
	endpoint := p.c.Endpoint.ResolveReference(PlatformEndpoints.loginAsTaxpayer)
	form := url.Values{}
	form.Add("client_id", clientId)
	form.Add("client_secret", clientSecret)
	form.Add("grant_type", DefaultGrantType)
	form.Add("scope", DefaultScope)

	res, err := p.c.httpClient.Post(endpoint.String(), "application/x-www-form-urlencoded", strings.NewReader(form.Encode()))
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrHttpRequestFailed, err)
	}
	defer res.Body.Close()

	var token OAuth2Token
	err = json.NewDecoder(res.Body).Decode(&token)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrReadBodyFailed, err)
	}

	return &token, nil
}
