package controllers

import "net/http"

type Response interface {
	Header() http.Header
	Write([]byte) (int, error)
	Flush()
}

const (
	charsetUTF8 = "charser=UTF-8"
	HeaderContentType = "Content-Type"
	MIMEApplicationJSON = "application/json"
	MIMEApplicationJSONCharsetUTF8 = MIMEApplicationJSON + "; " + charsetUTF8
)
