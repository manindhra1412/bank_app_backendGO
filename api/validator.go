package api

import (
	"github.com/go-playground/validator/v10"
	"github.com/manindhra1412/simple_bank/util"
)

var validCurrency validator.Func = func(fl validator.FieldLevel) bool {
	if currency, ok := fl.Field().Interface().(string); ok {
		// check the currency is supported or not
		return util.IsSupportedCurrency(currency)
	}
	return false
}
