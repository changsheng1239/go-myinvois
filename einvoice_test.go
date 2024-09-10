package myinvois

import (
	"crypto"
	"crypto/rsa"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	fileTestInvoice  = "testdata/test.json"
	fileValidInvoice = "testdata/invoice-valid.json"
	fileValidConsoIV = "testdata/conso-invoice-valid.json"
)

func setupEInvoiceTest() *Client {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	cert, err := os.ReadFile(os.Getenv("CERT_PATH"))
	if err != nil {
		panic(err)
	}
	key, err := os.ReadFile(os.Getenv("PKEY_PATH"))
	if err != nil {
		panic(err)
	}

	return NewClient(ClientOption{
		// Environment: Production,
		Environment:  Sandbox,
		Timeout:      DefaultTimeout,
		ClientID:     os.Getenv("CLIENT_ID"),
		ClientSecret: os.Getenv("CLIENT_SECRET"),
		Cert:         cert,
		PrivKey:      key,
		PrivKeyPass:  []byte(os.Getenv("PKEY_PASSWORD")),
	})
}

func loadInvoice(filename string) Ubl21Invoice {
	b, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	var ublInvoice Ubl21Invoice
	err = json.Unmarshal(b, &ublInvoice)
	if err != nil {
		panic(err)
	}

	ublInvoice.Invoice[0].ID[0].Empty = uuid.NewString()
	ublInvoice.Invoice[0].IssueDate[0].Empty = time.Now().Format("2006-01-02")
	ublInvoice.Invoice[0].IssueTime[0].Empty = time.Now().UTC().Format("15:04:05Z")

	return ublInvoice
}

func waitForDocumentStatus(t *testing.T, client *Client, uuid string, status string) (*GetDocumentDetailsResponse, error) {
	token := login(client)
	assert := assert.New(t)
	require := require.New(t)

	startTime := time.Now()
	for tick := range time.Tick(2 * time.Second) {
		t.Log("Getting document details for", uuid)
		res, err := client.GetDocumentDetails(token.AccessToken, uuid)
		// if document not found, continue polling in case LHDN server is having delay
		if err != nil {
			if strings.Contains(err.Error(), "404") {
				continue
			}
			return nil, err
		}

		t.Logf("Document %s's current status: %s, want: %s\n", uuid, res.Status, status)

		isDocumentSubmitted := (res.Status != stDocumentPending && res.Status != stDocumentSubmitted)
		// if status is what we want or not pending/submitted, return
		if isDocumentSubmitted {
			b, err := json.MarshalIndent(res, "", "  ")
			assert.Nil(err)

			require.Equal(status, res.Status,
				"Document status mismatch, want: %s, got: %s, body: \n%s",
				status, res.Status, string(b),
			)
			return res, nil
		} else if tick.After(startTime.Add(30 * time.Second)) {
			break
		}
	}
	return nil, fmt.Errorf("Timeout waiting for document status %s", status)
}

func printSubmissionResponse(t *testing.T, res *DocumentSubmissionResponse) {
	if res != nil {
		for _, i := range res.AcceptedDocuments {
			t.Log("Accepted:", i.UUID, i.InvoiceCodeNumber)
		}

		for _, i := range res.RejectedDocuments {
			t.Log("Rejected:", i.InvoiceCodeNumber)
			for _, e := range i.Error.Details {
				t.Log("Error:", e.Code, e.Message)
			}
		}
	}
}

func submitDocuments(client *Client, invoices []Ubl21Invoice) (*DocumentSubmissionResponse, error) {
	token := login(client)
	res, err := client.SubmitDocuments(token.AccessToken, invoices)

	return res, err
}

func submitAndAssert(t *testing.T, client *Client, doc Ubl21Invoice) AcceptedDocument {
	require := require.New(t)

	res, err := submitDocuments(client, []Ubl21Invoice{doc})
	require.Nil(err)
	require.NotNil(res)
	printSubmissionResponse(t, res)
	require.Equal(1, len(res.AcceptedDocuments))
	require.Equal(0, len(res.RejectedDocuments))

	return res.AcceptedDocuments[0]
}

