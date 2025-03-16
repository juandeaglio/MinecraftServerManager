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

func (m *MinecraftServer) IsAvailable(num int) bool {
	return m.rcon.IsAvailable(num)
}

func (m *MinecraftServer) GetStatus() bool {
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
