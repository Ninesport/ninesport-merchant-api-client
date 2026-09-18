package client

import (
	"fmt"
	"slices"
	"strings"
)

type APIStatusCode uint8

const (
	// Success
	API_STATUS_CODE_SUCCESS APIStatusCode = 1
	// Signature Error
	API_STATUS_CODE_SIGNATURE_ERROR APIStatusCode = 2
	// Invalid Parameters
	API_STATUS_CODE_INVALID_PARAMETERS APIStatusCode = 3
	// Merchant Not Found
	API_STATUS_CODE_MERCHANT_NOT_FOUND APIStatusCode = 4
	// Merchant Inactive
	API_STATUS_CODE_MERCHANT_INACTIVE APIStatusCode = 5
	// Player Not Found
	API_STATUS_CODE_PLAYER_NOT_FOUND APIStatusCode = 6
	// Player Inactive
	API_STATUS_CODE_PLAYER_INACTIVE APIStatusCode = 7
	// Player Already Exists
	API_STATUS_CODE_PLAYER_ALREADY_EXISTS APIStatusCode = 8
	// Create Player Failed
	API_STATUS_CODE_CREATE_PLAYER_FAILED APIStatusCode = 9
	// Order Not Found
	API_STATUS_CODE_ORDER_NOT_FOUND APIStatusCode = 10
	// Order ID Already Exists
	API_STATUS_CODE_ORDER_ID_ALREADY_EXISTS APIStatusCode = 11
	// Create Order Failed
	API_STATUS_CODE_CREATE_ORDER_FAILED APIStatusCode = 12
	// Transfer Not Found
	API_STATUS_CODE_TRANSFER_NOT_FOUND APIStatusCode = 13
	// Transfer ID Already Exists
	API_STATUS_CODE_TRANSFER_ID_ALREADY_EXISTS APIStatusCode = 14
	// Create Transfer Failed
	API_STATUS_CODE_CREATE_TRANSFER_FAILED APIStatusCode = 15
	// Merchant Insufficient Balance
	API_STATUS_CODE_MERCHANT_INSUFFICIENT_BALANCE APIStatusCode = 16
	// Player Insufficient Balance
	API_STATUS_CODE_PLAYER_INSUFFICIENT_BALANCE APIStatusCode = 17
	// Network Insufficient Fee
	API_STATUS_CODE_NETWORK_INSUFFICIENT_FEE APIStatusCode = 18
	// Network Unsupported
	API_STATUS_CODE_NETWORK_UNSUPPORTED APIStatusCode = 19
	// Asset Unsupported on this Network
	API_STATUS_CODE_ASSET_UNSUPPORTED_ON_THIS_NETWORK APIStatusCode = 20
	// Too Frequent Requests
	API_STATUS_CODE_TOO_FREQUENT_REQUESTS APIStatusCode = 21
	// Amount Must Be Positive
	API_STATUS_CODE_AMOUNT_MUST_BE_POSITIVE APIStatusCode = 22
	// Signature Not Found
	API_STATUS_CODE_SIGNATURE_NOT_FOUND APIStatusCode = 23
	// Merchant Cannot Deposit
	API_STATUS_CODE_MERCHANT_CANNOT_DEPOSIT APIStatusCode = 24
	// Merchant Cannot Withdraw
	API_STATUS_CODE_MERCHANT_CANNOT_WITHDRAW APIStatusCode = 25
	// IP Not in Whitelist
	API_STATUS_CODE_IP_NOT_IN_WHITELIST APIStatusCode = 26
	// Parse Input Error
	API_STATUS_CODE_PARSE_INPUT_ERROR APIStatusCode = 27
	// Invalid Timestamp
	API_STATUS_CODE_INVALID_TIMESTAMP APIStatusCode = 28
	// Invalid Merchant Code
	API_STATUS_CODE_INVALID_MERCHANT_CODE APIStatusCode = 29
	// Deposit Amount Too Low
	API_STATUS_CODE_DEPOSIT_AMOUNT_TOO_LOW APIStatusCode = 30
	// Withdrawal Amount Too Low
	API_STATUS_CODE_WITHDRAWAL_AMOUNT_TOO_LOW APIStatusCode = 31
	// Withdrawal Amount Too High
	API_STATUS_CODE_WITHDRAWAL_AMOUNT_TOO_HIGH APIStatusCode = 32
	// Invalid Time Interval
	API_STATUS_CODE_INVALID_TIME_INTERVAL APIStatusCode = 33
	// Wallet Not Found
	API_STATUS_CODE_WALLET_NOT_FOUND APIStatusCode = 34
	// Wallet Inactive
	API_STATUS_CODE_WALLET_INACTIVE APIStatusCode = 35
	// Invalid Sign Type
	API_STATUS_CODE_INVALID_SIGN_TYPE APIStatusCode = 36
	// Sign Type Not Supported
	API_STATUS_CODE_SIGN_TYPE_NOT_SUPPORTED APIStatusCode = 37
	// Server Is Under Maintenance
	API_STATUS_CODE_SERVER_IS_UNDER_MAINTENANCE APIStatusCode = 253
	// Internal Server Error
	API_STATUS_CODE_INTERNAL_SERVER_ERROR APIStatusCode = 254
	// Unknown Error
	API_STATUS_CODE_UNKNOWN_ERROR APIStatusCode = 255
)

