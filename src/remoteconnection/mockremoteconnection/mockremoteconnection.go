package mockremoteconnection

import (
	"server/src/remoteconnection"
)

type MockRemoteConnection struct {
}

func (m *MockRemoteConnection) IsAvailable(portNum int) bool {
	return portNum == 25565
}

func (m *MockRemoteConnection) PollServer() bool {
	return true
}

func NewMockRemoteConnection() *MockRemoteConnection {
	return &MockRemoteConnection{}
}

var _ remoteconnection.RemoteConnection = (*MockRemoteConnection)(nil)
