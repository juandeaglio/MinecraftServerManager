package servertest

import (
	"server/src/controls"
	"server/src/remoteconnection/mockremoteconnection"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestServerStatusWhileActive(t *testing.T) {
	manager := controls.NewServer(mockremoteconnection.NewMockRemoteConnection())
	assert.Truef(t, manager.IsAvailable(25565), "The server should be enabled here, but the manager says it is inactive.")
}

func TestServerStatusWhileOff(t *testing.T) {
	manager := controls.NewServer(mockremoteconnection.NewMockRemoteConnection())
	assert.Falsef(t, manager.IsAvailable(25560), "The server should be disabled here, but the manager says it is active.")
}
