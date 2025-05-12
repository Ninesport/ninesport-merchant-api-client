package client_test

import (
	"testing"
	"time"

	"github.com/Ninesport/ninesport-merchant-api-client/go/client"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func PointerBool(b bool) *bool {
	return &b
}

func TestTransfers(t *testing.T) {
	assert := assert.New(t)
	randomAmount := decimal.NewFromFloat(gofakeit.Float64Range(1, 1000000))
	currencyType := client.CurrencyTypes[gofakeit.IntRange(0, len(client.CurrencyTypes)-1)]

	startedAt := time.Now().Add(-time.Second * 2)

	depositID := gofakeit.UUID()
	withdrawID := gofakeit.UUID()
	overWithdrawID := gofakeit.UUID()

	c := MustNewClient(assert)

	overWithdrawResp, err := c.Withdraw(client.WithdrawInput{
		Account:            randomAccount,
		CurrencyType:       &currencyType,
		MerchantTransferID: overWithdrawID,
		Amount:             randomAmount.String(),
	})
	if !assert.Nil(err) {
		return
	}
	assert.Equal(client.API_STATUS_CODE_PLAYER_INSUFFICIENT_BALANCE, overWithdrawResp.Code, "expect insufficient balance, but got: [%d] %s", overWithdrawResp.Code, overWithdrawResp.Msg)
	assert.Nil(overWithdrawResp.Data)

	overWithdrawTransfer, err := c.GetTransfer(client.GetTransferInput{
		MerchantTransferID: overWithdrawID,
	})
	if !assert.Nil(err) {
		return
	}
	assert.Equal(client.API_STATUS_CODE_SUCCESS, overWithdrawTransfer.Code, "GetTransfer failed: [%d] %s", overWithdrawTransfer.Code, overWithdrawTransfer.Msg)
	assert.NotNil(overWithdrawTransfer.Data)
	assert.Equal(client.TRANSFER_STATUS_INSUFFICIENT_BALANCE, overWithdrawTransfer.Data.Status, "expect insufficient balance, but got: %d", overWithdrawTransfer.Data.Status)

	depositResp, err := c.Deposit(client.DepositInput{
		Account:            randomAccount,
		CurrencyType:       &currencyType,
		MerchantTransferID: depositID,
		Amount:             randomAmount.String(),
	})
	if !assert.Nil(err) {
		return
	}
	assert.Equal(client.API_STATUS_CODE_SUCCESS, depositResp.Code, "Deposit failed: [%d] %s", depositResp.Code, depositResp.Msg)
	assert.True(depositResp.Data.IsDeposit)

	withdrawResp, err := c.Withdraw(client.WithdrawInput{
		Account:            randomAccount,
		CurrencyType:       &currencyType,
		MerchantTransferID: withdrawID,
		Amount:             randomAmount.String(),
	})
	if !assert.Nil(err) {
		return
	}
	assert.Equal(client.API_STATUS_CODE_SUCCESS, withdrawResp.Code, "Withdraw failed: [%d] %s", withdrawResp.Code, withdrawResp.Msg)
	assert.False(withdrawResp.Data.IsDeposit)

	endedAt := time.Now().Add(time.Second * 2)

	pageSize := 1000
	listResp, err := c.ListTransfers(client.ListTransfersInput{
		PagenateInput: client.PagenateInput{
			Page:     1,
			PageSize: &pageSize,
		},
		From: startedAt.UnixMilli(),
		To:   endedAt.UnixMilli(),
	})
	if !assert.Nil(err) {
		return
	}
	assert.Equal(client.API_STATUS_CODE_SUCCESS, listResp.Code, "ListTransfers failed: [%d] %s", listResp.Code, listResp.Msg)
	if !assert.NotNil(listResp.Data) {
		return
	}
	assert.Equal(3, int(listResp.Data.Total))
	assert.Equal(1, listResp.Data.Page)

	foundDeposit := false
	foundWithdraw := false
	foundOverWithdraw := false
	for _, transfer := range listResp.Data.Records {
		switch transfer.MerchantTransferID {
		case depositID:
			foundDeposit = true
			assert.Equal(depositResp.Data.ID, transfer.ID)
		case withdrawID:
			foundWithdraw = true
			assert.Equal(withdrawResp.Data.ID, transfer.ID)
		case overWithdrawID:
			foundOverWithdraw = true
			assert.Equal(overWithdrawTransfer.Data.ID, transfer.ID)
		default:
			t.Errorf("found unknown transfer: %d\n", transfer.ID)
		}
	}
	assert.True(foundDeposit, "not found deposit")
	assert.True(foundWithdraw, "not found withdraw")
	assert.True(foundOverWithdraw, "not found overWithdraw")

	for _, isDeposit := range []*bool{PointerBool(true), PointerBool(false), nil} {
		for _, isSettled := range []*bool{PointerBool(true), PointerBool(false), nil} {
			for _, account := range []*string{&randomAccount, nil} {
				listResp, err := c.ListTransfers(client.ListTransfersInput{
					PagenateInput: client.PagenateInput{
						Page: 1,
					},
					From:              startedAt.UnixMilli(),
					To:                endedAt.UnixMilli(),
					IsDeposit:         isDeposit,
					IsSettledInterval: isSettled,
					Account:           account,
				})
				if !assert.Nil(err) {
					return
				}
				assert.Equal(client.API_STATUS_CODE_SUCCESS, listResp.Code)
				assert.NotNil(listResp.Data)
			}
		}
	}
}
