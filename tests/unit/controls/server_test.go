package servertest

import (
	"minecraftremote/src/server/mcservercontrols"
	"minecraftremote/tests/unit/httpdriver/cannedrequests"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestSuite struct holds the MinecraftServer instance
type TestSuite struct {
	controls *mcservercontrols.MinecraftServer
}

// setup initializes a fresh MinecraftServer instance for each test
func (ts *TestSuite) setup() {
	ts.controls = mcservercontrols.NewServer()
}

func TestMain(m *testing.M) {
	// Run all tests
	code := m.Run()
	os.Exit(code)
}

type TestCase struct {
	name string
	test func(t *testing.T, ts *TestSuite)
}

func TestMinecraftServer(t *testing.T) {
	// Create a single test suite instance
	suite := &TestSuite{}

	// Define test cases
	tests := []TestCase{
		{
			name: "StartServer",
			test: func(t *testing.T, ts *TestSuite) {
				resp := ts.controls.HandleHttp(cannedrequests.NewStartRequest().ToHTTPRequest())
				assert.Equal(t, 200, resp.StatusCode, "Server did not start successfully")
			},
		},
		{
			name: "StopServer",
			test: func(t *testing.T, ts *TestSuite) {
				ts.controls.HandleHttp(cannedrequests.NewStartRequest().ToHTTPRequest())
				resp := ts.controls.HandleHttp(cannedrequests.NewStopRequest().ToHTTPRequest())
				assert.Equal(t, 200, resp.StatusCode, "Server did not stop successfully, maybe it did not start?")
			},
		},
		{
			name: "ServerStatistics",
			test: func(t *testing.T, ts *TestSuite) {
				// Implement the test logic here
			},
		},
	}

	// Run each test case in a subtest
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			suite.setup() // Fresh instance for each test case
			tc.test(t, suite)
		})
	}
}
