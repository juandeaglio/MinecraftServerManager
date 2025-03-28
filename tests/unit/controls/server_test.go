package servertest

import (
	"minecraftremote/src/controls/mcservercontrols"
	"minecraftremote/src/process"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmptyServerHas0Players(t *testing.T) {
	controls := mcservercontrols.NewControls()
	assert.Equalf(t, 0, controls.Status().Players, "Got more than 0 players on an empty server.")
}

func TestStartServer(t *testing.T) {
	controls := mcservercontrols.NewControls()
	controls.Start(&process.FakeProcess{})
	assert.Truef(t, controls.IsStarted(), "The server process failed to start.")
}

func TestOfflineServerStatus(t *testing.T) {
	controls := mcservercontrols.NewControls()

	// Assert that the server status shows offline
	assert.Falsef(t, controls.Status().Online, "Server with no PID should report as offline.")
}

func TestStartAndStopServer(t *testing.T) {
	// Create server controls
	controls := mcservercontrols.NewControls()

	// Start the server with a fake process
	fakeProcess := &process.FakeProcess{}
	controls.Start(fakeProcess)

	// Assert that the server is started
	assert.Truef(t, controls.IsStarted(), "The server process failed to start.")

	// Simulate the process stopping on its own
	fakeProcess.Stop() // This should simulate the process terminating

	// Assert that controls still thinks server is running (revealing the bug)
	// This assertion should fail since controls isn't monitoring the process state
	assert.Falsef(t, controls.IsStarted(), "Controls thinks server is running when process has died")
}
