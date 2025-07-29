package process_steps

import (
	"errors"
	"github.com/cucumber/godog"
	"minecraftremote/src/os_api_adapter"
	"minecraftremote/src/os_api_adapter/real_os_ops"
	"minecraftremote/src/rcon"
	"minecraftremote/src/windowsconstants"
	"minecraftremote/tests/integration/godogs/test_infrastructure"
)

type startProcessFeature struct {
	testContext *test_infrastructure.TestContext
}

func StartProcess(s *godog.ScenarioContext) {
	c := &startProcessFeature{}
	osOps := &real_os_ops.RealOsOperations{}
	c.testContext = test_infrastructure.NewTestContext(
		rcon.NewStubRCONAdapter(),
		os_api_adapter.NewProcessHandler(osOps, "notepad.exe", ""),
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
	pid := c.testContext.ProcessContext.PID()
	ps, err := c.testContext.ProcessContext.GetProcessStatus(pid)
	if err != nil {
		return err
	}

	if pid > 0 && ps.Status != windowsconstants.InvalidProcessStatus {
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
