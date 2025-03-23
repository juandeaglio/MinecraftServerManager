package mcservercontrols

import (
	"bytes"
	"encoding/json"
	"io"
	server "minecraftremote/src/server"
	"net/http"
)

type MinecraftServer struct {
	started bool
}

// HandleHTTP processes incoming HTTP requests and routes them to appropriate handlers
func (m *MinecraftServer) HandleHTTP(req *http.Request) *http.Response {
	// Define route handlers mapping
	handlers := map[string]func(*http.Request) *http.Response{
		"/start":  m.handleStart,
		"/stop":   m.handleStop,
		"/status": m.handleStatus,
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
func (m *MinecraftServer) handleStart(req *http.Request) *http.Response {
	if m.Start() {
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
func (m *MinecraftServer) handleStop(req *http.Request) *http.Response {
	if m.Stop() {
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
func (m *MinecraftServer) handleStatus(req *http.Request) *http.Response {
	body, err := json.Marshal(m.Status())
	if err != nil {
		return &http.Response{
			StatusCode: 500,
			Status:     "Internal Server Error",
		}
	}

	return &http.Response{
		StatusCode: 200,
		Status:     "OK",
		Body:       io.NopCloser(bytes.NewReader(body)),
	}
}

func (m *MinecraftServer) Status() *server.Status {
	return &server.Status{Players: 1}
}

func (m *MinecraftServer) Start() bool {
	if !m.started {
		m.started = true
		return m.started
	}
	return false
}
func (m *MinecraftServer) Stop() bool {
	if m.started {
		m.started = !m.started
		return !m.started
	}
	return false
}

func NewServer() *MinecraftServer {
	return &MinecraftServer{}
}

var _ server.Server = (*MinecraftServer)(nil)
