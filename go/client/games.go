package client

type LoginPlayerInput struct {
	BaseInput `json:",inline"`
	Account   string  `json:"account" url:"account"`
	ReturnUrl *string `json:"returnUrl,omitempty" url:"returnUrl,omitempty"`
	Ip        *string `json:"ip,omitempty" url:"ip,omitempty"`
}

type LoginPlayerData struct {
	URL string `json:"url"`
}

type LoginPlayerResponse struct {
	BaseResponse `json:",inline"`
	Data         *LoginPlayerData `json:"data,omitempty"`
}

func (c *Client) LoginPlayer(input LoginPlayerInput) (*LoginPlayerResponse, error) {
	var output LoginPlayerResponse
	if err := c.request("login", &input, &output); err != nil {
		return nil, err
	}
	return &output, nil
}

type LogoutPlayerInput struct {
	BaseInput `json:",inline"`
	Account   string `json:"account" url:"account"`
}

type LogoutPlayerResponse struct {
	BaseResponse `json:",inline"`
	Data         *string `json:"data,omitempty"`
}

func (c *Client) LogoutPlayer(input LogoutPlayerInput) (*LogoutPlayerResponse, error) {
	var output LogoutPlayerResponse
	if err := c.request("logout", &input, &output); err != nil {
		return nil, err
	}
	return &output, nil
}

type LogoutAllPlayersInput struct {
	BaseInput `json:",inline"`
}

type LogoutAllPlayerResponse struct {
	BaseResponse `json:",inline"`
	Data         *string `json:"data,omitempty"`
}

func (c *Client) LogoutAllPlayers(input LogoutAllPlayersInput) (*LogoutAllPlayerResponse, error) {
	var output LogoutAllPlayerResponse
	if err := c.request("logoutAll", &input, &output); err != nil {
		return nil, err
	}
	return &output, nil
}
