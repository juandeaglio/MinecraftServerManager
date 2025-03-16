package mcservercontrols

import (
	remoteconnection "minecraftremote/src/remoteconnection"
)

type Server interface {
	Start()
	Stop()
	Restart()
	GetStatus() remoteconnection.StatusResponse
}
