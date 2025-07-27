package control_server_steps

import (
	"fmt"
	"net/http"
	"time"

	"minecraftremote/src/os_api_adapter"
	"minecraftremote/src/rcon"
	"minecraftremote/tests/integration/godogs/constants"
	"minecraftremote/tests/integration/godogs/test_infrastructure"

	"github.com/cucumber/godog"
)

type startServerFeature struct {
	testContext *test_infrastructure.TestContext
	resp        *http.Response
}

const port = "8081"
const getRunningURL = constants.BaseURL + port + constants.RunningURL
const startURL = constants.BaseURL + port + constants.StartURL

func ClientStartsServer(s *godog.ScenarioContext) {
	rconAdapter := rcon.NewMinecraftRCONAdapter()
	rconAdapter.WithTimeout(1 * time.Second)
	osOps := &os_api_adapter.WindowsOsOperations{}
	tc := test_infrastructure.NewTestContext(
		rconAdapter,
		os_api_adapter.NewProcessHandler(osOps, "notepad.exe", ""))

	c := &startServerFeature{testContext: tc}

	// Register hooks with common infrastructure
	s.Before(test_infrastructure.BeforeScenarioHook(c.testContext, port))
	s.After(test_infrastructure.AfterScenarioHook(c.testContext))

	// Register step definitions
	s.Given(`^the Minecraft server isn't started$`, c.theMinecraftProcessIsNotRunning)
	s.When(`^a client starts the server$`, c.aClientRequestsToStartMinecraftProcess)
	s.Then(`^the server starts$`, c.theMinecraftProcessShouldBeRunning)
}

func (c *startServerFeature) theMinecraftProcessIsNotRunning() error {
	// Check status endpoint of our HTTP API server
	resp, _ := http.Get(getRunningURL)
	if resp.StatusCode == 200 {
		return fmt.Errorf("minecraft process should be stopped, but status endpoint reports a successful status but instead reports %v", resp.StatusCode)
	}
	return nil
}

func (c *startServerFeature) aClientRequestsToStartMinecraftProcess() error {
	c.resp, _ = http.Get(startURL)
	return nil
}

func (c *startServerFeature) theMinecraftProcessShouldBeRunning() error {
	if c.resp.StatusCode == 200 {
		return nil
	}
	return fmt.Errorf("failed to start Minecraft process, status code: %v", c.resp.StatusCode)
}
