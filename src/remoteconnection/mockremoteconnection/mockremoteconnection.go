package mockremoteconnection

import (
	"math/rand"
	"server/src/remoteconnection"
)

type MockStatusResponse struct {
	totalPlayers int
}

func (mr *MockStatusResponse) TotalPlayers() int {
	return (rand.Int())%9 + 1 // max of 10 playrs, minimum of 1
}

func NewMockStatusResponse() *MockStatusResponse {
	return &MockStatusResponse{}
}

type MockRemoteConnection struct {
	portNum int
}

func (m *MockRemoteConnection) IsAvailable() bool {
	return m.portNum == 25565
}

func (m *MockRemoteConnection) PollServer() remoteconnection.StatusResponse {
	return NewMockStatusResponse()
}

func NewMockRemoteConnection(portNum int) *MockRemoteConnection {
	return &MockRemoteConnection{portNum: portNum}
}

var _ remoteconnection.RemoteConnection = (*MockRemoteConnection)(nil)
