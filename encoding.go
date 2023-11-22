package uule

import (
	"encoding/base64"
	"fmt"
)

var lengthPrefixKey = "EFGHIjKLMN0PQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789- ABCDEFGHIJKLMNOPQRSTUVWXYZL"

func getLengthPrefix(length int) (string, bool) {
	if length < 4 || length > 89 {
		return "", false
	}
	return string(lengthPrefixKey[length-4]), true
}

func encodeUULEString(uuleString string) (string, error) {
	encodedLocation := base64.StdEncoding.EncodeToString([]byte(uuleString))

	lengthPrefix, exists := getLengthPrefix(len(uuleString))
	if !exists {
		return "", fmt.Errorf("no length prefix found for length %d", len(uuleString))
	}

	// "w+CAIQICI" is the standard prefix for all UULEs
	return "w+CAIQICI" + lengthPrefix + encodedLocation, nil
}
