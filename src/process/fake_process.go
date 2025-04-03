package process

import (
	"os"
	"os/exec"
	"syscall"
)

type FakeProcess struct {
	started bool
	pid     int
	osOps   OsOperations
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
	if !f.started {
		return nil
	}

	proc, err := f.osOps.FindProcess(f.pid)
	if err != nil {
		return err
	}

	if err := f.osOps.KillProcess(proc); err != nil {
		return err
	}

	f.started = false
	return nil
}

// NewFakeProcess creates a new FakeProcess with default OsOperations
func NewFakeProcess(osOps OsOperations) *FakeProcess {

	return &FakeProcess{
		osOps:   osOps,
		started: false,
		pid:     -1,
	}
}

var _ Process = (*FakeProcess)(nil)

type FakeOsOperations struct{}

func (f *FakeOsOperations) FindProcess(pid int) (*os.Process, error) {
	return nil, nil
}

func (f *FakeOsOperations) Signal(process *os.Process, signal syscall.Signal) error {
	return nil
}

func (f *FakeOsOperations) CreateCommand(program string, args ...string) *exec.Cmd {
	return nil
}

func (f *FakeOsOperations) SetSysProcAttr(cmd *exec.Cmd) {
}

func (f *FakeOsOperations) StartCmd(cmd *exec.Cmd) error {
	return nil
}

func (f *FakeOsOperations) KillProcess(process *os.Process) error {
	return nil
}
