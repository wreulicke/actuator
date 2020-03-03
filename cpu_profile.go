package actuator

import (
	"net/http"
	"net/http/pprof"
)

var CpuProfile http.HandlerFunc = http.HandlerFunc(pprof.Profile).ServeHTTP
