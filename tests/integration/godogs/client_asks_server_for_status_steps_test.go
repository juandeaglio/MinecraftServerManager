package integrationtest

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"minecraftremote/src/controls/mcservercontrols"
	"minecraftremote/src/httprouter"
	"minecraftremote/src/httprouteradapter"
	"minecraftremote/src/process"
	"net/http"
	"strings"
	"time"

	"github.com/cucumber/godog"
)

type checkServerFeature struct {
	state *TestState
	resp  *http.Response
}

const statusURL = "http://localhost:8080/status"

func (c *checkServerFeature) theServerIsStarted() error {
	c.resp, _ = http.Get(statusURL)
	if c.resp.StatusCode == 200 {
		return nil
	}
	return godog.ErrAmbiguous
}

func (c *checkServerFeature) aClientAsksTheStatus() error {
	c.resp, _ = http.Get(statusURL)
	statusCode := c.resp.StatusCode
	if statusCode == 200 {
		return nil
	}
	return fmt.Errorf("The client was unable to get the status correctly")
}

func (c *checkServerFeature) iShouldTellTheClientTheStatus() error {
	contentType := c.resp.Header.Get("Content-Type")
	if !strings.Contains(contentType, "application/json") {
		return fmt.Errorf("expected content type to be application/json but got %s", contentType)
	}

	defer c.resp.Body.Close()
	body, err := io.ReadAll(c.resp.Body)
	if err != nil {
		return fmt.Errorf("error reading response body: %v", err)
	}

	var response map[string]interface{}
	if err := json.Unmarshal(body, &response); err != nil {
		return fmt.Errorf("error parsing JSON response: %v", err)
	}

	players := response["Players"]

	playersValue, _ := players.(float64) // JSON numbers are parsed as float64 by default

	if playersValue < 0 {
		return fmt.Errorf("expected 'Players' value to be 0 but got %v", playersValue)
	}

	return nil
}

func ClientAsksTheServerForTheStatusScenarioContext(s *godog.ScenarioContext) {
	controls := mcservercontrols.NewControls()
	router := httprouter.NewHTTPRouter(controls, &process.WinProcess{})
	routerAdapter := &httprouteradapter.HTTPRouterAdapter{Router: router}
	testState := &TestState{
		Controls: controls,
	}

	s.Before(func(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
		log.Printf("Running scenario: %s", sc.Name)
		testState.Server = *startServerWithRouter(routerAdapter)

		testState.Process = testState.Controls.Start(process.NewWinProcess("notepad.exe"))

		// Set up router and server

		waitForServerReady("http://localhost:8080/status", 5*time.Second)

		return ctx, nil
	})
	c := &checkServerFeature{state: testState} // Pass the shared state

	s.Given(`^the Minecraft server is running$`, c.theServerIsStarted)
	s.When(`^a client requests the server status$`, c.aClientAsksTheStatus)
	s.Then(`^the system returns a status response indicating "online" along with the current player count$`, c.iShouldTellTheClientTheStatus)

	s.After(func(ctx context.Context, sc *godog.Scenario, e error) (context.Context, error) {
		if e != nil {
			log.Printf("Scenario %s failed due to: %s", sc.Name, e.Error())
		}

		testState.Server.Close()

		if testState.Process != nil {
			log.Printf("Explicitly stopping notepad process")
			testState.Process.Stop()
		}

		if testState.Controls != nil {
			testState.Controls.Stop()
		}
		return ctx, nil
	})
}
