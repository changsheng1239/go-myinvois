package myinvois

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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

	return NewClient(ClientOption{
		Environment:  Sandbox,
		Timeout:      DefaultTimeout,
		ClientID:     os.Getenv("CLIENT_ID"),
		ClientSecret: os.Getenv("CLIENT_SECRET"),
		Cert:         cert,
		PrivKey:      key,
		PrivKeyPass:  []byte(os.Getenv("PKEY_PASSWORD")),
	})
}

func login(p *Client) *OAuth2Token {
	token, err := p.LoginAsTaxpayer()
	if err != nil {
		panic("login failed: " + err.Error())
	}

	return token
}

func loginAsIntermediary(p *Client, onbehalfof string) *OAuth2Token {
	token, err := p.LoginAsIntermediaries(onbehalfof)
	if err != nil {
		panic("login failed: " + err.Error())
	}

	return token
}

func TestDecodeToken(t *testing.T) {
	assert := assert.New(t)

	tokenString := `eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJodHRwczovL3ByZXByb2QtaWRlbnRpdHkubXlpbnZvaXMuaGFzaWwuZ292Lm15IiwibmJmIjoxNzIwNDkxMTA2LCJpYXQiOjE3MjA0OTExMDYsImV4cCI6MTcyMDQ5NDcwNiwiYXVkIjpbIkludm9pY2luZ0FQSSIsImh0dHBzOi8vcHJlcHJvZC1pZGVudGl0eS5teWludm9pcy5oYXNpbC5nb3YubXkvcmVzb3VyY2VzIl0sInNjb3BlIjpbIkludm9pY2luZ0FQSSJdLCJjbGllbnRfaWQiOiJkNjc5MDJkMi01OGE2LTRiODItOWU2OC0wNWIxZTQ2MzVmMDMiLCJJc1RheFJlcHJlcyI6IjEiLCJJc0ludGVybWVkaWFyeSI6IjAiLCJJbnRlcm1lZElkIjoiMCIsIkludGVybWVkVElOIjoiIiwiSW50ZXJtZWRFbmZvcmNlZCI6IjIiLCJuYW1lIjoiQzEyMzQ1Njc4OTAwOmQ2NzkwMmQyLTU4YTYtNGI4Mi05ZTY4LTA1YjFlNDYzNWYwMyIsIlNTSWQiOiJhZWNhN2Y1OC1iZmQwLTQ2ZTMtODZhYi01ZDRhMDYyYjdlNzEiLCJwcmVmZXJyZWRfdXNlcm5hbWUiOiJUZXN0IiwiVGF4SWQiOiIxMjM0IiwiVGF4cGF5ZXJUSU4iOiJDMTIzNDU2Nzg5MDAiLCJQcm9mSWQiOiIxMjM0IiwiSXNUYXhBZG1pbiI6IjAiLCJJc1N5c3RlbSI6IjEiLCJOYXRJZCI6IiJ9.YSKL4GyT_OMOmncF2PsB8QiMeDuLhY857w-WH0m-3SdjbZ7NfFut6MeXvlznECMGKgymib8yb10v4ahd1KH73dG-aCRv2azxaMRnyeHgn1TbF0W8x0iKS1yegz6uRFOuuvqpnqLvCOPuG-x3iQEmlsswkPI2v8m11Wrb0ppb5QhXXWdoQSs28YMpKOfyQvpGCtKkEQAOClg9rGQIjuB3Zfprbl-u6ljl3izGuttv9M0d_7w43KLdzNEodOtajmu6058FZhUWlX6YPGAkCootTFmB5J1siDPBIc46NGUzAnG6SPnc_4fxO5MlX5TrOsWpLgU-R3eCrV9aVo2YTgUaHQ`

	token, err := DecodeToken(tokenString)
	assert.Nil(err)
	assert.Equal("https://preprod-identity.myinvois.hasil.gov.my", token.Iss)
	assert.Equal(1720491106, token.Nbf)
	assert.Equal(1720491106, token.Iat)
	assert.Equal(1720494706, token.Exp)
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
	assert.Equal("C12345678900", token.TaxpayerTIN)
	assert.Equal("1234", token.ProfID)
	assert.Equal("0", token.IsTaxAdmin)
	assert.Equal("1", token.IsSystem)
	assert.Equal("", token.NatID)
}

