package servertest

import (
	"server/src/controls"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestServerStart(t *testing.T) {
	manager := controls.NewMinecraftServer()
	assert.NotNil(t, manager)
}
