package myinvois

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

var currentUuid string

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

	return SandboxClient(os.Getenv("CLIENT_ID"), os.Getenv("CLIENT_SECRET"), cert, key, []byte(os.Getenv("PKEY_PASSWORD")))
}

func validInvoice() Ubl21Invoice {
	b, err := os.ReadFile("testdata/submission_v1.1-unsigned-valid.json")
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
func validInvoiceOnBehalf() Ubl21Invoice {
	b, err := os.ReadFile("testdata/submission_v1.1-unsigned_2.json")
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

func invalidInvoice() Ubl21Invoice {
	b, err := os.ReadFile("testdata/submission_v1.1-unsigned-invalid.json")
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

func waitForDocumentStatus(t *testing.T, client *Client, token string, uuid string, status string) (*GetDocumentDetailsResponse, error) {
	startTime := time.Now()
	for tick := range time.Tick(2 * time.Second) {
		res, err := client.GetDocumentDetails(token, uuid)
		if err != nil {
			return nil, err
		}

		if res != nil {
			t.Logf("Document %s status: %s\n", uuid, res.Status)
		} else {
			continue
		}

		if res.Status == status {
			return res, nil
		} else if tick.After(startTime.Add(30 * time.Second)) {
			return nil, fmt.Errorf("Timeout waiting for document status %s", status)
		}
	}
	return nil, nil
}

func TestValidateTaxpayerTIN(t *testing.T) {
	client := setupEInvoiceTest()
	assert := assert.New(t)

	token, err := client.LoginAsTaxpayer()
	assert.Nil(err)
	assert.NotNil(token)

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

func TestSubmitDocuments(t *testing.T) {
	client := setupEInvoiceTest()
	assert := assert.New(t)

	token, err := client.LoginAsTaxpayer()
	assert.Nil(err)
	assert.NotNil(token)

	ublInvoice := validInvoice()

	res, err := client.SubmitDocuments(token.AccessToken, []Ubl21Invoice{ublInvoice})
	assert.Nil(err)
	assert.NotNil(res)

	if res != nil {
		for _, i := range res.AcceptedDocuments {
			t.Log("Accepted:", i.UUID, i.InvoiceCodeNumber)
			currentUuid = i.UUID
		}

		for _, i := range res.RejectedDocuments {
			t.Log("Rejected:", i.InvoiceCodeNumber)
			for _, e := range i.Error.Details {
				t.Log("Error:", e.Code, e.Message)
			}
		}
	}
}

func TestSubmitValidDocument(t *testing.T) {
	client := setupEInvoiceTest()
	assert := assert.New(t)

	token, err := client.LoginAsTaxpayer()
	assert.Nil(err)
	assert.NotNil(token)

	ublInvoice := validInvoice()

	res, err := client.SubmitDocuments(token.AccessToken, []Ubl21Invoice{ublInvoice})
	assert.Nil(err)
	assert.NotNil(res)
	if res != nil {
		assert.Greater(len(res.AcceptedDocuments), 0)
		assert.Equal(len(res.RejectedDocuments), 0)
	}

	if res != nil {
		for _, i := range res.AcceptedDocuments {
			t.Log("Accepted:", i.UUID, i.InvoiceCodeNumber)
			currentUuid = i.UUID
		}

		for _, i := range res.RejectedDocuments {
			t.Log("Rejected:", i.InvoiceCodeNumber)
			for _, e := range i.Error.Details {
				t.Log("Error:", e.Code, e.Message)
			}
		}
	}
}

func TestSubmitInvalidDocuments(t *testing.T) {
	client := setupEInvoiceTest()
	assert := assert.New(t)

	token, err := client.LoginAsTaxpayer()
	assert.Nil(err)
	assert.NotNil(token)

	ublInvoice := invalidInvoice()

	res, err := client.SubmitDocuments(token.AccessToken, []Ubl21Invoice{ublInvoice})
	assert.Nil(err)
	assert.NotNil(res)
	if res != nil {
		assert.Equal(0, len(res.AcceptedDocuments))
		assert.Greater(0, len(res.RejectedDocuments))
	}

	if res != nil {
		for _, i := range res.AcceptedDocuments {
			t.Log("Accepted:", i.UUID, i.InvoiceCodeNumber)
			currentUuid = i.UUID
		}

		for _, i := range res.RejectedDocuments {
			t.Log("Rejected:", i.InvoiceCodeNumber)
			for _, e := range i.Error.Details {
				t.Log("Error:", e.Code, e.Message)
			}
		}
	}
}

func TestGetDocumentDetails(t *testing.T) {
	client := setupEInvoiceTest()
	assert := assert.New(t)

	token, err := client.LoginAsTaxpayer()
	assert.Nil(err)
	assert.NotNil(token)

	ublInvoice := validInvoice()

	res, err := client.SubmitDocuments(token.AccessToken, []Ubl21Invoice{ublInvoice})
	assert.Nil(err)
	assert.NotNil(res)

	var currentUUID string
	if res != nil {
		for _, i := range res.AcceptedDocuments {
			t.Log("Accepted:", i.UUID, i.InvoiceCodeNumber)
			currentUUID = i.UUID
		}

		for _, i := range res.RejectedDocuments {
			t.Log("Rejected:", i.InvoiceCodeNumber)
			t.Log("Error:", i.Error.Code, i.Error.Message)
			for _, e := range i.Error.Details {
				t.Log("Error:", e.Code, e.Message)
			}
		}
	}

	if currentUUID == "" {
		t.Skip("No document to get details")
	}

	startTime := time.Now()
	for tick := range time.Tick(2 * time.Second) {
		t.Log("Getting document details for", currentUUID)
		res2, err := client.GetDocumentDetails(token.AccessToken, currentUUID)
		assert.Nil(err)
		assert.NotNil(res2)
		if res2 != nil {
			b, err := json.MarshalIndent(res2, "", "  ")
			assert.Nil(err)
			t.Log(string(b))
		} else {
			continue
		}

		if res2.Status != "Submitted" {
			if tick.After(startTime.Add(30 * time.Second)) {
				t.Fail()
			}
			break
		}
	}
}
func TestGetDocumentDetailsOnBehalf(t *testing.T) {
	client := setupEInvoiceTest()
	assert := assert.New(t)

	token, err := client.LoginAsIntermediaries(os.Getenv("TIN"))
	assert.Nil(err)
	assert.NotNil(token)
	assert.NotEmpty(token.AccessToken)

	ublInvoice := validInvoiceOnBehalf()

	res, err := client.SubmitDocuments(token.AccessToken, []Ubl21Invoice{ublInvoice})
	assert.Nil(err)
	assert.NotNil(res)

	var currentUUID string
	if res != nil {
		for _, i := range res.AcceptedDocuments {
			t.Log("Accepted:", i.UUID, i.InvoiceCodeNumber)
			currentUUID = i.UUID
		}

		for _, i := range res.RejectedDocuments {
			t.Log("Rejected:", i.InvoiceCodeNumber)
			t.Log("Error:", i.Error.Code, i.Error.Message)
			for _, e := range i.Error.Details {
				t.Log("Error:", e.Code, e.Message)
			}
		}
	}

	if currentUUID == "" {
		t.Skip("No document to get details")
	}

	startTime := time.Now()
	for tick := range time.Tick(2 * time.Second) {
		t.Log("Getting document details for", currentUUID)
		res2, err := client.GetDocumentDetails(token.AccessToken, currentUUID)
		assert.Nil(err)
		assert.NotNil(res2)
		if res2 != nil {
			b, err := json.MarshalIndent(res2, "", "  ")
			assert.Nil(err)
			t.Log(string(b))
		} else {
			continue
		}

		if res2.Status != "Submitted" {
			if tick.After(startTime.Add(30 * time.Second)) {
				t.Fail()
			}
			break
		}
	}
}

func TestCancelDocument(t *testing.T) {
	client := setupEInvoiceTest()
	assert := assert.New(t)

	token, err := client.LoginAsTaxpayer()
	assert.Nil(err)
	assert.NotNil(token)

	ublInvoice := validInvoice()

	res, err := client.SubmitDocuments(token.AccessToken, []Ubl21Invoice{ublInvoice})
	assert.Nil(err)
	assert.NotNil(res)

	var currentUUID string
	if res != nil {
		currentUUID = res.AcceptedDocuments[0].UUID
	}

	if currentUUID == "" {
		t.Log("No document to cancel")
		t.Fail()
	}

	details, err := waitForDocumentStatus(t, client, token.AccessToken, currentUUID, stDocumentValid)
	assert.Nil(err)
	assert.NotNil(details)

	t.Log("Cancelling document ", currentUUID)
	res2, err := client.CancelDocument(token.AccessToken, currentUUID, "Cancelled by tests")
	assert.Nil(err)
	assert.NotNil(res2)

	details, err = client.GetDocumentDetails(token.AccessToken, currentUUID)
	assert.Nil(err)
	assert.NotNil(details)
	assert.Equal(stDocumentCancelled, details.Status)
}

func TestRejectDocument(t *testing.T) {
	client := setupEInvoiceTest()
	assert := assert.New(t)

	token, err := client.LoginAsTaxpayer()
	assert.Nil(err)
	assert.NotNil(token)

	ublInvoice := validInvoice()

	res, err := client.SubmitDocuments(token.AccessToken, []Ubl21Invoice{ublInvoice})
	assert.Nil(err)
	assert.NotNil(res)

	var currentUUID string
	if res != nil {
		currentUUID = res.AcceptedDocuments[0].UUID
	}

	if currentUUID == "" {
		t.Log("No document to cancel")
		t.Fail()
	}

	details, err := waitForDocumentStatus(t, client, token.AccessToken, currentUUID, stDocumentValid)
	assert.Nil(err)
	assert.NotNil(details)

	t.Log("Rejecting document ", currentUUID)
	res2, err := client.RejectDocument(token.AccessToken, currentUUID, "Cancelled by tests")
	assert.Nil(err)
	assert.NotNil(res2)

	details, err = client.GetDocumentDetails(token.AccessToken, currentUUID)
	assert.Nil(err)
	assert.NotNil(details)
	assert.Equal(stDocumentRejected, details.Status)
}

func TestRsaSHA256Sign(t *testing.T) {
	c := setupEInvoiceTest()
	assert := assert.New(t)

	digest := `RqhvXUQSRlfxagJP2KXLYrXXY7D/indPZWBFcup0vyg=`
	digestDecoded, err := base64.StdEncoding.DecodeString(digest)
	assert.Nil(err)

	signature, err := rsaSHA256Sign(c.privKey, digestDecoded)
	assert.Nil(err)
	t.Log(base64.StdEncoding.EncodeToString(signature))
}
