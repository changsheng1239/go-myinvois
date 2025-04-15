package myinvois

import (
	"bytes"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/antchfx/xmlquery"
)

var (
	ErrAuthenticatedTinMismatch = errors.New("The authenticated TIN and documents TIN is not matching")
	ErrMultipleTinMatched       = errors.New("Search criteria in not conclusive and more than one TIN can be found to match the search criteria provided, please revise the search criteria.")
	ErrInvalidInput             = errors.New("invalid input")
	ErrTinMismatch              = errors.New("ID value and TIN mismatch")
	ErrMarshalFailed            = errors.New("failed to marshal struct")
	ErrIncorrectSubmitter       = errors.New("incorrect submitter")
	ErrDuplicateSubmission      = errors.New("duplicate submission, please retry later")
	ErrDecodeCertificate        = errors.New("failed to decode certificate")
	ErrParseCertificate         = errors.New("failed to parse certificate")
	ErrSignFailed               = errors.New("failed to sign document")
)

type EInvoiceAPI struct {
	baseURL    MyInvoisBaseURL
	httpClient *http.Client
	cert       x509CertWrapper
	privKey    *rsa.PrivateKey
}

func newEInvoiceClient(baseURL MyInvoisBaseURL, httpClient *http.Client, cert x509CertWrapper, pk *rsa.PrivateKey) EInvoiceAPI {
	return EInvoiceAPI{
		baseURL:    baseURL,
		httpClient: httpClient,
		cert:       cert,
		privKey:    pk,
	}
}

type DocumentSubmission struct {
	Documents []Document `json:"documents"`
}

type Document struct {
	Format       string `json:"format"`
	DocumentHash string `json:"documentHash"`
	CodeNumber   string `json:"codeNumber"`
	Document     string `json:"document"`
}

type DocumentSubmissionResponse struct {
	SubmissionUid     string             `json:"submissionUid"`
	AcceptedDocuments []AcceptedDocument `json:"acceptedDocuments"`
	RejectedDocuments []RejectedDocument `json:"rejectedDocuments"`
}

type AcceptedDocument struct {
	UUID              string `json:"uuid"`
	InvoiceCodeNumber string `json:"invoiceCodeNumber"`
}

type RejectedDocument struct {
	InvoiceCodeNumber string `json:"invoiceCodeNumber"`
	StandardErrorResponse
}

type GetDocumentDetailsResponse struct {
	UUID                  string            `json:"uuid"`
	SubmissionUid         string            `json:"submissionUid"`
	LongID                string            `json:"longId"`
	TypeName              string            `json:"typeName"`
	TypeVersionName       string            `json:"typeVersionName"`
	IssuerTin             string            `json:"issuerTin"`
	IssuerName            string            `json:"issuerName"`
	ReceiverID            string            `json:"receiverId"`
	ReceiverName          string            `json:"receiverName"`
	DateTimeReceived      time.Time         `json:"dateTimeReceived"`
	DateTimeValidated     time.Time         `json:"dateTimeValidated"`
	TotalExcludingTax     float64           `json:"totalExcludingTax"`
	TotalDiscount         float64           `json:"totalDiscount"`
	TotalNetAmount        float64           `json:"totalNetAmount"`
	TotalPayableAmount    float64           `json:"totalPayableAmount"`
	Status                string            `json:"status"`
	CreatedByUserID       string            `json:"createdByUserId"`
	DocumentStatusReason  string            `json:"documentStatusReason"`
	CancelDateTime        time.Time         `json:"cancelDateTime"`
	RejectRequestDateTime time.Time         `json:"rejectRequestDateTime"`
	ValidationResults     ValidationResults `json:"validationResults"`
	InternalID            string            `json:"internalId"`
	DateTimeIssued        time.Time         `json:"dateTimeIssued"`
	Document              json.RawMessage   `json:"document"`
}

type ValidationResults struct {
	Status          string           `json:"status"`
	ValidationSteps []ValidationStep `json:"validationSteps"`
}

type ValidationStep struct {
	Status string              `json:"status"`
	Name   string              `json:"name"`
	Error  ValidationStepError `json:"error,omitempty"`
}

