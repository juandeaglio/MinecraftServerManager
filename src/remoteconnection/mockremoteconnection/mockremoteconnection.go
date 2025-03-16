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
}

func (m *MockRemoteConnection) IsAvailable(portNum int) bool {
	return portNum == 25565
}

func (m *MockRemoteConnection) PollServer() remoteconnection.StatusResponse {
	return NewMockStatusResponse()
}

func NewMockRemoteConnection() *MockRemoteConnection {
	return &MockRemoteConnection{}
}

var _ remoteconnection.RemoteConnection = (*MockRemoteConnection)(nil)
