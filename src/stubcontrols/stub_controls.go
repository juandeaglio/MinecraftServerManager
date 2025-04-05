package stubcontrols

import (
	"minecraftremote/src/controls"
	"minecraftremote/src/process"
)

type StubControls struct {
	started bool
}

func (m *StubControls) Start(process process.Process) process.Process {
	m.started = process.Start() == nil
	return process
}

func (m *StubControls) Stop() bool {
	return true
}

func (m *StubControls) Status() *controls.Status {
	return &controls.Status{Players: 1}
}

func (m *StubControls) IsStarted() bool {
	return m.started
}

func NewStubControls() *StubControls {
	return &StubControls{}
}

var _ controls.Controls = (*StubControls)(nil)
