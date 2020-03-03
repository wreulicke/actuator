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
	http.Handle("/stack", http.HandlerFunc(actuator.StackTrace))
	http.Handle("/trace", http.HandlerFunc(actuator.Trace))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
```