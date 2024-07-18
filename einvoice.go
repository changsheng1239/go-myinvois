package myinvois

import (
	"bytes"
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

var (
	ErrInvalidInput        = errors.New("invalid input")
	ErrTinMismatch         = errors.New("ID value and TIN mismatch")
	ErrMarshalFailed       = errors.New("failed to marshal struct")
	ErrIncorrectSubmitter  = errors.New("incorrect submitter")
	ErrDuplicateSubmission = errors.New("duplicate submission, please retry later")
	ErrDecodeCertificate   = errors.New("failed to decode certificate")
	ErrParseCertificate    = errors.New("failed to parse certificate")
	ErrSignFailed          = errors.New("failed to sign document")
)

type EInvoiceAPI struct {
	myInvoisEndpoint *url.URL
	httpClient       *http.Client
	cert             x509CertWrapper
	privKey          *rsa.PrivateKey
}

func NewEInvoiceClient(endpoint *url.URL, httpClient *http.Client, cert x509CertWrapper, pk *rsa.PrivateKey) EInvoiceAPI {
	return EInvoiceAPI{
		myInvoisEndpoint: endpoint,
		httpClient:       httpClient,
		cert:             cert,
		privKey:          pk,
	}
}

type x509CertWrapper struct {
	base64       string
	digest       string
	issuer       string
	serialNumber string
	subject      string
}

func NewCertWrapper(cert []byte) (*x509CertWrapper, error) {
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

const (
	stDocumentValid     = "Valid"
	stDocumentCancelled = "Cancelled"
	stDocumentRejected  = "Rejected"
)

// ValidateTaxpayerTIN validates the taxpayer TIN
// api signature: GET /api/v1.0/taxpayer/validate/{tin}?idType={idType}&idValue={idValue}
func (e *EInvoiceAPI) ValidateTaxpayerTIN(token, tin, idType, idValue string) (bool, error) {
	endpoint := e.myInvoisEndpoint.ResolveReference(EinvoiceEndpoints.validateTaxpayerTIN)
	endpoint.Path = endpoint.Path + fmt.Sprintf("/%s", tin)

	q := endpoint.Query()
	q.Set("idType", idType)
	q.Set("idValue", idValue)
	endpoint.RawQuery = q.Encode()

	req, err := newAuthenticatedRequest("GET", endpoint.String(), token, nil)
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
			return false, fmt.Errorf("%w: %v", ErrTinMismatch, res.Status)
		default:
			return false, fmt.Errorf("%w: %v", ErrHttpRequestFailed, res.Status)
		}
	}

	return true, nil
}

