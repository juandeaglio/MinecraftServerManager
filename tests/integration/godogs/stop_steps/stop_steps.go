package stop_steps

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

const statusURL = constants.BaseURL + "8082" + constants.StatusURL
const stopURL = constants.BaseURL + "8082" + constants.StopURL

// func ServerStatusScenarioContext(s *godog.ScenarioContext) {
// 	tc := test_infrastructure.NewTestContext(rcon.NewStubRCONAdapter(), &process.WindowsOsOperations{}, process.NewProcess(&process.WindowsOsOperations{}, "notepad.exe", ""))
// 	c := &checkServerFeature{testContext: tc}

// 	baseHook := test_infrastructure.BeforeScenarioWithNotepadHook(tc, "8080")
// 	customHook := func(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
// 		tc.Process = tc.Controls.Start(process.NewProcess(&process.WindowsOsOperations{}, "notepad.exe"))
// 		return ctx, nil
// 	}
// 	s.Before(test_infrastructure.CombineBeforeHooks(baseHook, customHook))
// 	s.After(test_infrastructure.AfterScenarioHook(tc))

//		s.Given(`^the Minecraft server is running and ready$`, c.serverIsRunning)
//		s.When(`^a client requests the server status$`, c.GetProcessStatus)
//		s.Then(`^the system returns a status response$`, c.ProcessStatusIsSuccessful)
//	}
func ClientStopsServer(s *godog.ScenarioContext) {
	tc := test_infrastructure.NewTestContext(rcon.NewStubRCONAdapter(), &process.WindowsOsOperations{}, process.NewProcess(&process.WindowsOsOperations{}, "notepad.exe", ""))
	c := &checkServerFeature{testContext: tc}

	baseHook := test_infrastructure.BeforeScenarioWithNotepadHook(tc, "8082")
	customHook := func(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
		tc.Process = tc.Controls.Start(process.NewProcess(&process.WindowsOsOperations{}, "notepad.exe", ""))
		return ctx, nil
	}
	s.Before(test_infrastructure.CombineBeforeHooks(baseHook, customHook))
	s.After(test_infrastructure.AfterScenarioHook(tc))

	s.Given(`^a Minecraft server is running$`, c.serverIsRunning)
	s.When(`^a client stops the server$`, c.clientSendsStopRequest)
	s.Then(`^the server stops$`, c.serverProcessTerminates)
}

func (c *checkServerFeature) serverIsRunning() error {
	return test_infrastructure.CheckServerRunning("8082")
}

func (c *checkServerFeature) clientSendsStopRequest() error {
	return fmt.Errorf("not implemented")
}

func (c *checkServerFeature) serverProcessTerminates() error {
	return godog.ErrPending
}
