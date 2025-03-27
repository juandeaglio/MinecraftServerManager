package process

import (
	"fmt"
	"os/exec"
	"syscall"
)

type WinProcess struct {
	cmd     *exec.Cmd
	program string
	args    []string
}

// NewRealProcess initializes a process with a given command.
func NewWinProcess(command string, args ...string) *WinProcess {
	return &WinProcess{
		program: command, args: args,
	}
}

// PID implements Process.
func (w *WinProcess) PID() int {
	return -1 // Simulate no process running
}

// Start implements Process.
func (w *WinProcess) Start() error {
	if w.program == "" {
		return fmt.Errorf("program is empty")
	}
	w.cmd = exec.Command(w.program, w.args...)
	w.cmd.SysProcAttr = &syscall.SysProcAttr{
		CreationFlags: syscall.CREATE_NEW_PROCESS_GROUP, // Create a new process group
		HideWindow:    false,                            // Prevents the Notepad window from showing

	}
	return w.cmd.Start()
}

// Stop implements Process.
func (w *WinProcess) Stop() error {
	if w.cmd != nil && w.cmd.Process != nil {
		return w.cmd.Process.Kill()
	}
	return nil
}

var _ Process = (*FakeProcess)(nil)
