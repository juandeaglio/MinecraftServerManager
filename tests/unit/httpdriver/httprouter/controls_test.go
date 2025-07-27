package unit_test

import (
	"minecraftremote/src/controls"
	"minecraftremote/src/os_api_adapter"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmptyServerHas0Players(t *testing.T) {
	serverControls := controls.NewControls(nil, &os_api_adapter.ProcessImpl{})
	assert.Equalf(t, 0, serverControls.Status().Players, "Got more than 0 players on an empty server.")
}

func TestStartServerControls(t *testing.T) {
	serverControls := controls.NewControls(nil, &os_api_adapter.ProcessImpl{})
	serverControls.Start(os_api_adapter.NewProcessHandler(&os_api_adapter.FakeOsOperations{}, "fake", "args"))
	assert.Truef(t, serverControls.IsStarted(), "The server process failed to start.")
}

func TestStopServerControls(t *testing.T) {
	serverControls := controls.NewControls(nil, &os_api_adapter.ProcessImpl{})
	serverControls.Start(os_api_adapter.NewProcessHandler(&os_api_adapter.FakeOsOperations{}, "fake", "args"))
	assert.Truef(t, serverControls.Stop(), "The server process failed to stop.")
}

func TestOfflineServerStatus(t *testing.T) {
	serverControls := controls.NewControls(nil, &os_api_adapter.ProcessImpl{})
	assert.Falsef(t, serverControls.Status().Online, "Server with no PID should report as offline.")
}
