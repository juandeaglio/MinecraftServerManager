package httprouter

import "minecraftremote/src/server/mcservercontrols"

type HTTPRouter interface {
	handleStart()
	handleStop()
	handleStatus()
	HandleHTTP()
}

type HTTPServer struct {
	mcServer *mcservercontrols.MinecraftServer
}
