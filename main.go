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

	checker := NewAuthChecker(config)

	http.HandleFunc("*", func(response http.ResponseWriter, request *http.Request) {

		// Get the bearer token from the header
		authHeader := request.Header.Get("Authentication")

		err := checker.CheckToken(authHeader)
		if err != nil {
			response.WriteHeader(403)
			response.Write([]byte("Forbidden"))
		}

		// client := &http.Client{}
		// resp, _ := client.Do(request)

		// body, _ := ioutil.ReadAll(resp.Body)
		// response.Write(body)
		response.Write([]byte("Done"))
	})

	err := http.ListenAndServe(config.Address, nil)
	if err != nil {
		log.Fatal(err)
	}
}