func TestValidateTaxpayerTIN(t *testing.T) {
	client := setupEInvoiceTest()
	token := login(client)
	assert := assert.New(t)

	var tests = []struct {
		tin      string
		idType   string
		idValue  string
		expected bool
	}{
		{"D26020043030", "BRN", "KT0031276-X", true},
		{"C20127289100", "BRN", "200601007071", true},
		{"D26020043030", "BRN", "198403011803", false},
		{"C12345678900", "BRN", "200810101012", false},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("TIN: %s, %s:%s", test.tin, test.idType, test.idValue), func(t *testing.T) {
			t.Parallel()
			valid, _ := client.ValidateTaxpayerTIN(token.AccessToken, test.tin, test.idType, test.idValue)
			assert.Equal(test.expected, valid)
		})
	}

	var errorTests = []struct {
		tin         string
		idType      string
		idValue     string
		expectedErr error
	}{
		{"D26020043030", "BRN", "KT0031276-X", nil},
		{"C20127289100", "BRN", "200601007071", nil},
		{"C20127289100", "BRN", "123", ErrTinMismatch},
		{"D26020043030", "BRN", "198403011803", ErrTinMismatch},
		{"C12345678900", "BRN", "200810101012", ErrTinMismatch},
	}

	for _, test := range errorTests {
		t.Run(fmt.Sprintf("TIN: %s, %s:%s", test.tin, test.idType, test.idValue), func(t *testing.T) {
			t.Parallel()
			_, err := client.ValidateTaxpayerTIN(token.AccessToken, test.tin, test.idType, test.idValue)
			assert.ErrorIs(err, test.expectedErr, "expected error mismatch")
		})
	}
}

func TestSubmitValidDocument(t *testing.T) {
	client := setupEInvoiceTest()

	t.Run("Submit valid invoice", func(t *testing.T) {
		acceptedDocument := submitAndAssert(t, client, loadInvoice(fileValidInvoice))
		waitForDocumentStatus(t, client, acceptedDocument.UUID, stDocumentValid)
	})

	t.Run("Submit valid consolidated invoice", func(t *testing.T) {
		acceptedDocument := submitAndAssert(t, client, loadInvoice(fileValidConsoIV))
		waitForDocumentStatus(t, client, acceptedDocument.UUID, stDocumentValid)
	})
}

