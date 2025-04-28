package cannedrequests

import (
	"net/http"
	"net/url"
)

type ServerRequest struct {
	*http.Request
}

func NewStartRequest() *ServerRequest {
	url, _ := url.Parse("http://localhost/start")
	req := &http.Request{
		Method: "GET",
		URL:    url,
		Header: make(http.Header),
	}
	return &ServerRequest{Request: req}
}

func NewStopRequest() *ServerRequest {
	url, _ := url.Parse("http://localhost/stop")
	req := &http.Request{
		Method: "GET",
		URL:    url,
		Header: make(http.Header),
	}
	return &ServerRequest{Request: req}
}

func NewStatusRequest() *ServerRequest {
	url, _ := url.Parse("http://localhost/status")
	req := &http.Request{
		Method: "GET",
		URL:    url,
		Header: make(http.Header),
	}
	return &ServerRequest{Request: req}
}

func NewRunningRequest() *ServerRequest {
	url, _ := url.Parse("http://localhost/running")
	req := &http.Request{
		Method: "GET",
		URL:    url,
		Header: make(http.Header),
	}
	return &ServerRequest{Request: req}
}

// ToHTTPRequest extracts the underlying *http.Request
func (s *ServerRequest) ToHTTPRequest() *http.Request {
	return s.Request
}
