package validator

import (
	"github.com/go-playground/validator"
)

func New() *validator.Validate {
	v := validator.New()

	err := v.RegisterValidation("validPWD", passwordValidation)
	if err != nil {
		return nil
	}
	v.RegisterAlias("cPassword", "validPWD")

	return v
}
