package actuator

import (
	"net/http"
	"net/http/pprof"
)

var Trace http.HandlerFunc = pprof.Trace
