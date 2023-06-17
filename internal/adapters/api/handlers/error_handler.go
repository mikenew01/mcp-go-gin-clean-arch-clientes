package handlers

import (
	"encoding/json"
	"net/http"
)

type ErrorResponseHandler struct {
	statusCode int
	Errors     []string `json:"errors"`
}

func NewErrorResponseHandler(err error, status int) *ErrorResponseHandler {
	return &ErrorResponseHandler{
		Errors:     []string{err.Error()},
		statusCode: status,
	}
}

func (errResponseHandler ErrorResponseHandler) Send(response http.ResponseWriter) error {
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(errResponseHandler.statusCode)
	return json.NewEncoder(response).Encode(errResponseHandler)
}
