package mcservercontrols

import (
	remoteconnection "server/src/remoteconnection"
)

type Server interface {
	Start()
	Stop()
	Restart()
	GetStatus() remoteconnection.StatusResponse
}
