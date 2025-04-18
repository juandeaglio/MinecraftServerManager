package integrationtest

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"minecraftremote/src/process"
	"minecraftremote/src/rcon"
	"minecraftremote/tests/integration/godogs/constants"
	"net/http"
	"strings"

	"github.com/cucumber/godog"
)

type checkServerFeature struct {
	testContext *TestContext
	resp        *http.Response
}

func (c *checkServerFeature) theMinecraftProcessIsRunning() error {
	c.resp, _ = http.Get(constants.StatusURL)
	if c.resp.StatusCode == 200 {
		return nil
	}
	return fmt.Errorf("server is not running - status code: %d", c.resp.StatusCode)
}

func (c *checkServerFeature) aClientRequestsMinecraftProcessStatus() error {
	c.resp, _ = http.Get(constants.StatusURL)
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

	players := response["Players"]

	playersValue, _ := players.(float64) // JSON numbers are parsed as float64 by default

	if playersValue < 0 {
		return fmt.Errorf("expected 'Players' value to be 0 but got %v", playersValue)
	}

	return nil
}

func ClientAsksTheServerForTheStatusScenarioContext(s *godog.ScenarioContext) {
	tc := NewTestContext(rcon.NewStubRCONAdapter())
	c := &checkServerFeature{testContext: tc}

	// Register hooks with common infrastructure
	baseHook := BeforeScenarioWithNotepadHook(tc, "8080")
	customHook := func(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
		tc.Process = tc.Controls.Start(process.NewProcess(&process.FakeOsOperations{}, "notepad.exe"))
		return ctx, nil
	}
	s.Before(CombineBeforeHooks(baseHook, customHook))

	s.After(AfterScenarioHook(tc))

	// Register step definitions
	s.Given(`^the Minecraft server is running and ready$`, c.theMinecraftProcessIsRunning)
	s.When(`^a client requests the server status$`, c.aClientRequestsMinecraftProcessStatus)
	s.Then(`^the system returns a status response indicating "online" along with the current player count$`, c.theAPIReturnsMinecraftProcessStatusWithPlayerCount)
}
