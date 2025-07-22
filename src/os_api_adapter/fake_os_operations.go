package os_api_adapter

import (
	"fmt"
	"minecraftremote/src/windowsconstants"
	"os"
	"os/exec"
	"syscall"
)

type FakeOsOperations struct {
	pid    int
	killed bool
}

func (f *FakeOsOperations) FindProcess(pid int) (*os.Process, error) {
	if f.pid <= 0 || f.killed {
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
	f.killed = true
	return nil
}

func (f *FakeOsOperations) ProcessStatus(pid int) (*ProcessStatus, error) {
	if f.pid <= 0 || f.killed {
		return &ProcessStatus{Status: windowsconstants.ParentKilledChildStatus}, fmt.Errorf("process not found")
	}
	return &ProcessStatus{Status: windowsconstants.RunningStatus}, nil
}

var _ OsOperations = (*FakeOsOperations)(nil)
