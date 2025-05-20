package process_context

import (
	"fmt"
	"minecraftremote/src/windowsconstants"
	"os"
	"os/exec"
	"syscall"
	"unsafe"
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

func (w *WindowsOsOperations) ProcessStatus(pid int) (*ProcessStatus, error) {
	status, err := getStatus(uint32(pid))
	if err != nil {
		return nil, err
	}
	return &ProcessStatus{
		Status: status,
	}, nil
}

const (
	ProcessBasicInformation = 0
)

type PROCESS_BASIC_INFORMATION struct {
	ExitStatus                   uintptr
	PebBaseAddress               uintptr
	AffinityMask                 uintptr
	BasePriority                 uintptr
	UniqueProcessID              uintptr
	InheritedFromUniqueProcessID uintptr
}

func getStatus(pid uint32) (uintptr, error) {
	handle, err := syscall.OpenProcess(syscall.PROCESS_QUERY_INFORMATION, false, pid)
	if err != nil {
		errno, ok := err.(syscall.Errno)
		if !ok {
			return 0, err // Unknown error, bail
		}

		switch errno {
		case windowsconstants.InvalidProcessStatus:
			return windowsconstants.InvalidProcessStatus, nil
		}
	}
	defer syscall.CloseHandle(handle)

	ntdll := syscall.NewLazyDLL("ntdll.dll")
	procNtQueryInformationProcess := ntdll.NewProc("NtQueryInformationProcess")

	var pbi PROCESS_BASIC_INFORMATION
	var returnLength uintptr

	status, _, _ := procNtQueryInformationProcess.Call(
		uintptr(handle),
		uintptr(ProcessBasicInformation),
		uintptr(unsafe.Pointer(&pbi)),
		unsafe.Sizeof(pbi),
		uintptr(unsafe.Pointer(&returnLength)),
	)

	if status != 0 {
		return 0, fmt.Errorf("NtQueryInformationProcess failed with status 0x%x", status)
	}

	// STATUS_PENDING (0x103) means the process is still running
	return pbi.ExitStatus, nil
}

var _ OsOperations = (*WindowsOsOperations)(nil)
