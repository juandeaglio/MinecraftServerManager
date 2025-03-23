package servertest

import (
	"minecraftremote/src/server/mcservercontrols"
	"minecraftremote/tests/unit/httpdriver/cannedrequests"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStartServer(t *testing.T) {
	controls := mcservercontrols.NewServer()
	resp := controls.HandleHTTP(cannedrequests.NewStartRequest().ToHTTPRequest())
	assert.Truef(t, resp.StatusCode == 200, "Server did not start successfully")
}

func TestStopServer(t *testing.T) {
	controls := mcservercontrols.NewServer()
	controls.HandleHTTP(cannedrequests.NewStartRequest().ToHTTPRequest())
	resp := controls.HandleHTTP(cannedrequests.NewStopRequest().ToHTTPRequest())
	assert.Truef(t, resp.StatusCode == 200, "Server did not stop successfully, maybe it did not start?")
}

func TestServerStatistics(t *testing.T) {
	controls := mcservercontrols.NewServer()
	controls.HandleHTTP(cannedrequests.NewStartRequest().ToHTTPRequest())
	resp := controls.HandleHTTP(cannedrequests.NewStatusRequest().ToHTTPRequest())
	assert.Truef(t, resp.StatusCode == 200, "Server did not get status successfully, maybe it did not start?")
}
