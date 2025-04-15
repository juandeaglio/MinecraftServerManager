package httproutertests

import (
	"minecraftremote/src/controls"
	"minecraftremote/src/httprouter"
	"minecraftremote/src/process"
	"minecraftremote/tests/unit/httpdriver/cannedrequests"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStartServer(t *testing.T) {
	router := httprouter.NewHTTPRouter(controls.NewControls(), process.NewProcess(&process.FakeOsOperations{}, "fake", "args"))
	resp := router.HandleHTTP(cannedrequests.NewStartRequest().ToHTTPRequest())
	assert.Equalf(t, resp.StatusCode, 200, "Server did not start successfully")
}

func TestStopServer(t *testing.T) {
	router := httprouter.NewHTTPRouter(controls.NewControls(), process.NewProcess(&process.FakeOsOperations{}, "fake", "args"))
	router.HandleHTTP(cannedrequests.NewStartRequest().ToHTTPRequest())
	resp := router.HandleHTTP(cannedrequests.NewStopRequest().ToHTTPRequest())
	assert.Equalf(t, resp.StatusCode, 200, "Server did not stop successfully, maybe it did not start?")
}

func TestServerStatistics(t *testing.T) {
	router := httprouter.NewHTTPRouter(controls.NewControls(), process.NewProcess(&process.FakeOsOperations{}, "fake", "args"))
	router.HandleHTTP(cannedrequests.NewStartRequest().ToHTTPRequest())
	resp := router.HandleHTTP(cannedrequests.NewStatusRequest().ToHTTPRequest())
	assert.Equalf(t, resp.StatusCode, 200, "Server did not get status successfully, maybe it did not start?")
}

// func TestServerStatisticsWhenServerIsOffline(t *testing.T) {
// 	router := httprouter.NewHTTPRouter(controls.NewControls(), nil)
// 	resp := router.HandleHTTP(cannedrequests.NewStatusRequest().ToHTTPRequest())
// 	assert.Equalf(t, resp.StatusCode, 500, "Server is offline, but status endpoint returned 200")
// }
