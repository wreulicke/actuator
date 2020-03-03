package actuator

import (
	"net/http"
	"runtime/trace"
	"time"
)

func Trace(w http.ResponseWriter, req *http.Request) {
	trace.Start(w)
	time.Sleep(5 * time.Second)
	trace.Stop()
}
