package main

import (
	"errors"
	"minecraftremote/src/remoteconnection"
	"minecraftremote/src/remoteconnection/mockremoteconnection"
	"testing"

	"github.com/cucumber/godog"
)

// processFeature acts as our test fixture holding any shared state.
type controlServerFeature struct {
	portStatus bool
	portNum    int
	connection remoteconnection.RemoteConnection
}

// Given step: set up the process name.
func (c *controlServerFeature) theServerIsRunningOnPort(port int) error {
	c.portNum = port
	return nil
}

// When step: poll for the process.
func (c *controlServerFeature) iQueryThePort() error {
	c.connection = mockremoteconnection.NewMockRemoteConnection(25565)
	c.portStatus = c.connection.IsAvailable()
	if c.portStatus != false {
		return nil
	}
	return errors.New("failed to query the port: server status is false")
}

// Then step: assert that the process was found.
func (c *controlServerFeature) iShouldSeeThatAResponseFromTheServer() error {
	return nil
}

func TestFeatures(t *testing.T) {
	suite := godog.TestSuite{
		ScenarioInitializer: FeatureContext,
		Options: &godog.Options{
			Format:   "pretty",
			Paths:    []string{"features"},
			TestingT: t, // Testing instance that will run subtests.
		},
	}

	if suite.Run() != 0 {
		t.Fatal("non-zero status returned, failed to run feature tests")
	}
}

// FeatureContext registers the step definitions with Godog.
func FeatureContext(s *godog.ScenarioContext) {
	c := &controlServerFeature{}

	// Register each of the steps defined in the feature file
	s.Given(`the server is running RCON on port (\d+)`, c.theServerIsRunningOnPort)
	s.When(`^I query the port$`, c.iQueryThePort)
	s.Step(`^I should see that a response from the server$`, c.iShouldSeeThatAResponseFromTheServer)
}
