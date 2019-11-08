package main

import (
	"context"
	"fmt"
	"io/ioutil"
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

	checker := NewAuthChecker(config)

	http.HandleFunc("/", func(response http.ResponseWriter, request *http.Request) {

		// Get the bearer token from the header
		authHeader := request.Header.Get("Authorization")

		ctx := context.Background()

		err := checker.CheckToken(ctx, authHeader)
		if err != nil {
			log.Printf("Error in request: %s", err.Error())
			response.WriteHeader(403)
			response.Write([]byte("Forbidden"))
			return
		}

		client := &http.Client{}

		backendURL := fmt.Sprintf("http://%s%s", config.Backend, request.URL.Path)

		backendRequest, err := http.NewRequest(request.Method, backendURL, request.Body)
		if err != nil {
			response.WriteHeader(500)
			response.Write([]byte(err.Error()))
			return
		}

		resp, err := client.Do(backendRequest)
		if err != nil {
			response.WriteHeader(500)
			response.Write([]byte(err.Error()))
			return
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			response.WriteHeader(500)
			response.Write([]byte(err.Error()))
			return
		}

		response.Write(body)

		//response.Write([]byte("Done"))
	})

	log.Printf("Starting HTTP server on address %s", config.Address)
	err := http.ListenAndServe(config.Address, nil)
	if err != nil {
		log.Fatal(err)
	}
}
