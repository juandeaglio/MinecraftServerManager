package stubcontrols

import "minecraftremote/src/controls"

type StubControls struct {
}

func (m *StubControls) Start() bool {
	return true
}
func (m *StubControls) Stop() bool {
	return true
}

func (m *StubControls) Status() *controls.Status {
	return &controls.Status{Players: 1}
}

func NewStubControls() *StubControls {
	return &StubControls{}
}

var _ controls.Controls = (*StubControls)(nil)
