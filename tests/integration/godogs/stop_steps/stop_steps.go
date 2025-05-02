package stop_steps

import (
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

func ClientStopsServer(s *godog.ScenarioContext) {
	tc := test_infrastructure.NewTestContext(rcon.NewStubRCONAdapter(), &process.WindowsOsOperations{}, process.NewProcess(&process.WindowsOsOperations{}, "notepad.exe", ""))
	c := &checkServerFeature{testContext: tc}

	baseHook := test_infrastructure.BeforeScenarioWithNotepadHook(tc, "8082")
	s.Before(baseHook)
	s.After(test_infrastructure.AfterScenarioHook(tc))

	s.Given(`^a Minecraft server is running$`, c.serverIsRunning)
	s.When(`^a client stops the server$`, c.clientSendsStopRequest)
	s.Then(`^the server stops$`, c.serverProcessTerminates)
}

func (c *checkServerFeature) serverIsRunning() error {
	return test_infrastructure.CheckServerRunning("8082")
}

func (c *checkServerFeature) clientSendsStopRequest() error {
	resp, err := http.Get(constants.BaseURL + "8082" + constants.StopURL)
	if err != nil {
		return fmt.Errorf("failed to send stop request: %v", err)
	}
	c.resp = resp
	return nil
}

func (c *checkServerFeature) serverProcessTerminates() error {
	if c.resp.StatusCode == 200 {
		return nil
	}
	return fmt.Errorf("server did not stop - status code: %d", c.resp.StatusCode)
}
