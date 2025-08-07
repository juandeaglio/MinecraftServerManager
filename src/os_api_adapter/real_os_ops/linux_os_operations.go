//go:build linux

package real_os_ops

import (
	"minecraftremote/src/os_api_adapter"
	"os"
	"os/exec"
	"syscall"
)

// RealOsOperations implements OsOperations for Windows
type RealOsOperations struct{}

func (l RealOsOperations) FindProcess(pid int) (*os.Process, error) {
	return os.FindProcess(pid)
}

func (l RealOsOperations) Signal(process *os.Process, signal syscall.Signal) error {
	return process.Signal(signal)
}

func (l RealOsOperations) CreateCommand(program string, args ...string) *exec.Cmd {
	cmd := exec.Command(program, args...)
	return cmd
}

func (l RealOsOperations) StartCmd(cmd *exec.Cmd) error {
	return cmd.Start()
}

func (l RealOsOperations) KillProcess(process *os.Process) error {
	return process.Kill()
}

func (l RealOsOperations) ProcessStatus(pid int) (*os_api_adapter.ProcessStatus, error) {
	// TODO: implement me
	panic("not implemented")
}

var _ os_api_adapter.OsOperations = (*RealOsOperations)(nil)
