package stubcontrols

import "minecraftremote/src/controls"

type StubControls struct {
	started bool
}

func (m *StubControls) Start() bool {
	m.started = true
	return m.started
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
