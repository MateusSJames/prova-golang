package validator

import (
	"unicode"
)

func IsPasswordValid(password string) bool {
	var (
		hasMinLen  = false
		hasMaxLen  = false
		hasUpper   = false
		hasLower   = false
		hasNumber  = false
		hasSpecial = false
	)
	
	if len(password) >= 8 {
		hasMinLen = true
	}

	if len(password) <= 30 {
		hasMaxLen = true
	}

	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsNumber(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}

	return hasMinLen && hasMaxLen && hasUpper && hasLower && hasNumber && hasSpecial
}
