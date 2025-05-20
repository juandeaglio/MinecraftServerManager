package process_context

import (
	"os"
	"os/exec"
	"syscall"
)

var (
	kernel32                 = syscall.NewLazyDLL("kernel32.dll")
	procOpenProcess          = kernel32.NewProc("OpenProcess")
	procGetProcessMemoryInfo = psapi.NewProc("GetProcessMemoryInfo")
	psapi                    = syscall.NewLazyDLL("psapi.dll")
)

// WindowsOsOperations implements OsOperations for Windows
type WindowsOsOperations struct{}

func (w *WindowsOsOperations) FindProcess(pid int) (*os.Process, error) {
	return os.FindProcess(pid)
}

func (w *WindowsOsOperations) Signal(process *os.Process, signal syscall.Signal) error {
	return process.Signal(signal)
}

func (w *WindowsOsOperations) CreateCommand(program string, args ...string) *exec.Cmd {
	return exec.Command(program, args...)
}

func (w *WindowsOsOperations) SetSysProcAttr(cmd *exec.Cmd) {
	cmd.SysProcAttr = &syscall.SysProcAttr{
		CreationFlags: syscall.CREATE_NEW_PROCESS_GROUP,
		HideWindow:    false,
	}
}

func (w *WindowsOsOperations) StartCmd(cmd *exec.Cmd) error {
	return cmd.Start()
}

func (w *WindowsOsOperations) KillProcess(process *os.Process) error {
	return process.Kill()
}

func (w *WindowsOsOperations) ProcessStatus(pid int) (ProcessStatus, error) {

	return ProcessStatus{}, nil
}

var _ OsOperations = (*WindowsOsOperations)(nil)
