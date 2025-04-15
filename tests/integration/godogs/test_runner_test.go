package integrationtest

import (
	"testing"

	"github.com/cucumber/godog"
)

func TestScenariosWithStartedServer(t *testing.T) {

	// Combine both scenario contexts into one initializer
	combinedScenarioInitializer := func(s *godog.ScenarioContext) {
		ClientAsksTheServerForTheStatusScenarioContext(s)
		ClientStartsServer(s)
	}

	suite := runScenario(t, combinedScenarioInitializer)

	if status := suite.Run(); status != 0 {
		t.Fatalf("Feature tests failed with status: %d", status)
	}
}

func runScenario(t *testing.T, scenarioFeature ScenarioContextFunc) godog.TestSuite {
	return godog.TestSuite{
		ScenarioInitializer: func(s *godog.ScenarioContext) {
			scenarioFeature(s)
		},
		Options: &godog.Options{
			Format:   "pretty",
			Paths:    []string{"features"},
			TestingT: t,
			Strict:   false,
			NoColors: false,   // Ensure colors are enabled for clarity
			Tags:     "~@wip", // Exclude work-in-progress tests
		},
	}
}

type ScenarioContextFunc func(*godog.ScenarioContext)
