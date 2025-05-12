package client

type TestInput struct {
	BaseInput `json:",inline"`
	Param     string  `json:"param" url:"param"`
	Option    *string `json:"option,omitempty" url:"option,omitempty"`
}

type TestEchoData struct {
	Param  string  `json:"param"`
	Option *string `json:"option,omitempty"`
}

type TestResponse struct {
	BaseResponse `json:",inline"`
	Data         *TestEchoData `json:"data,omitempty"`
}

func (c *Client) Test(input TestInput) (*TestResponse, error) {
	var output TestResponse
	if err := c.request("test", &input, &output); err != nil {
		return nil, err
	}
	return &output, nil
}
