package httproutertests

import (
	"minecraftremote/src/httprouter"
	"minecraftremote/src/process"
	"minecraftremote/src/stubcontrols"
	"minecraftremote/tests/unit/httpdriver/cannedrequests"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStartServer(t *testing.T) {
	router := httprouter.NewHTTPRouter(stubcontrols.NewStubControls(), &process.FakeProcess{})
	resp := router.HandleHTTP(cannedrequests.NewStartRequest().ToHTTPRequest())
	assert.Equalf(t, resp.StatusCode, 200, "Server did not start successfully")
}

func TestStopServer(t *testing.T) {
	router := httprouter.NewHTTPRouter(stubcontrols.NewStubControls(), &process.FakeProcess{})
	router.HandleHTTP(cannedrequests.NewStartRequest().ToHTTPRequest())
	resp := router.HandleHTTP(cannedrequests.NewStopRequest().ToHTTPRequest())
	assert.Equalf(t, resp.StatusCode, 200, "Server did not stop successfully, maybe it did not start?")
}

func TestServerStatistics(t *testing.T) {
	router := httprouter.NewHTTPRouter(stubcontrols.NewStubControls(), &process.FakeProcess{})
	router.HandleHTTP(cannedrequests.NewStartRequest().ToHTTPRequest())
	resp := router.HandleHTTP(cannedrequests.NewStatusRequest().ToHTTPRequest())
	assert.Equalf(t, resp.StatusCode, 200, "Server did not get status successfully, maybe it did not start?")
}
