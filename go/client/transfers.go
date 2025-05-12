package client

import "time"

type Transfer struct {
	ID                 int64          `json:"id"`
	CreatedAt          time.Time      `json:"createdAt"`
	SettledAt          *time.Time     `json:"settledAt,omitempty"`
	Account            string         `json:"account"`
	MerchantTransferID string         `json:"merchantTransferId"`
	IsDeposit          bool           `json:"isDeposit"`
	CurrencyType       CurrencyType   `json:"currencyType"`
	Amount             string         `json:"amount"`
	AfterBalance       *string        `json:"afterBalance,omitempty"`
	Status             TransferStatus `json:"status"`
}

type DepositInput struct {
	BaseInput          `json:",inline"`
	MerchantTransferID string
	Account            string        `json:"account" url:"account"`
	CurrencyType       *CurrencyType `json:"currencyType" url:"currencyType"`
	Amount             string        `json:"amount" url:"amount"`
}

type DepositResponse struct {
	BaseResponse `json:",inline"`
	Data         *Transfer `json:"data,omitempty"`
}

func (c *Client) Deposit(input DepositInput) (*DepositResponse, error) {
	var output DepositResponse
	if err := c.request("deposit", &input, &output); err != nil {
		return nil, err
	}
	return &output, nil
}

type WithdrawInput = DepositInput

type WithdrawResponse = DepositResponse

func (c *Client) Withdraw(input WithdrawInput) (*WithdrawResponse, error) {
	var output WithdrawResponse
	if err := c.request("withdraw", &input, &output); err != nil {
		return nil, err
	}
	return &output, nil
}
