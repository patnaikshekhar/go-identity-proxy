package main

import (
	"context"
	"fmt"
	"log"

	"github.com/coreos/go-oidc"
)

// AuthChecker validates a JWT
type AuthChecker struct {
	verifier *oidc.IDTokenVerifier
}

// NewAuthChecker creates a new AuthChecker
func NewAuthChecker(config *Config) AuthChecker {
	ctx := context.Background()
	provider, err := oidc.NewProvider(ctx, config.Issuer)
	if err != nil {
		log.Println(err)

	}

	oidcConfig := oidc.Config{
		ClientID: config.ClientID,
	}

	verifier := provider.Verifier(&oidcConfig)

	return AuthChecker{verifier: verifier}
}

func getTokenFromAuthHeader(authHeader string) (string, error) {

	if len(authHeader) <= 6 {
		return "", fmt.Errorf("Invalid Authorization header")
	}

	if authHeader[:6] != "Bearer" {
		return "", fmt.Errorf("Missing bearer token in header")
	}

	tokenString := authHeader[7:]

	return tokenString, nil
}

func CheckToken(authHeader string) error {
	_, err := getTokenFromAuthHeader(authHeader)
	if err != nil {
		return err
	}

	return nil
}
