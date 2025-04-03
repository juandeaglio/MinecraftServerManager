package processtest

import (
	"minecraftremote/src/process"

	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProcessStartAndStop(t *testing.T) {
	fakeProcess := process.NewFakeProcess(&process.FakeOsOperations{})

	fakeProcess.Start()
	assert.Truef(t, fakeProcess.Started(), "Process failed to start.")

	fakeProcess.Stop()
	assert.Falsef(t, fakeProcess.Started(), "Process is still marked as started after stopping")
}

func TestStartABrokenProcess(t *testing.T) {
	fakeProcess := process.NewFakeProcess(&brokenOsOperations{})

	err := fakeProcess.Start()
	assert.Error(t, err, "Process should not start if OsOperations fails")
}
