package main

import (
	"minecraftremote/src/server/mcservercontrols"
	"testing"

	"github.com/cucumber/godog"
)

type checkServerFeature struct {
	mc mcservercontrols.MinecraftServer
}

func (c *checkServerFeature) theServerIsStarted() error {
	c.mc = *mcservercontrols.NewServer()
	c.mc.Start()
	return nil
}

func (c *checkServerFeature) aClientAsksTheStatus() error {
	return godog.ErrPending
}

func (c *checkServerFeature) iShouldSeeAResponseFromTheServer() error {
	return godog.ErrPending
}

func TestFeatures(t *testing.T) {
	suite := runFeature(t, CheckServerRunningFeatureContext)
	if suite.Run() != 0 {
		t.Fatal("non-zero status returned, failed to run feature tests")
	}
}

func runFeature(t *testing.T, scenarioFeature func(*godog.ScenarioContext)) godog.TestSuite {
	suite := godog.TestSuite{
		ScenarioInitializer: func(s *godog.ScenarioContext) {
			scenarioFeature(s)
		},
		Options: &godog.Options{
			Format:   "pretty",
			Paths:    []string{"features"},
			TestingT: t,
			Strict:   true,
		},
	}
	return suite
}

func CheckServerRunningFeatureContext(s *godog.ScenarioContext) {
	c := &checkServerFeature{}
	s.Given(`the server is started`, c.theServerIsStarted)
	s.When(`a client asks the status`, c.aClientAsksTheStatus)
	s.Then(`I should see a response from the server`, c.iShouldSeeAResponseFromTheServer)
}
