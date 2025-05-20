package contract_test

import (
	"minecraftremote/src/process_context"
	"testing"
)

func TestProcessAPIContract(t *testing.T) {
	t.Run("Windows contract", func(t *testing.T) {
		pc := process_context.NewProcessInvoker(&process_context.WindowsOsOperations{}, "notepad.exe", "")
		err := pc.Start()
		if err != nil {
			return
		}

		stopAfterTest(pc)

		ps := pc.GetProcessStatus(pc.PID())

		if ps.Status == "" {
			t.Errorf("Expected process status, but got empty.")
		}

		t.Logf("Process status: %s", ps.Status)
	})
}

func stopAfterTest(pc *process_context.ProcessImpl) {
	defer func(pc *process_context.ProcessImpl) {
		_ = pc.Stop()
	}(pc)
}
