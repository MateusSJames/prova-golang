package request

import (
	"Golang/src/core/domain/password"
	"Golang/src/core/errors"
)

type Password struct {
	Password string  `json:"password"`
	Rules    []Rules `json:"rules"`
}

func (dto Password) AsDomain() (*password.Password, errors.Error) {

	var rules []password.Rules
	for _, rule := range dto.Rules {
		rules = append(rules, *password.NewRuleReference(rule.Rule, rule.Value))
	}

	return password.NewForSignIn(dto.Password, rules)
}
