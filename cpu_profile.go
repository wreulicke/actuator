package actuator

import (
	"net/http"
	"runtime/pprof"
	"time"
)

func CpuProfile(w http.ResponseWriter, req *http.Request) {
	if err := pprof.StartCPUProfile(w); err != nil {
	}
	time.Sleep(30 * time.Second)
	pprof.StopCPUProfile()
}
