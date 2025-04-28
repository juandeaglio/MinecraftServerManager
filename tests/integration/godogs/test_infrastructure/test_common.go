package test_infrastructure

import (
	"log"
	"minecraftremote/src/controls"
	"minecraftremote/src/httprouteradapter"
	"minecraftremote/src/process"
	"net/http"
	"time"
)

// TestState holds shared state for tests
type TestState struct {
	Controls controls.Controls
	Server   http.Server
	Process  process.Process
}

func startServerWithRouter(adapter *httprouteradapter.HTTPRouterAdapter, port string) *http.Server {
	server := &http.Server{
		Addr:         ":" + port,
		Handler:      adapter, // Use your router's HandleHTTP as the handler
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
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
