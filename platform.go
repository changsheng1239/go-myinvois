package myinvois

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
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

var (
	ErrInvalidTokenStructure = errors.New("invalid token structure")
	ErrNewHttpRequestFailed  = errors.New("failed to create new http request")
	ErrHttpRequestFailed     = errors.New("http request failed")
	ErrReadBodyFailed        = errors.New("failed to read response body")
	ErrRequestError          = errors.New("http request status not OK")
)

type PlatformAPI struct {
	myInvoisEndpoint *url.URL
	httpClient       *http.Client
	clientID         string
	clientSecret     string
}

func NewPlatformClient(endpoint *url.URL, httpClient *http.Client, clientID, clientSecret string) PlatformAPI {
	return PlatformAPI{
		myInvoisEndpoint: endpoint,
		httpClient:       httpClient,
		clientID:         clientID,
		clientSecret:     clientSecret,
	}
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

// LoginAsTaxpayer logs in as a taxpayer
// api signature: POST /connect/token
func (p *PlatformAPI) LoginAsTaxpayer() (*OAuth2Token, error) {
	endpoint := p.myInvoisEndpoint.ResolveReference(PlatformEndpoints.loginAsTaxpayer)
	form := url.Values{}
	form.Add("client_id", p.clientID)
	form.Add("client_secret", p.clientSecret)
	form.Add("grant_type", defaultGrantType)
	form.Add("scope", defaultScope)

	res, err := p.httpClient.Post(endpoint.String(), "application/x-www-form-urlencoded", strings.NewReader(form.Encode()))
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

// LoginAsIntermediaries logs in as an intermediary
// api signature: POST /connect/token
// require header: onbehalfof (taxpayer tin)
func (p *PlatformAPI) LoginAsIntermediaries(onbehalfof string) (*OAuth2Token, error) {
	endpoint := p.myInvoisEndpoint.ResolveReference(PlatformEndpoints.loginAsTaxpayer)

	form := url.Values{}
	form.Add("client_id", p.clientID)
	form.Add("client_secret", p.clientSecret)
	form.Add("grant_type", defaultGrantType)
	form.Add("scope", defaultScope)

	req, err := http.NewRequest("POST", endpoint.String(), strings.NewReader(form.Encode()))
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrNewHttpRequestFailed, err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("onbehalfof", onbehalfof)

	res, err := p.httpClient.Do(req)
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

// GetAllDocumentTypes retrieves all document types
// api signature: GET /api/v1.0/documenttypes
func (p *PlatformAPI) GetAllDocumentTypes(token string) (*DocumentTypes, error) {
	endpoint := p.myInvoisEndpoint.ResolveReference(PlatformEndpoints.getAllDocumentTypes)
	req, err := newAuthenticatedRequest("GET", endpoint.String(), token, nil)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrNewHttpRequestFailed, err)
	}

	res, err := p.httpClient.Do(req)
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

// GetDocumentType retrieves a document type by id
// api signature: GET /api/v1.0/documenttypes/{id}
func (p *PlatformAPI) GetDocumentType(token string, id int) (*DocumentType, error) {
	endpoint := p.myInvoisEndpoint.ResolveReference(PlatformEndpoints.getDocumentType)
	endpoint.Path = endpoint.Path + fmt.Sprintf("/%d", id)
	req, err := newAuthenticatedRequest("GET", endpoint.String(), token, nil)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrNewHttpRequestFailed, err)
	}

	res, err := p.httpClient.Do(req)
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
