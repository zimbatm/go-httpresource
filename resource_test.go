package httpresource_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/zimbatm/go-httpresource"
)

func TestHTTPResource(t *testing.T) {

	res := httpresource.GET(func(w http.ResponseWriter, r *http.Request) {
		return
	})

	fmt.Println(res)
}
