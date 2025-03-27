package servertest

import (
	"minecraftremote/src/backgroundserver"
)

type FakeProcess struct {
}

// StartProcess implements backgroundserver.BackgroundServer.
func (f *FakeProcess) StartProcess() {
	panic("unimplemented")
}

var _ backgroundserver.BackgroundServer = (*FakeProcess)(nil)
