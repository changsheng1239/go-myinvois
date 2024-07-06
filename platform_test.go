package myinvois

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

var PlatformAPIClient = NewPlatformClient(SandboxClient())
var clientID = ""
var clientSecret = ""

func setup() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	clientID = os.Getenv("CLIENT_ID")
	clientSecret = os.Getenv("CLIENT_SECRET")
}
func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	os.Exit(code)
}

func TestLoginAsTaxpayer(t *testing.T) {
	assert := assert.New(t)

	token, err := PlatformAPIClient.LoginAsTaxpayer(clientID, clientSecret)
	assert.Nil(err)
	assert.NotEmpty(token.AccessToken)
	assert.Equal("Bearer", token.TokenType)
	assert.Equal(DefaultScope, token.Scope)
	assert.Equal(3600, token.ExpiresIn)
}