type ValidationStepError struct {
	PropertyName *string               `json:"propertyName,omitempty"`
	PropertyPath *string               `json:"propertyPath,omitempty"`
	ErrorCode    string                `json:"errorCode,omitempty"`
	Error        string                `json:"error,omitempty"`
	ErrorMS      string                `json:"errorMs,omitempty"`
	InnerError   []ValidationStepError `json:"innerError,omitempty"`
}

type UpdateStatusRequest struct {
	Status string `json:"status"`
	Reason string `json:"reason"`
}

type UpdateStatusResponse struct {
	UUID   string `json:"uuid"`
	Status string `json:"status"`
	StandardErrorResponse
}

type GetDocumentsOption struct {
	PageNo             int
	PageSize           int
	SubmissionDateFrom time.Time
	SubmissionDateTo   time.Time
	IssueDateFrom      time.Time
	IssueDateTo        time.Time
	Direction          string
	Status             string
	DocumentType       string
	ReceiverIdType     string
	ReceiverId         string
	IssuerIdType       string
	IssuerId           string
	ReceiverTin        string
	IssuerTin          string
}

type TaxpayerInfo struct {
	Name                          string `json:"name"`
	TIN                           string `json:"tin"`
	IDType                        string `json:"idType"`
	IDNumber                      string `json:"idNumber"`
	SST                           string `json:"sst"`
	Email                         string `json:"email"`
	ContactNumber                 string `json:"contactNumber"`
	TTX                           string `json:"ttx"`
	BusinessActivityDescriptionBM string `json:"businessActivityDescriptionBM"`
	BusinessActivityDescriptionEN string `json:"businessActivityDescriptionEN"`
	MSIC                          string `json:"msic"`
	AddressLine0                  string `json:"addressLine0"`
	AddressLine1                  string `json:"addressLine1"`
	AddressLine2                  string `json:"addressLine2"`
	PostalZone                    string `json:"postalZone"`
	City                          string `json:"city"`
	State                         string `json:"state"`
	Country                       string `json:"country"`
	GeneratedTimestamp            string `json:"generatedTimestamp"`
}

const (
	stDocumentPending   = "Pending"
	stDocumentSubmitted = "Submitted"
	stDocumentValid     = "Valid"
	stDocumentInvalid   = "Invalid"
	stDocumentCancelled = "Cancelled"
	stDocumentRejected  = "Rejected"
)

// ValidateTaxpayerTIN validates the taxpayer TIN
// api signature: GET /api/v1.0/taxpayer/validate/{tin}?idType={idType}&idValue={idValue}
func (e *EInvoiceAPI) ValidateTaxpayerTIN(accessToken, tin, idType, idValue string) (bool, error) {
	endpoint := e.baseURL.API.ResolveReference(EinvoiceEndpoints.validateTaxpayerTIN)
	endpoint.Path = endpoint.Path + fmt.Sprintf("/%s", tin)

	q := endpoint.Query()
	q.Set("idType", idType)
	q.Set("idValue", idValue)
	endpoint.RawQuery = q.Encode()

	req, err := newRequestWithToken(accessToken, http.MethodGet, endpoint.String(), nil)
	if err != nil {
		return false, fmt.Errorf("%w: %v", ErrNewHttpRequestFailed, err)
	}

	res, err := e.httpClient.Do(req)
	if err != nil {
		return false, fmt.Errorf("%w: %v", ErrHttpRequestFailed, err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		switch res.StatusCode {
		case 400:
			return false, fmt.Errorf("%w: %v", ErrInvalidInput, res.Status)
		case 404:
			// 404 means the TIN & ID combination does not match
			return false, nil
		default:
			return false, fmt.Errorf("%w: %v", ErrHttpRequestFailed, res.Status)
		}
	}

	return true, nil
}

// SearchTaxpayerTIN searches the taxpayer TIN
// api signature: GET /api/v1.0/taxpayer/search/tin?idType={idType}&idValue={idValue}&taxpayerName={taxpayerName}
func (e *EInvoiceAPI) SearchTaxpayerTIN(accessToken, idType, idValue, taxpayerName string) (string, error) {
	endpoint := e.baseURL.API.ResolveReference(EinvoiceEndpoints.searchTaxpayerTIN)

	q := endpoint.Query()
	if idType != "" && idValue != "" {
		q.Set("idType", idType)
		q.Set("idValue", idValue)
	} else if taxpayerName != "" {
		q.Set("taxpayerName", taxpayerName)
	} else {
		return "", ErrInvalidInput
	}
	endpoint.RawQuery = q.Encode()

	req, err := newRequestWithToken(accessToken, http.MethodGet, endpoint.String(), nil)
	if err != nil {
		return "", fmt.Errorf("%w: %v", ErrNewHttpRequestFailed, err)
	}

	res, err := e.httpClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("%w: %v", ErrHttpRequestFailed, err)
	}
	defer res.Body.Close()

	b, err := io.ReadAll(res.Body)
	if err != nil {
		return "", fmt.Errorf("%w: %v", ErrReadBodyFailed, err)
	}

	if res.StatusCode != 200 {
		var r string
		switch res.StatusCode {
		case 400:
			err = json.Unmarshal(b, &r)
			if err != nil {
				return "", fmt.Errorf("%w: %v", ErrReadBodyFailed, err)
			}
			if r == ErrMultipleTinMatched.Error() {
				return "", ErrMultipleTinMatched
			}
			return "", fmt.Errorf("%w: %v", ErrInvalidInput, res.Status)
		case 404:
			// 404 means the TIN & ID combination does not match
			return "", nil
		default:
			return "", fmt.Errorf("%w: %v", ErrHttpRequestFailed, res.Status)
		}
	}

	var r struct {
		TIN string `json:"tin"`
	}
	err = json.Unmarshal(b, &r)
	if err != nil {
		return "", fmt.Errorf("%w: %v", ErrReadBodyFailed, err)
	}

	return r.TIN, nil
}

