# go-httpresource - simple HTTP end-point

The go stdlib `net/http` library is quite good for providing a simple routing
experience but some times you want just a bit more. This library is designed
to work in conjuction and add a dispatch on HTTP methods to the mix. Some
times it's nice to have APIs being a bit more REST-ful-y.

## Example

[$ examples/simple.go](examples/simple.go)

```go
package main

import (
	"net/http"
	"github.com/zimbatm/go-httpresource"
)

func main() {
	// Only reply to GET and HEAD requests 
	http.Handle("/hello", httpresource.GET(func (w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	}))

	// A more complex example
	http.Handle("/users", httpresource.Resource {
		GET: http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
			// Return the list of all users
		}),
		POST: http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
			// Create a new user
		}),
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
```

## License

MIT (c) 2019 - zimbatm and contributors
