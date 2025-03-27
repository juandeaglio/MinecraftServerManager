package servertest

import (
	"minecraftremote/src/controls/mcservercontrols"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmptyServerHas0Players(t *testing.T) {
	controls := mcservercontrols.NewControls(&FakeProcess{})
	assert.Equalf(t, 0, controls.Status().Players, "Got more than 0 players on an empty server.")
}

func TestStartServer(t *testing.T) {
	controls := mcservercontrols.NewControls(&FakeProcess{})
	controls.Start()
	assert.Truef(t, controls.IsStarted(), "The server failed to start.")
}

func TestStopServer(t *testing.T) {
	controls := mcservercontrols.NewControls(&FakeProcess{})
	controls.Start()
	controls.Stop()
	assert.Truef(t, controls.IsStarted(), "The server failed to start.")
}
