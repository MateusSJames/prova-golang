package response

type ErrorMessage struct {
	StatusCode       int            `json:"status_code"`
	Msg              string         `json:"message"`
	InvalidFields    []InvalidField `json:"invalid_fields,omitempty"`
}

type InvalidField struct {
	FieldName   string `json:"field_name"`
	Description string `json:"description"`
}
