package integrationtest

import (
	"context"
	"fmt"
	"log"

	"github.com/cucumber/godog"
)

func (c *checkServerFeature) theServerIsStartedWithPlayers() error {
	log.Println("Step 'the server is started with players' is not implemented!")
	return fmt.Errorf("failed to start server with players")
}

func (c *checkServerFeature) aClientAsksTheStatusWithPlayers() error {

	return fmt.Errorf("The client was unable to get the status correctly")
}

func (c *checkServerFeature) iShouldTellTheClientTheStatusWithPlayers() error {
	log.Println("Step 'I should tell the client the status with players' is not implemented!")
	return fmt.Errorf("failed to get players")
}

func ClientAsksThePopulatedServerForTheStatusScenarioContext(s *godog.ScenarioContext, state *TestState) {
	s.Before(func(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
		log.Printf("Running scenario: %s", sc.Name)
		return ctx, nil
	})
	c := &checkServerFeature{state: state} // Pass the shared state
	s.Given(`^the server is started with players$`, c.theServerIsStartedWithPlayers)
	s.When(`^a client asks the status with players$`, c.aClientAsksTheStatusWithPlayers)
	s.Then(`^the server should tell the client the status with players$`, c.iShouldTellTheClientTheStatusWithPlayers)
}