var ApiStatusCodeMap map[APIStatusCode]string = map[APIStatusCode]string{
	API_STATUS_CODE_SUCCESS:                           "Success",
	API_STATUS_CODE_SIGNATURE_ERROR:                   "Signature Error",
	API_STATUS_CODE_INVALID_PARAMETERS:                "Invalid Parameters",
	API_STATUS_CODE_MERCHANT_NOT_FOUND:                "Merchant Not Found",
	API_STATUS_CODE_MERCHANT_INACTIVE:                 "Merchant Inactive",
	API_STATUS_CODE_PLAYER_NOT_FOUND:                  "Player Not Found",
	API_STATUS_CODE_PLAYER_INACTIVE:                   "Player Inactive",
	API_STATUS_CODE_PLAYER_ALREADY_EXISTS:             "Player Already Exists",
	API_STATUS_CODE_CREATE_PLAYER_FAILED:              "Create Player Failed",
	API_STATUS_CODE_ORDER_NOT_FOUND:                   "Order Not Found",
	API_STATUS_CODE_ORDER_ID_ALREADY_EXISTS:           "Order ID Already Exists",
	API_STATUS_CODE_CREATE_ORDER_FAILED:               "Create Order Failed",
	API_STATUS_CODE_TRANSFER_NOT_FOUND:                "Transfer Not Found",
	API_STATUS_CODE_TRANSFER_ID_ALREADY_EXISTS:        "Transfer ID Already Exists",
	API_STATUS_CODE_CREATE_TRANSFER_FAILED:            "Create Transfer Failed",
	API_STATUS_CODE_MERCHANT_INSUFFICIENT_BALANCE:     "Merchant Insufficient Balance",
	API_STATUS_CODE_PLAYER_INSUFFICIENT_BALANCE:       "Player Insufficient Balance",
	API_STATUS_CODE_NETWORK_INSUFFICIENT_FEE:          "Network Insufficient Fee",
	API_STATUS_CODE_NETWORK_UNSUPPORTED:               "Network Unsupported",
	API_STATUS_CODE_ASSET_UNSUPPORTED_ON_THIS_NETWORK: "Asset Unsupported on this Network",
	API_STATUS_CODE_TOO_FREQUENT_REQUESTS:             "Too Frequent Requests",
	API_STATUS_CODE_AMOUNT_MUST_BE_POSITIVE:           "Amount Must Be Positive",
	API_STATUS_CODE_SIGNATURE_NOT_FOUND:               "Signature Not Found",
	API_STATUS_CODE_MERCHANT_CANNOT_DEPOSIT:           "Merchant Cannot Deposit",
	API_STATUS_CODE_MERCHANT_CANNOT_WITHDRAW:          "Merchant Cannot Withdraw",
	API_STATUS_CODE_IP_NOT_IN_WHITELIST:               "IP Not in Whitelist",
	API_STATUS_CODE_PARSE_INPUT_ERROR:                 "Parse Input Error",
	API_STATUS_CODE_INVALID_TIMESTAMP:                 "Invalid Timestamp",
	API_STATUS_CODE_INVALID_MERCHANT_CODE:             "Invalid Merchant Code",
	API_STATUS_CODE_DEPOSIT_AMOUNT_TOO_LOW:            "Deposit Amount Too Low",
	API_STATUS_CODE_WITHDRAWAL_AMOUNT_TOO_LOW:         "Withdrawal Amount Too Low",
	API_STATUS_CODE_WITHDRAWAL_AMOUNT_TOO_HIGH:        "Withdrawal Amount Too High",
	API_STATUS_CODE_INVALID_TIME_INTERVAL:             "Invalid Time Interval",
	API_STATUS_CODE_WALLET_NOT_FOUND:                  "Wallet Not Found",
	API_STATUS_CODE_WALLET_INACTIVE:                   "Wallet Inactive",
	API_STATUS_CODE_INVALID_SIGN_TYPE:                 "Invalid Sign Type",
	API_STATUS_CODE_SIGN_TYPE_NOT_SUPPORTED:           "Sign Type Not Supported",
	API_STATUS_CODE_SERVER_IS_UNDER_MAINTENANCE:       "Server Is Under Maintenance",
	API_STATUS_CODE_INTERNAL_SERVER_ERROR:             "Internal Server Error",
	API_STATUS_CODE_UNKNOWN_ERROR:                     "Unknown Error",
}

