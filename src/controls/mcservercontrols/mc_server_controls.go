package mcservercontrols

import (
	background "minecraftremote/src/backgroundserver"
	server "minecraftremote/src/controls"
)

type MinecraftServer struct {
	serverInBackground background.BackgroundServer
	started            bool
}

// IsStarted implements controls.Controls.
func (m *MinecraftServer) IsStarted() bool {
	return m.started
}

func (m *MinecraftServer) Status() *server.Status {
	return &server.Status{Players: m.getPlayers()}
}

func (m *MinecraftServer) getPlayers() int {
	return 0
}

func (m *MinecraftServer) Start() bool {
	if !m.started {
		m.started = true
		return m.started
	}
	return m.started
}
func (m *MinecraftServer) Stop() bool {
	if m.started {
		m.started = !m.started
		return !m.started
	}
	return false
}

func NewControls(minecraftServer background.BackgroundServer) *MinecraftServer {
	return &MinecraftServer{serverInBackground: minecraftServer}
}

var _ server.Controls = (*MinecraftServer)(nil)
