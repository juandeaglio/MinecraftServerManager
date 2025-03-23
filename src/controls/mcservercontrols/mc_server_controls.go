package mcservercontrols

import (
	server "minecraftremote/src/controls"
)

type MinecraftServer struct {
	started bool
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

func NewControls() *MinecraftServer {
	return &MinecraftServer{}
}

var _ server.Controls = (*MinecraftServer)(nil)