// SubmitDocuments submits documents to the LHDN MyInvois API with Digital Signature
// Support JSON format only
// api signature: POST /api/v1.0/documentsubmissions/
func (e *EInvoiceAPI) SubmitDocuments(accessToken string, docs []Ubl21Invoice) (*DocumentSubmissionResponse, error) {
	endpoint := e.baseURL.API.ResolveReference(EinvoiceEndpoints.submitDocuments)

	var ds DocumentSubmission
	for _, doc := range docs {
		if len(doc.Invoice) == 0 {
			return nil, ErrInvalidInput
		}

		d := &doc
		if doc.Invoice[0].InvoiceTypeCode[0].ListVersionID != "1.0" {
			signedDoc, err := signDocument(e.privKey, doc, e.cert)
			if err != nil {
				return nil, fmt.Errorf("%w: %v", ErrSignFailed, err)
			}
			d = signedDoc
		}
		b, err := json.Marshal(d)
		if err != nil {
			return nil, fmt.Errorf("%w: %v", ErrMarshalFailed, err)
		}

		h := sha256.New()
		h.Write(b)

		ds.Documents = append(ds.Documents, Document{
			Format:       "json",
			DocumentHash: hex.EncodeToString(h.Sum(nil)),
			CodeNumber:   doc.Invoice[0].ID[0].Empty,
			Document:     base64.StdEncoding.EncodeToString(b),
		})
	}

	b, err := json.Marshal(ds)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrMarshalFailed, err)
	}

	req, err := newRequestWithToken(accessToken, http.MethodPost, endpoint.String(), bytes.NewReader(b))
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrNewHttpRequestFailed, err)
	}

	res, err := e.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrHttpRequestFailed, err)
	}
	defer res.Body.Close()

	if res.StatusCode != 202 {
		if res.StatusCode == http.StatusBadRequest {
			var sErr StandardErrorResponse
			err = json.NewDecoder(res.Body).Decode(&sErr)
			if err != nil {
				return nil, fmt.Errorf("%w: %v", ErrReadBodyFailed, err)
			}
			errMsg := sErr.Error.Code + ": "
			for _, e := range sErr.Error.Details {
				errMsg += e.Message + "\n"
			}
			return nil, errors.New(errMsg)
		} else if res.StatusCode == http.StatusForbidden {
			return nil, ErrIncorrectSubmitter
		} else if res.StatusCode == http.StatusUnprocessableEntity {
			return nil, ErrDuplicateSubmission
		} else {
			return nil, fmt.Errorf("%w: %v", ErrHttpRequestFailed, res.Status)
		}
	}

	var r DocumentSubmissionResponse
	err = json.NewDecoder(res.Body).Decode(&r)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrReadBodyFailed, err)
	}

	return &r, nil
}

