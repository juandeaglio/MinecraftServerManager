package contract_test

import (
	"minecraftremote/src/process_context"
	"minecraftremote/src/windowsconstants"
	"testing"
)

func TestProcessAPIContract(t *testing.T) {
	t.Run("Windows contract", func(t *testing.T) {
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
	})
	t.Run("Windows non-existent process contract", func(t *testing.T) {
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

	})

}
