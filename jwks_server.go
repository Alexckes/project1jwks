package main

import (
	"encoding/json"
	"net/http"
)

// JWKS structure
type JWKS struct {
	Keys []Key `json:"keys"`
}

// Key structure
type Key struct {
	Kty       string `json:"kty"`
	Alg       string `json:"alg"`
	Use       string `json:"use"`
	Kid       string `json:"kid"`
	N         string `json:"n"`
	E         string `json:"e"`
}

// Sample JWKS data
var jwks = JWKS{
	Keys: []Key{
		{
			Kty: "RSA",
			Alg: "RS256",
			Use: "sig",
			Kid: "example-key-id", //Replace
			N:   "your-modulus-here", // Replace with your RSA modulus
			E:   "AQAB",               
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
