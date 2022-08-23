package errors

import "fmt"

type ValidationError struct {
	Fields       string `json:"fields"`
	MessageError string `json:"message"`
}

func PrepareValidationErrorToString(validError []ValidationError) (strError string) {
	for _, validationError := range validError {
		strError += fmt.Sprintf("%s ", validationError.MessageError)
	}
	return
}
