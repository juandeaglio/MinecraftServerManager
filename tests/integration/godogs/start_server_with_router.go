package integrationtest

import (
	"fmt"
	"log"
	"minecraftremote/src/httprouteradapter"
	"net/http"
	"time"
)

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
