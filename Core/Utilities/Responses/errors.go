package Responses

import (
	"Backend/Core/Utilities/Model"
	"net/http"
)

func InvokeBadRequest(w http.ResponseWriter, message string) {
	item := generateErrorItem(message, http.StatusBadRequest)
	w.WriteHeader(http.StatusBadRequest)
	w.Write(item.JsonEncode())
}

func InvokeInternalServerError(w http.ResponseWriter, message string) {
	item := generateErrorItem(message, http.StatusInternalServerError)
	w.WriteHeader(http.StatusInternalServerError)
	w.Write(item.JsonEncode())
}

func InvokeMethodNotAllowed(w http.ResponseWriter, message string) {
	item := generateErrorItem(message, http.StatusMethodNotAllowed)
	w.WriteHeader(http.StatusMethodNotAllowed)
	w.Write(item.JsonEncode())
}

func InvokeUnAuthorized(w http.ResponseWriter, message string) {
	item := generateErrorItem(message, http.StatusUnauthorized)
	w.WriteHeader(item.StatusCode)
	w.Write(item.JsonEncode())
}

func generateErrorItem(message string, statusCode int) *ResponseModel.BaseErrorModel {
	return &ResponseModel.BaseErrorModel{
		Success:    false,
		StatusCode: statusCode,
		Message:    message,
	}

}
