package os_api_adapter

import (
	"os"
	"os/exec"
	"syscall"
)

// OsOperations abstracts operating system level operations
type OsOperations interface {
	FindProcess(pid int) (*os.Process, error)
	Signal(process *os.Process, signal syscall.Signal) error
	CreateCommand(program string, args ...string) *exec.Cmd
	StartCmd(cmd *exec.Cmd) error
	KillProcess(process *os.Process) error
	ProcessStatus(pid int) (*ProcessStatus, error)
}

type ProcessStatus struct {
	Status uintptr
}
