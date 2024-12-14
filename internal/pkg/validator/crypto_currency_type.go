package validator

import (
	"github.com/go-playground/validator"
	"orderbook/internal/core/model"
)

func isValidCryptoCurrency(fl validator.FieldLevel) bool {
	currency := fl.Field().Interface().(model.CryptoCurrency)
	switch currency {
	case model.BTC, model.ETH:
		return true
	default:
		return false
	}
}
