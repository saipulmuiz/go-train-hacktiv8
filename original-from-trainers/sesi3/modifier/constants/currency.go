package constants

type CurrencyName string

const (
	IDR CurrencyName = "IDR"
	USD CurrencyName = "USD"
)

func (c CurrencyName) IsValid() bool {
	switch c {
	case IDR:
		return true
	case USD:
		return true
	default:
		return false
	}
}
