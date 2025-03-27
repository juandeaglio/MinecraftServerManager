package integrationtest

import (
	"minecraftremote/src/controls"
	"minecraftremote/src/process"
	"net/http"
)

// Shared test state
type TestState struct {
	Controls controls.Controls
	Server   *http.Server
	Process  *process.Process
}
