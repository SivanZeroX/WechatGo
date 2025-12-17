package crypto

import (
	"regexp"
	"testing"
)

func TestRandomString(t *testing.T) {
	length := 16
	result := RandomString(length)

	// Check length
	if len(result) != length {
		t.Fatalf("Expected length %d, got %d", length, len(result))
	}

	// Check if it contains only alphanumeric characters
	matched, err := regexp.MatchString("^[a-zA-Z0-9]+$", result)
	if err != nil {
		t.Fatalf("Regex match failed: %v", err)
	}
	if !matched {
		t.Fatalf("Random string contains invalid characters: %s", result)
	}

	// Check if different calls produce different results
	result2 := RandomString(length)
	if result == result2 {
		t.Fatalf("Two different calls produced the same result: %s", result)
	}
}

func TestRandomString_DifferentLengths(t *testing.T) {
	lengths := []int{8, 16, 32, 64}

	for _, length := range lengths {
		result := RandomString(length)
		if len(result) != length {
			t.Fatalf("Expected length %d, got %d", length, len(result))
		}
	}
}

func TestRandomString_EmptyLength(t *testing.T) {
	result := RandomString(0)
	if result != "" {
		t.Fatalf("Expected empty string for length 0, got %s", result)
	}
}
