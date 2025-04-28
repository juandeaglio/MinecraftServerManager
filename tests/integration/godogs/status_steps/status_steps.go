package status_steps

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"minecraftremote/src/process"
	"minecraftremote/src/rcon"
	"minecraftremote/tests/integration/godogs/constants"
	"minecraftremote/tests/integration/godogs/test_infrastructure"
	"net/http"
	"strings"

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
	statusCode := c.resp.StatusCode
	if statusCode == 200 {
		return nil
	}
	return fmt.Errorf("The client was unable to get the Minecraft process status correctly")
}

func (c *checkServerFeature) theAPIReturnsMinecraftProcessStatusWithPlayerCount() error {
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

	return nil
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
	s.Then(`^the system returns a status response indicating "online" along with the current player count$`, c.theAPIReturnsMinecraftProcessStatusWithPlayerCount)
}
