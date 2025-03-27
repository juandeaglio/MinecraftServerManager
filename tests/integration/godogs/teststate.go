package integrationtest

import (
	"minecraftremote/src/controls"
	"net/http"
)

// Shared test state
type TestState struct {
	Controls controls.Controls
	Server   *http.Server
}
