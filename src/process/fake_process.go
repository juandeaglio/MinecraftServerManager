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
	cmd := f.osOps.CreateCommand("fake", "args")
	f.osOps.SetSysProcAttr(cmd)

	if err := f.osOps.StartCmd(cmd); err != nil {
		return err
	}

	f.started = true
	f.pid = cmd.Process.Pid
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
	cmd := exec.Command(program, args...)
	cmd.Process = &os.Process{Pid: 12345} // Fake PID
	return cmd
}

func (f *FakeOsOperations) SetSysProcAttr(cmd *exec.Cmd) {
}

func (f *FakeOsOperations) StartCmd(cmd *exec.Cmd) error {
	return nil
}

func (f *FakeOsOperations) KillProcess(process *os.Process) error {
	return nil
}
