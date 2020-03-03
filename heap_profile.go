package actuator

import (
	"net/http"
	"runtime/pprof"
)

func HeapProfile(w http.ResponseWriter, req *http.Request) {
	pprof.WriteHeapProfile(w)
}
