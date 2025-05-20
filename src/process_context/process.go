package process_context

import (
	"fmt"

	"os/exec"
	"syscall"
)

type Process interface {
	Start() error
	Stop() error
	PID() int
	Started() bool
	GetProcessStatus(pid int) (*ProcessStatus, error)
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

		// Verify the process is actually dead
		if p.isProcessRunning(pid) {
			return fmt.Errorf("process %s (PID: %d) is still running after kill attempt", p.program, pid)
		}

		p.cmd = nil
		return nil
	}
	return nil
}

func (p *ProcessImpl) GetProcessStatus(pid int) (*ProcessStatus, error) {
	return p.osOps.ProcessStatus(pid)
}

func NewProcessInvoker(osOps OsOperations, program string, args ...string) *ProcessImpl {
	return &ProcessImpl{
		program: program,
		args:    args,
		osOps:   osOps,
	}
}
