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

const statusRequestURL = constants.BaseURL + "8080" + constants.StatusURL
const runningURL = constants.BaseURL + "8080" + constants.RunningURL

func ServerStatusScenarioContext(s *godog.ScenarioContext) {
	tc := test_infrastructure.NewTestContext(rcon.NewStubRCONAdapter(), &process.WindowsOsOperations{}, process.NewProcess(&process.WindowsOsOperations{}, "notepad.exe", ""))
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

func (c *checkServerFeature) serverIsRunning() error {
	resp, err := http.Get(runningURL)
	if err != nil {
		return fmt.Errorf("failed to get server running status: %v", err)
	}

	if resp.StatusCode == 200 {
		return nil
	}
	return fmt.Errorf("server is not running - status code: %d", resp.StatusCode)
}

func (c *checkServerFeature) GetProcessStatus() error {
	resp, err := http.Get(statusRequestURL)
	if err != nil {
		return fmt.Errorf("failed to get server status: %v", err)
	}
	c.resp = resp
	return nil
}

func (c *checkServerFeature) ProcessStatusIsSuccessful() error {
	if c.resp.StatusCode == 200 {
		return nil
	}
	return fmt.Errorf("the client was unable to get the Minecraft process status correctly")
}
