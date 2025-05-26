package client

import "time"

type BaseInputer interface {
	SetSign(sign string)
	SetTimestamp()
	SetSignType(signType SignType)
	// 回傳的值其實是UUID格式
	SetMerchantCode(merchantCode string)
}

type BaseInput struct {
	Timestamp    int64    `json:"timestamp" url:"timestamp"`
	MerchantCode string   `json:"merchantCode" url:"merchantCode"`
	SignType     SignType `json:"signType" url:"signType"`
	Sign         string   `json:"sign" url:"-"`
}

func (b *BaseInput) SetSign(sign string) {
	b.Sign = sign
}

func (b *BaseInput) SetTimestamp() {
	b.Timestamp = time.Now().UnixMilli()
}

func (b *BaseInput) SetMerchantCode(merchantCode string) {
	b.MerchantCode = merchantCode
}
func (b *BaseInput) SetSignType(signType SignType) {
	b.SignType = signType
}

type BaseResponse struct {
	Code      APIStatusCode `json:"code"`
	Msg       string        `json:"msg"`
	Timestamp int64         `json:"timestamp"`
}

type PagenateInput struct {
	Page     int  `json:"page" url:"page"`
	PageSize *int `json:"pageSize,omitempty" url:"pageSize,omitempty"`
}
type PagenateResponse struct {
	Page        int   `json:"page"`
	PageSize    int   `json:"pageSize"`
	Total       int64 `json:"total"`
	HasNext     bool  `json:"hasNext"`
	HasPrevious bool  `json:"hasPrevious"`
}
