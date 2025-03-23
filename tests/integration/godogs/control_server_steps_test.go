package main

import (
	"minecraftremote/src/controls/mcservercontrols"
	"minecraftremote/src/httprouter"
	"testing"

	"github.com/cucumber/godog"
)

type checkServerFeature struct {
	router *httprouter.HTTPServer
}

func (c *checkServerFeature) theServerIsStarted() error {
	controls := mcservercontrols.NewControls()
	controls.Start()
	c.router = httprouter.NewHTTPServer(controls)

	return nil
}

func (c *checkServerFeature) aClientAsksTheStatus() error {

	return godog.ErrPending
}

func (c *checkServerFeature) iShouldTellTheClientTheStatus() error {
	return godog.ErrPending
}

func TestFeatures(t *testing.T) {
	suite := runFeature(t, ClientAsksTheServerForTheStatusFeatureContext)
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

func ClientAsksTheServerForTheStatusFeatureContext(s *godog.ScenarioContext) {
	c := &checkServerFeature{}
	s.Given(`the server is started`, c.theServerIsStarted)
	s.When(`a client asks the status`, c.aClientAsksTheStatus)
	s.Then(`I should tell the client the status`, c.iShouldTellTheClientTheStatus)
}
