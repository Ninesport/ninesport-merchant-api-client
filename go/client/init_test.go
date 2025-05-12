package client_test

import "github.com/brianvoe/gofakeit/v6"

var randomAccount string

func init() {
	randomAccount = gofakeit.AchAccount()
}
