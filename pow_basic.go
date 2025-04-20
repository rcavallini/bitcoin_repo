package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func main() {
	prefix := "0000" // The prefix we want to find, increased from 00 to 0000 (16x harder)
	base := "hello"
	nonce := 0
	var hash string

	start := time.Now()

	for {
		input := base + strconv.Itoa(nonce)
		sum := sha256.Sum256([]byte(input))
		hash = hex.EncodeToString(sum[:])

		if strings.HasPrefix(hash, prefix) {
			duration := time.Since(start)
			fmt.Println("✅ Success!")
			fmt.Printf("Nonce: %d\n", nonce)
			fmt.Printf("Input: %s\n", input)
			fmt.Printf("Hash:  %s\n", hash)
			fmt.Printf("⏱️ Time: %.4f seconds\n", duration.Seconds())
			break
		}

		nonce++
	}
}
