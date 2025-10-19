package main

import (
	"minecraftremote/src/controls"
	"minecraftremote/src/httprouter"
	"minecraftremote/src/httprouteradapter"
	"minecraftremote/src/os_api_adapter"
	"minecraftremote/src/os_api_adapter/real_os_ops"
	"minecraftremote/src/rcon"
	"minecraftremote/tests/integration/godogs/test_infrastructure"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	osOps := &real_os_ops.RealOsOperations{}
	newProcess := os_api_adapter.NewProcessHandler(osOps, `cmd.exe`, "")
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
		tryCloseServer(server)
	}
	if newProcess != nil {
		tryStopProcess(newProcess)
	}
}

func tryStopProcess(newProcess *os_api_adapter.ProcessImpl) bool {
	err := newProcess.Stop()
	return err != nil
}

func tryCloseServer(server *http.Server) bool {
	err := server.Close()
	return err != nil
}
