package utils

import (
	"math/rand"
	"strings"
	"time"
)

// SeedRandom seeds the random number generator.
func SeedRandom() {
	rand.Seed(time.Now().UnixNano())
}

// RandomInt returns a random integer between min and max.
func RandomInt(min, max int) int {
	return rand.Intn(max-min) + min
}

// RandomString generates a random string of the specified length.
func RandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var result strings.Builder
	for i := 0; i < length; i++ {
		result.WriteByte(charset[rand.Intn(len(charset))])
	}
	return result.String()
}

// ToUpper converts a rune to uppercase if it's a lowercase letter.
func ToUpper(r rune) rune {
	if r >= 'a' && r <= 'z' {
		return r - ('a' - 'A')
	}
	return r
}

// ToLower converts a rune to lowercase if it's an uppercase letter.
func ToLower(r rune) rune {
	if r >= 'A' && r <= 'Z' {
		return r + ('a' - 'A')
	}
	return r
}
