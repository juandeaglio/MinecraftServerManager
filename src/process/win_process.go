package process

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

// OsOperations abstracts operating system level operations
type OsOperations interface {
	FindProcess(pid int) (*os.Process, error)
	Signal(process *os.Process, signal syscall.Signal) error
	CreateCommand(program string, args ...string) *exec.Cmd
	SetSysProcAttr(cmd *exec.Cmd)
	StartCmd(cmd *exec.Cmd) error
	KillProcess(process *os.Process) error
}

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

type WinProcess struct {
	cmd     *exec.Cmd
	program string
	args    []string
	osOps   OsOperations
}

// Started implements Process.
func (w *WinProcess) Started() bool {
	// TODO: needs unit test
	if w.cmd == nil || w.cmd.Process == nil {
		return false
	}

	return w.isProcessRunning(w.cmd.Process.Pid)
}

// isProcessRunning checks if a process with the given PID is running
func (w *WinProcess) isProcessRunning(pid int) bool {
	process, err := w.osOps.FindProcess(pid)
	if err != nil {
		return false
	}

	err = w.osOps.Signal(process, syscall.Signal(0))
	return err == nil
}

// NewWinProcess initializes a process with a given command.
func NewWinProcess(command string, args ...string) *WinProcess {
	return &WinProcess{
		program: command,
		args:    args,
		osOps:   &WindowsOsOperations{},
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
	w.cmd = w.osOps.CreateCommand(w.program, w.args...)
	w.osOps.SetSysProcAttr(w.cmd)

	err := w.osOps.StartCmd(w.cmd)
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
		err := w.osOps.KillProcess(w.cmd.Process)
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
