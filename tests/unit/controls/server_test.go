package servertest

import (
	"server/src/controls"
	"server/src/remoteconnection/mockremoteconnection"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestServerAvailableWhileActive(t *testing.T) {
	manager := controls.NewServer(mockremoteconnection.NewMockRemoteConnection())
	assert.Truef(t, manager.IsAvailable(25565), "The server should be enabled here, but the manager says it is inactive.")
}

func TestServerAvailableWhileOff(t *testing.T) {
	manager := controls.NewServer(mockremoteconnection.NewMockRemoteConnection())
	assert.Falsef(t, manager.IsAvailable(25560), "The server should be disabled here, but the manager says it is active.")
}

func TestServerStatus(t *testing.T) {
	manager := controls.NewServer(mockremoteconnection.NewMockRemoteConnection())
	actual := manager.GetStatus()
	assert.NotNil(t, actual)
	expected := 0
	assert.Greater(t, actual.TotalPlayers(), expected)
}
