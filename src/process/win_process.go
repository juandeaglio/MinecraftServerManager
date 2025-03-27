package process

type WinProcess struct {
	started bool
	pid     int
}

// PID implements Process.
func (f *WinProcess) PID() int {
	return -1 // Simulate no process running
}

// Start implements Process.
func (f *WinProcess) Start() error {
	panic("Unimplemented")
}

// Stop implements Process.
func (f *WinProcess) Stop() error {
	panic("Unimplemented")
}

var _ Process = (*FakeProcess)(nil)
