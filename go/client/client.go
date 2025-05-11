package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

type Client struct {
	merchantCode string
	secretKey    string
	baseURL      *url.URL
	timeout      time.Duration
}

type NewClientInput struct {
	MerchantCode string
	SecretKey    string
	// Ex: https://xxxxx/api/v1
	BaseURL string
	Timeout time.Duration
}

func NewClient(input NewClientInput) (*Client, error) {
	u, err := url.Parse(input.BaseURL)
	if err != nil {
		return nil, err
	}

	timeout := time.Second * 5
	if input.Timeout > 0 {
		timeout = input.Timeout
	}

	c := &Client{
		merchantCode: input.MerchantCode,
		secretKey:    input.SecretKey,
		baseURL:      u,
		timeout:      timeout,
	}
	resp, err := c.Test(&TestInput{})
	if err != nil {
		return nil, err
	}
	if resp.Code != API_STATUS_CODE_SUCCESS {
		return nil, fmt.Errorf("[%d] NewClient test failed: %s", resp.Code, resp.Msg)
	}
	return c, nil
}

func (c *Client) request(path string, input BaseInputer, output any) error {
	input.SetTimestamp()
	input.SetMerchantCode(c.merchantCode)
	queryString, err := ToQueryString(input, true)
	if err != nil {
		return err
	}
	sign, err := Sign(c.secretKey, queryString)
	if err != nil {
		return err
	}
	input.SetSign(sign)

	body, err := json.Marshal(input)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), c.timeout)
	defer cancel()

	buf := bytes.NewBuffer(body)
	urlString := c.baseURL.JoinPath(path).String()
	req, err := http.NewRequestWithContext(ctx, "POST", urlString, buf)
	if err != nil {
		return err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	responseData, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("url: %s, status: [%d] %s, body: %s",
			urlString, resp.StatusCode, resp.Status, responseData)
	}
	err = json.Unmarshal(responseData, output)
	if err != nil {
		return err
	}

	return nil
}
