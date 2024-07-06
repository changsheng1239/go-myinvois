package myinvois

// type EinvoiceAPI interface {
// 	ValidateTaxpayerTIN()
// 	SubmitDocuments()
// 	CancelDocument()
// 	RejectDocument()
// 	GetRecentDocuments()
// 	GetSubmission()
// 	GetDocument()
// 	GetDocumentDetails()
// 	SeachDocuments()
// }

type EInvoiceAPI struct {
	c *Client
}

func NewEInvoiceClient(c *Client) EInvoiceAPI {
	return EInvoiceAPI{c: c}
}
