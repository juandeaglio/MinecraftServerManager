package integrationtest

import (
	"fmt"
	"log"
	"net/http"

	"github.com/cucumber/godog"
)

type startServerFeature struct {
	testContext *TestContext
	resp        *http.Response
}

func (c *startServerFeature) theServerIsStartedWithPlayers() error {
	if !c.testContext.Controls.IsStarted() {
		return fmt.Errorf("the server was unable to start")
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
	s.Given(`^the Minecraft server isn't started$`, c.theServerIsStartedWithPlayers)
	s.When(`^a client starts the server$`, c.aClientAsksTheStatusWithPlayers)
	s.Then(`^the server starts$`, c.iShouldTellTheClientTheStatusWithPlayers)
}
