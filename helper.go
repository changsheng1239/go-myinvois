package myinvois

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"os"
	"time"
)

func signDocument(pkey *rsa.PrivateKey, iv Ubl21Invoice, cert x509CertWrapper) (*Ubl21Invoice, error) {
	signatureObject := computeSignatureObject(cert)
	docBytes, err := json.Marshal(iv)
	if err != nil {
		return nil, err
	}

	_ = os.WriteFile("response/raw.json", docBytes, 0644)

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
