package servertest

import (
	"minecraftremote/src/controls/mcservercontrols"
	"minecraftremote/src/process"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmptyServerHas0Players(t *testing.T) {
	controls := mcservercontrols.NewControls(&process.ProcessImpl{})
	assert.Equalf(t, 0, controls.Status().Players, "Got more than 0 players on an empty server.")
}

func TestStartServer(t *testing.T) {
	controls := mcservercontrols.NewControls()
	controls.Start(process.NewProcess(&process.FakeOsOperations{}, "fake", "args"))
	assert.Truef(t, controls.IsStarted(), "The server process failed to start.")
}

func TestOfflineServerStatus(t *testing.T) {
	controls := mcservercontrols.NewControls(&process.ProcessImpl{})

	// Assert that the server status shows offline
	assert.Falsef(t, controls.Status().Online, "Server with no PID should report as offline.")
}
