package main

import (
	"fmt"
	"strings"
)

const ETH string = "ETH"
const BTC string = "BTC"

func main() {
	var currentYear = 2022

	if age := currentYear - 1999; age > 17 {
		fmt.Println("Sudah punya KTP", age)
	} else if age > 25 {
		fmt.Println("Sudah menikah", age)
	} else {
		fmt.Println("Blm punya KTP dan belum menikah")
	}

	score := 9.2
	scoreAlpha := "F"

	coin := "doge"

	reqCoin := ""
	switch strings.ToUpper(coin) {
	case ETH:
		reqCoin = ETH
	case BTC:
		reqCoin = BTC
	}

	fmt.Println(reqCoin)

	// switch score {
	// case 9: // score == 9
	// 	scoreAlpha = "A"
	// case 8: // score == 8
	// 	scoreAlpha = "AB"
	// case 7:
	// 	scoreAlpha = "B"
	// default:
	// 	scoreAlpha = "F"
	// }

	switch {
	case score >= 9:
		scoreAlpha += "A"
		fallthrough
	case score >= 8 && score < 9:
		scoreAlpha += "AB"
	case score >= 7 && score < 8:
		scoreAlpha = "B"
	}

	fmt.Println(scoreAlpha)
}
