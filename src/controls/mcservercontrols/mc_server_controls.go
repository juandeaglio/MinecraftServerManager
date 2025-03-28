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
	if m.serverInBackground == nil {
		return false
	}
	return m.serverInBackground.Started()
}

func (m *MinecraftServer) Status() *server.Status {
	return &server.Status{
		Players: m.getPlayers(),
		Online:  m.IsStarted(),
	}
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

// NewControls creates a new MinecraftServer instance.
// If a process is provided, it will be used as the server process.
func NewControls(process ...process.Process) *MinecraftServer {
	server := &MinecraftServer{}
	if len(process) > 0 {
		server.serverInBackground = process[0]
	}
	return server
}

var _ server.Controls = (*MinecraftServer)(nil)
