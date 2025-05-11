package client

import "time"

type Player struct {
	ID          int64      `json:"id"`
	CreatedAt   time.Time  `json:"createdAt"`
	Account     string     `json:"account"`
	LastLoginAt *time.Time `json:"lastLoginAt,omitempty"`
	IsOnline    bool       `json:"isOnline"`
	IsActive    bool       `json:"isActive"`
	Nickname    *string    `json:"nickName,omitempty"`
}

type CreatePlayerInput struct {
	BaseInput `json:",inline"`
	Account   string `json:"account" url:"account"`
}

type CreatePlayerResponse struct {
	BaseResponse `json:",inline"`
	Data         *Player `json:"data,omitempty"`
}

func (c *Client) CreatePlayer(input CreatePlayerInput) (*CreatePlayerResponse, error) {
	var output CreatePlayerResponse
	if err := c.request("createPlayer", &input, &output); err != nil {
		return nil, err
	}
	return &output, nil
}

type GetPlayerInput struct {
	BaseInput `json:",inline"`
	Account   string `json:"account" url:"account"`
}

type GetPlayerResponse struct {
	BaseResponse `json:",inline"`
	Data         *Player `json:"data,omitempty"`
}

func (c *Client) GetPlayer(input GetPlayerInput) (*GetPlayerResponse, error) {
	var output GetPlayerResponse
	if err := c.request("getPlayer", &input, &output); err != nil {
		return nil, err
	}
	return &output, nil
}

type ListOnlinePlayersInput struct {
	BaseInput     `json:",inline"`
	PagenateInput `json:",inline"`
}

type ListOnlinePlayersData struct {
	PagenateResponse `json:",inline"`
	Records          []*Player `json:"records"`
}

type ListOnlinePlayersResponse struct {
	BaseResponse `json:",inline"`
	Data         *ListOnlinePlayersData `json:"data,omitempty"`
}

func (c *Client) ListOnlinePlayers(input ListOnlinePlayersInput) (*ListOnlinePlayersResponse, error) {
	var output ListOnlinePlayersResponse
	if err := c.request("listOnlinePlayers", &input, &output); err != nil {
		return nil, err
	}
	return &output, nil
}
