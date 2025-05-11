package client

type ListTransfersInput struct {
	BaseInput     `json:",inline"`
	PagenateInput `json:",inline"`
	From          int64 `json:"from" url:"from"`
	To            int64 `json:"to" url:"to"`
}

type ListTransfersData struct {
	PagenateResponse `json:",inline"`
	Records          []*Transfer `json:"records"`
}

type ListTransfersResponse struct {
	BaseResponse `json:",inline"`
	Data         *ListTransfersData `json:"data,omitempty"`
}

func (c *Client) ListTransfers(input *ListTransfersInput) (*ListTransfersResponse, error) {
	var output ListTransfersResponse
	if err := c.request("listTransfers", input, &output); err != nil {
		return nil, err
	}
	return &output, nil
}

type GetTransferInput struct {
	BaseInput          `json:",inline"`
	MerchantTransferID string
}

type GetTransferResponse struct {
	BaseResponse `json:",inline"`
	Data         *Transfer `json:"data,omitempty"`
}

func (c *Client) GetTransfer(input *GetTransferInput) (*GetTransferResponse, error) {
	var output GetTransferResponse
	if err := c.request("getTransfer", input, &output); err != nil {
		return nil, err
	}
	return &output, nil
}
