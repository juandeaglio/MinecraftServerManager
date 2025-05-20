package process_steps

import (
	"context"
	"github.com/cucumber/godog"
	"minecraftremote/src/process_context"
	"minecraftremote/src/rcon"
	"minecraftremote/tests/integration/godogs/test_infrastructure"
)

type stopProcessFeature struct {
	testContext *test_infrastructure.TestContext
}

func BeforeScenarioHook(tc *test_infrastructure.TestContext) func(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
	return func(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
		err := tc.ProcessContext.Start()
		if err != nil {
			return nil, err
		}

		return ctx, nil
	}
}

func StopProcess(s *godog.ScenarioContext) {
	c := &startProcessFeature{}
	osOps := &process_context.WindowsOsOperations{}
	c.testContext = test_infrastructure.NewTestContext(
		rcon.NewStubRCONAdapter(),
		process_context.NewProcessInvoker(osOps, "notepad.exe", ""),
	)

	s.Before(BeforeScenarioHook(c.testContext))

	s.Given(`^a process is running$`, c.processIsRunning)
	s.When(`^the process stops$`, c.stopProcess)
	s.Then(`^the process should not be running$`, c.processIsNotRunning)

	s.After(test_infrastructure.AfterScenarioHook(c.testContext))
}

func (c *startProcessFeature) stopProcess() error {
	err := c.testContext.ProcessContext.Start()
	if err != nil {
		return err
	}
	return nil
}
