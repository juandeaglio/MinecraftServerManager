package httprouteradapter

import (
	"io"
	"minecraftremote/src/httprouter"
	"net/http"
)

type HTTPRouterAdapter struct {
	Router *httprouter.ServerRouter
}

func (a *HTTPRouterAdapter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Call your existing HandleHTTP method
	response := a.Router.HandleHTTP(r)

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
