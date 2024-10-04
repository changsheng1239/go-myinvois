package myinvois

import (
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
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
	var baseURL MyInvoisBaseURL
	if opt.Environment == "" || opt.Environment == Sandbox {
		baseURL = SandboxBaseURL
	} else {
		baseURL = ProdBaseURL
	}

	// create a httpClient with timeout
	httpClient := &http.Client{
		Timeout: opt.Timeout,
		Transport: &http.Transport{
			// Example timeouts, adjust as needed
			IdleConnTimeout:     90 * time.Second,
			TLSHandshakeTimeout: 10 * time.Second,
		},
	}

	certWrapper, err := newCertWrapper(opt.Cert)
	if err != nil {
		log.Fatalf("NewCertWrapper failed: %v", err)
	}

	c := &Client{
		PlatformAPI: newPlatformClient(baseURL, httpClient, opt.ClientID, opt.ClientSecret),
		EInvoiceAPI: newEInvoiceClient(baseURL, httpClient, *certWrapper, mustParsePrivateKey(opt.PrivKey, opt.PrivKeyPass)),
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

func mustParsePrivateKey(privKey, passphrase []byte) *rsa.PrivateKey {
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

type x509CertWrapper struct {
	base64       string
	digest       string
	issuer       string
	serialNumber string
	subject      string
}

func newCertWrapper(cert []byte) (*x509CertWrapper, error) {
	block, _ := pem.Decode(cert)
	if block == nil {
		return nil, ErrDecodeCertificate
	}

	c, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrParseCertificate, err)
	}

	h := sha256.New()
	h.Write(block.Bytes)

	return &x509CertWrapper{
		base64:       base64.StdEncoding.EncodeToString(block.Bytes),
		digest:       base64.StdEncoding.EncodeToString(h.Sum(nil)),
		issuer:       strings.Join(strings.Split(c.Issuer.String(), ","), ", "),
		serialNumber: c.SerialNumber.String(),
		subject:      c.Subject.String(),
	}, nil
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
