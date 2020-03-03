package actuator

import (
	"net/http"
	"runtime"
)

func Gc(w http.ResponseWriter, req *http.Request) {
	runtime.GC()
	w.Write([]byte("ok"))
}
