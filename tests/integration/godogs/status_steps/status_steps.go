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

func (c *checkServerFeature) serverIsRunning() error {
	c.resp, _ = http.Get(url)
	if c.resp.StatusCode == 200 {
		return nil
	}
	return fmt.Errorf("server is not running - status code: %d", c.resp.StatusCode)
}

func (c *checkServerFeature) GetProcessStatus() error {
	c.resp, _ = http.Get(url)
	return nil
}

func (c *checkServerFeature) ProcessStatusIsSuccessful() error {
	if c.resp.StatusCode == 200 {
		return nil
	}
	return fmt.Errorf("the client was unable to get the Minecraft process status correctly")
}

// We use a lot of specific stubs/fakes here. The question is why?
// Can we simply define something to describe the context required for these fakes/stubs?
// The test might not need real impls to test them, but WHY is it written this way?
// Needs a name to describe this.

// What we test:
// HTTP request check
func ServerStatusScenarioContext(s *godog.ScenarioContext) {
	tc := test_infrastructure.NewTestContext(rcon.NewStubRCONAdapter(), &process.WindowsOsOperations{}, "notepad.exe")
	c := &checkServerFeature{testContext: tc}

	baseHook := test_infrastructure.BeforeScenarioWithNotepadHook(tc, "8080")
	customHook := func(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
		tc.Process = tc.Controls.Start(process.NewProcess(&process.WindowsOsOperations{}, "notepad.exe"))
		return ctx, nil
	}
	s.Before(test_infrastructure.CombineBeforeHooks(baseHook, customHook))
	s.After(test_infrastructure.AfterScenarioHook(tc))

	s.Given(`^the Minecraft server is running and ready$`, c.serverIsRunning)
	s.When(`^a client requests the server status$`, c.GetProcessStatus)
	s.Then(`^the system returns a status response$`, c.ProcessStatusIsSuccessful)
}
