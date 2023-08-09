package Responses

import (
	"Backend/Core/Utilities/Model"
	"net/http"
)

func InvokeSuccess(w http.ResponseWriter, data interface{}) {
	item := generateSuccessItem(data)
	w.WriteHeader(http.StatusOK)
	w.Write(item.JsonEncode())
}

func generateSuccessItem(data interface{}) ResponseModel.BaseResponseModel {
	item := ResponseModel.BaseResponseModel{
		Success:    true,
		StatusCode: http.StatusOK,
		Data:       data,
	}
	return item
}
