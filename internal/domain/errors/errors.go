package errors

import (
	"encoding/json"
	"net/http"
)

type APIError struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func NewAPIError(message string, code int) *APIError {
	return &APIError{
		Message: message,
		Code:    code,
	}
}

func (e *APIError) Error() string {
	return e.Message
}

func WriteErrorResponse(w http.ResponseWriter, apiErr *APIError) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(apiErr.Code)
	json.NewEncoder(w).Encode(apiErr)
}
