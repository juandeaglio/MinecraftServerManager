package servertest

import (
	"server/src/remoteconnection/mockremoteconnection"
	"server/src/server/mcservercontrols"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestServerAvailableWhileActive(t *testing.T) {
	manager := mcservercontrols.NewServer(mockremoteconnection.NewMockRemoteConnection(25565))
	assert.Truef(t, manager.IsAvailable(), "The server should be enabled here, but the manager says it is inactive.")
}

func TestServerAvailableWhileOff(t *testing.T) {
	manager := mcservercontrols.NewServer(mockremoteconnection.NewMockRemoteConnection(25560))
	assert.Falsef(t, manager.IsAvailable(), "The server should be disabled here, but the manager says it is active.")
}

func TestServerWhileRunningStatus(t *testing.T) {
	manager := mcservercontrols.NewServer(mockremoteconnection.NewMockRemoteConnection(25565))
	actual := manager.GetStatus()
	assert.NotNil(t, actual)
	expected := 0
	assert.Greater(t, actual.TotalPlayers(), expected)
}
