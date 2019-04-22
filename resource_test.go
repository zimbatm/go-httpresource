package httpresource_test

import (
	"testing"
	"net/http"
	"fmt"

	"github.com/zimbatm/go-httpresource"
)


func TestHTTPResource(t *testing.T) {

	res := httpresource.GET(func (w http.ResponseWriter, r *http.Request) {
		return
	})

	fmt.Println(res)
}
