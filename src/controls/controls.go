package controls

import (
	"server/src/systeminterface"
)

type Server interface {
	Start()
	Stop()
	Restart()
	GetStatus() bool
}

type MinecraftServer struct {
	systemInterface systeminterface.SystemInterface
}

func (s *MinecraftServer) GetStatus() bool {
	return s.systemInterface.PollExecutable()
}

func (s *MinecraftServer) Start() {

}
func (s *MinecraftServer) Stop() {

}
func (s *MinecraftServer) Restart() {

}

func NewServer(inter systeminterface.SystemInterface) *MinecraftServer {
	return &MinecraftServer{systemInterface: inter}
}
