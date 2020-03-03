package actuator

import (
	"net/http"
	"net/http/pprof"
)

var StackTrace http.Handler = pprof.Handler("goroutine")
