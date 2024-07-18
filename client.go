package myinvois

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
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
	PlatformAPI
	EInvoiceAPI
}

type Option struct {
	Environment  Environment
	Timeout      time.Duration
	ClientID     string
	ClientSecret string
	Cert         []byte
	PrivKey      []byte
	PrivKeyPass  []byte
}

func newClient(opt Option) *Client {
	var apiURL string
	if opt.Environment == "" || opt.Environment == Sandbox {
		apiURL = DefaultSandboxURL
	} else {
		apiURL = DefaultProductionURL
	}
	u, _ := url.Parse(apiURL)
	httpClient := &http.Client{Timeout: opt.Timeout}

	certWrapper, err := NewCertWrapper(opt.Cert)
	if err != nil {
		log.Fatalf("NewCertWrapper failed: %v", err)
	}

	c := &Client{
		PlatformAPI: NewPlatformClient(u, httpClient, opt.ClientID, opt.ClientSecret),
		EInvoiceAPI: NewEInvoiceClient(u, httpClient, *certWrapper, MustParsePrivateKey(opt.PrivKey, opt.PrivKeyPass)),
	}
	return c
}

func SandboxClient(clientID, clientSecret string, cert, pk, pkPass []byte) *Client {
	return newClient(Option{
		Environment:  Sandbox,
		Timeout:      DefaultTimeout,
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Cert:         cert,
		PrivKey:      pk,
		PrivKeyPass:  pkPass,
	})
}

func ProductionClient(clientID, clientSecret string, cert, pk, pkPass []byte) *Client {
	return newClient(Option{
		Environment:  Production,
		Timeout:      DefaultTimeout,
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Cert:         cert,
		PrivKey:      pk,
		PrivKeyPass:  pkPass,
	})
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
