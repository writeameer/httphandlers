package handlers

import (
	"fmt"
	"net/http"
	"time"
)

// TimeHandler outputs time
func TimeHandler(format string) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tm := time.Now().Format(format)
		fmt.Fprintf(w, "The time is:"+tm)
	})

}
