package process_steps

import (
	"context"
	"errors"
	"fmt"
	"github.com/cucumber/godog"
	"minecraftremote/src/os_api_adapter"
	"minecraftremote/src/os_api_adapter/real_os_ops"
	"minecraftremote/src/rcon"
	"minecraftremote/src/windowsconstants"
	"minecraftremote/tests/integration/godogs/test_infrastructure"
)

type stopProcessFeature struct {
	testContext *test_infrastructure.TestContext
	pid         int
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
	c := &stopProcessFeature{}

	osOps := &real_os_ops.RealOsOperations{}
	c.testContext = test_infrastructure.NewTestContext(
		rcon.NewStubRCONAdapter(),
		os_api_adapter.NewProcessHandler(osOps, "notepad.exe", ""),
	)

	s.Before(BeforeScenarioHook(c.testContext))

	s.Given(`^a process is running$`, c.processIsRunning)
	s.When(`^the process stops$`, c.stopProcess)
	s.Then(`^the process should not be running$`, c.processIsNotRunning)

	s.After(test_infrastructure.AfterScenarioHook(c.testContext))
}

func (c *stopProcessFeature) processIsRunning() error {
	c.pid = c.testContext.ProcessContext.PID()
	if c.pid > 0 {
		return nil
	}

	return errors.New("process is not running")
}

func (c *stopProcessFeature) stopProcess() error {
	err := c.testContext.ProcessContext.Stop()
	if err != nil {
		return err
	}
	return nil
}

func (c *stopProcessFeature) processIsNotRunning() error {
	ps, err := c.testContext.ProcessContext.GetProcessStatus(c.pid)

	if err != nil {
		return err
	}

	if ps.Status != windowsconstants.ParentKilledChildStatus {
		return fmt.Errorf("process did not exit, instead it shows exit code %d", ps.Status)
	}
	return nil
}
