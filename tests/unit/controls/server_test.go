package servertest

import (
	"minecraftremote/src/controls/mcservercontrols"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmptyServerHas0Players(t *testing.T) {
	server := mcservercontrols.NewControls()
	assert.Equalf(t, 0, server.Status().Players, "Got more than 0 players on an empty server.")
}
