package main

import (
	"log"
	"net/http"

	handlers "github.com/writeameer/httphandlers/handlers"
)

func main() {

	// Create new router
	mux := http.NewServeMux()

	originHost := "google.com"
	mux.Handle("/", handlers.AuthenticationHandler(handlers.ReverseProxyHandler(originHost)))

	// Listen and Server
	port := ":8080"
	log.Println("Server started on port" + port)
	log.Fatal(http.ListenAndServe(port, mux))
}
