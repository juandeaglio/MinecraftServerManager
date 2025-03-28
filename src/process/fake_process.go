package process

type FakeProcess struct {
	started bool
	pid     int
}

// Started implements Process.
func (f *FakeProcess) Started() bool {
	return f.started
}

// PID implements Process.
func (f *FakeProcess) PID() int {
	if f.started {
		return f.pid
	}
	return -1 // Simulate no process running
}

// Start implements Process.
func (f *FakeProcess) Start() error {
	f.started = true
	f.pid = 1234 // simulate a PID
	return nil
}

// Stop implements Process.
func (f *FakeProcess) Stop() error {
	f.started = false
	return nil
}

var _ Process = (*FakeProcess)(nil)
