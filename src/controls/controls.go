package controls

import (
	"minecraftremote/src/os_api_adapter"
	"minecraftremote/src/rcon"
	"minecraftremote/src/windowsconstants"
)

type Controls struct {
	processHandler os_api_adapter.Process
	started        bool
	rcon           rcon.RCONAdapter
}

// IsStarted implements controls.Controls.
func (m *Controls) IsStarted() bool {
	if m.processHandler == nil {
		return false
	}
	return m.processHandler.Started()
}

func (m *Controls) Status() *rcon.Status {
	if m.rcon == nil {
		return &rcon.Status{
			Online: m.IsStarted(),
		}
	}
	return m.rcon.GetStatus()
}

func (m *Controls) Start(serverProcess os_api_adapter.Process) os_api_adapter.Process {
	m.processHandler = serverProcess

	err := serverProcess.Start()
	if err != nil {
		return nil
	}

	pid := serverProcess.PID()
	ps, err := serverProcess.GetProcessStatus(pid)

	if err != nil {
		return nil
	}
	m.started = ps.Status == windowsconstants.RunningStatus
	return serverProcess
}

func (m *Controls) Stop() bool {
	if m.started {
		err := m.processHandler.Stop()
		if err != nil {
			return false
		}

		// Verify process is actually stopped
		if m.processHandler.Started() {
			return false
		}

		m.started = false
		return true
	}
	return false
}

func NewControls(rcon rcon.RCONAdapter, process ...os_api_adapter.Process) *Controls {
	controls := &Controls{}
	if len(process) > 0 {
		controls.processHandler = process[0]
	}
	controls.rcon = rcon
	return controls
}
