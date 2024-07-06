package myinvois

import "encoding/json"

func UnmarshalUbl21Invoice(data []byte) (Ubl21Invoice, error) {
	var r Ubl21Invoice
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Ubl21Invoice) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

// Library:           OASIS Universal Business Language (UBL) 2.1
// http://docs.oasis-open.org/ubl/UBL-2.1-JSON/v2.0//
// Release Date:      10 April 2020
// Module:            json-schema/maindoc/UBL-Invoice-2.1.json
// Generated on:      2020-04-13 17:59z
// Copyright (c) OASIS Open 2020. All Rights Reserved.
//
// OASIS takes no position regarding the validity or scope of any
// intellectual property or other rights that might be claimed to pertain
// to the implementation or use of the technology described in this
// document or the extent to which any license under such rights
// might or might not be available; neither does it represent that it has
// made any effort to identify any such rights. Information on OASIS's
// procedures with respect to rights in OASIS specifications can be
// found at the OASIS website. Copies of claims of rights made
// available for publication and any assurances of licenses to be made
// available, or the result of an attempt made to obtain a general
// license or permission for the use of such proprietary rights by
// implementors or users of this specification, can be obtained from
// the OASIS Executive Director.
//
// OASIS invites any interested party to bring to its attention any
// copyrights, patents or patent applications, or other proprietary
// rights which may cover technology that may be required to
// implement this specification. Please address the information to the
// OASIS Executive Director.
//
// This document and translations of it may be copied and furnished to
// others, and derivative works that comment on or otherwise explain
// it or assist in its implementation may be prepared, copied,
// published and distributed, in whole or in part, without restriction of
// any kind, provided that the above copyright notice and this
// paragraph are included on all such copies and derivative works.
// However, this document itself may not be modified in any way,
// such as by removing the copyright notice or references to OASIS,
// except as needed for the purpose of developing OASIS
// specifications, in which case the procedures for copyrights defined
// in the OASIS Intellectual Property Rights document must be
// followed, or as required to translate it into languages other than
// English.
//
// The limited permissions granted above are perpetual and will not be
// revoked by OASIS or its successors or assigns.
//
// This document and the information contained herein is provided on
// an "AS IS" basis and OASIS DISCLAIMS ALL WARRANTIES,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO ANY
// WARRANTY THAT THE USE OF THE INFORMATION HEREIN
// WILL NOT INFRINGE ANY RIGHTS OR ANY IMPLIED
// WARRANTIES OF MERCHANTABILITY OR FITNESS FOR A
// PARTICULAR PURPOSE.
type Ubl21Invoice struct {
	// Library ABIE XML namespace string (for ASBIEs)
	A *string `json:"_A,omitempty"`
	// BBIE XML namespace string
	B *string `json:"_B,omitempty"`
	// Document ABIE XML namespace string
	D *string `json:"_D,omitempty"`
	// Extension scaffolding XML namespace string
	E *string `json:"_E,omitempty"`
	// A document used to request payment.
	Invoice []InvoiceDetails `json:"Invoice"`
}

