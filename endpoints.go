package myinvois

import (
	"fmt"
	"net/url"
)

const (
	apiPrefixV10 = "/api/v1.0"

	// platform
	loginEndpoint         = "/connect/token"
	documentTypesEndpoint = "/documenttypes"

	// einvoice
	validateTaxpayerTinEndpoint  = "/taxpayer/validate"
	submitDocumentsEndpoint      = "/documentsubmissions"
	getDocumentsEndpoint         = "/documents"
	updateDocumentStatusEndpoint = "/documents/state/%s/state"
	getRecentDocumentsEndpoint   = "/documents/recent"

	defaultGrantType = "client_credentials"
	defaultScope     = "InvoicingAPI"
)

type MyInvoisBaseURL struct {
	API    *url.URL
	Portal *url.URL
}

type platformEndpoints struct {
	login               *url.URL
	getAllDocumentTypes *url.URL
	getDocumentType     *url.URL
}

type einvoiceEndpoints struct {
	validateTaxpayerTIN *url.URL
	submitDocuments     *url.URL
	getDocuments        *url.URL
	getRecentDocuments  *url.URL
}

var (
	SandboxBaseURL = MyInvoisBaseURL{
		API: &url.URL{
			Scheme: "https",
			Host:   "preprod-api.myinvois.hasil.gov.my",
		},
		Portal: &url.URL{
			Scheme: "https",
			Host:   "preprod.myinvois.hasil.gov.my",
		},
	}

	ProdBaseURL = MyInvoisBaseURL{
		API: &url.URL{
			Scheme: "https",
			Host:   "api.myinvois.hasil.gov.my",
		},
		Portal: &url.URL{
			Scheme: "https",
			Host:   "myinvois.hasil.gov.my",
		},
	}

	PlatformEndpoints = &platformEndpoints{
		login:               &url.URL{Path: loginEndpoint},
		getAllDocumentTypes: &url.URL{Path: apiPrefixV10 + documentTypesEndpoint},
		getDocumentType:     &url.URL{Path: apiPrefixV10 + documentTypesEndpoint},
	}

	EinvoiceEndpoints = &einvoiceEndpoints{
		validateTaxpayerTIN: &url.URL{Path: apiPrefixV10 + validateTaxpayerTinEndpoint},
		submitDocuments:     &url.URL{Path: apiPrefixV10 + submitDocumentsEndpoint},
		getDocuments:        &url.URL{Path: apiPrefixV10 + getDocumentsEndpoint},
		getRecentDocuments:  &url.URL{Path: apiPrefixV10 + getRecentDocumentsEndpoint},
	}
)

func (e *einvoiceEndpoints) updateDocumentStatus(uuid string) *url.URL {
	return &url.URL{Path: apiPrefixV10 + fmt.Sprintf(updateDocumentStatusEndpoint, uuid)}
}
