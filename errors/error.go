package errors

import (
	"encoding/json"
	"io"
	"log"
)

var Errors = map[int8]string{
	1: "Database query error",
	2: "Error adding data to database",
	3: "Database update error",
	4: "Error deleting database data",
	5: "Parse JSON error",
	6: "Validation error",
}

type typeErrors struct {
	Code    int8   `json:"code"`
	Name    string `json:"name"`
	Message string `json:"errors"`
}

func WriteErrorInResponse(w io.Writer, code int8, name, message string) {
	errMsg := typeErrors{
		Code:    code,
		Name:    name,
		Message: message,
	}
	if err := json.NewEncoder(w).Encode(errMsg); err != nil {
		log.Printf(err.Error())
	}
}
