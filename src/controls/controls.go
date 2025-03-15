package controls

type Server interface {
	Start()
	Stop()
	Restart()
}

type MinecraftServer struct {
}

func (s *MinecraftServer) Start() {

}

func NewMinecraftServer() *MinecraftServer {
	return &MinecraftServer{}
}
