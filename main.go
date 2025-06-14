package main

import (
	"minecraftremote/src/controls"
	"minecraftremote/src/httprouter"
	"minecraftremote/src/httprouteradapter"
	"minecraftremote/src/process_context"
	"minecraftremote/src/rcon"
	"minecraftremote/tests/integration/godogs/test_infrastructure"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	osOps := &process_context.WindowsOsOperations{}
	newProcess := process_context.NewProcessInvoker(osOps, "notepad.exe", "")
	mcControls := controls.NewControls(rcon.NewStubRCONAdapter(), newProcess)
	router := httprouter.NewHTTPRouter(mcControls, newProcess)
	adapter := &httprouteradapter.HTTPRouterAdapter{Router: router}

	server := test_infrastructure.StartServerWithRouter(adapter, "8080")

	// Wait for interrupt signal
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	// Cleanup
	if server != nil {
		err := server.Close()
		if err != nil {
			return
		}
	}
	if newProcess != nil {
		err := newProcess.Stop()
		if err != nil {
			return
		}
	}
}
