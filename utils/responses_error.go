package utils

// omitempty means that the field `fields` returns null or empty dont show in the JSON
type ErrorResponse struct {
	Message string          `json:"message"`
	Fields  []FieldsMessage `json:"fields,omitempty"`
}

type FieldsMessage struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}
