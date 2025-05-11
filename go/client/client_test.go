package client_test

import (
	"os"
	"testing"

	"github.com/Ninesport/ninesport-merchant-api-client/go/client"
)

func MustNewClient(t *testing.T) *client.Client {
	merchantCode := os.Getenv("NS_MERCHANT_CODE")
	secretKey := os.Getenv("NS_SECRET_KEY")
	baseUrl := os.Getenv("NS_BASE_URL")
	c, err := client.NewClient(client.NewClientInput{
		MerchantCode: merchantCode,
		SecretKey:    secretKey,
		BaseURL:      baseUrl,
	})
	if err != nil {
		t.Error(err)
		panic(err)
	}
	return c
}

func TestNewClient(t *testing.T) {
	MustNewClient(t)
}
