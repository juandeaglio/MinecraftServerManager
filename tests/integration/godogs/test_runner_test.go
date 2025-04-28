package test_runner

import (
	"minecraftremote/tests/integration/godogs/start_steps"
	"minecraftremote/tests/integration/godogs/status_steps"
	"testing"

	"github.com/cucumber/godog"
)

func TestServerStatusScenarios(t *testing.T) {
	suite := runScenario(t, status_steps.ServerStatusScenarioContext, "features/control_server.feature:3")
	if status := suite.Run(); status != 0 {
		t.Fatalf("Server status feature tests failed with status: %d", status)
	}
}

func TestServerStartScenarios(t *testing.T) {
	suite := runScenario(t, start_steps.ClientStartsServer, "features/control_server.feature:8")
	if status := suite.Run(); status != 0 {
		t.Fatalf("Server start feature tests failed with status: %d", status)
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
			Strict:   false,
			NoColors: false,   // Ensure colors are enabled for clarity
			Tags:     "~@wip", // Exclude work-in-progress tests
		},
	}
}

type ScenarioContextFunc func(*godog.ScenarioContext)