// SubmitDocuments submits documents to the LHDN MyInvois API with Digital Signature
// api signature: POST /api/v1.0/documentsubmissions/
func (e *EInvoiceAPI) SubmitDocuments(token string, docs []Ubl21Invoice) (*DocumentSubmissionResponse, error) {
	endpoint := e.myInvoisEndpoint.ResolveReference(EinvoiceEndpoints.submitDocuments)

	var d DocumentSubmission
	for _, doc := range docs {
		signedDoc, err := signDocument(e.privKey, doc, e.cert)
		if err != nil {
			return nil, fmt.Errorf("%w: %v", ErrSignFailed, err)
		}
		b, err := json.Marshal(signedDoc)
		if err != nil {
			return nil, fmt.Errorf("%w: %v", ErrMarshalFailed, err)
		}

		h := sha256.New()
		h.Write(b)

		d.Documents = append(d.Documents, Document{
			Format:       "json",
			DocumentHash: hex.EncodeToString(h.Sum(nil)),
			CodeNumber:   doc.Invoice[0].ID[0].Empty,
			Document:     base64.StdEncoding.EncodeToString(b),
		})
	}

	b, err := json.Marshal(d)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrMarshalFailed, err)
	}

	req, err := newAuthenticatedRequest("POST", endpoint.String(), token, bytes.NewReader(b))
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
// api signature: GET /api/v1.0/documents/{uuid}/details
func (e *EInvoiceAPI) GetDocumentDetails(token string, uuid string) (*GetDocumentDetailsResponse, error) {
	endpoint := e.myInvoisEndpoint.ResolveReference(EinvoiceEndpoints.getDocuments)
	endpoint.Path = endpoint.Path + fmt.Sprintf("/%s/details", uuid)

	req, err := newAuthenticatedRequest("GET", endpoint.String(), token, nil)
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
func (e *EInvoiceAPI) CancelDocument(token, uuid, reason string) (*UpdateStatusResponse, error) {
	return e.updateDocumentStatus(token, uuid, stDocumentCancelled, reason)
}

// RejectDocument rejects a document
// api signature: PUT /api/v1.0/documents/state/{UUID}/state
func (e *EInvoiceAPI) RejectDocument(token, uuid, reason string) (*UpdateStatusResponse, error) {
	return e.updateDocumentStatus(token, uuid, stDocumentRejected, reason)
}

func (e *EInvoiceAPI) updateDocumentStatus(token, uuid, status, reason string) (*UpdateStatusResponse, error) {
	endpoint := e.myInvoisEndpoint.ResolveReference(EinvoiceEndpoints.updateDocumentStatus(uuid))

	body := UpdateStatusRequest{
		Status: status,
		Reason: reason,
	}
	b, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrMarshalFailed, err)
	}

	req, err := newAuthenticatedRequest("PUT", endpoint.String(), token, bytes.NewReader(b))
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
func (e *EInvoiceAPI) GetRecentDocuments(token string, limit int) ([]GetDocumentDetailsResponse, error) {
	endpoint := e.myInvoisEndpoint.ResolveReference(EinvoiceEndpoints.getRecentDocuments)

	q := endpoint.Query()
	q.Set("pageSize", fmt.Sprintf("%d", limit))

	return nil, nil
}

func signDocument(pkey *rsa.PrivateKey, iv Ubl21Invoice, cert x509CertWrapper) (*Ubl21Invoice, error) {
	signatureObject := computeSignatureObject(cert)
	docBytes, err := json.Marshal(iv)
	if err != nil {
		return nil, err
	}

	docDigest := computeDigest(docBytes)
	signedDocDigest, err := rsa.SignPKCS1v15(nil, pkey, crypto.SHA256, sha256Hash(docBytes))
	if err != nil {
		return nil, err
	}

	signedPropBytes, err := json.Marshal(signatureObject[0].QualifyingProperties[0])
	if err != nil {
		return nil, err
	}
	signedPropDigest := computeDigest(signedPropBytes)

	signature := SignatureDetails{
		ID: []IdentifierType{
			{
				Empty: "urn:oasis:names:specification:ubl:signature:Invoice",
			},
		},
		SignatureMethod: []TextType{
			{
				Empty: "urn:oasis:names:specification:ubl:dsig:enveloped:xades",
			},
		},
	}
	iv.Invoice[0].Signature = append(iv.Invoice[0].Signature, signature)

	ublDocumentSignature := []UBLDocumentSignature{
		{
			SignatureInformation: []SignatureInformation{
				{
					ID: []IdentifierType{
						{
							Empty: "urn:oasis:names:specification:ubl:signature:1",
						},
					},
					ReferencedSignatureID: []IdentifierType{
						{
							Empty: "urn:oasis:names:specification:ubl:signature:Invoice",
						},
					},
					Signature: []Signature{
						{
							ID:      "signature",
							Object:  signatureObject,
							KeyInfo: computeKeyInfo(cert),
							SignatureValue: []IdentifierType{
								{
									Empty: base64.StdEncoding.EncodeToString(signedDocDigest), // TODO compute signature value
								},
							},
							SignedInfo: computeSignedInfo(docDigest, signedPropDigest),
						},
					},
				},
			},
		},
	}

	ublExtension := UBLExtension{
		ExtensionURI: []IdentifierType{
			{
				Empty: "urn:oasis:names:specification:ubl:dsig:enveloped:xades",
			},
		},
		ExtensionContent: []map[string]interface{}{
			{
				"UBLDocumentSignatures": ublDocumentSignature,
			},
		},
	}

	iv.Invoice[0].UBLExtensions = []UBLExtensions{
		{
			UBLExtension: []UBLExtension{
				ublExtension,
			},
		},
	}

	return &iv, nil
}

