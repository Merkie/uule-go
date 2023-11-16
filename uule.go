// Package uule provides a simple and efficient library for generating URL-encoded
// Unicode Location Elements (UULE) in Go.
package uule

import (
	"encoding/base64"
	"fmt"
)

// UULEParams represents the core parameters that constitute a UULE string.
// It includes the city, region, country, and an optional canonical name.
type UULEParams struct {
	City          string
	Region        string
	Country       string
	CanonicalName string
}

// CreateUULE initializes a new UULEParams object with the specified city, region,
// and country. This object can then be used to generate a UULE string.
func CreateUULE(city, region, country string) *UULEParams {
	return &UULEParams{
		City:          city,
		Region:        region,
		Country:       country,
		CanonicalName: "",
	}
}

// Encode converts the UULEParams object into an encoded UULE string. The method
// first constructs a comma-separated string from the city, region, and country fields,
// or uses the CanonicalName if provided. This string is then base64-encoded.
//
// The method also calculates a length prefix, which is a single character representing
// the length of the encoded location string minus 4. The final UULE string is composed
// of a standard UULE prefix, the length prefix, and the encoded location.
//
// Returns the complete UULE string or an error if the length prefix cannot be determined.
func (uule *UULEParams) Encode() (string, error) {
	location := uule.City + "," + uule.Region + "," + uule.Country
	if uule.CanonicalName != "" {
		location = uule.CanonicalName
	}

	encodedLocation := base64.StdEncoding.EncodeToString([]byte(location))

	lengthPrefix, exists := getLengthPrefix(len(location))
	if !exists {
		return "", fmt.Errorf("no length prefix found for length %d", len(location))
	}

	// "w+CAIQICI" is the standard prefix for all UULEs
	return "w+CAIQICI" + lengthPrefix + encodedLocation, nil
}
