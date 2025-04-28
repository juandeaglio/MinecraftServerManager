package status_steps

import (
	"context"
	"fmt"
	"minecraftremote/src/process"
	"minecraftremote/src/rcon"
	"minecraftremote/tests/integration/godogs/constants"
	"minecraftremote/tests/integration/godogs/test_infrastructure"
	"net/http"

	"github.com/cucumber/godog"
)

type checkServerFeature struct {
	testContext *test_infrastructure.TestContext
	resp        *http.Response
}

const url = constants.BaseURL + "8080" + constants.StatusURL

// We are testing the HTTP server, but the process underneath is faked.
// Maybe we should actually use a real process, but it is just the contract test.
// So, this integration test should actually be more of a contract test.
func (c *checkServerFeature) theMinecraftProcessIsRunning() error {
	c.resp, _ = http.Get(url)
	if c.resp.StatusCode == 200 {
		return nil
	}
	return fmt.Errorf("server is not running - status code: %d", c.resp.StatusCode)
}

func (c *checkServerFeature) aClientRequestsMinecraftProcessStatus() error {
	c.resp, _ = http.Get(url)
	return nil
}

func (c *checkServerFeature) theAPIReturnsMinecraftProcessStatus() error {
	if c.resp.StatusCode == 200 {
		return nil
	}
	return fmt.Errorf("The client was unable to get the Minecraft process status correctly")
}

// We use a lot of specific stubs/fakes here. The question is why?
// Can we simply define something to describe the context required for these fakes/stubs?
// The test might not need real impls to test them, but WHY is it written this way?
// Needs a name to describe this.

// What we test:
// HTTP request check
// Potential failure paths
// Dependency on port bindings, stub adapters, etc.

// vs. what the steps are describing:
func ClientAsksTheServerForTheStatusScenarioContext(s *godog.ScenarioContext) {
	tc := test_infrastructure.NewTestContext(rcon.NewStubRCONAdapter(), &process.FakeOsOperations{}, "notepad.exe")
	c := &checkServerFeature{testContext: tc}

	baseHook := test_infrastructure.BeforeScenarioWithNotepadHook(tc, "8080")
	customHook := func(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
		tc.Process = tc.Controls.Start(process.NewProcess(&process.FakeOsOperations{}, "notepad.exe"))
		return ctx, nil
	}
	s.Before(test_infrastructure.CombineBeforeHooks(baseHook, customHook))

	s.After(test_infrastructure.AfterScenarioHook(tc))

	s.Given(`^the Minecraft server is running and ready$`, c.theMinecraftProcessIsRunning)
	s.When(`^a client requests the server status$`, c.aClientRequestsMinecraftProcessStatus)
	s.Then(`^the system returns a status response$`, c.theAPIReturnsMinecraftProcessStatus)
}
