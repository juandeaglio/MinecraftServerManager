package mcservercontrols

import (
	remoteconnection "minecraftremote/src/remoteconnection"
	server "minecraftremote/src/server"
)

type MinecraftServer struct {
	rcon remoteconnection.RemoteConnection
}

func (m *MinecraftServer) IsAvailable() bool {
	return m.rcon.IsAvailable()
}

func (m *MinecraftServer) GetStatus() remoteconnection.StatusResponse {
	return m.rcon.PollServer()
}

func (m *MinecraftServer) Start() {

}
func (m *MinecraftServer) Stop() {

}
func (m *MinecraftServer) Restart() {

}

func NewServer(conn remoteconnection.RemoteConnection) *MinecraftServer {
	return &MinecraftServer{rcon: conn}
}

var _ server.Server = (*MinecraftServer)(nil)
