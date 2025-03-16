package interfaces

import (
	"net/http"
	"net/url"
)

type Context struct {
	Request     http.Request
	Response    http.ResponseWriter
	QueryParams url.Values
	GetParam    func(key string) string
	Set         func(key string, value interface{})
}
