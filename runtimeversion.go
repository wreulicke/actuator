package actuator

import (
	"fmt"
	"net/http"
	"runtime"
)

func RuntimeVersion(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "%v\n", runtime.Version())
}
