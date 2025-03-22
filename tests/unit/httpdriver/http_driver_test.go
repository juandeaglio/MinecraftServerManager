package httpdrivertests

import (
	"minecraftremote/src/httpdriver"
	"minecraftremote/src/httplib"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDefineHTTPDriver(t *testing.T) {
	mockClient := httplib.NewMockHTTPLib()
	driver := httpdriver.NewHTTPDriver(mockClient)
	require.NotNil(t, driver)
}