func TestSubmitInvalidDocument(t *testing.T) {
	client := setupEInvoiceTest()
	assert := assert.New(t)

	submitAndAssertInvalid := func(t *testing.T, ublInvoice Ubl21Invoice) *DocumentSubmissionResponse {
		res, err := submitDocuments(client, []Ubl21Invoice{ublInvoice})
		printSubmissionResponse(t, res)
		assert.Nil(err)
		assert.NotNil(res)
		assert.Equal(0, len(res.AcceptedDocuments), "expected accepted documents to be 0")
		assert.Equal(1, len(res.RejectedDocuments), "expected rejected documents to be 1")
		return res
	}

	t.Run("Submit invalid invoice with missing Classification code", func(t *testing.T) {
		ublInvoice := loadInvoice(fileValidInvoice)
		ublInvoice.Invoice[0].InvoiceLine[0].Item[0].CommodityClassification = nil
		res := submitAndAssertInvalid(t, ublInvoice)
		assert.Equal("CF364", res.RejectedDocuments[0].Error.Details[0].Code)
		assert.Equal("Line classification is required", res.RejectedDocuments[0].Error.Details[0].Message)
	})
	t.Run("Submit invalid invoice with missing Supplier Name", func(t *testing.T) {
		ublInvoice := loadInvoice(fileValidInvoice)
		ublInvoice.Invoice[0].AccountingSupplierParty[0].Party[0].PartyLegalEntity[0].RegistrationName = nil
		res := submitAndAssertInvalid(t, ublInvoice)
		assert.Equal("CF334", res.RejectedDocuments[0].Error.Details[0].Code)
		assert.Equal("Name is not valid - SUPPLIER", res.RejectedDocuments[0].Error.Details[0].Message)
	})
	t.Run("Submit invalid invoice with missing Supplier Country", func(t *testing.T) {
		ublInvoice := loadInvoice(fileValidInvoice)
		ublInvoice.Invoice[0].AccountingSupplierParty[0].Party[0].PostalAddress[0].Country = nil
		res := submitAndAssertInvalid(t, ublInvoice)
		assert.Equal("CF336", res.RejectedDocuments[0].Error.Details[0].Code)
		assert.Equal("Country is required - SUPPLIER", res.RejectedDocuments[0].Error.Details[0].Message)
	})
	t.Run("Submit invalid invoice with missing Supplier City", func(t *testing.T) {
		ublInvoice := loadInvoice(fileValidInvoice)
		ublInvoice.Invoice[0].AccountingSupplierParty[0].Party[0].PostalAddress[0].CityName = nil
		res := submitAndAssertInvalid(t, ublInvoice)
		assert.Equal("CF339", res.RejectedDocuments[0].Error.Details[0].Code)
		assert.Equal("City is required - SUPPLIER", res.RejectedDocuments[0].Error.Details[0].Message)
	})
	t.Run("Submit invalid invoice with missing Supplier State", func(t *testing.T) {
		ublInvoice := loadInvoice(fileValidInvoice)
		ublInvoice.Invoice[0].AccountingSupplierParty[0].Party[0].PostalAddress[0].CountrySubentityCode = nil
		res := submitAndAssertInvalid(t, ublInvoice)
		assert.Equal("CF342", res.RejectedDocuments[0].Error.Details[0].Code)
		assert.Equal("State is required - SUPPLIER", res.RejectedDocuments[0].Error.Details[0].Message)
	})
	t.Run("Submit invalid invoice with missing Supplier Address line 1", func(t *testing.T) {
		ublInvoice := loadInvoice(fileValidInvoice)
		ublInvoice.Invoice[0].AccountingSupplierParty[0].Party[0].PostalAddress[0].AddressLine[0].Line[0].Empty = ""
		res := submitAndAssertInvalid(t, ublInvoice)
		assert.Equal("CF345", res.RejectedDocuments[0].Error.Details[0].Code)
		assert.Equal("Address line 1 is required - SUPPLIER", res.RejectedDocuments[0].Error.Details[0].Message)
	})
	t.Run("Submit invalid invoice with missing Supplier Contact number", func(t *testing.T) {
		ublInvoice := loadInvoice(fileValidInvoice)
		ublInvoice.Invoice[0].AccountingSupplierParty[0].Party[0].Contact[0].Telephone = nil
		res := submitAndAssertInvalid(t, ublInvoice)
		assert.Equal("CF348", res.RejectedDocuments[0].Error.Details[0].Code)
		assert.Equal("Contact number is required - SUPPLIER", res.RejectedDocuments[0].Error.Details[0].Message)
	})
	t.Run("Submit invalid invoice with missing DocumentCurrencyCode", func(t *testing.T) {
		ublInvoice := loadInvoice(fileValidInvoice)
		ublInvoice.Invoice[0].DocumentCurrencyCode = nil
		res := submitAndAssertInvalid(t, ublInvoice)
		assert.Equal("ArrayItemNotValid: #/Invoice[0]\n{\n  ArrayExpected: #/Invoice[0].DocumentCurrencyCode\n}\n", res.RejectedDocuments[0].Error.Details[0].Message)
	})

	t.Run("Submit empty invoice", func(t *testing.T) {
		var ublInvoice Ubl21Invoice
		res, err := submitDocuments(client, []Ubl21Invoice{ublInvoice})
		if assert.Error(err) {
			assert.Equal(ErrInvalidInput, err, "expected error mismatch")
		}
		assert.Nil(res)
	})
}

func TestSubmitCreditNote(t *testing.T) {
	client := setupEInvoiceTest()
	assert := assert.New(t)
	require := require.New(t)

	acceptedDocument := submitAndAssert(t, client, loadInvoice(fileValidInvoice))
	currentUUID := acceptedDocument.UUID
	internalID := acceptedDocument.InvoiceCodeNumber

	details, err := waitForDocumentStatus(t, client, currentUUID, stDocumentValid)
	assert.Nil(err)
	assert.NotNil(details)

	t.Log("Submitting credit note for ", currentUUID)
	ublCreditNote := loadInvoice(fileValidInvoice)
	ublCreditNote.Invoice[0].InvoiceTypeCode[0].Empty = "02"
	ublCreditNote.Invoice[0].BillingReference = []BillingReferenceDetails{
		{
			InvoiceDocumentReference: []DocumentReferenceDetails{{
				ID:   []IdentifierType{{Empty: internalID}},
				UUID: []IdentifierType{{Empty: currentUUID}},
			}},
		},
	}

	acceptedCN := submitAndAssert(t, client, ublCreditNote)
	currentUUID = acceptedCN.UUID
	require.NotEmpty(currentUUID)

	details, err = waitForDocumentStatus(t, client, currentUUID, stDocumentValid)
	assert.Nil(err)
	assert.NotNil(details)
}

