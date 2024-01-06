package utils

import (
	"encoding/json"
	"net/http"
)

type APIError struct {
	StatusCode int `json:code`
	Message string `json:message`
}

func (apiErr APIError) FillAndStringify(code int, message string) []byte {
	apiErr.StatusCode = code
	apiErr.Message = message
	err, _ := json.Marshal(apiErr);
	return err;
}

func FillErrorResponse(w *http.ResponseWriter, code int, message string) {
	(*w).Header().Set("Content-Type", "application/json")
	(*w).WriteHeader(code)
	var apiErr APIError;
	apiErr.StatusCode = code
	apiErr.Message = message
	err, _ := json.Marshal(apiErr)
	(*w).Write(err)
}