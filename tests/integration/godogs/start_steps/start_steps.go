package start_steps

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"minecraftremote/src/process"
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
const getStatusURL = constants.BaseURL + port + constants.StatusURL
const startURL = constants.BaseURL + port + constants.StartURL

func (c *startServerFeature) theMinecraftProcessIsNotRunning() error {
	// Check status endpoint of our HTTP API server
	resp, err := http.Get(getStatusURL)
	if err != nil {
		return fmt.Errorf("failed to connect to HTTP API status endpoint: %v", err)
	}
	defer resp.Body.Close()

	// Check if the response code indicates server not running (4xx or 5xx status)
	if resp.StatusCode != 404 {
		return fmt.Errorf("minecraft process should be stopped, but status endpoint reports a successful status but instead reports %v", resp.StatusCode)
	}

	return nil
}

func (c *startServerFeature) aClientRequestsToStartMinecraftProcess() error {
	resp, err := http.Get(startURL)
	if err != nil {
		return fmt.Errorf("failed to connect to HTTP API start endpoint: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		return nil
	}

	return fmt.Errorf("failed to start Minecraft process, status code: %v", resp.StatusCode)
}

func (c *startServerFeature) theMinecraftProcessShouldBeRunning() error {
	return godog.ErrPending
	log.Println("Step 'the Minecraft process should be running' is not implemented!")
	return fmt.Errorf("failed to verify Minecraft process is running")
}

func ClientStartsServer(s *godog.ScenarioContext) {
	rconAdapter := rcon.NewMinecraftRCONAdapter()
	rconAdapter.WithTimeout(1 * time.Second)
	tc := test_infrastructure.NewTestContext(rconAdapter, &process.WindowsOsOperations{}, "notepad.exe")
	c := &startServerFeature{testContext: tc}

	// Register hooks with common infrastructure
	s.Before(test_infrastructure.BeforeScenarioHook(tc, port))
	s.After(test_infrastructure.AfterScenarioHook(tc))

	// Register step definitions
	s.Given(`^the Minecraft server isn't started$`, c.theMinecraftProcessIsNotRunning)
	s.When(`^a client starts the server$`, c.aClientRequestsToStartMinecraftProcess)
	s.Then(`^the server starts$`, c.theMinecraftProcessShouldBeRunning)
}
