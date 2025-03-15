package httpdrivertests

import (
	"server/src/httpdriver"
	"server/src/httplib"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDefineHTTPDriver(t *testing.T) {
	mockClient := httplib.NewMockHTTPLib()
	driver := httpdriver.NewHTTPDriver(mockClient)
	assert.NotNil(t, driver)
}
