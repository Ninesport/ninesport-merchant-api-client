package client_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/Ninesport/ninesport-merchant-api-client/go/client"
	"github.com/stretchr/testify/assert"
)

func TestBetRecords(t *testing.T) {
	assert := assert.New(t)
	c := MustNewClient(assert)

	from := time.Now().Add(-30 * 24 * time.Hour).UnixMilli()
	to := time.Now().UnixMilli()
	pageSize := 10

	listResp, err := c.ListBetRecords(client.ListBetRecordsInput{
		From: from,
		To:   to,
		PagenateInput: client.PagenateInput{
			Page:     1,
			PageSize: &pageSize,
		},
	})

	if !assert.Nil(err, "listBetRecords must success, but got error: %v", err) {
		return
	}
	assert.Equal(client.API_STATUS_CODE_SUCCESS, listResp.Code, "listBetRecords must success, but got: [%d] %s", listResp.Code, listResp.Msg)
	if !assert.NotNil(listResp.Data, "listBetRecords data should not be nil") {
		return
	}
	if !assert.NotEmpty(listResp.Data.Records, "No bet records found, cannot test GetBetRecord") {
		return
	}
	// 測試 GetBetRecord
	firstRecord := listResp.Data.Records[0]
	betID := fmt.Sprintf("%d", firstRecord.ID)

	t.Logf("Found %d bet records, testing GetBetRecord with ID: %s", len(listResp.Data.Records), betID)

	getResp, err := c.GetBetRecord(client.GetBetRecordInput{
		BetID: betID,
	})

	if !assert.Nil(err, "getBetRecord must success, but got error: %v", err) {
		return
	}
	assert.Equal(client.API_STATUS_CODE_SUCCESS, getResp.Code, "getBetRecord must success, but got: [%d] %s", getResp.Code, getResp.Msg)
	if !assert.NotNil(getResp.Data, "getBetRecord data should not be nil") {
		return
	}
	assert.Equal(firstRecord.ID, getResp.Data.ID, "getBetRecord should return the same bet record")
	assert.Equal(firstRecord.Account, getResp.Data.Account, "account should match")
}
