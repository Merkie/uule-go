// A simple and efficient UULE (URL-encoded Unicode Location Element) generator library for Go.
package uule

import (
	"encoding/base64"
	"fmt"
)

// Creates a UULE string from the specified location parameters.
func CreateUULE(city, region, country string) (string, error) {
	location := fmt.Sprintf("%s,%s,%s", city, region, country)

	uuleString, err := encodeUULEString(location)
	if err != nil {
		return "", err
	}

	return uuleString, nil
}

// Creates a UULE string from the specified location string.
func CreateUULEFromString(location string) (string, error) {
	uuleString, err := encodeUULEString(location)
	if err != nil {
		return "", err
	}

	return uuleString, nil
}

// Decodes a UULE string returning the location string encoded within.
func DecodeUULEString(uuleString string) (string, error) {
	// Remove prefix and length prefix
	// This will leave us with the base64 encoded location string
	encodedLocation := uuleString[10:]

	decodedLocation, err := base64.StdEncoding.DecodeString(encodedLocation)
	if err != nil {
		return "", err
	}

	return string(decodedLocation), nil
}
