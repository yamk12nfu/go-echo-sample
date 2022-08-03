package controllers

import "net/http"

type Context interface {
	Bind(interface{}) error
	JSON(int, interface{}) error
	Param(string) string
	Request() *http.Request
	CustomResponse() Response
}
