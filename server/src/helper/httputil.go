package helper

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type HTTPError struct {
	Status int
	Err    error
}

func (err *HTTPError) Error() string {
	if err.Err != nil {
		return fmt.Sprintf("status %d, reason %s", err.Status, err.Err.Error())
	}
	return fmt.Sprintf("Status %d", err.Status)
}

func CreateHTTPError(status int, errorStr string) *HTTPError {
	return &HTTPError{Status: status, Err: errors.New(errorStr)}
}

func JSON(w http.ResponseWriter, code int, data interface{}) error {
	w.WriteHeader(code)
	if data == nil {
		data = map[string]interface{}{}
	}
	return json.NewEncoder(w).Encode(data)
}
