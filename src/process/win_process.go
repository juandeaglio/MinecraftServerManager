package process

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

type WinProcess struct {
	cmd     *exec.Cmd
	program string
	args    []string
}

// Started implements Process.
func (w *WinProcess) Started() bool {
	// TODO: needs unit test
	if w.cmd == nil || w.cmd.Process == nil {
		return false
	}

	process, err := os.FindProcess(w.cmd.Process.Pid)
	if err != nil {
		return false
	}

	err = process.Signal(syscall.Signal(0))
	return err == nil
}

// NewRealProcess initializes a process with a given command.
func NewWinProcess(command string, args ...string) *WinProcess {
	return &WinProcess{
		program: command, args: args,
	}
}

func (w *WinProcess) PID() int {
	if w.cmd != nil && w.cmd.Process != nil {
		return w.cmd.Process.Pid
	}
	return -1 // Indicates no process running
}

func (w *WinProcess) Start() error {
	if w.program == "" {
		return fmt.Errorf("program is empty")
	}
	w.cmd = exec.Command(w.program, w.args...)
	w.cmd.SysProcAttr = &syscall.SysProcAttr{
		CreationFlags: syscall.CREATE_NEW_PROCESS_GROUP, // Create a new process group
		HideWindow:    false,                            // Prevents the Notepad window from showing

	}
	err := w.cmd.Start()
	if err != nil {
		return fmt.Errorf("failed to start process %s: %w", w.program, err)
	}

	// Log successful start with PID
	fmt.Printf("Process started: %s (PID: %d)\n", w.program, w.cmd.Process.Pid)
	return nil
}

func (w *WinProcess) Stop() error {
	if w.cmd != nil && w.cmd.Process != nil {
		pid := w.cmd.Process.Pid
		err := w.cmd.Process.Kill()
		if err != nil {
			return fmt.Errorf("failed to kill process %s (PID: %d): %w", w.program, pid, err)
		}
		// consider removing printf
		fmt.Printf("Process killed: %s (PID: %d)\n", w.program, pid)
		// Ensure we clean up the reference
		w.cmd = nil
		return nil
	}
	return nil
}

var _ Process = (*WinProcess)(nil)
