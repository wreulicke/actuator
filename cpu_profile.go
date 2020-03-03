package actuator

import (
	"net/http"
	"net/http/pprof"
)

var CpuProfile http.Handler = http.HandlerFunc(pprof.Profile)
