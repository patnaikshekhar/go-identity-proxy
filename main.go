package main

import (
	"log"
	"net/http"
)

func main() {
	config, err := ReadConfig()
	if err != nil {
		log.Fatal(err)
	}

	startProxy(config)
}

func startProxy(config *Config) {

	http.HandleFunc("*", func(response http.ResponseWriter, request *http.Request) {
		// Validate the JWT
		authHeader := request.Header.Get("Authentication")
		if authHeader == "" {
			response.WriteHeader(403)
			response.Write([]byte("Forbidden"))
		}

		// Send Request to local machine
	})

	http.ListenAndServe(config.Address, nil)
}
