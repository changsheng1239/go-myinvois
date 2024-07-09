package myinvois

import (
	"encoding/base64"
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
	ErrInvalidTokenStructure = errors.New("invalid token structure")
	ErrHttpRequestFailed     = errors.New("http request failed")
	ErrReadBodyFailed        = errors.New("failed to read response body")
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

type TokenPayload struct {
	Iss               string   `json:"iss"`
	Nbf               int      `json:"nbf"`
	Iat               int      `json:"iat"`
	Exp               int      `json:"exp"`
	Aud               []string `json:"aud"`
	Scope             []string `json:"scope"`
	ClientID          string   `json:"client_id"`
	IsTaxRepres       string   `json:"IsTaxRepres"`
	IsIntermediary    string   `json:"IsIntermediary"`
	IntermedID        string   `json:"IntermedId"`
	IntermedTIN       string   `json:"IntermedTIN"`
	IntermedEnforced  string   `json:"IntermedEnforced"`
	Name              string   `json:"name"`
	SSID              string   `json:"SSId"`
	PreferredUsername string   `json:"preferred_username"`
	TaxID             string   `json:"TaxId"`
	TaxTin            string   `json:"TaxTin"`
	ProfID            string   `json:"ProfId"`
	IsTaxAdmin        string   `json:"IsTaxAdmin"`
	IsSystem          string   `json:"IsSystem"`
	NatID             string   `json:"NatId"`
}

func DecodeToken(token string) (*TokenPayload, error) {
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		return nil, ErrInvalidTokenStructure
	}

	// decode base64 parts[1]
	claims := parts[1]
	decodedPayload, err := decodeSegment(claims)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrInvalidTokenStructure, err)
	}

	var payload TokenPayload
	err = json.Unmarshal(decodedPayload, &payload)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrInvalidTokenStructure, err)
	}

	return &payload, nil
}

func decodeSegment(seg string) ([]byte, error) {
	if l := len(seg) % 4; l > 0 {
		seg += strings.Repeat("=", 4-l)
	}
	encoding := base64.URLEncoding

	return encoding.DecodeString(seg)
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
