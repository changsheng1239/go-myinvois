package myinvois

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"
)

const (
	DefaultProductionURL = "https://api.myinvois.hasil.gov.my"
	DefaultSandboxURL    = "https://preprod-api.myinvois.hasil.gov.my"
	DefaultTimeout       = 30 * time.Second
)

const (
	Sandbox    Environment = "sandbox"
	Production Environment = "production"
)

type Environment string

type Client struct {
	ClientID     string
	ClientSecret string
	AccessToken  *OAuth2Token
	PlatformAPI
	EInvoiceAPI
}

type ClientOption struct {
	Environment  Environment
	Timeout      time.Duration
	ClientID     string
	ClientSecret string
	Cert         []byte
	PrivKey      []byte
	PrivKeyPass  []byte
}

func newClient(opt ClientOption) *Client {
	var apiURL string
	if opt.Environment == "" || opt.Environment == Sandbox {
		apiURL = DefaultSandboxURL
	} else {
		apiURL = DefaultProductionURL
	}
	u, _ := url.Parse(apiURL)

	// create a httpClient with timeout
	httpClient := &http.Client{
		Timeout: opt.Timeout,
		Transport: &http.Transport{
			// Example timeouts, adjust as needed
			IdleConnTimeout:     90 * time.Second,
			TLSHandshakeTimeout: 10 * time.Second,
		},
	}

	certWrapper, err := NewCertWrapper(opt.Cert)
	if err != nil {
		log.Fatalf("NewCertWrapper failed: %v", err)
	}

	c := &Client{
		PlatformAPI: newPlatformClient(u, httpClient, opt.ClientID, opt.ClientSecret),
		EInvoiceAPI: newEInvoiceClient(u, httpClient, *certWrapper, MustParsePrivateKey(opt.PrivKey, opt.PrivKeyPass)),
	}

	return c
}

func NewClient(opt ClientOption) *Client {
	if opt.Environment == "" || (opt.Environment != Sandbox && opt.Environment != Production) {
		log.Fatalf("Invalid environment: %v", opt.Environment)
	}

	if opt.ClientID == "" || opt.ClientSecret == "" {
		log.Fatalf("ClientID and ClientSecret are required")
	}

	if opt.Cert == nil || opt.PrivKey == nil || opt.PrivKeyPass == nil {
		log.Fatalf("Cert, PrivKey, and PrivKeyPass are required")
	}

	if opt.Timeout == 0 {
		opt.Timeout = DefaultTimeout
	}

	return newClient(opt)
}

func MustParsePrivateKey(privKey, passphrase []byte) *rsa.PrivateKey {
	block, _ := pem.Decode(privKey)
	der, err := x509.DecryptPEMBlock(block, passphrase)
	if err != nil {
		log.Fatalf("Decrypt failed: %v", err)
	}
	key, err := x509.ParsePKCS1PrivateKey(der)
	if err != nil {
		log.Fatalf("ParsePKCS1PrivateKey failed: %v", err)
	}
	return key
}

func newRequestWithToken(token string, httpMethod, endpoint string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(httpMethod, endpoint, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Accept-Language", "en")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	return req, nil
}
