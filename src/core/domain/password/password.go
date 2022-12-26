package password

import (
	"Golang/src/core/errors"
	"Golang/src/core/messages"
	"Golang/src/utils/validation"
)

type Password struct {
	password string
	rules    []Rules
}

func (p Password) Password() string {
	return p.password
}

func (p Password) Rules() []Rules {
	return p.rules
}

func NewForSignIn(password string, rule []Rules) (*Password, errors.Error) {
	var invalidFields []errors.InvalidField

	if !validator.IsPasswordValid(password) {
		invalidFields = append(invalidFields, errors.InvalidField{
			Name:        messages.AccountPassword,
			Description: messages.InvalidAccountPasswordMsg,
		})
	}

	return &Password{
		password: password,
		rules:    rule,
	}, nil

}
