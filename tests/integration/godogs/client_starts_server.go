package integrationtest

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"minecraftremote/tests/integration/godogs/constants"

	"github.com/cucumber/godog"
)

type startServerFeature struct {
	testContext *TestContext
	resp        *http.Response
}

func (c *startServerFeature) theServerIsNotStarted() error {
	// Check status endpoint
	resp, err := http.Get(constants.StatusURL)
	if err != nil {
		return fmt.Errorf("failed to connect to status endpoint: %v", err)
	}
	defer resp.Body.Close()

	var statusResponse map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&statusResponse); err != nil {
		return fmt.Errorf("failed to decode status response: %v", err)
	}

	if online, exists := statusResponse["Online"].(bool); exists && online {
		return fmt.Errorf("server should be stopped, but status endpoint reports Online=true")
	}

	return nil
}

func (c *startServerFeature) aClientAsksTheStatusWithPlayers() error {
	return fmt.Errorf("the client was unable to get the status correctly")
}

func (c *startServerFeature) iShouldTellTheClientTheStatusWithPlayers() error {
	log.Println("Step 'I should tell the client the status with players' is not implemented!")
	return fmt.Errorf("failed to get players")
}

func ClientStartsServer(s *godog.ScenarioContext) {
	tc := NewTestContext()
	c := &startServerFeature{testContext: tc}

	// Register hooks with common infrastructure
	s.Before(BeforeScenarioHook(tc, "8081"))
	s.After(AfterScenarioHook(tc))

	// Register step definitions
	s.Given(`^the Minecraft server isn't started$`, c.theServerIsNotStarted)
	s.When(`^a client starts the server$`, c.aClientAsksTheStatusWithPlayers)
	s.Then(`^the server starts$`, c.iShouldTellTheClientTheStatusWithPlayers)
}
