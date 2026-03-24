package util

import "github.com/google/uuid"

// NewUUID returns a new UUID v4 string.
func NewUUID() string {
	return uuid.New().String()
}
