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
