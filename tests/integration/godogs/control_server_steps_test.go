package main

import (
	"fmt"
	"minecraftremote/src/remoteconnection"
	"minecraftremote/src/remoteconnection/mockremoteconnection"
	"testing"

	"github.com/cucumber/godog"
)

type checkServerFeature struct {
	portStatus bool
	portNum    int
	connection remoteconnection.RemoteConnection
}

// Given step: Define the port where the server is running.
func (c *checkServerFeature) theServerIsRunningOnPort(port int) error {
	c.portNum = port
	return nil
}

// When step: poll for the process.
func (c *checkServerFeature) iQueryThePort() error {
	c.connection = mockremoteconnection.NewMockRemoteConnection(c.portNum)
	c.portStatus = c.connection.IsAvailable()
	return nil
}

// Then step: assert that the process was found.
func (c *checkServerFeature) iShouldSeeAResponseFromTheServer() error {
	if c.portStatus {
		return nil
	}
	return fmt.Errorf("server on port %d is unavailable", c.portNum)
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
	s.Given(`the server is running RCON on port (\d+)`, c.theServerIsRunningOnPort)
	s.When(`^I query the port$`, c.iQueryThePort)
	s.Then(`^I should see a response from the server$`, c.iShouldSeeAResponseFromTheServer)
}
