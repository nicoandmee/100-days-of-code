package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func main() {
	secret := "supersecret"
	data := "data"
	fmt.Printf("Secret: %s Data: %s\n", secret, data)

	// Init a new HMAC, defining the hash type and the key as a byte array
	h := hmac.New(sha256.New, []byte(secret))

	// Write some secret data
	h.Write([]byte(data))

	// Get result and encode
	sha := hex.EncodeToString(h.Sum(nil))

	fmt.Println("Result: " + sha)
}