// SubmitDocuments submits documents to the LHDN MyInvois API
// Support XML format without Digital Signature (Invoice v1.0)
// api signature: POST /api/v1.0/documentsubmissions/
func (e *EInvoiceAPI) SubmitRawXML(accessToken string, docXML []byte) (*DocumentSubmissionResponse, error) {
	endpoint := e.baseURL.API.ResolveReference(EinvoiceEndpoints.submitDocuments)

	h := sha256.New()
	h.Write(docXML)

	doc, err := xmlquery.Parse(bytes.NewReader(docXML))
	if err != nil {
		panic(err)
	}
	invoiceNode := xmlquery.FindOne(doc, "//Invoice")
	var invoiceID string
	if n := invoiceNode.SelectElement("cbc:ID"); n != nil {
		invoiceID = n.InnerText()
	}

	ds := DocumentSubmission{
		Documents: []Document{
			{
				Format:       "xml",
				DocumentHash: hex.EncodeToString(h.Sum(nil)),
				CodeNumber:   invoiceID,
				Document:     base64.StdEncoding.EncodeToString(docXML),
			},
		},
	}

	b, err := json.Marshal(ds)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrMarshalFailed, err)
	}

	req, err := newRequestWithToken(accessToken, http.MethodPost, endpoint.String(), bytes.NewReader(b))
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrNewHttpRequestFailed, err)
	}

	res, err := e.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrHttpRequestFailed, err)
	}

	if res.StatusCode != 202 {
		if res.StatusCode == http.StatusBadRequest {
			var sErr StandardErrorResponse
			err = json.NewDecoder(res.Body).Decode(&sErr)
			if err != nil {
				return nil, fmt.Errorf("%w: %v", ErrReadBodyFailed, err)
			}
			errMsg := sErr.Error.Code + ": "
			for _, e := range sErr.Error.Details {
				errMsg += e.Message + "\n"
			}
			return nil, errors.New(errMsg)
		} else if res.StatusCode == http.StatusForbidden {
			return nil, ErrIncorrectSubmitter
		} else if res.StatusCode == http.StatusUnprocessableEntity {
			return nil, ErrDuplicateSubmission
		} else {
			return nil, fmt.Errorf("%w: %v", ErrHttpRequestFailed, res.Status)
		}
	}

	var r DocumentSubmissionResponse
	err = json.NewDecoder(res.Body).Decode(&r)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrReadBodyFailed, err)
	}

	return &r, nil
}

// GetDocumentDetails retrieves the status & details of a document
// api signature: GET /api/v1.0/documents/{uuid}/raw
func (e *EInvoiceAPI) GetDocument(accessToken, uuid string) (*GetDocumentDetailsResponse, error) {
	endpoint := e.baseURL.API.ResolveReference(EinvoiceEndpoints.getDocuments)
	endpoint.Path = endpoint.Path + fmt.Sprintf("/%s/raw", uuid)

	req, err := newRequestWithToken(accessToken, http.MethodGet, endpoint.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrNewHttpRequestFailed, err)
	}

	res, err := e.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrHttpRequestFailed, err)
	}

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("%w: %v", ErrHttpRequestFailed, res.Status)
	}

	var r GetDocumentDetailsResponse
	err = json.NewDecoder(res.Body).Decode(&r)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrReadBodyFailed, err)
	}

	return &r, nil
}

// GetDocumentDetails retrieves the status & details of a document
// api signature: GET /api/v1.0/documents/{uuid}/details
func (e *EInvoiceAPI) GetDocumentDetails(accessToken, uuid string) (*GetDocumentDetailsResponse, error) {
	endpoint := e.baseURL.API.ResolveReference(EinvoiceEndpoints.getDocuments)
	endpoint.Path = endpoint.Path + fmt.Sprintf("/%s/details", uuid)

	req, err := newRequestWithToken(accessToken, http.MethodGet, endpoint.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrNewHttpRequestFailed, err)
	}

	res, err := e.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrHttpRequestFailed, err)
	}

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("%w: %v", ErrHttpRequestFailed, res.Status)
	}

	var r GetDocumentDetailsResponse
	err = json.NewDecoder(res.Body).Decode(&r)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrReadBodyFailed, err)
	}

	return &r, nil
}

