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
	c.connection = mockremoteconnection.NewMockRemoteConnection(c.portNum)
	c.portStatus = c.connection.IsAvailable()
	return nil
}

// Then step: assert that the process was found.
func (c *controlServerFeature) iShouldSeeAResponseFromTheServer() error {
	if c.portStatus {
		return nil
	}
	return errors.New("failed to query the port: server status is false")
}

func TestFeatures(t *testing.T) {
	suite := godog.TestSuite{
		ScenarioInitializer: CheckServerRunningFeature,
		Options: &godog.Options{
			Format:   "pretty",
			Paths:    []string{"features"},
			TestingT: t,
			Strict:   true,
		},
	}

	if suite.Run() != 0 {
		t.Fatal("non-zero status returned, failed to run feature tests")
	}
}

func CheckServerRunningFeature(s *godog.ScenarioContext) {
	c := &controlServerFeature{}
	s.Given(`the server is running RCON on port (\d+)`, c.theServerIsRunningOnPort)
	s.When(`^I query the port$`, c.iQueryThePort)
	s.Then(`^I should see that a response from the server$`, c.iShouldSeeAResponseFromTheServer)
}
