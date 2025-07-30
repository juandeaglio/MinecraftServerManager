//go:build windows

package real_os_ops

import (
	"fmt"
	"minecraftremote/src/os_api_adapter"
	"os"
	"os/exec"
	"syscall"
	"unsafe"
)

// RealOsOperations implements OsOperations for Windows
type RealOsOperations struct{}

func (w *RealOsOperations) FindProcess(pid int) (*os.Process, error) {
	return os.FindProcess(pid)
}

func (w *RealOsOperations) Signal(process *os.Process, signal syscall.Signal) error {
	return process.Signal(signal)
}

func (w *RealOsOperations) CreateCommand(program string, args ...string) *exec.Cmd {
	cmd := exec.Command(program, args...)
	w.SetSysProcAttr(cmd)

	return cmd
}

func (w *RealOsOperations) SetSysProcAttr(cmd *exec.Cmd) {
	cmd.SysProcAttr = &syscall.SysProcAttr{
		CreationFlags: syscall.CREATE_NEW_PROCESS_GROUP,
		HideWindow:    false,
	}
}

func (w *RealOsOperations) StartCmd(cmd *exec.Cmd) error {
	return cmd.Start()
}

func (w *RealOsOperations) KillProcess(process *os.Process) error {
	return process.Kill()
}

func (w *RealOsOperations) ProcessStatus(pid int) (*os_api_adapter.ProcessStatus, error) {
	status, err := getStatus(uint32(pid))
	if err != nil {
		return nil, err
	}
	return &os_api_adapter.ProcessStatus{
		Status: status,
	}, nil
}

func getStatus(pid uint32) (uintptr, error) {
	handle, err := syscall.OpenProcess(syscall.PROCESS_QUERY_INFORMATION, false, pid)
	if err != nil {
		errno, ok := err.(syscall.Errno)
		if !ok {
			return 0, err // Unknown error, bail
		}
		return uintptr(errno), nil

	}
	defer func(handle syscall.Handle) {
		_ = syscall.CloseHandle(handle)
	}(handle)

	ntdll := syscall.NewLazyDLL("ntdll.dll")
	procNtQueryInformationProcess := ntdll.NewProc("NtQueryInformationProcess")

	pbi, status := getStatusCode(procNtQueryInformationProcess, handle)

	if status != 0 {
		return 0, fmt.Errorf("NtQueryInformationProcess failed with status 0x%x", status)
	}

	return pbi.ExitStatus, nil
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

func getStatusCode(procNtQueryInformationProcess *syscall.LazyProc, handle syscall.Handle) (PROCESS_BASIC_INFORMATION, uintptr) {
	var pbi PROCESS_BASIC_INFORMATION
	var returnLength uintptr

	status, _, _ := procNtQueryInformationProcess.Call(
		uintptr(handle),
		uintptr(ProcessBasicInformation),
		uintptr(unsafe.Pointer(&pbi)),
		unsafe.Sizeof(pbi),
		uintptr(unsafe.Pointer(&returnLength)),
	)
	return pbi, status
}

var _ os_api_adapter.OsOperations = (*RealOsOperations)(nil)
