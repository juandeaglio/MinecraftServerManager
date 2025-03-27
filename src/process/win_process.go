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
		HideWindow: true, // Prevents the Notepad window from showing
	}
	return w.cmd.Start()
}

// Stop implements Process.
func (w *WinProcess) Stop() error {
	panic("Unimplemented")
}

var _ Process = (*FakeProcess)(nil)
