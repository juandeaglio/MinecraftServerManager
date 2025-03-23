package main

import (
	"io"
	"log"
	"minecraftremote/src/controls/mcservercontrols"
	"minecraftremote/src/httprouter"
	"net/http"
	"testing"
	"time"

	"github.com/cucumber/godog"
)

type checkServerFeature struct {
	router *httprouter.HTTPServer
}

// First, create an adapter type
type HTTPHandlerAdapter struct {
	router *httprouter.HTTPServer
}

// Implement the http.Handler interface
func (a *HTTPHandlerAdapter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Call your existing HandleHTTP method
	response := a.router.HandleHTTP(r)

	// Transfer the response to the http.ResponseWriter
	for key, values := range response.Header {
		for _, value := range values {
			w.Header().Add(key, value)
		}
	}

	w.WriteHeader(response.StatusCode)

	// If there's a response body, write it
	if response.Body != nil {
		io.Copy(w, response.Body)
		response.Body.Close()
	}
}

func (c *checkServerFeature) theServerIsStarted() error {
	controls := mcservercontrols.NewControls()
	controls.Start()
	c.router = httprouter.NewHTTPServer(controls)
	adapter := &HTTPHandlerAdapter{router: c.router}
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
