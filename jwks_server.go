package main

import (
	"encoding/json"
	"net/http"
)

// JWKS structure
type JWKS struct {
	keys []Key `json:"keys"`
}

// Key structure
type Key struct {
	kty       string `json:"kty"`
	alg       string `json:"alg"`
	use       string `json:"use"`
	kid       string `json:"kid"`
	n         string `json:"n"`
	e         string `json:"e"`
}

// Sample JWKS data
var jwks = JWKS{
	keys: []Key{
		{
			kty: "RSA",
			alg: "RS256",
			use: "sig",
			kid: "example-key-id", //Replace
			n:   "your-modulus-here", // Replace with your RSA modulus
			e:   "AQAB",               
		},
	},
}

func jwksHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(jwks)
}

func main() {
	http.HandleFunc("/.well-known/jwks.json", jwksHandler)

	port := "8080"
	println("JWKS server is running on port " + port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		println("Failed to start server: ", err)
	}
}
