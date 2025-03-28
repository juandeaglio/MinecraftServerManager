package integrationtest

import (
	"context"
	"fmt"
	"log"
	"minecraftremote/src/controls/mcservercontrols"
	"minecraftremote/src/httprouter"
	"minecraftremote/src/httprouteradapter"
	"minecraftremote/src/process"
	"net/http"
	"time"

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

func ClientStartsServer(s *godog.ScenarioContext) {
	controls := mcservercontrols.NewControls()
	router := httprouter.NewHTTPRouter(controls, &process.WinProcess{})
	routerAdapter := &httprouteradapter.HTTPRouterAdapter{Router: router}
	testState := &TestState{
		Controls: controls,
	}

	s.Before(func(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
		log.Printf("Running scenario: %s", sc.Name)
		testState.Server = *startServerWithRouter(routerAdapter)

		waitForServerReady("http://localhost:8080/status", 5*time.Second)

		return ctx, nil
	})
	c := &startServerFeature{state: testState} // Pass the shared state
	s.Given(`^the Minecraft server isn't started$`, c.theServerIsStartedWithPlayers)
	s.When(`^a client starts the server$`, c.aClientAsksTheStatusWithPlayers)
	s.Then(`^the server starts$`, c.iShouldTellTheClientTheStatusWithPlayers)

	s.After(func(ctx context.Context, sc *godog.Scenario, e error) (context.Context, error) {
		if e != nil {
			log.Printf("Scenario %s failed due to: %s", sc.Name, e.Error())
		}
		testState.Server.Close()

		if testState.Process != nil {
			log.Printf("Explicitly stopping process in ClientStartsServer")
			testState.Process.Stop()
		}

		if testState.Controls != nil {
			testState.Controls.Stop()
		}
		return ctx, nil
	})
}
