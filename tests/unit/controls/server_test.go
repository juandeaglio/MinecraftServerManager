package servertest

import (
	"minecraftremote/src/server/mcservercontrols"
	"net/http"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

type StatusRequest struct {
	*http.Request
}

func NewStatusRequest() *StatusRequest {
	url, _ := url.Parse("http://localhost/status")
	req := &http.Request{
		Method: "GET",
		URL:    url,
		Header: make(http.Header),
	}
	return &StatusRequest{Request: req}
}

// ToHTTPRequest extracts the underlying *http.Request
func (s *StatusRequest) ToHTTPRequest() *http.Request {
	return s.Request
}

func TestStartServer(t *testing.T) {
	controls := mcservercontrols.NewServer()
	assert.Truef(t, controls.HandleHttp(NewStatusRequest().ToHTTPRequest()).StatusCode == 200, "Server did not start successfully")
}

// func TestStopServer(t *testing.T) {
// 	controls := mcservercontrols.NewServer()
// 	controls.Start()
// 	assert.Truef(t, controls.Start().StatusCode == 200, "Server did not start successfully")
// }
