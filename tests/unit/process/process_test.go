package processtest

import (
	"testing"

	"minecraftremote/src/process_context"
	"minecraftremote/tests/unit/process/brokenosoperations"

	"github.com/stretchr/testify/assert"
)

func TestProcess(t *testing.T) {
	fakeProcess := process_context.NewProcessInvoker(&process_context.FakeOsOperations{}, "fake", "args")
	fakeProcess.Start()
	assert.Truef(t, fakeProcess.Started(), "ProcessContext failed to start.")

	fakeProcess.Stop()
	assert.Falsef(t, fakeProcess.Started(), "ProcessContext is still marked as started after stopping")
}

func TestProcessError(t *testing.T) {
	fakeProcess := process_context.NewProcessInvoker(&brokenosoperations.BrokenOsOperations{}, "fake", "args")
	err := fakeProcess.Start()
	assert.Error(t, err)
}
