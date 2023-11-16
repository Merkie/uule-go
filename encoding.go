package uule

import (
	"encoding/base64"
	"fmt"
)

// lengthPrefixKey is a lookup table for length prefixes.
var lengthPrefixKey = "EFGHIjKLMN0PQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789- ABCDEFGHIJKLMNOPQRSTUVWXYZL"

// getLengthPrefix returns the length prefix for the specified length.
func getLengthPrefix(length int) (string, bool) {
	if length < 4 || length > 89 {
		return "", false
	}
	return string(lengthPrefixKey[length-4]), true
}

// encodeUULEString encodes the specified location string into a UULE string.
func encodeUULEString(uuleString string) (string, error) {
	encodedLocation := base64.StdEncoding.EncodeToString([]byte(uuleString))

	lengthPrefix, exists := getLengthPrefix(len(uuleString))
	if !exists {
		return "", fmt.Errorf("no length prefix found for length %d", len(uuleString))
	}

	// "w+CAIQICI" is the standard prefix for all UULEs
	return "w+CAIQICI" + lengthPrefix + encodedLocation, nil
}
