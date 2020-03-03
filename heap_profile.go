package actuator

import (
	"net/http"
	"net/http/pprof"
)

var HeapProfile http.HandlerFunc = pprof.Handler("heap").ServeHTTP