func TestRequiredFieldsInvoice(t *testing.T) {
	client := setupEInvoiceTest()
	token := login(client)
	assert := assert.New(t)

	var ublInvoice Ubl21Invoice
	var invoice InvoiceDetails

	// required fields for structured invoice
	{
		invoice.ID = append(invoice.ID, IdentifierType{Empty: uuid.NewString()})
		invoice.IssueDate = append(invoice.IssueDate, DateType{Empty: time.Now().UTC().Format("2006-01-02")})
		invoice.IssueTime = append(invoice.IssueTime, TimeType{Empty: time.Now().UTC().Format("15:04:05Z")})
		invoice.InvoiceTypeCode = append(invoice.InvoiceTypeCode, CodeType{Empty: "01", ListVersionID: "1.0"})
		invoice.DocumentCurrencyCode = append(invoice.DocumentCurrencyCode, CodeType{Empty: "MYR"})
		invoice.BillingReference = append(invoice.BillingReference, BillingReferenceDetails{
			InvoiceDocumentReference: []DocumentReferenceDetails{{
				ID:   []IdentifierType{{Empty: ""}},
				UUID: []IdentifierType{{Empty: ""}},
			}},
		})
		invoice.AccountingSupplierParty = append(invoice.AccountingSupplierParty, SupplierPartyDetails{
			Party: []PartyDetails{{
				IndustryClassificationCode: []CodeType{{Empty: "46510", Name: "Wholesale of computers, computer peripheral equipment and software"}},
				PartyIdentification: []PartyIdentificationDetails{
					{ID: []IdentifierType{{Empty: "C24050894070", SchemeID: "TIN"}}},
					{ID: []IdentifierType{{Empty: "200801024110", SchemeID: "BRN"}}},
					{ID: []IdentifierType{{Empty: "NA", SchemeID: "TTX"}}},
				},
			}},
		})
		invoice.AccountingCustomerParty = append(invoice.AccountingCustomerParty, CustomerPartyDetails{
			Party: []PartyDetails{{
				PartyIdentification: []PartyIdentificationDetails{
					{ID: []IdentifierType{{Empty: "EI00000000010", SchemeID: "TIN"}}},
					{ID: []IdentifierType{{Empty: "NA", SchemeID: "BRN"}}},
				},
			}},
		})
		invoice.LegalMonetaryTotal = append(invoice.LegalMonetaryTotal, MonetaryTotalDetails{
			PayableAmount: []AmountType{{Empty: 0, CurrencyID: "MYR"}},
		})
		invoice.InvoiceLine = append(invoice.InvoiceLine, InvoiceLineDetails{
			ID:                  []IdentifierType{{Empty: "1"}},
			LineExtensionAmount: []AmountType{{Empty: 0, CurrencyID: "MYR"}},
			Item: []ItemDetails{{
				Description: []TextType{{Empty: "Test item"}},
			}},
		})
	}

	ublInvoice.D = "urn:oasis:names:specification:ubl:schema:xsd:Invoice-2"
	ublInvoice.A = "urn:oasis:names:specification:ubl:schema:xsd:CommonAggregateComponents-2"
	ublInvoice.B = "urn:oasis:names:specification:ubl:schema:xsd:CommonBasicComponents-2"
	ublInvoice.Invoice = append(ublInvoice.Invoice, invoice)

	res, err := client.SubmitDocuments(token.AccessToken, []Ubl21Invoice{ublInvoice})
	assert.Nil(err)
	assert.NotNil(res)
	if res != nil {
		assert.Equal(0, len(res.AcceptedDocuments))
		assert.Greater(len(res.RejectedDocuments), 0)
	}

	expectedErrors := []string{
		"CF364: Line classification is required",
		"CF334: Name is not valid - SUPPLIER",
		"CF336: Country is required - SUPPLIER",
		"CF339: City is required - SUPPLIER",
		"CF342: State is required - SUPPLIER",
		"CF345: Address line 1 is required - SUPPLIER",
		"CF348: Contact number is required - SUPPLIER",
		"CF333: Name is not valid - BUYER",
		"CF337: Country is required - BUYER",
		"CF340: City is required - BUYER",
		"CF343: State is required - BUYER",
		"CF346: Address line 1 is required - BUYER",
		"CF349: Contact number is required - BUYER",
	}

	if res != nil {
		for _, i := range res.AcceptedDocuments {
			t.Log("Accepted:", i.UUID, i.InvoiceCodeNumber)
		}

		for _, i := range res.RejectedDocuments {
			for _, e := range i.Error.Details {
				assert.Contains(expectedErrors, fmt.Sprintf("%s: %s", e.Code, e.Message), "expected error not found")
			}
		}
	}
}

