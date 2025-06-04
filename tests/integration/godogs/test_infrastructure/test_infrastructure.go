package test_infrastructure

import (
	"context"
	"fmt"
	"log"
	"minecraftremote/src/controls"
	"minecraftremote/src/httprouter"
	"minecraftremote/src/httprouteradapter"
	"minecraftremote/src/process_context"
	"minecraftremote/src/rcon"
	"net/http"
	"time"

	"github.com/cucumber/godog"
)

// TestContext holds all shared test state and dependencies
type TestContext struct {
	Controls       *controls.Controls
	ProcessContext process_context.Process
	Server         *http.Server
	Router         *httprouter.ServerRouter
	Adapter        *httprouteradapter.HTTPRouterAdapter
}

// NewTestContext creates a new test context with initialized dependencies
func NewTestContext(rconAdapter rcon.RCONAdapter, process process_context.Process) *TestContext {
	if rconAdapter == nil {
		rconAdapter = rcon.NewMinecraftRCONAdapter()
	}
	serverControls := controls.NewControls(rconAdapter, process)
	router := httprouter.NewHTTPRouter(serverControls, process)
	adapter := &httprouteradapter.HTTPRouterAdapter{Router: router}

	return &TestContext{
		Controls:       serverControls,
		Router:         router,
		Adapter:        adapter,
		ProcessContext: process,
	}
}

// BeforeScenarioHook sets up the test environment before each scenario
func BeforeScenarioHook(tc *TestContext, port string) func(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
	return func(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
		log.Printf("Running scenario: %s", sc.Name)
		tc.Server = StartServerWithRouter(tc.Adapter, port)
		WaitForServerReady("http://localhost:"+port+"/running", 5*time.Second)

		return ctx, nil
	}
}

// BeforeScenarioWithNotepadHook sets up the test environment and starts notepad
func BeforeScenarioWithNotepadHook(tc *TestContext, port string) func(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
	return func(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
		log.Printf("Running scenario: %s", sc.Name)
		tc.Server = StartServerWithRouter(tc.Adapter, port)
		_, _ = http.Get("http://localhost:" + port + "/start")
		fmt.Println("Sent start request to server")
		WaitForServerReady("http://localhost:"+port+"/running", 5*time.Second)

		return ctx, nil
	}
}

// AfterScenarioHook cleans up resources after each scenario
func AfterScenarioHook(tc *TestContext) func(ctx context.Context, sc *godog.Scenario, err error) (context.Context, error) {
	return func(ctx context.Context, sc *godog.Scenario, err error) (context.Context, error) {
		if err != nil {

			log.Printf("Scenario %s failed due to: %s", sc.Name, err.Error())
		}

		err = tc.ProcessContext.Stop()
		if err != nil {
			return nil, err
		}

		tc.Controls.Stop()

		if tc.Server != nil {
			err := tc.Server.Close()
			if err != nil {
				return nil, err
			}
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

func WaitForServerReady(url string, timeout time.Duration) {
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
