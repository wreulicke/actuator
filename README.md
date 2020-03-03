## actuator

HTTP interface like gops.

## Usage

```go
import (
	"github.com/wreulicke/actuator"
)

func main() {
	http.Handle("/gc", http.HandlerFunc(actuator.Gc))
	http.Handle("/version", http.HandlerFunc(actuator.RuntimeVersion))
	http.Handle("/stack", actuator.StackTrace)
	http.Handle("/trace", http.HandlerFunc(actuator.Trace))
	http.Handle("/cpu_profile", actuator.CpuProfile)
	http.Handle("/heap_profile", actuator.HeapProfile)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
```