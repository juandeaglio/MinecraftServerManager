package test_infrastructure

import (
	"context"
	"log"
	"minecraftremote/src/controls"
	"minecraftremote/src/httprouter"
	"minecraftremote/src/httprouteradapter"
	"minecraftremote/src/process"
	"minecraftremote/src/rcon"
	"net/http"
	"time"

	"github.com/cucumber/godog"
)

// TestContext holds all shared test state and dependencies
type TestContext struct {
	Controls *controls.Controls
	Process  process.Process
	Server   *http.Server
	Router   *httprouter.ServerRouter
	Adapter  *httprouteradapter.HTTPRouterAdapter
}

// NewTestContext creates a new test context with initialized dependencies
func NewTestContext(rconAdapter rcon.RCONAdapter, osOperations process.OsOperations, process process.Process) *TestContext {
	if rconAdapter == nil {
		rconAdapter = rcon.NewMinecraftRCONAdapter()
	}
	controls := controls.NewControls(rconAdapter, process)
	router := httprouter.NewHTTPRouter(controls, process)
	adapter := &httprouteradapter.HTTPRouterAdapter{Router: router}

	return &TestContext{
		Controls: controls,
		Router:   router,
		Adapter:  adapter,
		Process:  process,
	}
}

// BeforeScenarioHook sets up the test environment before each scenario
func BeforeScenarioHook(tc *TestContext, port string) func(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
	return func(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
		log.Printf("Running scenario: %s", sc.Name)
		tc.Server = startServerWithRouter(tc.Adapter, port)

		waitForServerReady("http://localhost:"+port+"/status", 5*time.Second)

		return ctx, nil
	}
}

// BeforeScenarioWithNotepadHook sets up the test environment and starts notepad
func BeforeScenarioWithNotepadHook(tc *TestContext, port string) func(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
	return func(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
		log.Printf("Running scenario: %s", sc.Name)
		tc.Server = startServerWithRouter(tc.Adapter, port)

		waitForServerReady("http://localhost:"+port+"/status", 5*time.Second)

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

func CombineBeforeHooks(hooks ...func(ctx context.Context, sc *godog.Scenario) (context.Context, error)) func(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
	return func(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
		var err error
		for _, hook := range hooks {
			ctx, err = hook(ctx, sc)
			if err != nil {
				return ctx, err
			}
		}
		return ctx, nil
	}
}

func waitForServerReady(url string, timeout time.Duration) {
	deadline := time.Now().Add(timeout)
	for time.Now().Before(deadline) {
		resp, err := http.Get(url)
		if err == nil {
			resp.Body.Close()
			return
		}
		time.Sleep(100 * time.Millisecond)
	}
	log.Printf("Warning: Server might not be ready after waiting %v", timeout)
}
