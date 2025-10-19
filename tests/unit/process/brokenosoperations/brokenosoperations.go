package brokenosoperations

import (
	"errors"
	"minecraftremote/src/os_api_adapter"
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

func (b *BrokenOsOperations) StartCmd(cmd *exec.Cmd) error {
	return errors.New("failed to start command")
}

func (b *BrokenOsOperations) KillProcess(process *os.Process) error {
	return errors.New("failed to kill process")
}

func (b *BrokenOsOperations) Signal(process *os.Process, signal syscall.Signal) error {
	return errors.New("failed to signal process")
}

func (b *BrokenOsOperations) ProcessStatus(pid int) (*os_api_adapter.ProcessStatus, error) {
	return &os_api_adapter.ProcessStatus{}, nil
}

var _ os_api_adapter.OsOperations = (*BrokenOsOperations)(nil)
