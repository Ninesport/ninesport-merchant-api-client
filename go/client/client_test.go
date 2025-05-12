package client_test

import (
	"os"
	"testing"

	"github.com/Ninesport/ninesport-merchant-api-client/go/client"
	"github.com/stretchr/testify/assert"
)

func MustNewClient(assert *assert.Assertions) *client.Client {
	merchantCode := os.Getenv("NS_MERCHANT_CODE")
	secretKey := os.Getenv("NS_SECRET_KEY")
	baseUrl := os.Getenv("NS_BASE_URL")
	c, err := client.NewClient(client.NewClientInput{
		MerchantCode: merchantCode,
		SecretKey:    secretKey,
		BaseURL:      baseUrl,
	})
	if !assert.Nil(err) {
		panic(err)
	}
	return c
}

func TestNewClient(t *testing.T) {
	assert := assert.New(t)
	MustNewClient(assert)
}
