package client_test

import (
	"testing"

	"github.com/Ninesport/ninesport-merchant-api-client/go/client"
	"github.com/stretchr/testify/assert"
)

func MustGetOrCreatePlayer(assert *assert.Assertions) *client.Player {
	c := MustNewClient(assert)
	getResp, err := c.GetPlayer(client.GetPlayerInput{
		Account: randomAccount,
	})
	if !assert.Nil(err) {
		panic(err)
	}
	switch getResp.Code {
	case client.API_STATUS_CODE_PLAYER_NOT_FOUND:
		createResp, err := c.CreatePlayer(client.CreatePlayerInput{
			Account: randomAccount,
		})
		if !assert.Nil(err, "createPlayer must success, but got error: %v", err) {
			panic(err)
		}
		assert.Equal(client.API_STATUS_CODE_SUCCESS, createResp.Code, "createPlayer must success, but got: [%d] %s", createResp.Code, createResp.Msg)
		if !assert.NotNil(createResp.Data) {
			panic("createPlayer must success, but got nil data")
		}
		assert.Equal(randomAccount, createResp.Data.Account)
		return createResp.Data
	case client.API_STATUS_CODE_SUCCESS:
		return getResp.Data
	default:
		panic(getResp.Msg)
	}

}

func TestPlayers(t *testing.T) {
	assert := assert.New(t)

	c := MustNewClient(assert)

	player := MustGetOrCreatePlayer(assert)
	assert.Equal(randomAccount, player.Account)

	pageSize := 1000
	listResp, err := c.ListOnlinePlayers(client.ListOnlinePlayersInput{
		PagenateInput: client.PagenateInput{
			Page:     1,
			PageSize: &pageSize,
		},
	})
	if !assert.Nil(err, "listOnlinePlayers must success, but got error: %v", err) {
		return
	}
	assert.Equal(client.API_STATUS_CODE_SUCCESS, listResp.Code, "listOnlinePlayers must success, but got: [%d] %s", listResp.Code, listResp.Msg)
	if !assert.NotNil(listResp.Data) {
		return
	}
}
