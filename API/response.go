package api

import (
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
