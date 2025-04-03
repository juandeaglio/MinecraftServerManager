package processtest

import (
	"errors"
	"os"
	"os/exec"
	"syscall"
)

type brokenOsOperations struct{}

func (b *brokenOsOperations) FindProcess(pid int) (*os.Process, error) {
	return nil, errors.New("failed to find process")
}

func (b *brokenOsOperations) CreateCommand(program string, args ...string) *exec.Cmd {
	return nil
}

func (b *brokenOsOperations) SetSysProcAttr(cmd *exec.Cmd) {
}

func (b *brokenOsOperations) StartCmd(cmd *exec.Cmd) error {
	return errors.New("failed to start command")
}

func (b *brokenOsOperations) KillProcess(process *os.Process) error {
	return errors.New("failed to kill process")
}

func (b *brokenOsOperations) Signal(process *os.Process, signal syscall.Signal) error {
	return errors.New("failed to signal process")
}
