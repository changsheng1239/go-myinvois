package myinvois

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func setupPlatformTest() *Client {
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

func TestDecodeToken(t *testing.T) {
	assert := assert.New(t)

	tokenString := `eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJodHRwczovL3ByZXByb2QtaWRlbnRpdHkubXlpbnZvaXMuaGFzaWwuZ292Lm15IiwibmJmIjoxNzIwNDkxMTA2LCJpYXQiOjE3MjA0OTExMDYsImV4cCI6MTcyMDQ5NDcwNiwiYXVkIjpbIkludm9pY2luZ0FQSSIsImh0dHBzOi8vcHJlcHJvZC1pZGVudGl0eS5teWludm9pcy5oYXNpbC5nb3YubXkvcmVzb3VyY2VzIl0sInNjb3BlIjpbIkludm9pY2luZ0FQSSJdLCJjbGllbnRfaWQiOiJkNjc5MDJkMi01OGE2LTRiODItOWU2OC0wNWIxZTQ2MzVmMDMiLCJJc1RheFJlcHJlcyI6IjEiLCJJc0ludGVybWVkaWFyeSI6IjAiLCJJbnRlcm1lZElkIjoiMCIsIkludGVybWVkVElOIjoiIiwiSW50ZXJtZWRFbmZvcmNlZCI6IjIiLCJuYW1lIjoiQzEyMzQ1Njc4OTAwOmQ2NzkwMmQyLTU4YTYtNGI4Mi05ZTY4LTA1YjFlNDYzNWYwMyIsIlNTSWQiOiJhZWNhN2Y1OC1iZmQwLTQ2ZTMtODZhYi01ZDRhMDYyYjdlNzEiLCJwcmVmZXJyZWRfdXNlcm5hbWUiOiJUZXN0IiwiVGF4SWQiOiIxMjM0IiwiVGF4VGluIjoiQzEyMzQ1Njc4OTAwIiwiUHJvZklkIjoiMTIzNCIsIklzVGF4QWRtaW4iOiIwIiwiSXNTeXN0ZW0iOiIxIiwiTmF0SWQiOiIifQ.YSKL4GyT_OMOmncF2PsB8QiMeDuLhY857w-WH0m-3SdjbZ7NfFut6MeXvlznECMGKgymib8yb10v4ahd1KH73dG-aCRv2azxaMRnyeHgn1TbF0W8x0iKS1yegz6uRFOuuvqpnqLvCOPuG-x3iQEmlsswkPI2v8m11Wrb0ppb5QhXXWdoQSs28YMpKOfyQvpGCtKkEQAOClg9rGQIjuB3Zfprbl-u6ljl3izGuttv9M0d_7w43KLdzNEodOtajmu6058FZhUWlX6YPGAkCootTFmB5J1siDPBIc46NGUzAnG6SPnc_4fxO5MlX5TrOsWpLgU-R3eCrV9aVo2YTgUaHQ`

	token, err := DecodeToken(tokenString)
	assert.Nil(err)
	assert.Equal("https://preprod-identity.myinvois.hasil.gov.my", token.Iss)
	assert.Equal(1720491106, token.Nbf)
	assert.Equal(1720491106, token.Iat)
	assert.Equal(1720494706, token.Exp)
	assert.Equal([]string{"InvoicingAPI", "https://preprod-identity.myinvois.hasil.gov.my/resources"}, token.Aud)
	assert.Equal([]string{"InvoicingAPI"}, token.Scope)
	assert.Equal("d67902d2-58a6-4b82-9e68-05b1e4635f03", token.ClientID)
	assert.Equal("1", token.IsTaxRepres)
	assert.Equal("0", token.IsIntermediary)
	assert.Equal("0", token.IntermedID)
	assert.Equal("2", token.IntermedEnforced)
	assert.Equal("C12345678900:d67902d2-58a6-4b82-9e68-05b1e4635f03", token.Name)
	assert.Equal("aeca7f58-bfd0-46e3-86ab-5d4a062b7e71", token.SSID)
	assert.Equal("Test", token.PreferredUsername)
	assert.Equal("1234", token.TaxID)
	assert.Equal("C12345678900", token.TaxTin)
	assert.Equal("1234", token.ProfID)
	assert.Equal("0", token.IsTaxAdmin)
	assert.Equal("1", token.IsSystem)
	assert.Equal("", token.NatID)
}

func TestLoginAsTaxpayer(t *testing.T) {
	p := setupPlatformTest()
	assert := assert.New(t)

	token, err := p.LoginAsTaxpayer()
	assert.Nil(err)
	assert.NotEmpty(token.AccessToken)
	assert.Equal("Bearer", token.TokenType)
	assert.Equal(defaultScope, token.Scope)
	assert.Equal(3600, token.ExpiresIn)
}

func TestLoginAsIntermediary(t *testing.T) {
	p := setupPlatformTest()
	assert := assert.New(t)

	token, err := p.LoginAsIntermediaries(os.Getenv("TIN"))
	assert.Nil(err)
	assert.NotEmpty(token.AccessToken)
	assert.Equal("Bearer", token.TokenType)
	assert.Equal(defaultScope, token.Scope)
	assert.Equal(3600, token.ExpiresIn)

	payload, err := DecodeToken(token.AccessToken)
	assert.Nil(err)
	b, err := json.MarshalIndent(payload, "", "  ")
	assert.Nil(err)
	t.Log(string(b))

}

func TestGetAllDocumentTypes(t *testing.T) {
	p := setupPlatformTest()
	assert := assert.New(t)

	token, err := p.LoginAsTaxpayer()
	assert.Nil(err)
	assert.NotEmpty(token.AccessToken)

	documentTypes, err := p.GetAllDocumentTypes(token.AccessToken)
	assert.Nil(err)
	assert.NotEmpty(documentTypes.Result)
}

func TestGetDocumentType(t *testing.T) {
	p := setupPlatformTest()
	assert := assert.New(t)

	token, err := p.LoginAsTaxpayer()
	assert.Nil(err)
	assert.NotEmpty(token.AccessToken)

	documentType, err := p.GetDocumentType(token.AccessToken, 1)
	assert.Nil(err)
	assert.NotEmpty(documentType)
	assert.Equal(int64(1), documentType.ID)
	assert.Equal(int64(1), documentType.InvoiceTypeCode)
	assert.Equal("Invoice", documentType.Description)
	assert.Greater(len(documentType.DocumentTypeVersions), 0)
	assert.NotEmpty(documentType.DocumentTypeVersions)
}
