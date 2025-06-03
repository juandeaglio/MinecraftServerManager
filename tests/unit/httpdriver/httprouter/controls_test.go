package unit_test

import (
	"minecraftremote/src/controls"
	"minecraftremote/src/process_context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmptyServerHas0Players(t *testing.T) {
	serverControls := controls.NewControls(nil, &process_context.ProcessImpl{})
	assert.Equalf(t, 0, serverControls.Status().Players, "Got more than 0 players on an empty server.")
}

func TestStartServerControls(t *testing.T) {
	serverControls := controls.NewControls(nil, &process_context.ProcessImpl{})
	serverControls.Start(process_context.NewProcessInvoker(&process_context.FakeOsOperations{}, "fake", "args"))
	assert.Truef(t, serverControls.IsStarted(), "The server process failed to start.")
}

func TestStopServerControls(t *testing.T) {
	serverControls := controls.NewControls(nil, &process_context.ProcessImpl{})
	serverControls.Start(process_context.NewProcessInvoker(&process_context.FakeOsOperations{}, "fake", "args"))
	serverControls.Stop()
	assert.Falsef(t, serverControls.IsStarted(), "The server process failed to stop.")
}

func TestOfflineServerStatus(t *testing.T) {
	controls := controls.NewControls(nil, &process_context.ProcessImpl{})

	assert.Falsef(t, controls.Status().Online, "Server with no PID should report as offline.")
}
