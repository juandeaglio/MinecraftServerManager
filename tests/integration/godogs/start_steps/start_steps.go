package start_steps

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"minecraftremote/src/rcon"
	"minecraftremote/tests/integration/godogs/constants"
	"minecraftremote/tests/integration/godogs/test_infrastructure"

	"github.com/cucumber/godog"
)

type startServerFeature struct {
	testContext *test_infrastructure.TestContext
	resp        *http.Response
}

const url = constants.BaseURL + "8081" + constants.StatusURL

func (c *startServerFeature) theMinecraftProcessIsNotRunning() error {
	// Check status endpoint of our HTTP API server
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("failed to connect to HTTP API status endpoint: %v", err)
	}
	defer resp.Body.Close()

	var statusResponse map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&statusResponse); err != nil {
		return fmt.Errorf("failed to decode status response: %v", err)
	}

	if online, exists := statusResponse["Online"].(bool); exists && online {
		return fmt.Errorf("Minecraft process should be stopped, but status endpoint reports Online=true")
	}

	return nil
}

func (c *startServerFeature) aClientRequestsToStartMinecraftProcess() error {
	return godog.ErrPending

	return fmt.Errorf("the client was unable to start the Minecraft process correctly")
}

func (c *startServerFeature) theMinecraftProcessShouldBeRunning() error {
	return godog.ErrPending

	log.Println("Step 'the Minecraft process should be running' is not implemented!")
	return fmt.Errorf("failed to verify Minecraft process is running")
}

func ClientStartsServer(s *godog.ScenarioContext) {
	rconAdapter := rcon.NewMinecraftRCONAdapter()
	rconAdapter.WithTimeout(1 * time.Second)
	tc := test_infrastructure.NewTestContext(rconAdapter)
	c := &startServerFeature{testContext: tc}

	// Register hooks with common infrastructure
	s.Before(test_infrastructure.BeforeScenarioHook(tc, "8081"))
	s.After(test_infrastructure.AfterScenarioHook(tc))

	// Register step definitions
	s.Given(`^the Minecraft server isn't started$`, c.theMinecraftProcessIsNotRunning)
	s.When(`^a client starts the server$`, c.aClientRequestsToStartMinecraftProcess)
	s.Then(`^the server starts$`, c.theMinecraftProcessShouldBeRunning)
}
