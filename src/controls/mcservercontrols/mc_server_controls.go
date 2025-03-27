package mcservercontrols

import (
	server "minecraftremote/src/controls"
	"minecraftremote/src/process"
)

type MinecraftServer struct {
	serverInBackground process.Process
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

func (m *MinecraftServer) Start(minecraftServer process.Process) process.Process {
	m.serverInBackground = minecraftServer
	m.serverInBackground.Start()
	if m.serverInBackground.PID() >= 0 {
		m.started = true
	}

	return minecraftServer
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
