package test_runner

import (
	"minecraftremote/tests/integration/godogs/controlplane/control_steps"
	"minecraftremote/tests/integration/godogs/controlplane/server_steps"
	"minecraftremote/tests/integration/godogs/process_handler/process_steps"
	"testing"

	"github.com/cucumber/godog"
)

func TestServerStatusScenarios(t *testing.T) {
	suite := runScenario(t, control_server_steps.ServerStatusScenarioContext, "controlplane/control_steps/control_server.feature:3")
	if status := suite.Run(); status != 0 {
		t.Fatalf("Server status feature tests failed with status: %d", status)
	}
}

func TestServerStartScenarios(t *testing.T) {
	suite := runScenario(t, control_server_steps.ClientStartsServer, "controlplane/control_steps/control_server.feature:8")
	if status := suite.Run(); status != 0 {
		t.Fatalf("Server start feature tests failed with status: %d", status)
	}
}

func TestServerStopScenarios(t *testing.T) {
	suite := runScenario(t, control_server_steps.ClientStopsServer, "controlplane/control_steps/control_server.feature:13")
	if status := suite.Run(); status != 0 {
		t.Fatalf("Server stop feature tests failed with status: %d", status)
	}
}

func TestStartServer(t *testing.T) {
	suite := runScenario(t, server_steps.StartServer, "controlplane/server_steps/server.feature:3")
	if status := suite.Run(); status != 0 {
		t.Fatalf("Server start feature tests failed with status: %d", status)
	}
}

func TestWindowsOpsStart(t *testing.T) {
	suite := runScenario(t, process_steps.StartProcess, "process_handler/process_steps/windows_ops.feature:3")
	if status := suite.Run(); status != 0 {
		t.Fatalf("ProcessContext start test failed with status: %d", status)
	}

}

func TestWindowsOpsStop(t *testing.T) {
	suite := runScenario(t, process_steps.StopProcess, "process_handler/process_steps/windows_ops.feature:8")
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
