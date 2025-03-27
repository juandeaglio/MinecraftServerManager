package integrationtest

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/cucumber/godog"
)

type startServerFeature struct {
	state *TestState
	resp  *http.Response
}

func (c *startServerFeature) theServerIsStartedWithPlayers() error {
	if !c.state.Controls.IsStarted() {
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

func ClientStartsServer(s *godog.ScenarioContext, state *TestState) {
	s.Before(func(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
		log.Printf("Running scenario: %s", sc.Name)
		return ctx, nil
	})
	c := &startServerFeature{state: state} // Pass the shared state
	s.Given(`^the Minecraft server isn't running$`, c.theServerIsStartedWithPlayers)
	s.When(`^a client starts the server$`, c.aClientAsksTheStatusWithPlayers)
	s.Then(`^the server starts$`, c.iShouldTellTheClientTheStatusWithPlayers)
}
