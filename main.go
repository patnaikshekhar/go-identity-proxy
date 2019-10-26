package main

import (
	"context"
	"log"
	"net/http"

	"github.com/coreos/go-oidc"
)

func main() {
	config, err := ReadConfig()
	if err != nil {
		log.Fatal(err)
	}

	startProxy(config)
}

func startProxy(config *Config) {

	ctx := context.Background()
	provider, err := oidc.NewProvider(ctx, config.Issuer)
	if err != nil {
		log.Println(err)

	}

	oidcConfig := oidc.Config{
		ClientID: config.ClientID,
	}

	http.HandleFunc("*", func(response http.ResponseWriter, request *http.Request) {

		// Get the bearer token from the header
		authHeader := request.Header.Get("Authentication")
		if authHeader == "" {
			response.WriteHeader(403)
			response.Write([]byte("Forbidden"))
			return
		}

		if authHeader[:6] != "Bearer" {
			response.WriteHeader(403)
			response.Write([]byte("Forbidden"))
			return
		}

		tokenString := authHeader[7:]
		verifier := provider.Verifier(&oidcConfig)

		verifier.Verify(ctx, tokenString)

		// tokenString := authHeader[7:]
		// // Test JWT
		// token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		// 	// Don't forget to validate the alg is what you expect:
		// 	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		// 		return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		// 	}

		// 	// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		// 	return hmacSampleSecret, nil
		// })

		// if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// 	// Send HTTP Request to destination
		// 	response.Write([]byte("Done"))
		// } else {
		// 	response.WriteHeader(403)
		// 	response.Write([]byte("Forbidden"))
		// }

	})

	err := http.ListenAndServe(config.Address, nil)
	if err != nil {
		log.Fatal(err)
	}
}
