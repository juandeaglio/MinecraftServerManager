package integrationtest

import (
	"context"
	"log"
	"minecraftremote/src/controls/mcservercontrols"
	"minecraftremote/src/httprouter"
	"minecraftremote/src/httprouteradapter"
	"minecraftremote/src/process"
	"net/http"
	"time"

	"github.com/cucumber/godog"
)

// TestContext holds all shared test state and dependencies
type TestContext struct {
	Controls *mcservercontrols.MinecraftServer
	Process  process.Process
	Server   *http.Server
	Router   *httprouter.ServerRouter
	Adapter  *httprouteradapter.HTTPRouterAdapter
}

// NewTestContext creates a new test context with initialized dependencies
func NewTestContext() *TestContext {
	controls := mcservercontrols.NewControls()
	router := httprouter.NewHTTPRouter(controls, &process.WinProcess{})
	adapter := &httprouteradapter.HTTPRouterAdapter{Router: router}

	return &TestContext{
		Controls: controls,
		Router:   router,
		Adapter:  adapter,
	}
}

// BeforeScenarioHook sets up the test environment before each scenario
func BeforeScenarioHook(tc *TestContext) func(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
	return func(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
		log.Printf("Running scenario: %s", sc.Name)
		tc.Server = startServerWithRouter(tc.Adapter)

		// Wait for server to be ready
		waitForServerReady("http://localhost:8080/status", 5*time.Second)

		return ctx, nil
	}
}

// BeforeScenarioWithNotepadHook sets up the test environment and starts notepad
func BeforeScenarioWithNotepadHook(tc *TestContext) func(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
	return func(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
		log.Printf("Running scenario: %s", sc.Name)
		tc.Server = startServerWithRouter(tc.Adapter)

		// Start notepad process for testing
		tc.Process = tc.Controls.Start(process.NewWinProcess("notepad.exe"))

		// Wait for server to be ready
		waitForServerReady("http://localhost:8080/status", 5*time.Second)

		return ctx, nil
	}
}

// AfterScenarioHook cleans up resources after each scenario
func AfterScenarioHook(tc *TestContext) func(ctx context.Context, sc *godog.Scenario, err error) (context.Context, error) {
	return func(ctx context.Context, sc *godog.Scenario, err error) (context.Context, error) {
		if err != nil {
			log.Printf("Scenario %s failed due to: %s", sc.Name, err.Error())
		}

		if tc.Server != nil {
			tc.Server.Close()
		}

		if tc.Process != nil {
			log.Printf("Explicitly stopping process")
			tc.Process.Stop()
		}

		if tc.Controls != nil {
			tc.Controls.Stop()
		}

		return ctx, nil
	}
}

// waitForServerReady polls the server until it's ready or timeout
func waitForServerReady(url string, timeout time.Duration) {
	deadline := time.Now().Add(timeout)
	for time.Now().Before(deadline) {
		resp, err := http.Get(url)
		if err == nil && resp.StatusCode == http.StatusOK {
			resp.Body.Close()
			return
		}
		time.Sleep(100 * time.Millisecond)
	}
	log.Printf("Warning: Server might not be ready after waiting %v", timeout)
}