func (code APIStatusCode) String(moreMessage string) string {
	moreMessage = strings.TrimSpace(moreMessage)
	msg := ApiStatusCodeMap[code]
	if moreMessage != "" {
		msg = fmt.Sprintf("[%s] %s", msg, moreMessage)
	}

	return msg
}

type CurrencyType string

const (
	CURRENCY_TYPE_ADA   CurrencyType = "ADA"
	CURRENCY_TYPE_AED   CurrencyType = "AED"
	CURRENCY_TYPE_ARS   CurrencyType = "ARS"
	CURRENCY_TYPE_ATOM  CurrencyType = "ATOM"
	CURRENCY_TYPE_AUD   CurrencyType = "AUD"
	CURRENCY_TYPE_AVAX  CurrencyType = "AVAX"
	CURRENCY_TYPE_BDT   CurrencyType = "BDT"
	CURRENCY_TYPE_BGN   CurrencyType = "BGN"
	CURRENCY_TYPE_BNB   CurrencyType = "BNB"
	CURRENCY_TYPE_BRL   CurrencyType = "BRL"
	CURRENCY_TYPE_BTC   CurrencyType = "BTC"
	CURRENCY_TYPE_CAD   CurrencyType = "CAD"
	CURRENCY_TYPE_CHF   CurrencyType = "CHF"
	CURRENCY_TYPE_CLP   CurrencyType = "CLP"
	CURRENCY_TYPE_CNY   CurrencyType = "CNY"
	CURRENCY_TYPE_COP   CurrencyType = "COP"
	CURRENCY_TYPE_CZK   CurrencyType = "CZK"
	CURRENCY_TYPE_DKK   CurrencyType = "DKK"
	CURRENCY_TYPE_DOGE  CurrencyType = "DOGE"
	CURRENCY_TYPE_DOT   CurrencyType = "DOT"
	CURRENCY_TYPE_DZD   CurrencyType = "DZD"
	CURRENCY_TYPE_EGP   CurrencyType = "EGP"
	CURRENCY_TYPE_ETH   CurrencyType = "ETH"
	CURRENCY_TYPE_EUR   CurrencyType = "EUR"
	CURRENCY_TYPE_GBP   CurrencyType = "GBP"
	CURRENCY_TYPE_HKD   CurrencyType = "HKD"
	CURRENCY_TYPE_HRK   CurrencyType = "HRK"
	CURRENCY_TYPE_HUF   CurrencyType = "HUF"
	CURRENCY_TYPE_HYPE  CurrencyType = "HYPE"
	CURRENCY_TYPE_IDR   CurrencyType = "IDR"
	CURRENCY_TYPE_ILS   CurrencyType = "ILS"
	CURRENCY_TYPE_INR   CurrencyType = "INR"
	CURRENCY_TYPE_JOD   CurrencyType = "JOD"
	CURRENCY_TYPE_JPY   CurrencyType = "JPY"
	CURRENCY_TYPE_KRW   CurrencyType = "KRW"
	CURRENCY_TYPE_KWD   CurrencyType = "KWD"
	CURRENCY_TYPE_LKR   CurrencyType = "LKR"
	CURRENCY_TYPE_LTC   CurrencyType = "LTC"
	CURRENCY_TYPE_MAD   CurrencyType = "MAD"
	CURRENCY_TYPE_MATIC CurrencyType = "MATIC"
	CURRENCY_TYPE_MXN   CurrencyType = "MXN"
	CURRENCY_TYPE_MYR   CurrencyType = "MYR"
	CURRENCY_TYPE_NGN   CurrencyType = "NGN"
	CURRENCY_TYPE_NOK   CurrencyType = "NOK"
	CURRENCY_TYPE_NZD   CurrencyType = "NZD"
	CURRENCY_TYPE_OMR   CurrencyType = "OMR"
	CURRENCY_TYPE_PHP   CurrencyType = "PHP"
	CURRENCY_TYPE_PKR   CurrencyType = "PKR"
	CURRENCY_TYPE_PLN   CurrencyType = "PLN"
	CURRENCY_TYPE_QAR   CurrencyType = "QAR"
	CURRENCY_TYPE_RON   CurrencyType = "RON"
	CURRENCY_TYPE_RUB   CurrencyType = "RUB"
	CURRENCY_TYPE_SAR   CurrencyType = "SAR"
	CURRENCY_TYPE_SEK   CurrencyType = "SEK"
	CURRENCY_TYPE_SGD   CurrencyType = "SGD"
	CURRENCY_TYPE_SOL   CurrencyType = "SOL"
	CURRENCY_TYPE_THB   CurrencyType = "THB"
	CURRENCY_TYPE_TON   CurrencyType = "TON"
	CURRENCY_TYPE_TRX   CurrencyType = "TRX"
	CURRENCY_TYPE_TRY   CurrencyType = "TRY"
	CURRENCY_TYPE_TWD   CurrencyType = "TWD"
	CURRENCY_TYPE_USD   CurrencyType = "USD"
	CURRENCY_TYPE_USDC  CurrencyType = "USDC"
	CURRENCY_TYPE_USDT  CurrencyType = "USDT"
	CURRENCY_TYPE_VND   CurrencyType = "VND"
	CURRENCY_TYPE_XRP   CurrencyType = "XRP"
	CURRENCY_TYPE_ZAR   CurrencyType = "ZAR"
)

