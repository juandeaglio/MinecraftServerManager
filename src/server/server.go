package mcservercontrols

import "net/http"

type Server interface {
	Start() bool
	Stop() bool
	HandleHTTP(*http.Request) *http.Response
	Status() *Status
}

type Status struct {
	Players int
}
