package session

import (
	"testing"
	"time"
)

func TestMemoryStorage_SetAndGet(t *testing.T) {
	storage := NewMemoryStorage()

	// Test Set and Get
	err := storage.Set("key1", "value1", 0)
	if err != nil {
		t.Fatalf("Set failed: %v", err)
	}

	value, err := storage.Get("key1")
	if err != nil {
		t.Fatalf("Get failed: %v", err)
	}
	if value != "value1" {
		t.Fatalf("Expected value1, got %s", value)
	}
}

func TestMemoryStorage_Delete(t *testing.T) {
	storage := NewMemoryStorage()

	// Set a value
	storage.Set("key1", "value1", 0)

	// Delete it
	err := storage.Delete("key1")
	if err != nil {
		t.Fatalf("Delete failed: %v", err)
	}

	// Try to get it (should return empty)
	value, err := storage.Get("key1")
	if err != nil {
		t.Fatalf("Get failed: %v", err)
	}
	if value != "" {
		t.Fatalf("Expected empty string after delete, got %s", value)
	}
}

func TestMemoryStorage_TTL(t *testing.T) {
	storage := NewMemoryStorage()

	// Set a value with short TTL
	ttl := 100 * time.Millisecond
	storage.Set("key1", "value1", ttl)

	// Get it immediately (should exist)
	value, err := storage.Get("key1")
	if err != nil {
		t.Fatalf("Get failed: %v", err)
	}
	if value != "value1" {
		t.Fatalf("Expected value1, got %s", value)
	}

	// Wait for expiration
	time.Sleep(150 * time.Millisecond)

	// Try to get it (should be expired)
	value, err = storage.Get("key1")
	if err != nil {
		t.Fatalf("Get failed: %v", err)
	}
	if value != "" {
		t.Fatalf("Expected empty string after expiration, got %s", value)
	}
}

func TestMemoryStorage_EmptyValue(t *testing.T) {
	storage := NewMemoryStorage()

	// Set empty value (should be ignored)
	err := storage.Set("key1", "", 0)
	if err != nil {
		t.Fatalf("Set failed: %v", err)
	}

	// Get should return empty
	value, err := storage.Get("key1")
	if err != nil {
		t.Fatalf("Get failed: %v", err)
	}
	if value != "" {
		t.Fatalf("Expected empty string, got %s", value)
	}
}

func TestMemoryStorage_NonExistentKey(t *testing.T) {
	storage := NewMemoryStorage()

	// Get non-existent key
	value, err := storage.Get("nonExistent")
	if err != nil {
		t.Fatalf("Get failed: %v", err)
	}
	if value != "" {
		t.Fatalf("Expected empty string for non-existent key, got %s", value)
	}
}
