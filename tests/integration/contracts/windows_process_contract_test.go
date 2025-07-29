package contract_test

import (
	"minecraftremote/src/os_api_adapter"
	"minecraftremote/src/os_api_adapter/real_os_ops"
	"minecraftremote/src/windowsconstants"
	"testing"
)

func TestWindowsRunningProcessContract(t *testing.T) {
	pc := os_api_adapter.NewProcessHandler(&real_os_ops.RealOsOperations{}, "notepad.exe", "")
	err := pc.Start()
	if err != nil {
		return
	}

	t.Logf("Process started successfully with PID %d", pc.PID())

	defer func(pc *os_api_adapter.ProcessImpl) {
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
	pc := os_api_adapter.NewProcessHandler(&real_os_ops.RealOsOperations{}, "notepad.exe", "")

	defer func(pc *os_api_adapter.ProcessImpl) {
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
	pc := os_api_adapter.NewProcessHandler(&real_os_ops.RealOsOperations{}, "notepad.exe", "")
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
