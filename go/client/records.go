package client

type ListTransfersInput struct {
	BaseInput         `json:",inline"`
	PagenateInput     `json:",inline"`
	From              int64   `json:"from" url:"from"`
	To                int64   `json:"to" url:"to"`
	Account           *string `json:"account,omitempty" url:"account,omitempty"`
	IsDeposit         *bool   `json:"isDeposit,omitempty" url:"isDeposit,omitempty"`
	IsSettledInterval *bool   `json:"isSettledInterval,omitempty" url:"isSettledInterval,omitempty"`
}

type ListTransfersData struct {
	PagenateResponse `json:",inline"`
	Records          []*Transfer `json:"records"`
}

type ListTransfersResponse struct {
	BaseResponse `json:",inline"`
	Data         *ListTransfersData `json:"data,omitempty"`
}

func (c *Client) ListTransfers(input ListTransfersInput) (*ListTransfersResponse, error) {
	var output ListTransfersResponse
	if err := c.request("listTransfers", &input, &output); err != nil {
		return nil, err
	}
	return &output, nil
}

type GetTransferInput struct {
	BaseInput          `json:",inline"`
	MerchantTransferID string `json:"merchantTransferID" url:"merchantTransferID"`
}

type GetTransferResponse struct {
	BaseResponse `json:",inline"`
	Data         *Transfer `json:"data,omitempty"`
}

func (c *Client) GetTransfer(input GetTransferInput) (*GetTransferResponse, error) {
	var output GetTransferResponse
	if err := c.request("getTransfer", &input, &output); err != nil {
		return nil, err
	}
	return &output, nil
}

type ListBetRecordsInput struct {
	BaseInput         `json:",inline"`
	PagenateInput     `json:",inline"`
	From              int64        `json:"from" url:"from"`
	To                int64        `json:"to" url:"to"`
	Account           *string      `json:"account,omitempty" url:"account,omitempty"`
	Status            *OrderStatus `json:"status,omitempty" url:"status,omitempty"`
	Language          *string      `json:"language,omitempty" url:"language,omitempty"`
	IsSettledInterval *bool        `json:"isSettledInterval,omitempty" url:"isSettledInterval,omitempty"`
	SportID           *int64       `json:"sportID,omitempty" url:"sportID,omitempty"`
}
type ListBetRecordsData struct {
	PagenateResponse `json:",inline"`
	Records          []*Order `json:"records"`
}

type ListBetRecordsResponse struct {
	BaseResponse `json:",inline"`
	Data         *ListBetRecordsData `json:"data,omitempty"`
}

func (c *Client) ListBetRecords(input ListBetRecordsInput) (*ListBetRecordsResponse, error) {
	var output ListBetRecordsResponse
	if err := c.request("listBetRecords", &input, &output); err != nil {
		return nil, err
	}
	return &output, nil
}

type GetBetRecordInput struct {
	BaseInput `json:",inline"`
	BetID     string  `json:"betId" url:"betId" binding:"required"`
	Language  *string `json:"language,omitempty" url:"language,omitempty"`
}

type GetBetRecordResponse struct {
	BaseResponse `json:",inline"`
	Data         *Order `json:"data,omitempty"`
}

func (c *Client) GetBetRecord(input GetBetRecordInput) (*GetBetRecordResponse, error) {
	var output GetBetRecordResponse
	if err := c.request("getBetRecord", &input, &output); err != nil {
		return nil, err
	}
	return &output, nil
}
