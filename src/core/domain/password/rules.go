package password

type Rules struct {
	rule  string
	value int
}

func NewRuleReference(rule string, value int) *Rules {
	return &Rules{
		rule:  rule,
		value: value,
	}
}