func computeSignatureObject(cert x509CertWrapper) []Object {
	return []Object{
		{
			QualifyingProperties: []QualifyingProperty{
				{
					Target: "signature",
					SignedProperties: []SignedProperty{
						{
							ID: "id-xades-signed-props",
							SignedSignatureProperties: []SignedSignatureProperty{
								{
									SigningTime: []IdentifierType{
										{
											Empty: time.Now().UTC().Format("2006-01-02T15:04:05Z"),
										},
									},
									SigningCertificate: []SigningCertificate{
										{
											Cert: []Cert{
												{
													CertDigest: []CertDigest{
														{
															DigestMethod: []Method{
																{
																	Empty:     "",
																	Algorithm: "http://www.w3.org/2001/04/xmlenc#sha256",
																},
															},
															DigestValue: []IdentifierType{
																{
																	Empty: cert.digest,
																},
															},
														},
													},
													IssuerSerial: []IssuerSerial{
														{
															X509IssuerName: []IdentifierType{
																{
																	Empty: cert.issuer,
																},
															},
															X509SerialNumber: []IdentifierType{
																{
																	Empty: cert.serialNumber,
																},
															},
														},
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func computeKeyInfo(cert x509CertWrapper) []KeyInfo {
	return []KeyInfo{
		{
			X509Data: []X509Datum{
				{
					X509Certificate: []IdentifierType{
						{
							Empty: cert.base64,
						},
					},
					X509SubjectName: []IdentifierType{
						{
							Empty: cert.subject,
						},
					},
					X509IssuerSerial: []IssuerSerial{
						{
							X509IssuerName: []IdentifierType{
								{
									Empty: cert.issuer,
								},
							},
							X509SerialNumber: []IdentifierType{
								{
									Empty: cert.serialNumber,
								},
							},
						},
					},
				},
			},
		},
	}
}

func computeSignedInfo(docDIgest, signedPropDigest string) []SignedInfo {
	return []SignedInfo{
		{
			SignatureMethod: []Method{
				{
					Empty:     "",
					Algorithm: "http://www.w3.org/2001/04/xmldsig-more#rsa-sha256",
				},
			},
			Reference: []Reference{
				{
					Type: "http://uri.etsi.org/01903/v1.3.2#SignedProperties",
					URI:  "#id-xades-signed-props",
					DigestMethod: []Method{
						{
							Empty:     "",
							Algorithm: "http://www.w3.org/2001/04/xmlenc#sha256",
						},
					},
					DigestValue: []IdentifierType{
						{
							Empty: signedPropDigest,
						},
					},
				},
				{
					Type: "",
					URI:  "",
					DigestMethod: []Method{
						{
							Empty:     "",
							Algorithm: "http://www.w3.org/2001/04/xmlenc#sha256",
						},
					},
					DigestValue: []IdentifierType{
						{
							Empty: docDIgest,
						},
					},
				},
			},
		},
	}
}

func sha256Hash(data []byte) []byte {
	h := sha256.New()
	h.Write(data)
	return h.Sum(nil)
}

func computeDigest(data []byte) string {
	h := sha256Hash(data)
	return base64.StdEncoding.EncodeToString(h)
}

func rsaSHA256Sign(pkey *rsa.PrivateKey, data []byte) ([]byte, error) {
	return rsa.SignPKCS1v15(nil, pkey, crypto.SHA256, data)
}
