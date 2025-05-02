package main

import (
	"minecraftremote/src/controls"
	"minecraftremote/src/httprouter"
	"minecraftremote/src/httprouteradapter"
	"minecraftremote/src/process"
	"minecraftremote/src/rcon"
	"minecraftremote/tests/integration/godogs/test_infrastructure"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	osOps := &process.WindowsOsOperations{}
	process := process.NewProcess(osOps, "notepad.exe", "")
	controls := controls.NewControls(rcon.NewStubRCONAdapter(), process)
	router := httprouter.NewHTTPRouter(controls, process)
	adapter := &httprouteradapter.HTTPRouterAdapter{Router: router}

	server := test_infrastructure.StartServerWithRouter(adapter, "8080")

	// Wait for interrupt signal
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	// Cleanup
	if server != nil {
		server.Close()
	}
	if process != nil {
		process.Stop()
	}
}