// A document used to request payment.
type InvoiceDetails struct {
	// The buyer's accounting code, applied to the Invoice as a whole, expressed as text.
	AccountingCost []TextType `json:"AccountingCost,omitempty"`
	// The buyer's accounting code, applied to the Invoice as a whole.
	AccountingCostCode []CodeType `json:"AccountingCostCode,omitempty"`
	// The accounting customer party.
	AccountingCustomerParty []CustomerPartyDetails `json:"AccountingCustomerParty"`
	// The accounting supplier party.
	AccountingSupplierParty []SupplierPartyDetails `json:"AccountingSupplierParty"`
	// A reference to an additional document associated with this document.
	AdditionalDocumentReference []DocumentReferenceDetails `json:"AdditionalDocumentReference,omitempty"`
	// A discount or charge that applies to a price component.
	AllowanceCharge []AllowanceChargeDetails `json:"AllowanceCharge,omitempty"`
	// A reference to a billing document associated with this document.
	BillingReference []BillingReferenceDetails `json:"BillingReference,omitempty"`
	// The buyer.
	BuyerCustomerParty []CustomerPartyDetails `json:"BuyerCustomerParty,omitempty"`
	// A reference provided by the buyer used for internal routing of the document.
	BuyerReference []TextType `json:"BuyerReference,omitempty"`
	// A reference to a contract associated with this document.
	ContractDocumentReference []DocumentReferenceDetails `json:"ContractDocumentReference,omitempty"`
	// Indicates whether this document is a copy (true) or not (false).
	CopyIndicator []IndicatorType `json:"CopyIndicator,omitempty"`
	// Identifies a user-defined customization of UBL for a specific use.
	CustomizationID []IdentifierType `json:"CustomizationID,omitempty"`
	// A delivery associated with this document.
	Delivery []DeliveryDetails `json:"Delivery,omitempty"`
	// A set of delivery terms associated with this document.
	DeliveryTerms []DeliveryTermsDetails `json:"DeliveryTerms,omitempty"`
	// A reference to a Despatch Advice associated with this document.
	DespatchDocumentReference []DocumentReferenceDetails `json:"DespatchDocumentReference,omitempty"`
	// A code signifying the default currency for this document.
	DocumentCurrencyCode []CodeType `json:"DocumentCurrencyCode,omitempty"`
	// The date on which Invoice is due.
	DueDate []DateType `json:"DueDate,omitempty"`
	// An identifier for this document, assigned by the sender.
	ID []IdentifierType `json:"ID"`
	// A line describing an invoice item.
	InvoiceLine []InvoiceLineDetails `json:"InvoiceLine"`
	// A period to which the Invoice applies.
	InvoicePeriod []PeriodDetails `json:"InvoicePeriod,omitempty"`
	// A code signifying the type of the Invoice.
	InvoiceTypeCode []CodeType `json:"InvoiceTypeCode,omitempty"`
	// The date, assigned by the sender, on which this document was issued.
	IssueDate []DateType `json:"IssueDate"`
	// The time, assigned by the sender, at which this document was issued.
	IssueTime []TimeType `json:"IssueTime,omitempty"`
	// The total amount payable on the Invoice, including Allowances, Charges, and Taxes.
	LegalMonetaryTotal []MonetaryTotalDetails `json:"LegalMonetaryTotal"`
	// The number of lines in the document.
	LineCountNumeric []NumericType `json:"LineCountNumeric,omitempty"`
	// Free-form text pertinent to this document, conveying information that is not contained
	// explicitly in other structures.
	Note []TextType `json:"Note,omitempty"`
	// A reference to the Order with which this Invoice is associated.
	OrderReference []OrderReferenceDetails `json:"OrderReference,omitempty"`
	// A reference to an originator document associated with this document.
	OriginatorDocumentReference []DocumentReferenceDetails `json:"OriginatorDocumentReference,omitempty"`
	// The payee.
	PayeeParty []PartyDetails `json:"PayeeParty,omitempty"`
	// A code signifying the alternative currency used for payment in the Invoice.
	PaymentAlternativeCurrencyCode []CodeType `json:"PaymentAlternativeCurrencyCode,omitempty"`
	// The exchange rate between the document currency and the payment alternative currency.
	PaymentAlternativeExchangeRate []ExchangeRateDetails `json:"PaymentAlternativeExchangeRate,omitempty"`
	// A code signifying the currency used for payment in the Invoice.
	PaymentCurrencyCode []CodeType `json:"PaymentCurrencyCode,omitempty"`
	// The exchange rate between the document currency and the payment currency.
	PaymentExchangeRate []ExchangeRateDetails `json:"PaymentExchangeRate,omitempty"`
	// Expected means of payment.
	PaymentMeans []PaymentMeansDetails `json:"PaymentMeans,omitempty"`
	// A set of payment terms associated with this document.
	PaymentTerms []PaymentTermsDetails `json:"PaymentTerms,omitempty"`
	// A prepaid payment.
	PrepaidPayment []PaymentDetails `json:"PrepaidPayment,omitempty"`
	// A code signifying the currency used for prices in the Invoice.
	PricingCurrencyCode []CodeType `json:"PricingCurrencyCode,omitempty"`
	// The exchange rate between the document currency and the pricing currency.
	PricingExchangeRate []ExchangeRateDetails `json:"PricingExchangeRate,omitempty"`
	// Identifies an instance of executing a profile, to associate all transactions in a
	// collaboration.
	ProfileExecutionID []IdentifierType `json:"ProfileExecutionID,omitempty"`
	// Identifies a user-defined profile of the customization of UBL being used.
	ProfileID []IdentifierType `json:"ProfileID,omitempty"`
	// Information about a project.
	ProjectReference []ProjectReferenceDetails `json:"ProjectReference,omitempty"`
	// A reference to a Receipt Advice associated with this document.
	ReceiptDocumentReference []DocumentReferenceDetails `json:"ReceiptDocumentReference,omitempty"`
	// The seller.
	SellerSupplierParty []SupplierPartyDetails `json:"SellerSupplierParty,omitempty"`
	// A signature applied to this document.
	Signature []SignatureDetails `json:"Signature,omitempty"`
	// A reference to a Statement associated with this document.
	StatementDocumentReference []DocumentReferenceDetails `json:"StatementDocumentReference,omitempty"`
	// A code signifying the currency used for tax amounts in the Invoice.
	TaxCurrencyCode []CodeType `json:"TaxCurrencyCode,omitempty"`
	// The exchange rate between the document currency and the tax currency.
	TaxExchangeRate []ExchangeRateDetails `json:"TaxExchangeRate,omitempty"`
	// The date of the Invoice, used to indicate the point at which tax becomes applicable.
	TaxPointDate []DateType `json:"TaxPointDate,omitempty"`
	// The tax representative.
	TaxRepresentativeParty []PartyDetails `json:"TaxRepresentativeParty,omitempty"`
	// The total amount of a specific type of tax.
	TaxTotal []TaxTotalDetails `json:"TaxTotal,omitempty"`
	// An optional set of extensions to the committee model
	UBLExtensions []UBLExtensions `json:"UBLExtensions,omitempty"`
	// Identifies the earliest version of the UBL 2 schema for this document type that defines
	// all of the elements that might be encountered in the current instance.
	UBLVersionID []IdentifierType `json:"UBLVersionID,omitempty"`
	// A universally unique identifier for an instance of this document.
	UUID []IdentifierType `json:"UUID,omitempty"`
	// The total withholding tax.
	WithholdingTaxTotal []TaxTotalDetails `json:"WithholdingTaxTotal,omitempty"`
}