func validateToken(t *testing.T, accessTokenString string) {
	assert := assert.New(t)

	token, err := DecodeToken(accessTokenString)
	assert.Nil(err)

	b, err := json.MarshalIndent(token, "", "  ")
	if err != nil {
		panic(err)
	}
	t.Log(string(b))

	assert.Greater(token.Nbf, 0)
	assert.Greater(token.Iat, 0)
	assert.Greater(token.Exp, 0)
	assert.Equal("https://preprod-identity.myinvois.hasil.gov.my", token.Iss)
	assert.Equal(CustomArray{"InvoicingAPI", "https://preprod-identity.myinvois.hasil.gov.my/resources"}, token.Aud)
	assert.Equal([]string{"InvoicingAPI"}, token.Scope)
	assert.NotEmpty(token.PreferredUsername)
	assert.NotEmpty(token.TaxID)
	assert.NotEmpty(token.TaxpayerTIN)
	assert.NotEmpty(token.ProfID)
	assert.NotEmpty(token.IsTaxAdmin)
	assert.NotEmpty(token.IsSystem)
	assert.NotEmpty(token.IntermedID)
	assert.Equal(36, len(token.ClientID)) // uuid format d67902d2-58a6-4b82-9e68-05b1e4635f03
	assert.Equal(36, len(token.SSID))
	assert.Equal(1, len(token.IsTaxRepres))
	assert.Equal(1, len(token.IsIntermediary))
	assert.Equal(1, len(token.IntermedEnforced))
	assert.Equal(token.TaxpayerTIN+":"+token.ClientID, token.Name)
}

func TestLoginAsTaxpayer(t *testing.T) {
	p := setupPlatformTest()
	assert := assert.New(t)
	require := require.New(t)

	t.Run("Valid Login", func(t *testing.T) {
		token, err := p.LoginAsTaxpayer()
		require.Nil(err)
		require.NotEmpty(token.AccessToken)
		assert.Equal("Bearer", token.TokenType)
		assert.Equal(defaultScope, token.Scope)
		assert.Equal(3600, token.ExpiresIn)
		validateToken(t, token.AccessToken)
	})

	t.Run("Invalid Login", func(t *testing.T) {
		p.PlatformAPI.clientID = "invalid-client-id"

		token, err := p.LoginAsTaxpayer()
		require.NotNil(err)
		require.ErrorIs(err, ErrInvalidCredential)
		assert.Nil(token)
	})
}

func TestLoginAsIntermediary(t *testing.T) {
	p := setupPlatformTest()
	assert := assert.New(t)
	require := require.New(t)

	t.Run("Valid Login As Intermediary", func(t *testing.T) {
		token, err := p.LoginAsIntermediaries(os.Getenv("TIN"))
		require.Nil(err)
		require.NotEmpty(token.AccessToken)
		assert.Equal("Bearer", token.TokenType)
		assert.Equal(defaultScope, token.Scope)
		assert.Equal(3600, token.ExpiresIn)
		validateToken(t, token.AccessToken)
	})

	t.Run("Invalid Login As Intermediary", func(t *testing.T) {
		token, err := p.LoginAsIntermediaries("random-onbehalfof")
		require.NotNil(err)
		require.ErrorIs(err, ErrUnauthorizedIntermediary)
		assert.Nil(token)
	})
}

func TestGetAllDocumentTypes(t *testing.T) {
	p := setupPlatformTest()
	token := login(p)
	require := require.New(t)

	documentTypes, err := p.GetAllDocumentTypes(token.AccessToken)
	require.Nil(err)
	require.NotEmpty(documentTypes.Result)
}

func TestGetDocumentType(t *testing.T) {
	p := setupPlatformTest()
	token := login(p)
	assert := assert.New(t)
	require := require.New(t)

	documentType, err := p.GetDocumentType(token.AccessToken, 1)
	require.Nil(err)
	require.NotEmpty(documentType)
	assert.Equal(int64(1), documentType.ID)
	assert.Equal("01", documentType.InvoiceTypeCode)
	assert.Equal("Invoice", documentType.Description)
	assert.Greater(len(documentType.DocumentTypeVersions), 0)
	assert.NotEmpty(documentType.DocumentTypeVersions)
}
