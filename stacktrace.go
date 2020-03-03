package actuator

import (
	"net/http"
	"net/http/pprof"
)

var StackTrace http.HandlerFunc = pprof.Handler("goroutine").ServeHTTP
