package utils

import (
	"encoding/hex"
	"math/rand"
)

// GenerateID generates a random unique id.
func GenerateID() string {
	b := make([]byte, 32)
	rand.Read(b)
	return hex.EncodeToString(b)
}
