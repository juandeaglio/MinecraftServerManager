package control_server_steps

import (
	"fmt"
	"minecraftremote/src/os_api_adapter"
	"minecraftremote/src/os_api_adapter/real_os_ops"
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

func ServerStatusScenarioContext(s *godog.ScenarioContext) {
	osOps := &real_os_ops.RealOsOperations{}
	tc := test_infrastructure.NewTestContext(
		rcon.NewStubRCONAdapter(),
		os_api_adapter.NewProcessHandler(osOps, "notepad.exe", ""),
	)
	c := &StatusServerFeature{testContext: tc}

	baseHook := test_infrastructure.BeforeScenarioWithNotepadHook(tc, "8080")
	s.Before(baseHook)
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
