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

type StatusServerFeature struct {
	testContext *test_infrastructure.TestContext
	resp        *http.Response
}

const statusRequestURL = constants.BaseURL + "8080" + constants.StatusURL
const runningURL = constants.BaseURL + "8080" + constants.RunningURL

func ServerStatusScenarioContext(s *godog.ScenarioContext) {
	tc := test_infrastructure.NewTestContext(rcon.NewStubRCONAdapter(), &process.WindowsOsOperations{}, process.NewProcess(&process.WindowsOsOperations{}, "notepad.exe", ""))
	c := &StatusServerFeature{testContext: tc}

	baseHook := test_infrastructure.BeforeScenarioWithNotepadHook(tc, "8080")
	customHook := func(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
		tc.Process = tc.Controls.Start(process.NewProcess(&process.WindowsOsOperations{}, "notepad.exe"))
		return ctx, nil
	}
	s.Before(test_infrastructure.CombineBeforeHooks(baseHook, customHook))
	s.After(test_infrastructure.AfterScenarioHook(tc))

	s.Given(`^the Minecraft server is running and ready$`, c.ServerIsRunning)
	s.When(`^a client requests the server status$`, c.GetProcessStatus)
	s.Then(`^the system returns a status response$`, c.ProcessStatusIsSuccessful)
}

// CheckServerRunning checks if the server is running at the specified port

func (c *StatusServerFeature) ServerIsRunning() error {
	return test_infrastructure.CheckServerRunning("8080")
}

func (c *StatusServerFeature) GetProcessStatus() error {
	resp, err := http.Get(statusRequestURL)
	if err != nil {
		return fmt.Errorf("failed to get server status: %v", err)
	}
	c.resp = resp
	return nil
}

func (c *StatusServerFeature) ProcessStatusIsSuccessful() error {
	if c.resp.StatusCode == 200 {
		return nil
	}
	return fmt.Errorf("the client was unable to get the Minecraft process status correctly")
}
