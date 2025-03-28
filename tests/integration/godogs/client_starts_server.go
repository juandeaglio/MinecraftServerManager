package integrationtest

import (
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
	// Check if the Controls indicate the server is not running
	if c.testContext.Controls.IsStarted() {
		return fmt.Errorf("expected server to be stopped, but it's running")
	}
	// Check the status endpoint
	// Additional verification: Try to reach the server status URL
	resp, err := http.Get(constants.StatusURL)
	if err != nil {
		// Connection error means server is likely not running - which is what we want
		return nil
	}
	defer resp.Body.Close()

	// If we get a successful status code, server might be running
	if resp.StatusCode == http.StatusOK {
		return fmt.Errorf("expected server to be stopped, but status endpoint is responding")
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
	s.Before(BeforeScenarioHook(tc))
	s.After(AfterScenarioHook(tc))

	// Register step definitions
	s.Given(`^the Minecraft server isn't started$`, c.theServerIsNotStarted)
	s.When(`^a client starts the server$`, c.aClientAsksTheStatusWithPlayers)
	s.Then(`^the server starts$`, c.iShouldTellTheClientTheStatusWithPlayers)
}
