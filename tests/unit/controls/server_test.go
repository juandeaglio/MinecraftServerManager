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
