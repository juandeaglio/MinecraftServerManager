package processtest

import (
	"testing"

	"minecraftremote/src/os_api_adapter"
	"minecraftremote/tests/unit/process/brokenosoperations"

	"github.com/stretchr/testify/assert"
)

func TestProcess(t *testing.T) {
	fakeProcess := os_api_adapter.NewProcessHandler(&os_api_adapter.FakeOsOperations{}, "fake", "args")
	_ = fakeProcess.Start()
	assert.Truef(t, fakeProcess.Started(), "ProcessContext failed to start.")

	_ = fakeProcess.Stop()
	assert.Falsef(t, fakeProcess.Started(), "ProcessContext is still marked as started after stopping")
}

func TestProcessError(t *testing.T) {
	fakeProcess := os_api_adapter.NewProcessHandler(&brokenosoperations.BrokenOsOperations{}, "fake", "args")
	err := fakeProcess.Start()
	assert.Error(t, err)
}
