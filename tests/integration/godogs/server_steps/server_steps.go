package server_steps

import (
	"fmt"
	"minecraftremote/src/process"
	"minecraftremote/src/rcon"
	"minecraftremote/tests/integration/godogs/test_infrastructure"
	"minecraftremote/tests/unit/httpdriver/cannedrequests"
	"net/http"
	"time"

	"github.com/cucumber/godog"
)

type startServerFeature struct {
	testContext *test_infrastructure.TestContext
}

func StartServer(s *godog.ScenarioContext) {
	c := &startServerFeature{}

	// Register step definitions
	s.Given(`^a process is not running$`, c.processIsNotRunning)
	s.When(`^the process starts$`, c.processStarts)
	s.Then(`^the process should be running$`, c.processIsRunning)

}

func (c *startServerFeature) processIsNotRunning() error {
	osOps := &process.WindowsOsOperations{}
	c.testContext = test_infrastructure.NewTestContext(
		rcon.NewStubRCONAdapter(),
		osOps,
		process.NewProcess(osOps, "notepad.exe", ""),
	)

	test_infrastructure.StartServerWithRouter(c.testContext.Adapter, "8083")

	test_infrastructure.WaitForServerReady("http://localhost:8083/running", 1*time.Second)

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
