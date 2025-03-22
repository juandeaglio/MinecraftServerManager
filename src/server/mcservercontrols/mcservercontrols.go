package mcservercontrols

import (
	server "minecraftremote/src/server"
	"net/http"
)

type MinecraftServer struct {
}

func (m *MinecraftServer) HandleHttp(req *http.Request) *http.Response {
	if req.URL.Path == "/status" && m.Start() {
		return &http.Response{StatusCode: 200}
	}
	return &http.Response{StatusCode: 400}
}

func (m *MinecraftServer) Start() bool {
	return true

}
func (m *MinecraftServer) Stop() bool {
	return false
}

func NewServer() *MinecraftServer {
	return &MinecraftServer{}
}

var _ server.Server = (*MinecraftServer)(nil)
