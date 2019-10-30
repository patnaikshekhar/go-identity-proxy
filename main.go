package main

import (
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

	http.HandleFunc("*", func(response http.ResponseWriter, request *http.Request) {

		// Get the bearer token from the header
		authHeader := request.Header.Get("Authentication")
		if authHeader == "" {
			response.WriteHeader(403)
			response.Write([]byte("Forbidden"))
			return
		}

		

		idToken, err := verifier.Verify(ctx, tokenString)
		if err != nil {
			log.Printf("Error: %s", err.Error())
			response.WriteHeader(403)
			response.Write([]byte("Forbidden"))
			return
		}

		log.Println(idToken.Audience)

		client := &http.Client{}
		resp, err := client.Do(request)

		body, _ := ioutil.ReadAll(resp.Body)
		response.Write(body)
	})

	err := http.ListenAndServe(config.Address, nil)
	if err != nil {
		log.Fatal(err)
	}
}
