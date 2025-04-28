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
	response := a.Router.HandleHTTP(r)

	for key, values := range response.Header {
		for _, value := range values {
			w.Header().Add(key, value)
		}
	}
	w.WriteHeader(response.StatusCode)

	if response.Body != nil {
		io.Copy(w, response.Body)
		response.Body.Close()
	}
}
