
# goform

A standalone library that automatically polyfills, marshals and unmarshals data structures
from HTML Web Forms to a Backend written in go.


## Opinions / Features

- Embraces static and semantic HTML5 syntax.
- Embraces go generics and struct marshalling into JSON formats.
- Embraces `struct per route` approach, where each API endpoint marshals to a predefined struct.
- Isolates and compiles all code into a `WebASM` container.
- Unifies integration for multiple pages via `goform.Init(schemas *goform.Schemas)`.
- Marshals HTML forms with `action=/api/endpoint` and `type=application/json`.


## Mapping Schemas to Endpoints

```go
import "github.com/cookiengineer/goform"
import "myapp/schemas"

// Example WebASM integration
func main() {

    schemas := goform.NewSchema(map[string]any{
        "/api/user/new":  schemas.User,
        "/api/user/save": schemas.User,
        "/api/car/new":   schemas.Car,
    })

    goform.Init(&schemas)

}
```
