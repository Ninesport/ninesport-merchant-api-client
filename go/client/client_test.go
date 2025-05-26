package client_test

import (
	"os"
	"strconv"
	"testing"

	"github.com/Ninesport/ninesport-merchant-api-client/go/client"
	"github.com/stretchr/testify/assert"
)

func MustNewClient(assert *assert.Assertions) *client.Client {
	merchantCode := os.Getenv("NS_MERCHANT_CODE")
	secretKey := os.Getenv("NS_SECRET_KEY")
	baseUrl := os.Getenv("NS_BASE_URL")
	signType := os.Getenv("NS_SIGN_TYPE")
	signTypeInt, err := strconv.ParseInt(signType, 10, 64)
	if err != nil {
		panic(err)
	}
	c, err := client.NewClient(client.NewClientInput{
		MerchantCode: merchantCode,
		SecretKey:    secretKey,
		SignType:     client.SignType(signTypeInt),
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
