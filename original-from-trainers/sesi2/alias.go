package main

import (
	"fmt"
	"strings"
)

type ConfigKey string

type CoinName string

const (
	ETH  CoinName = "ETH"
	BTC  CoinName = "BTC"
	DOGE CoinName = "DOGE"
)

const (
	DB_HOST ConfigKey = "localhost"
)

func main() {
	coinReq := "luna"

	fmt.Println(string([]byte{108, 117, 110, 97}))

	isValid := IsValid(CoinName(strings.ToUpper(coinReq)))
	fmt.Println(coinReq, isValid)

	generate(DB_HOST)
}

func generate(DB_HOST ConfigKey) {

}

func IsValid(coin CoinName) bool {
	switch coin {
	case ETH:
		return true
	case BTC:
		return true
	case DOGE:
		return true
	default:
		return false
	}
}
