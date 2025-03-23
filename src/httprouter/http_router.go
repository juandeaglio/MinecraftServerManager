package httprouter

import (
	"bytes"
	"encoding/json"
	"io"
	"minecraftremote/src/controls"
	"net/http"
)

type HTTPRouter interface {
	handleStart(*http.Request) *http.Response
	handleStop(*http.Request) *http.Response
	handleStatus(*http.Request) *http.Response
	HandleHTTP(*http.Request) *http.Response
}

type HTTPServer struct {
	handler controls.Controls
}

func NewHTTPServer(controls controls.Controls) *HTTPServer {
	return &HTTPServer{handler: controls}
}

// HandleHTTP processes incoming HTTP requests and routes them to appropriate handlers
func (h *HTTPServer) HandleHTTP(req *http.Request) *http.Response {
	// Define route handlers mapping
	handlers := map[string]func(*http.Request) *http.Response{
		"/start":  h.handleStart,
		"/stop":   h.handleStop,
		"/status": h.handleStatus,
	}

	// Look up the appropriate handler for the requested path
	if handler, exists := handlers[req.URL.Path]; exists {
		return handler(req)
	}

	// Return 404 Not Found for undefined routes
	return &http.Response{
		StatusCode: 404,
		Status:     "Not Found",
	}
}

// handleStart handles server start requests
func (h *HTTPServer) handleStart(req *http.Request) *http.Response {
	if h.handler.Start() {
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

// handleStop handles server stop requests
func (h *HTTPServer) handleStop(req *http.Request) *http.Response {
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

// handleStatus handles server status requests
func (h *HTTPServer) handleStatus(req *http.Request) *http.Response {
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

var _ HTTPRouter = (*HTTPServer)(nil)
