package myinvois

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"strings"
	"time"
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
	apiPrefixV10          = "/api/v1.0"
	loginEndpoint         = "/connect/token"
	documentTypesEndpoint = "/documenttypes"

	defaultGrantType = "client_credentials"
	defaultScope     = "InvoicingAPI"
)

var (
	ErrInvalidTokenStructure = errors.New("invalid token structure")
	ErrNewHttpRequestFailed  = errors.New("failed to create new http request")
	ErrHttpRequestFailed     = errors.New("http request failed")
	ErrReadBodyFailed        = errors.New("failed to read response body")
)

type platformEndpoints struct {
	loginAsTaxpayer       *url.URL
	loginAsIntermediaries *url.URL
	getAllDocumentTypes   *url.URL
	getDocumentType       *url.URL
}

var (
	PlatformEndpoints = &platformEndpoints{
		loginAsTaxpayer:       &url.URL{Path: loginEndpoint},
		loginAsIntermediaries: &url.URL{Path: loginEndpoint},
		getAllDocumentTypes:   &url.URL{Path: apiPrefixV10 + documentTypesEndpoint},
		getDocumentType:       &url.URL{Path: apiPrefixV10 + documentTypesEndpoint},
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

type DocumentTypes struct {
	Result []DocumentTypesResult `json:"result"`
}

type DocumentTypesResult struct {
	ID                   int64                 `json:"id"`
	InvoiceTypeCode      int64                 `json:"invoiceTypeCode"`
	Description          string                `json:"description"`
	ActiveFrom           time.Time             `json:"activeFrom"`
	ActiveTo             time.Time             `json:"activeTo"`
	DocumentTypeVersions []DocumentTypeVersion `json:"documentTypeVersions"`
}

type DocumentTypeVersion struct {
	ID            int64     `json:"id"`
	Name          string    `json:"name"`
	Description   string    `json:"description"`
	ActiveFrom    time.Time `json:"activeFrom"`
	ActiveTo      time.Time `json:"activeTo"`
	VersionNumber float64   `json:"versionNumber"`
	Status        string    `json:"status"`
}

type DocumentType struct {
	ID                   int64                 `json:"id"`
	InvoiceTypeCode      int64                 `json:"invoiceTypeCode"`
	Description          string                `json:"description"`
	ActiveFrom           time.Time             `json:"activeFrom"`
	ActiveTo             time.Time             `json:"activeTo"`
	DocumentTypeVersions []DocumentTypeVersion `json:"documentTypeVersions"`
	WorkflowParameters   []WorkflowParameter   `json:"workflowParameters"`
}

type WorkflowParameter struct {
	ID         int64     `json:"id"`
	Parameter  string    `json:"parameter"`
	Value      int64     `json:"value"`
	ActiveFrom time.Time `json:"activeFrom"`
	ActiveTo   time.Time `json:"activeTo"`
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
	form.Add("grant_type", defaultGrantType)
	form.Add("scope", defaultScope)

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
	form.Add("grant_type", defaultGrantType)
	form.Add("scope", defaultScope)

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

func (p *PlatformAPI) GetAllDocumentTypes(token string) (*DocumentTypes, error) {
	endpoint := p.c.Endpoint.ResolveReference(PlatformEndpoints.getAllDocumentTypes)
	req, err := newAuthenticatedRequest("GET", endpoint.String(), token, nil)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrNewHttpRequestFailed, err)
	}

	res, err := p.c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrHttpRequestFailed, err)
	}
	defer res.Body.Close()

	var documentTypes DocumentTypes
	err = json.NewDecoder(res.Body).Decode(&documentTypes)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrReadBodyFailed, err)
	}

	return &documentTypes, nil
}

func (p *PlatformAPI) GetDocumentType(token string, id int) (*DocumentType, error) {
	endpoint := p.c.Endpoint.ResolveReference(PlatformEndpoints.getDocumentType)
	endpoint.Path = endpoint.Path + fmt.Sprintf("/%d", id)
	req, err := newAuthenticatedRequest("GET", endpoint.String(), token, nil)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrNewHttpRequestFailed, err)
	}

	res, err := p.c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrHttpRequestFailed, err)
	}
	defer res.Body.Close()

	var documentType DocumentType
	err = json.NewDecoder(res.Body).Decode(&documentType)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrReadBodyFailed, err)
	}

	return &documentType, nil
}
