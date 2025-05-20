package brokenosoperations

import (
	"errors"
	"minecraftremote/src/process_context"
	"os"
	"os/exec"
	"syscall"
)

type BrokenOsOperations struct{}

func (b *BrokenOsOperations) FindProcess(pid int) (*os.Process, error) {
	return nil, errors.New("failed to find process")
}

func (b *BrokenOsOperations) CreateCommand(program string, args ...string) *exec.Cmd {
	return nil
}

func (b *BrokenOsOperations) SetSysProcAttr(cmd *exec.Cmd) {
}

func (b *BrokenOsOperations) StartCmd(cmd *exec.Cmd) error {
	return errors.New("failed to start command")
}

func (b *BrokenOsOperations) KillProcess(process *os.Process) error {
	return errors.New("failed to kill process")
}

func (b *BrokenOsOperations) Signal(process *os.Process, signal syscall.Signal) error {
	return errors.New("failed to signal process")
}

func (b *BrokenOsOperations) ProcessStatus(pid int) (*process_context.ProcessStatus, error) {
	return &process_context.ProcessStatus{}, nil
}

var _ process_context.OsOperations = (*BrokenOsOperations)(nil)
