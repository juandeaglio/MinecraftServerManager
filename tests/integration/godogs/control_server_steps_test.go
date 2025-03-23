package main

import (
	"fmt"
	"log"
	"minecraftremote/src/controls/mcservercontrols"
	"minecraftremote/src/httprouter"
	"minecraftremote/src/httprouteradapter"
	"net/http"
	"testing"
	"time"

	"github.com/cucumber/godog"
)

type checkServerFeature struct {
	router *httprouter.HTTPServer
	server *http.Server
}

func (c *checkServerFeature) theServerIsStarted() error {
	controls := mcservercontrols.NewControls()
	controls.Start()
	c.router = httprouter.NewHTTPServer(controls)
	routerAdapter := &httprouteradapter.HTTPRouterAdapter{Router: c.router}
	c.server = startServerWithRouter(routerAdapter)
	return waitForServerReady("http://localhost:8080/status", 5*time.Second)
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

func startServerWithRouter(adapter *httprouteradapter.HTTPRouterAdapter) *http.Server {
	server := &http.Server{
		Addr:         ":8080",
		Handler:      adapter, // Use your router's HandleHTTP as the handler
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	// Start the server in a goroutine
	go func() {
		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Printf("HTTP server error: %v", err)
		}
	}()

	return server
}

func waitForServerReady(url string, timeout time.Duration) error {
	deadline := time.Now().Add(timeout)
	for time.Now().Before(deadline) {
		_, err := http.Get(url)
		if err == nil {
			return nil // Server is up and responding
		}
		time.Sleep(100 * time.Millisecond)
	}
	return fmt.Errorf("server failed to become ready within %v", timeout)
}
