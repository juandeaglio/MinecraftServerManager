package main

import (
	"minecraftremote/src/remoteconnection/stubremoteconnection"
	"minecraftremote/src/server/mcservercontrols"
	"testing"

	"github.com/cucumber/godog"
)

type checkServerFeature struct {
	mc mcservercontrols.MinecraftServer
}

func (c *checkServerFeature) theClientAsksTheServerIfItsThere() error {
	conn := stubremoteconnection.NewMockRemoteConnection(25565).IsAvailable()
	c.mc = *mcservercontrols.NewServer(conn)
	return nil
}

func (c *checkServerFeature) iReceiveTheClientsRequest() error {
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
	s.Given(`the client asks the server if it's there`, c.theClientAsksTheServerIfItsThere)
	s.When(`I receive the client's request`, c.iReceiveTheClientsRequest)
	s.Then(`I should see a response from the server`, c.iShouldSeeAResponseFromTheServer)
}
