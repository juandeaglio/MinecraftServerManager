package controls

import (
	remoteconnection "server/src/remoteconnection"
)

type Server interface {
	Start()
	Stop()
	Restart()
	GetStatus() bool
}

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

func NewServer(inter remoteconnection.RemoteConnection) *MinecraftServer {
	return &MinecraftServer{rcon: inter}
}
