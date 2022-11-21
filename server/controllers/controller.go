package controllers

import "net/http"

type Controller struct {
}

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// InitResponse
func (ctrl Controller) InitResponse() Response {
	return Response{
		Code:    http.StatusOK,
		Message: "",
		Data:    nil,
	}
}
