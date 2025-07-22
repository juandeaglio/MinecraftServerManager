package server_steps

import (
	"fmt"
	"minecraftremote/src/os_api_adapter"
	"minecraftremote/src/rcon"
	"minecraftremote/tests/integration/godogs/test_infrastructure"
	"minecraftremote/tests/unit/httpdriver/cannedrequests"
	"net/http"

	"github.com/cucumber/godog"
)

type startServerFeature struct {
	testContext *test_infrastructure.TestContext
}

func StartServer(s *godog.ScenarioContext) {
	c := &startServerFeature{}
	osOps := &os_api_adapter.WindowsOsOperations{}
	c.testContext = test_infrastructure.NewTestContext(
		rcon.NewStubRCONAdapter(),
		os_api_adapter.NewProcessInvoker(osOps, "notepad.exe", ""),
	)
	s.Before(test_infrastructure.BeforeScenarioHook(c.testContext, "8083"))

	s.Given(`^the server does not have a process$`, c.processIsNotRunning)
	s.When(`^the server starts a process$`, c.processStarts)
	s.Then(`^the server process should be running$`, c.processIsRunning)

	s.After(test_infrastructure.AfterScenarioHook(c.testContext))
}

func (c *startServerFeature) processIsNotRunning() error {
	resp, err := http.Get("http://localhost:8083/running")
	if err != nil {
		return fmt.Errorf("failed to get server running status: %v", err)
	}
	if resp.StatusCode == 200 {
		return fmt.Errorf("server is running when it should not be")
	}
	return nil
}

func (c *startServerFeature) processStarts() error {
	c.testContext.Router.HandleHTTP(cannedrequests.NewStartRequest().ToHTTPRequest())
	return nil
}

func (c *startServerFeature) processIsRunning() error {
	return nil
}
