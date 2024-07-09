package myinvois

import (
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
	Endpoint   *url.URL
	httpClient *http.Client
	PlatformAPI
	EInvoiceAPI
}

type Option struct {
	Environment Environment
	Timeout     time.Duration
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
	c := &Client{
		Endpoint:   u,
		httpClient: httpClient,
	}
	c.PlatformAPI = NewPlatformClient(c)
	c.EInvoiceAPI = NewEInvoiceClient(c)
	return c
}

func SandboxClient() *Client {
	return newClient(Option{Environment: Sandbox, Timeout: DefaultTimeout})
}

func ProductionClient() *Client {
	return newClient(Option{Environment: Production, Timeout: DefaultTimeout})
}