// CancelDocument cancels a document
// api signature: PUT /api/v1.0/documents/state/{UUID}/state
func (e *EInvoiceAPI) CancelDocument(accessToken, uuid, reason string) (*UpdateStatusResponse, error) {
	return e.updateDocumentStatus(accessToken, uuid, stDocumentCancelled, reason)
}

// RejectDocument rejects a document
// api signature: PUT /api/v1.0/documents/state/{UUID}/state
func (e *EInvoiceAPI) RejectDocument(accessToken, uuid, reason string) (*UpdateStatusResponse, error) {
	return e.updateDocumentStatus(accessToken, uuid, stDocumentRejected, reason)
}

func (e *EInvoiceAPI) updateDocumentStatus(accessToken, uuid, status, reason string) (*UpdateStatusResponse, error) {
	endpoint := e.baseURL.API.ResolveReference(EinvoiceEndpoints.updateDocumentStatus(uuid))

	body := UpdateStatusRequest{
		Status: status,
		Reason: reason,
	}
	b, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrMarshalFailed, err)
	}

	req, err := newRequestWithToken(accessToken, http.MethodPut, endpoint.String(), bytes.NewReader(b))
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrNewHttpRequestFailed, err)
	}

	res, err := e.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrHttpRequestFailed, err)
	}

	if res.StatusCode != 200 {
		if res.StatusCode == http.StatusBadRequest {
			var sErr StandardErrorResponse
			err = json.NewDecoder(res.Body).Decode(&sErr)
			if err != nil {
				return nil, fmt.Errorf("%w: %v", ErrReadBodyFailed, err)
			}
			errMsg := sErr.Error.Code + ": "
			for _, e := range sErr.Error.Details {
				errMsg += e.Message + "\n"
			}
			return nil, errors.New(errMsg)
		}
		return nil, fmt.Errorf("%w: %v", ErrHttpRequestFailed, res.Status)
	}

	var r UpdateStatusResponse
	err = json.NewDecoder(res.Body).Decode(&r)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrReadBodyFailed, err)
	}

	return &r, nil
}

// GetRecentDocuments retrieves the recent documents (maximum 30 days ago)
// api signature: GET /api/v1.0/documents/recent?
// query params:
//
//	pageNo={pageNo}&
//	pageSize={pageSize}&
//	submissionDateFrom={submissionDateFrom}&
//	submissionDateTo={submissionDateTo}&
//	issueDateFrom={issueDateFrom}&
//	issueDateTo={IssueDateTo}&
//	direction={direction}&
//	status={status}&
//	documentType={documentType}&
//	receiverIdType={receiverIdType}&
//	receiverId={receiverId}&
//	issuerIdType={issuerIdType}&
//	issuerId={issuerId}&
//	receiverTin={receiverTin}&
//	issuerTin={issuerTin}
func (e *EInvoiceAPI) GetRecentDocuments(limit int) ([]GetDocumentDetailsResponse, error) {
	endpoint := e.baseURL.API.ResolveReference(EinvoiceEndpoints.getRecentDocuments)

	q := endpoint.Query()
	q.Set("pageSize", fmt.Sprintf("%d", limit))

	return nil, nil
}

// PublicLink returns the public link for a document
// {envbaseurl}/uuid-of-document/share/longid
func (e *EInvoiceAPI) PublicLink(uuid, longid string) string {
	return fmt.Sprintf("%s/%s/share/%s", e.baseURL.Portal.String(), uuid, longid)
}

// TaxpayerQrCode allows taxpayer’s ERP system to retrieve the information for a specific Taxpayer based on the Base64 formatted string obtained from scanning the respective QR code.
// api signature: GET /api/v1.0/taxpayers/qrcodeinfo/{qrCodeText}
func (e *EInvoiceAPI) TaxpayerQrCode(accessToken string, id string) (*TaxpayerInfo, error) {
	endpoint := e.baseURL.API.ResolveReference(EinvoiceEndpoints.taxpayerQrCodeInfo)
	endpoint.Path = endpoint.Path + fmt.Sprintf("/%s", id)

	req, err := newRequestWithToken(accessToken, http.MethodGet, endpoint.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrNewHttpRequestFailed, err)
	}
	res, err := e.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrHttpRequestFailed, err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("%w: %v", ErrHttpRequestFailed, res.Status)
	}
	var r TaxpayerInfo
	err = json.NewDecoder(res.Body).Decode(&r)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrReadBodyFailed, err)
	}

	return &r, nil
}