var CurrencyTypes []CurrencyType = []CurrencyType{
	CURRENCY_TYPE_ADA,
	CURRENCY_TYPE_AED,
	CURRENCY_TYPE_ARS,
	CURRENCY_TYPE_ATOM,
	CURRENCY_TYPE_AUD,
	CURRENCY_TYPE_AVAX,
	CURRENCY_TYPE_BDT,
	CURRENCY_TYPE_BGN,
	CURRENCY_TYPE_BNB,
	CURRENCY_TYPE_BRL,
	CURRENCY_TYPE_BTC,
	CURRENCY_TYPE_CAD,
	CURRENCY_TYPE_CHF,
	CURRENCY_TYPE_CLP,
	CURRENCY_TYPE_CNY,
	CURRENCY_TYPE_COP,
	CURRENCY_TYPE_CZK,
	CURRENCY_TYPE_DKK,
	CURRENCY_TYPE_DOGE,
	CURRENCY_TYPE_DOT,
	CURRENCY_TYPE_DZD,
	CURRENCY_TYPE_EGP,
	CURRENCY_TYPE_ETH,
	CURRENCY_TYPE_EUR,
	CURRENCY_TYPE_GBP,
	CURRENCY_TYPE_HKD,
	CURRENCY_TYPE_HRK,
	CURRENCY_TYPE_HUF,
	CURRENCY_TYPE_HYPE,
	CURRENCY_TYPE_IDR,
	CURRENCY_TYPE_ILS,
	CURRENCY_TYPE_INR,
	CURRENCY_TYPE_JOD,
	CURRENCY_TYPE_JPY,
	CURRENCY_TYPE_KRW,
	CURRENCY_TYPE_KWD,
	CURRENCY_TYPE_LKR,
	CURRENCY_TYPE_LTC,
	CURRENCY_TYPE_MAD,
	CURRENCY_TYPE_MATIC,
	CURRENCY_TYPE_MXN,
	CURRENCY_TYPE_MYR,
	CURRENCY_TYPE_NGN,
	CURRENCY_TYPE_NOK,
	CURRENCY_TYPE_NZD,
	CURRENCY_TYPE_OMR,
	CURRENCY_TYPE_PHP,
	CURRENCY_TYPE_PKR,
	CURRENCY_TYPE_PLN,
	CURRENCY_TYPE_QAR,
	CURRENCY_TYPE_RON,
	CURRENCY_TYPE_RUB,
	CURRENCY_TYPE_SAR,
	CURRENCY_TYPE_SEK,
	CURRENCY_TYPE_SGD,
	CURRENCY_TYPE_SOL,
	CURRENCY_TYPE_THB,
	CURRENCY_TYPE_TON,
	CURRENCY_TYPE_TRX,
	CURRENCY_TYPE_TRY,
	CURRENCY_TYPE_TWD,
	CURRENCY_TYPE_USD,
	CURRENCY_TYPE_USDC,
	CURRENCY_TYPE_USDT,
	CURRENCY_TYPE_VND,
	CURRENCY_TYPE_XRP,
	CURRENCY_TYPE_ZAR,
}

func (c *CurrencyType) IsValid() bool {
	return slices.Contains(CurrencyTypes, *c)
}
func PointerCurrencyType(value CurrencyType) *CurrencyType {
	return &value
}

type TransferStatus uint8

