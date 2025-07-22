package test_infrastructure

import (
	"fmt"
	"log"
	"minecraftremote/src/controls"
	"minecraftremote/src/httprouteradapter"
	"minecraftremote/src/os_api_adapter"
	"minecraftremote/tests/integration/godogs/constants"
	"net/http"
	"time"
)

// TestState holds shared state for tests
type TestState struct {
	Controls controls.Controls
	Server   http.Server
	Process  os_api_adapter.Process
}

func StartServerWithRouter(adapter *httprouteradapter.HTTPRouterAdapter, port string) *http.Server {
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

func CheckServerRunning(port string) error {
	checkURL := constants.BaseURL + port + constants.RunningURL
	resp, err := http.Get(checkURL)
	if err != nil {
		return fmt.Errorf("failed to get server running status: %v", err)
	}
	if resp.StatusCode == 200 {
		return nil
	}
	return fmt.Errorf("server is not running - status code: %d", resp.StatusCode)
}
