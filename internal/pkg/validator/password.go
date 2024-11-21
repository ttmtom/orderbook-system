package validator

import (
	"github.com/go-playground/validator"
	"regexp"
)

func passwordValidation(fl validator.FieldLevel) bool {
	password := fl.Field().String()

	hasUppercase := regexp.MustCompile(`[A-Z]`).MatchString(password)
	hasLowercase := regexp.MustCompile(`[a-z]`).MatchString(password)

	hasDigit := regexp.MustCompile(`\d`).MatchString(password)

	specialChars := `~!@#$%^&*()\-\+={}\[\]\\\|;:"<>,./?`
	hasSpecialChar := regexp.MustCompile(`[` + regexp.QuoteMeta(specialChars) + `]`).MatchString(password)

	return hasUppercase && hasLowercase && hasDigit && hasSpecialChar
}
