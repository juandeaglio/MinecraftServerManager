package integrationtest

import (
	"fmt"
	"log"
	"minecraftremote/src/controls"
	"minecraftremote/src/controls/mcservercontrols"
	"minecraftremote/src/httprouter"
	"minecraftremote/src/httprouteradapter"
	"net/http"
	"testing"
	"time"

	"github.com/cucumber/godog"
)

// Shared test state
type TestState struct {
	Controls controls.Controls
	Server   *http.Server
}

func TestScenarios(t *testing.T) {
	testState := &TestState{
		Controls: mcservercontrols.NewControls(),
	}
	testState.Controls.Start()

	// Set up router and server
	router := httprouter.NewHTTPServer(testState.Controls)
	routerAdapter := &httprouteradapter.HTTPRouterAdapter{Router: router}
	testState.Server = startServerWithRouter(routerAdapter)

	waitForServerReady("http://localhost:8080/status", 5*time.Second)

	// Combine both scenario contexts into one initializer
	combinedScenarioInitializer := func(s *godog.ScenarioContext) {
		ClientAsksTheServerForTheStatusScenarioContext(s, testState)
		ClientAsksThePopulatedServerForTheStatusScenarioContext(s, testState)
	}

	suite := godog.TestSuite{
		ScenarioInitializer: combinedScenarioInitializer,
		Options: &godog.Options{
			Format:   "pretty",
			Paths:    []string{"features"},
			TestingT: t,
			Strict:   true,
			NoColors: false,   // Ensure colors are enabled for clarity
			Tags:     "~@wip", // Exclude work-in-progress tests
		},
	}

	if status := suite.Run(); status != 0 {
		t.Fatalf("Feature tests failed with status: %d", status)
	}

	testState.Server.Close()
}

func runScenario(t *testing.T, scenarioFeature ScenarioContextFunc) godog.TestSuite {
	return godog.TestSuite{
		ScenarioInitializer: func(s *godog.ScenarioContext) {
			scenarioFeature(s)
		},
		Options: &godog.Options{
			Format:   "pretty",
			Paths:    []string{"features"},
			TestingT: t,
			Strict:   true,
			NoColors: false,   // Ensure colors are enabled for clarity
			Tags:     "~@wip", // Exclude work-in-progress tests
		},
	}
}

func startServerWithRouter(adapter *httprouteradapter.HTTPRouterAdapter) *http.Server {
	server := &http.Server{
		Addr:         ":8080",
		Handler:      adapter, // Use your router's HandleHTTP as the handler
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	// Start the server in a goroutine
	go func() {
		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Printf("HTTP server error: %v", err)
		}
	}()

	return server
}

func waitForServerReady(url string, timeout time.Duration) error {
	deadline := time.Now().Add(timeout)
	for time.Now().Before(deadline) {
		_, err := http.Get(url)
		if err == nil {
			return nil // Server is up and responding
		}
		time.Sleep(100 * time.Millisecond)
	}
	return fmt.Errorf("server failed to become ready within %v", timeout)
}

type ScenarioContextFunc func(*godog.ScenarioContext)
