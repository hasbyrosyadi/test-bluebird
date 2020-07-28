package api

import (
	"encoding/json"
	"goldminer/model"
	"net/http"
)

func Success(data interface{}) model.Success {
	e := model.Success{
		HTTPStatusCode: http.StatusOK,
		Success:        true,
		StateCode:      "SUCCESS",
		Message:        "Successfully executed",
		Data:           data,
	}
	return e
}

func ErrorMessage(err error) model.Error {
	e := model.Error{
		HTTPStatusCode: http.StatusInternalServerError,
		Success:        false,
		StateCode:      "Internal Server Error",
		Message:        err.Error(),
	}
	return e
}

func ErrorClient(err error) model.Error {
	e := model.Error{
		HTTPStatusCode: http.StatusBadRequest,
		Success:        false,
		StateCode:      "Bad Request",
		Message:        err.Error(),
	}
	return e
}

func HttpResponseJson(w http.ResponseWriter, responseBody interface{}, httpStatus int) {
	respBody, _ := json.Marshal(responseBody)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatus)
	w.Write(respBody)
}
