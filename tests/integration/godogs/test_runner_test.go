package test_runner

import (
	"minecraftremote/tests/integration/godogs/control_steps"
	"minecraftremote/tests/integration/godogs/process_steps"
	"minecraftremote/tests/integration/godogs/server_steps"
	"testing"

	"github.com/cucumber/godog"
)

func TestServerStatusScenarios(t *testing.T) {
	suite := runScenario(t, control_server_steps.ServerStatusScenarioContext, "control_steps/control_server.feature:3")
	if status := suite.Run(); status != 0 {
		t.Fatalf("Server status feature tests failed with status: %d", status)
	}
}

func TestServerStartScenarios(t *testing.T) {
	suite := runScenario(t, control_server_steps.ClientStartsServer, "control_steps/control_server.feature:8")
	if status := suite.Run(); status != 0 {
		t.Fatalf("Server start feature tests failed with status: %d", status)
	}
}

func TestServerStopScenarios(t *testing.T) {
	suite := runScenario(t, control_server_steps.ClientStopsServer, "control_steps/control_server.feature:13")
	if status := suite.Run(); status != 0 {
		t.Fatalf("Server stop feature tests failed with status: %d", status)
	}
}

func TestStartServer(t *testing.T) {
	suite := runScenario(t, server_steps.StartServer, "server_steps/server.feature:3")
	if status := suite.Run(); status != 0 {
		t.Fatalf("Server start feature tests failed with status: %d", status)
	}
}

func TestWindowsOps(t *testing.T) {
	suite := runScenario(t, process_steps.StartProcess, "process_steps/windows_ops.feature:3")
	if status := suite.Run(); status != 0 {
		t.Fatalf("ProcessContext start test failed with status: %d", status)
	}

	suite = runScenario(t, process_steps.StopProcess, "process_steps/windows_ops.feature:8")
	if status := suite.Run(); status != 0 {
		t.Fatalf("ProcessContext stop test failed with status: %d", status)
	}
}

func runScenario(t *testing.T, scenarioFeature ScenarioContextFunc, featurePath string) godog.TestSuite {
	return godog.TestSuite{
		ScenarioInitializer: func(s *godog.ScenarioContext) {
			scenarioFeature(s)
		},
		Options: &godog.Options{
			Format:   "pretty",
			Paths:    []string{featurePath},
			TestingT: t,
			Strict:   true,
			NoColors: false,   // Ensure colors are enabled for clarity
			Tags:     "~@wip", // Exclude work-in-progress tests
		},
	}
}

type ScenarioContextFunc func(*godog.ScenarioContext)
