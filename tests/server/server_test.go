package servertest

import (
	"server/src/controls"
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockSystemInterface struct {
	onOrOff bool
}

func (m *MockSystemInterface) PollExecutable() bool {
	return m.onOrOff
}

func NewMockSystemInterface(onOrOff bool) *MockSystemInterface {
	return &MockSystemInterface{onOrOff: onOrOff}
}

func TestServerStatusWhileActive(t *testing.T) {
	manager := controls.NewServer(NewMockSystemInterface(true))
	assert.Truef(t, manager.GetStatus(), "The server should be enabled here, but the manager says it is inactive.")
}

func TestServerStatusWhileOff(t *testing.T) {
	manager := controls.NewServer(NewMockSystemInterface(false))
	assert.Falsef(t, manager.GetStatus(), "The server should be disabled here, but the manager says it is active.")
}
