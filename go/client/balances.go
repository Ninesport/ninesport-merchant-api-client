package client

type Wallet struct {
	ID             int64        `json:"id"`
	Account        string       `json:"account"`
	CurrencyType   CurrencyType `json:"currencyType"`
	IsActive       bool         `json:"isActive"`
	Balance        string       `json:"balance"`
	FreezedBalance string       `json:"freezedBalance"`
}

type GetBalanceInput struct {
	BaseInput    `json:",inline"`
	Account      string        `json:"account" url:"account"`
	CurrencyType *CurrencyType `json:"currencyType,omitempty" url:"currencyType,omitempty"`
}

type GetBalanceResponse struct {
	BaseResponse `json:",inline"`
	Data         *Wallet `json:"data,omitempty"`
}

func (c *Client) GetBalance(input GetBalanceInput) (*GetBalanceResponse, error) {
	var output GetBalanceResponse
	if err := c.request("getBalance", &input, &output); err != nil {
		return nil, err
	}
	return &output, nil
}

type ListBalancesInput struct {
	BaseInput `json:",inline"`
	Account   string `json:"account" url:"account"`
}

type ListBalancesResponse struct {
	BaseResponse `json:",inline"`
	Data         []*Wallet `json:"data,omitempty"`
}

func (c *Client) ListBalances(input ListBalancesInput) (*ListBalancesResponse, error) {
	var output ListBalancesResponse
	if err := c.request("listBalances", &input, &output); err != nil {
		return nil, err
	}
	return &output, nil
}