func TestGetDocumentDetails(t *testing.T) {
	client := setupEInvoiceTest()
	assert := assert.New(t)

	acceptedDocument := submitAndAssert(t, client, loadInvoice(fileValidInvoice))
	currentUUID := acceptedDocument.UUID

	if currentUUID == "" {
		t.Skip("No document to get details")
	}

	details, err := waitForDocumentStatus(t, client, currentUUID, stDocumentValid)
	assert.Nil(err)
	assert.NotNil(details)

	b, err := json.MarshalIndent(details, "", "  ")
	assert.Nil(err)
	t.Log(string(b))
}

func TestCancelDocument(t *testing.T) {
	client := setupEInvoiceTest()
	token := login(client)
	assert := assert.New(t)
	require := require.New(t)

	acceptedDocument := submitAndAssert(t, client, loadInvoice(fileValidInvoice))
	currentUUID := acceptedDocument.UUID
	require.NotEmpty(currentUUID)

	details, err := waitForDocumentStatus(t, client, currentUUID, stDocumentValid)
	assert.Nil(err)
	assert.NotNil(details)

	t.Log("Cancelling document ", currentUUID)
	res2, err := client.CancelDocument(token.AccessToken, currentUUID, "Cancelled by tests")
	assert.Nil(err)
	assert.NotNil(res2)

	details, err = waitForDocumentStatus(t, client, currentUUID, stDocumentCancelled)
	assert.Nil(err)
	assert.NotNil(details)
}

// not working currently, need to reject using buyer's credentials
func TestRejectDocument(t *testing.T) {
	client := setupEInvoiceTest()
	token := login(client)
	assert := assert.New(t)
	require := require.New(t)

	acceptedDocument := submitAndAssert(t, client, loadInvoice(fileValidInvoice))
	currentUUID := acceptedDocument.UUID
	require.NotEmpty(currentUUID)

	t.Log("Waiting for document to be valid, uuid: ", currentUUID)
	details, err := waitForDocumentStatus(t, client, currentUUID, stDocumentValid)
	assert.Nil(err)
	assert.NotNil(details)

	t.Log("Rejecting document ", currentUUID)
	res2, err := client.RejectDocument(token.AccessToken, currentUUID, "Cancelled by tests")
	assert.Nil(err)
	assert.NotNil(res2)

	// todo: need to check if document is rejected using the reason
	details, err = waitForDocumentStatus(t, client, currentUUID, stDocumentValid)
	assert.Nil(err)
	assert.NotNil(details)
}

func TestRsaSHA256Sign(t *testing.T) {
	c := setupEInvoiceTest()
	assert := assert.New(t)

	digest := `RqhvXUQSRlfxagJP2KXLYrXXY7D/indPZWBFcup0vyg=`
	digestDecoded, err := base64.StdEncoding.DecodeString(digest)
	assert.Nil(err)

	signature, err := rsa.SignPKCS1v15(nil, c.privKey, crypto.SHA256, digestDecoded)
	assert.Nil(err)
	t.Log(base64.StdEncoding.EncodeToString(signature))
}
