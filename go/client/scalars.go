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
	CURRENCY_TYPE_BNB  CurrencyType = "BNB"
	CURRENCY_TYPE_BTC  CurrencyType = "BTC"
	CURRENCY_TYPE_ETH  CurrencyType = "ETH"
	CURRENCY_TYPE_JPY  CurrencyType = "JPY"
	CURRENCY_TYPE_TRX  CurrencyType = "TRX"
	CURRENCY_TYPE_TWD  CurrencyType = "TWD"
	CURRENCY_TYPE_USD  CurrencyType = "USD"
	CURRENCY_TYPE_USDC CurrencyType = "USDC"
	CURRENCY_TYPE_USDT CurrencyType = "USDT"
)

var CurrencyTypes []CurrencyType = []CurrencyType{
	CURRENCY_TYPE_BNB,
	CURRENCY_TYPE_BTC,
	CURRENCY_TYPE_ETH,
	CURRENCY_TYPE_JPY,
	CURRENCY_TYPE_TRX,
	CURRENCY_TYPE_TWD,
	CURRENCY_TYPE_USD,
	CURRENCY_TYPE_USDC,
	CURRENCY_TYPE_USDT,
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
	TRANSFER_STATUS_INSUFFICIENT_BALANCE
	TRANSFER_STATUS_INSUFFICIENT_APPLY_AMOUNT
	TRANSFER_STATUS_FAILED
	TRANSFER_STATUS_SUCCESS
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
