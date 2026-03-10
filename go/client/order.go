package client

import (
	"time"
)

type Order struct {
	ID                            int64       `json:"id"`
	Status                        OrderStatus `json:"status"`
	CreatedAt                     time.Time   `json:"createdAt"`
	SettledAt                     *time.Time  `json:"settledAt,omitempty"`
	Account                       string      `json:"account"`
	BillCurrencyType              string      `json:"billCurrencyType"`
	CurrencyType                  string      `json:"currencyType"`
	ExchangeRate                  string      `json:"exchangeRate"`
	CombinationBetOptionID        int64       `json:"combinationBetOptionID"`
	CombinationBetOptionLocalName string      `json:"combinationBetOptionLocalName"`
	FoldSize                      uint8       `json:"foldSize"`
	PerStakeAmount                string      `json:"perStakeAmount"`
	CombinationCount              int         `json:"combinationCount"`
	// 下注額
	StakeAmount           string  `json:"stakeAmount"`
	MaxPriceWithoutDiv    string  `json:"maxPriceWithoutDiv"`
	MaxInexactPrice       float64 `json:"maxInexactPriced"`
	MaxPayout             string  `json:"maxPayout"`
	PayoutPriceWithoutDiv *string `json:"payoutPriceWithoutDiv"`
	// 實際賠率
	ActualInexactPrice *float64 `json:"actualInexactPrice"`
	// 實際賠付額
	PayoutAmount *string `json:"payoutAmount"`
	// 有效投注额
	Turnover *string `json:"turnover"`
	IsSingle bool    `json:"isSingle"`
	// 子注
	Legs []*Leg `json:"legs"`
}

type Leg struct {
	SportID            int64      `json:"sportID"`
	SportName          string     `json:"sportName"`
	LeagueName         string     `json:"leagueName"`
	ParticipantName    []string   `json:"participantName"`
	MarketTypeName     string     `json:"marketTypeName"`
	BetName            string     `json:"betName"`
	ActualPrice        string     `json:"actualPrice"`
	PayoutPrice        *string    `json:"payoutPrice"`
	StartedAt          *time.Time `json:"startedAt"`
	FinishedAt         *time.Time `json:"finishedAt"`
	LastUpdatedAt      time.Time  `json:"lastUpdatedAt"`
	SettlementTypeName *string    `json:"settlementTypeName"`
	IsOutrightLeague   bool       `json:"isOutrightLeague"`
}
