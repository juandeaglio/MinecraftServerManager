package controls

import (
	"minecraftremote/src/process"
	"minecraftremote/src/rcon"
	"time"
)

type Controls struct {
	serverInBackground process.Process
	started            bool
	rcon               rcon.RCONAdapter
}

// IsStarted implements controls.Controls.
func (m *Controls) IsStarted() bool {
	if m.serverInBackground == nil {
		return false
	}
	return m.serverInBackground.Started()
}

func (m *Controls) Status() *rcon.Status {
	if m.rcon == nil {
		return &rcon.Status{
			Online: m.IsStarted(),
		}
	}
	return m.rcon.GetStatus()
}

func (m *Controls) getPlayers() int {
	return 0
}

func (m *Controls) Start(minecraftServer process.Process) process.Process {
	m.serverInBackground = minecraftServer

	err := minecraftServer.Start()
	if err != nil {
		return nil
	}

	pid := minecraftServer.PID()
	m.started = isPIDValid(pid)

	return minecraftServer
}

func isPIDValid(pid int) bool {
	return pid > 0
}

func (m *Controls) Stop() bool {
	if m.started {
		err := m.serverInBackground.Stop()
		if err != nil {
			return false
		}

		// Give a small window for the process to die
		time.Sleep(100 * time.Millisecond)

		// Verify process is actually stopped
		if m.serverInBackground.Started() {
			return false
		}

		m.started = false
		return true
	}
	return false
}

func NewControls(rcon rcon.RCONAdapter, process ...process.Process) *Controls {
	controls := &Controls{}
	if len(process) > 0 {
		controls.serverInBackground = process[0]
	}
	controls.rcon = rcon
	return controls
}
