package controls

import (
	"minecraftremote/src/process_context"
	"minecraftremote/src/rcon"
	"minecraftremote/src/windowsconstants"
	"time"
)

type Controls struct {
	processInvoker process_context.Process
	started        bool
	rcon           rcon.RCONAdapter
}

// IsStarted implements controls.Controls.
func (m *Controls) IsStarted() bool {
	if m.processInvoker == nil {
		return false
	}
	return m.processInvoker.Started()
}

func (m *Controls) Status() *rcon.Status {
	if m.rcon == nil {
		return &rcon.Status{
			Online: m.IsStarted(),
		}
	}
	return m.rcon.GetStatus()
}

func (m *Controls) Start(serverProcess process_context.Process) process_context.Process {
	m.processInvoker = serverProcess

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

func isPIDValid(pid int) bool {
	return pid > 0
}

func (m *Controls) Stop() bool {
	if m.started {
		err := m.processInvoker.Stop()
		if err != nil {
			return false
		}

		// Give a small window for the process to die
		time.Sleep(100 * time.Millisecond)

		// Verify process is actually stopped
		if m.processInvoker.Started() {
			return false
		}

		m.started = false
		return true
	}
	return false
}

func NewControls(rcon rcon.RCONAdapter, process ...process_context.Process) *Controls {
	controls := &Controls{}
	if len(process) > 0 {
		controls.processInvoker = process[0]
	}
	controls.rcon = rcon
	return controls
}
