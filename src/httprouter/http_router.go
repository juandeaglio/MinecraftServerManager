package httprouter

import (
	"bytes"
	"encoding/json"
	"io"
	"minecraftremote/src/controls"
	"minecraftremote/src/process"
	"net/http"
)

type HTTPRouter interface {
	handleStart(*http.Request) *http.Response
	handleStop(*http.Request) *http.Response
	handleStatus(*http.Request) *http.Response
	HandleHTTP(*http.Request) *http.Response
}

type ServerRouter struct {
	proc    process.Process
	handler *controls.Controls
}

func NewHTTPRouter(controls *controls.Controls, proc process.Process) *ServerRouter {
	return &ServerRouter{handler: controls, proc: proc}
}

// HandleHTTP processes incoming HTTP requests and routes them to appropriate handlers
func (h *ServerRouter) HandleHTTP(req *http.Request) *http.Response {
	// Define route handlers mapping
	handlers := map[string]func(*http.Request) *http.Response{
		"/start":   h.handleStart,
		"/stop":    h.handleStop,
		"/status":  h.handleStatus,
		"/running": h.handleRunning,
	}

	// Look up the appropriate handler for the requested path
	if handler, exists := handlers[req.URL.Path]; exists {
		return handler(req)
	}

	// Return 404 Not Found for undefined routes
	return &http.Response{
		StatusCode: 501,
		Status:     "Not Implemented",
	}
}

func (h *ServerRouter) handleStart(req *http.Request) *http.Response {
	process := h.handler.Start(h.proc)
	if process == nil {
		return &http.Response{
			StatusCode: 500,
			Status:     "Internal Server Error",
		}
	}
	if process.PID() != -1 {
		return &http.Response{
			StatusCode: 200,
			Status:     "OK",
		}
	}
	return &http.Response{
		StatusCode: 500,
		Status:     "Internal Server Error",
	}
}

func (h *ServerRouter) handleStop(req *http.Request) *http.Response {
	if h.handler.Stop() {
		return &http.Response{
			StatusCode: 200,
			Status:     "OK",
		}
	}
	return &http.Response{
		StatusCode: 500,
		Status:     "Internal Server Error",
	}
}

func (h *ServerRouter) handleStatus(req *http.Request) *http.Response {
	if h.handler.Status() == nil || h.proc == nil {
		return &http.Response{
			StatusCode: 404,
			Status:     "Server is not running",
		}
	}
	// do we need to validate this json contract?
	// TODO: review comment and remove if not needed
	body, err := json.Marshal(h.handler.Status())
	if err != nil {
		return &http.Response{
			StatusCode: 500,
			Status:     "Internal Server Error",
		}
	}

	resp := &http.Response{
		StatusCode: 200,
		Status:     "OK",
		Header:     http.Header{},
		Body:       io.NopCloser(bytes.NewReader(body)),
	}
	resp.Header.Set("Content-Type", "application/json")
	return resp
}

func (h *ServerRouter) handleRunning(req *http.Request) *http.Response {
	if h.proc == nil {
		return &http.Response{
			StatusCode: 404,
			Status:     "Server is not running",
		}
	}
	return &http.Response{
		StatusCode: 200,
		Status:     "OK",
	}
}

var _ HTTPRouter = (*ServerRouter)(nil)
