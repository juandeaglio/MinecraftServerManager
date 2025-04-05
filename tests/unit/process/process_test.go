package processtest

import (
	"testing"

	"minecraftremote/src/process"

	"github.com/stretchr/testify/assert"
)

func TestProcess(t *testing.T) {
	fakeProcess := process.NewProcess(&process.FakeOsOperations{}, "fake", "args")
	fakeProcess.Start()
	assert.Truef(t, fakeProcess.Started(), "Process failed to start.")

	fakeProcess.Stop()
	assert.Falsef(t, fakeProcess.Started(), "Process is still marked as started after stopping")
}

func TestProcessError(t *testing.T) {
	fakeProcess := process.NewProcess(&brokenOsOperations{}, "fake", "args")
	err := fakeProcess.Start()
	assert.Error(t, err)
}
