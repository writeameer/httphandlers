package handlers

import (
	"fmt"
	"net/http"
)

// HelloWorldHandler says hello world
func HelloWorldHandler(next ...http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello ,  %s! \n", r.URL.Path[1:])

		if next != nil {
			next[0].ServeHTTP(w, r)
		}
	})
}
