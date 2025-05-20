package process_context

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

type FakeOsOperations struct {
	pid int
}

func (f *FakeOsOperations) FindProcess(pid int) (*os.Process, error) {
	if f.pid <= 0 {
		return nil, fmt.Errorf("process not found")
	}
	return &os.Process{Pid: f.pid}, nil
}

func (f *FakeOsOperations) Signal(process *os.Process, signal syscall.Signal) error {
	return nil
}

func (f *FakeOsOperations) CreateCommand(program string, args ...string) *exec.Cmd {
	cmd := exec.Command(program, args...)
	f.pid = 12345
	cmd.Process = &os.Process{Pid: f.pid} // Fake PID
	return cmd
}

func (f *FakeOsOperations) SetSysProcAttr(cmd *exec.Cmd) {
}

func (f *FakeOsOperations) StartCmd(cmd *exec.Cmd) error {
	return nil
}

func (f *FakeOsOperations) KillProcess(process *os.Process) error {
	process.Pid = 0
	f.pid = 0
	return nil
}

func (f *FakeOsOperations) ProcessStatus(pid int) (ProcessStatus, error) {
	return ProcessStatus{}, nil
}

var _ OsOperations = (*FakeOsOperations)(nil)
