package controls

import "minecraftremote/src/process"

type Controls interface {
	Start(process.Process) process.Process
	Stop() bool
	Status() *Status
	IsStarted() bool
}

type Status struct {
	Players int
}
