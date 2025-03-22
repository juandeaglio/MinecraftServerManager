package mcservercontrols

import "net/http"

type Server interface {
	Start() bool
	Stop() bool
	HandleHttp(*http.Request) *http.Response
}
