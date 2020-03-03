package actuator

import (
	"net/http"
	"runtime/pprof"
)

func StackTrace(w http.ResponseWriter, req *http.Request) {
	pprof.Lookup("goroutine").WriteTo(w, 2)
}
