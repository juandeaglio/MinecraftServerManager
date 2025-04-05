package process

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

type Process interface {
	Start() error
	Stop() error
	PID() int
	Started() bool
}

type ProcessImpl struct {
	cmd     *exec.Cmd
	program string
	args    []string
	osOps   OsOperations
}

func (p *ProcessImpl) Started() bool {
	if p.cmd == nil || p.cmd.Process == nil {
		return false
	}
	return p.isProcessRunning(p.cmd.Process.Pid)
}

func (p *ProcessImpl) isProcessRunning(pid int) bool {
	process, err := p.osOps.FindProcess(pid)
	if err != nil {
		return false
	}
	err = p.osOps.Signal(process, syscall.Signal(0))
	return err == nil
}

func (p *ProcessImpl) PID() int {
	if p.cmd != nil && p.cmd.Process != nil {
		return p.cmd.Process.Pid
	}
	return -1
}

func (p *ProcessImpl) Start() error {
	if p.program == "" {
		return fmt.Errorf("program is empty")
	}
	p.cmd = p.osOps.CreateCommand(p.program, p.args...)
	p.osOps.SetSysProcAttr(p.cmd)

	err := p.osOps.StartCmd(p.cmd)
	if err != nil {
		return fmt.Errorf("failed to start process %s: %w", p.program, err)
	}
	return nil
}

func (p *ProcessImpl) Stop() error {
	if p.cmd != nil && p.cmd.Process != nil {
		pid := p.cmd.Process.Pid
		err := p.osOps.KillProcess(p.cmd.Process)
		if err != nil {
			return fmt.Errorf("failed to kill process %s (PID: %d): %w", p.program, pid, err)
		}
		p.cmd = nil
		return nil
	}
	return nil
}

func NewProcess(osOps OsOperations, program string, args ...string) *ProcessImpl {
	return &ProcessImpl{
		program: program,
		args:    args,
		osOps:   osOps,
	}
}

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
