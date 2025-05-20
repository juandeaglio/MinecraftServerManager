package process_steps

import (
	"errors"
	"github.com/cucumber/godog"
	"minecraftremote/src/process_context"
	"minecraftremote/src/rcon"
	"minecraftremote/tests/integration/godogs/test_infrastructure"
)

type startProcessFeature struct {
	testContext *test_infrastructure.TestContext
}

func StartProcess(s *godog.ScenarioContext) {
	c := &startProcessFeature{}
	osOps := &process_context.WindowsOsOperations{}
	c.testContext = test_infrastructure.NewTestContext(
		rcon.NewStubRCONAdapter(),
		process_context.NewProcessInvoker(osOps, "notepad.exe", ""),
	)

	s.Given(`^a process is not running$`, c.processIsNotRunning)
	s.When(`^the process starts$`, c.processStarts)
	s.Then(`^the process should be running$`, c.processIsRunning)

	s.After(test_infrastructure.AfterScenarioHook(c.testContext))
}

func (c *startProcessFeature) processIsNotRunning() error {
	// pid := c.testContext.ProcessContext.PID()
	// processStatus := c.testContext.ProcessContext.GetProcessStatus(pid)
	// if processStatus.Status == "Running" && processStatus.User != "" {
	if c.testContext.ProcessContext.PID() > 0 {
		return errors.New("process is running")
	}
	return nil
}

func (c *startProcessFeature) processStarts() error {
	err := c.testContext.ProcessContext.Start()
	if err != nil {
		return err
	}
	return nil
}

func (c *startProcessFeature) processIsRunning() error {
	if c.testContext.ProcessContext.PID() > 0 {
		return nil
	}
	return errors.New("process is not running")
}
