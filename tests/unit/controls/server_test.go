package servertest

import (
	"minecraftremote/src/remoteconnection/mockremoteconnection"
	"minecraftremote/src/server/mcservercontrols"
	"testing"

	"github.com/stretchr/testify/assert"
)

// func TestServerAvailableWhileActive(t *testing.T) {
// 	conn := stubremoteconnection.NewMockRemoteConnection(25565).IsAvailable()
// 	manager := mcservercontrols.NewServer()
// 	resp := manager.Accept(conn)
// 	assert.Truef(t, resp.StatusCode == "200", "The server should be enabled here, but the manager says it is inactive.")
// }

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
