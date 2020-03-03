package actuator

import (
	"net/http"
	"net/http/pprof"
)

var HeapProfile http.Handler = pprof.Handler("heap")
