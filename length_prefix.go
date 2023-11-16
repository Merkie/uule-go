package uule

// lengthPrefixKey is a lookup table for length prefixes.
var lengthPrefixKey = "EFGHIjKLMN0PQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789- ABCDEFGHIJKLMNOPQRSTUVWXYZL"

// getLengthPrefix returns the length prefix for the specified length.
func getLengthPrefix(length int) (string, bool) {
	if length < 4 || length > 89 {
		return "", false
	}
	return string(lengthPrefixKey[length-4]), true
}
