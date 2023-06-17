package handlers

import (
	"encoding/json"
	"net/http"
)

type SuccessResponseHandler struct {
	statusCode int
	result     interface{}
}

func NewSuccessResponseHandler(result interface{}, statusCode int) *SuccessResponseHandler {
	return &SuccessResponseHandler{statusCode: statusCode, result: result}
}

func (successResponse *SuccessResponseHandler) Send(response http.ResponseWriter) error {
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(successResponse.statusCode)
	return json.NewEncoder(response).Encode(successResponse.result)
}
