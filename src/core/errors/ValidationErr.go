package errors

import "reflect"

var _ Error = (*ValidationErr)(nil)

type ValidationErr struct {
	*superErr
	msg           string
	invalidFields []InvalidField
}

type InvalidField struct {
	Name, Description string
}

func (e *ValidationErr) Error() string {
	return e.msg
}

func (e *ValidationErr) FriendlyMsg() string {
	return e.msg
}

func (e *ValidationErr) InvalidFields() []InvalidField {
	return e.invalidFields
}

func (e *ValidationErr) Equals(err Error) bool {
	errToCompare, ok := err.(*ValidationErr)

	return ok &&
		e.msg == errToCompare.msg &&
		reflect.DeepEqual(e.invalidFields, errToCompare.invalidFields)
}

func NewValidationErr(msg string, invalidFields ...InvalidField) *ValidationErr {
	return &ValidationErr{
		superErr:      newSuperErr(),
		msg:           msg,
		invalidFields: invalidFields,
	}
}