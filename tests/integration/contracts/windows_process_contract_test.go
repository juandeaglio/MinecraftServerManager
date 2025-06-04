package contract_test

import (
	"minecraftremote/src/process_context"
	"minecraftremote/src/windowsconstants"
	"testing"
)

func TestWindowsRunningProcessContract(t *testing.T) {
	pc := process_context.NewProcessInvoker(&process_context.WindowsOsOperations{}, "notepad.exe", "")
	err := pc.Start()
	if err != nil {
		return
	}

	t.Logf("Process started successfully with PID %d", pc.PID())

	defer func(pc *process_context.ProcessImpl) {
		_ = pc.Stop()
	}(pc)

	ps, err := pc.GetProcessStatus(pc.PID())

	if err != nil {
		t.Fatal(err)
	}

	if ps.Status != windowsconstants.RunningStatus {
		t.Errorf("Expected process to be running, but got some other status.")
	}

	t.Logf("Process status: %d", ps.Status)
}
func TestWindowsProcessNonExistentContract(t *testing.T) {
	pc := process_context.NewProcessInvoker(&process_context.WindowsOsOperations{}, "notepad.exe", "")

	defer func(pc *process_context.ProcessImpl) {
		_ = pc.Stop()
	}(pc)

	ps, err := pc.GetProcessStatus(9999)

	if err != nil {
		t.Fatal(err)
	}

	if ps.Status != windowsconstants.InvalidProcessStatus {
		t.Errorf("Expected process to not be running, but got some other status.")
	}

	t.Logf("Process status: %d", ps.Status)

}

func TestWindowsKilledProcessContract(t *testing.T) {
	pc := process_context.NewProcessInvoker(&process_context.WindowsOsOperations{}, "notepad.exe", "")
	err := pc.Start()
	if err != nil {
		return
	}

	t.Logf("Process started successfully with PID %d", pc.PID())

	pid := pc.PID()

	err = pc.Stop()
	if err != nil {
		return
	}
	t.Logf("Process stopped successfully with PID %d", pc.PID())

	ps, err := pc.GetProcessStatus(pid)

	if err != nil {
		t.Fatal(err)
	}

	if ps.Status != windowsconstants.ParentKilledChildStatus {
		t.Errorf("Expected process to have exited successfully, but got some other status %d.", ps.Status)
	}

	t.Logf("Process status: %d", ps.Status)
}
