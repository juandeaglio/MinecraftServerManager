package unit_test

import (
	"encoding/json"
	"minecraftremote/src/controls"
	"minecraftremote/src/httprouter"
	"minecraftremote/src/process_context"
	"minecraftremote/src/rcon"
	"minecraftremote/tests/unit/httpdriver/cannedrequests"
	"minecraftremote/tests/unit/process/brokenosoperations"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStartServer(t *testing.T) {
	router := httprouter.NewHTTPRouter(
		controls.NewControls(nil),
		process_context.NewProcessInvoker(&process_context.FakeOsOperations{}, "fake", "args"),
	)
	resp := router.HandleHTTP(cannedrequests.NewStartRequest().ToHTTPRequest())
	assert.Equalf(t, 200, resp.StatusCode, "Server did not start successfully")
}

func TestFailtoStartServer(t *testing.T) {
	router := httprouter.NewHTTPRouter(
		controls.NewControls(nil),
		process_context.NewProcessInvoker(&brokenosoperations.BrokenOsOperations{}, "fake", "args"),
	)
	resp := router.HandleHTTP(cannedrequests.NewStartRequest().ToHTTPRequest())
	assert.Equalf(t, 500, resp.StatusCode, "Server did not start successfully")
}

func TestStopServer(t *testing.T) {
	router := httprouter.NewHTTPRouter(
		controls.NewControls(nil),
		process_context.NewProcessInvoker(&process_context.FakeOsOperations{}, "fake", "args"),
	)
	router.HandleHTTP(cannedrequests.NewStartRequest().ToHTTPRequest())
	router.HandleHTTP(cannedrequests.NewStopRequest().ToHTTPRequest())
	resp := router.HandleHTTP(cannedrequests.NewRunningRequest().ToHTTPRequest())

	assert.Equalf(t, 404, resp.StatusCode, "Server did not stop successfully, maybe it did not stop?")
}

func TestServerStatus(t *testing.T) {
	router := httprouter.NewHTTPRouter(
		controls.NewControls(nil),
		process_context.NewProcessInvoker(&process_context.FakeOsOperations{}, "fake", "args"),
	)
	router.HandleHTTP(cannedrequests.NewStartRequest().ToHTTPRequest())
	resp := router.HandleHTTP(cannedrequests.NewStatusRequest().ToHTTPRequest())
	assert.Equalf(t, 200, resp.StatusCode, "Server did not get status successfully, maybe it did not start?")

	var status map[string]interface{}
	err := json.NewDecoder(resp.Body).Decode(&status)
	assert.NoError(t, err)
	assert.Truef(t, status["Online"].(bool), "Server should be online")
}

func TestServerStatusWhenServerIsOffline(t *testing.T) {
	router := httprouter.NewHTTPRouter(controls.NewControls(rcon.NewStubRCONAdapter()), nil)
	resp := router.HandleHTTP(cannedrequests.NewStatusRequest().ToHTTPRequest())
	assert.Equalf(t, 404, resp.StatusCode, "Server is offline, but status endpoint returned 200")
}

func TestServerRunning(t *testing.T) {
	router := httprouter.NewHTTPRouter(controls.NewControls(
		rcon.NewStubRCONAdapter()),
		process_context.NewProcessInvoker(&process_context.FakeOsOperations{}, "fake", "args"),
	)
	router.HandleHTTP(cannedrequests.NewStartRequest().ToHTTPRequest())
	resp := router.HandleHTTP(cannedrequests.NewRunningRequest().ToHTTPRequest())
	assert.Equalf(t, 200, resp.StatusCode, "Server is running, but running endpoint returned 404")
}

func TestServerNotRunning(t *testing.T) {
	router := httprouter.NewHTTPRouter(
		controls.NewControls(rcon.NewStubRCONAdapter()),
		process_context.NewProcessInvoker(&process_context.FakeOsOperations{}, "fake", "args"),
	)
	resp := router.HandleHTTP(cannedrequests.NewRunningRequest().ToHTTPRequest())
	assert.Equalf(t, 404, resp.StatusCode, "Server is not running, but running endpoint returned 200")
}