const (
	TRANSFER_STATUS_INIT TransferStatus = iota + 1
	TRANSFER_STATUS_SUCCESS
	TRANSFER_STATUS_INSUFFICIENT_BALANCE
	TRANSFER_STATUS_INSUFFICIENT_APPLY_AMOUNT
	TRANSFER_STATUS_FAILED
	TRANSFER_STATUS_PROCESSING
	TRANSFER_STATUS_UNPAID
)

type SignType uint8

const (
	API_SIGN_TYPE_SHA256 SignType = iota + 1
	API_SIGN_TYPE_SHA512
	API_SIGN_TYPE_DILITHIUM2
	API_SIGN_TYPE_DILITHIUM3
	API_SIGN_TYPE_DILITHIUM5
	API_SIGN_TYPE_FALCON512
)

var SignTypes []SignType = []SignType{
	API_SIGN_TYPE_SHA256,
	API_SIGN_TYPE_SHA512,
	API_SIGN_TYPE_DILITHIUM2,
	API_SIGN_TYPE_DILITHIUM3,
	API_SIGN_TYPE_DILITHIUM5,
	API_SIGN_TYPE_FALCON512,
}

func (s SignType) IsValid() bool {
	return slices.Contains(SignTypes, s)
}
func PointerSignType(value SignType) *SignType {
	return &value
}

type OrderStatus uint8

func PointerOrderStatus(status OrderStatus) *OrderStatus {
	return &status
}

const (
	// 初始化: 建立訂單時只要輸入參數沒有錯誤，訂單都會被建立，之後才會開始直接扣款，並改變狀態
	ORDER_STATUS_INIT OrderStatus = iota + 1
	// 下注(扣款)成功(等待結算)
	ORDER_STATUS_STAKED
	// 退款(整張Order不管Bet內容為何，把STAKED AMOUNT還回) - 已結算
	ORDER_STATUS_REFUND
	// 沒收(整張Order不管Bet內容為何，把STAKED AMOUNT收走) - 已結算
	ORDER_STATUS_CONFISCATED
	// 標準的注單終點(輸贏都是) - 已結算
	ORDER_STATUS_SETTLED
	// 即使訂單成立了，也有機會餘額不足 - 死單
	ORDER_STATUS_INSUFFICIENT_BALANCE
	// 指定注單目前無法下注 - 死單
	ORDER_STATUS_BETS_INACTIVE
	// 低於投注限額 - 死單
	ORDER_STATUS_UNDER_BET_LIMIT
	// 高於投注限額 - 死單
	ORDER_STATUS_OVER_BET_LIMIT
	// 賠率已變動 - 死單
	ORDER_STATUS_PRICE_CHANGED
	// 串關組合禁止 - 死單
	ORDER_STATUS_FOBIDDEN_COMBINATION
	// 提前結算(Bet還沒結算就提前派彩了) - 已結算
	ORDER_STATUS_CASHOUT
	// 退組(行為同退款) - 已結算
	ORDER_STATUS_VOID
	// 手動全贏 - 已結算
	ORDER_STATUS_MANUAL_WON
	// 手動半贏 - 已結算
	ORDER_STATUS_MANUAL_HALF_WON
	// 手動全輸 - 已結算
	ORDER_STATUS_MANUAL_LOST
	// 手動半輸 - 已結算
	ORDER_STATUS_MANUAL_HALF_LOST
	// 回滾，狀態同STAKED(等待結算)
	ORDER_STATUS_ROLLBACKED
	// 未知錯誤 - 死單
	ORDER_STATUS_UNKNOWN = 255
)

type EventStatus uint8

const (
	// The event has not started yet
	EVENT_STATUS_NOT_STARTED_YET EventStatus = 1
	// The event is live
	EVENT_STATUS_IN_PROGRESS EventStatus = 2
	// The event is finished
	EVENT_STATUS_FINISHED EventStatus = 3
	// The event has been canceled
	EVENT_STATUS_CANCELED EventStatus = 4
	// The event has been postponed
	EVENT_STATUS_POSTPONED EventStatus = 5
	// The event has been interrupted
	EVENT_STATUS_INTERRUPTED EventStatus = 6
	// The event has been abandoned
	EVENT_STATUS_ABANDONED EventStatus = 7
	// The coverage for this event has been lost
	EVENT_STATUS_COVERAGE_LOST EventStatus = 8
	// The event has not started but is about to., NOTE: This status will be shown up to 30 minutes before the event has started
	EVENT_STATUS_ABOUT_TO_START EventStatus = 9
	// Timeout Closed
	EVENT_STATUS_TIMEOUT_CLOSED EventStatus = 10
)
