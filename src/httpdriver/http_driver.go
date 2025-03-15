package httpdriver

import "server/src/httplib"

type HTTPDriver struct {
	httpLib httplib.HTTPLib
}

func NewHTTPDriver(httpLib httplib.HTTPLib) *HTTPDriver {
	return &HTTPDriver{httpLib: httpLib}
}
