package httproutertests

import (
	"minecraftremote/src/httprouter"
	"minecraftremote/tests/unit/httpdriver/cannedrequests"
	"minecraftremote/tests/unit/httpdriver/stubcontrols"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStartServer(t *testing.T) {
	server := httprouter.NewHTTPServer(stubcontrols.NewStubControls())
	resp := server.HandleHTTP(cannedrequests.NewStartRequest().ToHTTPRequest())
	assert.Truef(t, resp.StatusCode == 200, "Server did not start successfully")
}

func TestStopServer(t *testing.T) {
	server := httprouter.NewHTTPServer(stubcontrols.NewStubControls())
	server.HandleHTTP(cannedrequests.NewStartRequest().ToHTTPRequest())
	resp := server.HandleHTTP(cannedrequests.NewStopRequest().ToHTTPRequest())
	assert.Truef(t, resp.StatusCode == 200, "Server did not stop successfully, maybe it did not start?")
}

func TestServerStatistics(t *testing.T) {
	server := httprouter.NewHTTPServer(stubcontrols.NewStubControls())
	server.HandleHTTP(cannedrequests.NewStartRequest().ToHTTPRequest())
	resp := server.HandleHTTP(cannedrequests.NewStatusRequest().ToHTTPRequest())
	assert.Truef(t, resp.StatusCode == 200, "Server did not get status successfully, maybe it did not start?")
}
