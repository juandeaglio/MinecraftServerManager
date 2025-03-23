package cannedrequests

import (
	"net/http"
	"net/url"
)

type StatusRequest struct {
	*http.Request
}

func NewStartRequest() *StatusRequest {
	url, _ := url.Parse("http://localhost/start")
	req := &http.Request{
		Method: "GET",
		URL:    url,
		Header: make(http.Header),
	}
	return &StatusRequest{Request: req}
}

func NewStopRequest() *StatusRequest {
	url, _ := url.Parse("http://localhost/stop")
	req := &http.Request{
		Method: "GET",
		URL:    url,
		Header: make(http.Header),
	}
	return &StatusRequest{Request: req}
}

// ToHTTPRequest extracts the underlying *http.Request
func (s *StatusRequest) ToHTTPRequest() *http.Request {
	return s.Request
}
