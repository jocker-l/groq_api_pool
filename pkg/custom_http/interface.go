package custom_http

import (
	"io"
	"net/http"
)

type StreamHandlerInter interface {
	StreamHandler()
}

type HTTPClient interface {
	Request(method string, url string, headers Headers, cookies []*http.Cookie, body io.Reader) (*http.Response, error)
	SetProxy(proxy string)
}

type Headers map[string]string

func (h Headers) Set(key string, value string) {
	h[key] = value
}

func NewHeader() Headers {
	return make(Headers)
}
