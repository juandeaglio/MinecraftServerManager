package httpdrivertests

import (
	"minecraftremote/src/httpdriver"
	"minecraftremote/src/httplib"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDefineHTTPDriver(t *testing.T) {
	mockClient := httplib.NewMockHTTPLib()
	driver := httpdriver.NewHTTPDriver(mockClient)
	assert.NotNil(t, driver)
}
