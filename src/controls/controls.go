package controls

import (
	"minecraftremote/src/process"
	"minecraftremote/src/rcon"
)

type Controls interface {
	Start(process.Process) process.Process
	Stop() bool
	Status() *rcon.Status
	IsStarted() bool
}
type ControlsImpl struct {
	rconAdapter rcon.RCONAdapter
	started     bool
}

func (c *ControlsImpl) Start(p process.Process) process.Process {
	c.started = true
	p.Start()
	return p
}
func (c *ControlsImpl) Stop() bool {
	c.started = false
	return true
}
func (c *ControlsImpl) Status() *rcon.Status {
	return c.rconAdapter.GetStatus()
}
func (c *ControlsImpl) IsStarted() bool {
	return
}
