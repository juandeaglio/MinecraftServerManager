package mcservercontrols

import (
	server "minecraftremote/src/server"
	"net/http"
)

type MinecraftServer struct {
	started bool
}

func (m *MinecraftServer) HandleHttp(req *http.Request) *http.Response {
	if req.URL.Path == "/start" && m.Start() {
		return &http.Response{StatusCode: 200}
	}
	if req.URL.Path == "/stop" && m.Stop() {
		return &http.Response{StatusCode: 200}
	}
	return &http.Response{StatusCode: 400}
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
