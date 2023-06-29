package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"testing"
)

func TestSecureRandom(t *testing.T) {
	// Step 2: Generate random bytes
	randomBytes := make([]byte, 32) // 32 bytes = 256 bits
	_, err := rand.Read(randomBytes)
	if err != nil {
		panic(err)
	}

	// Step 3: Convert bytes to string
	apiKey := hex.EncodeToString(randomBytes)

	// Output the API key
	fmt.Println("Generated API Key:", apiKey)
}
