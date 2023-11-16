// A simple and efficient UULE (URL-encoded Unicode Location Element) generator library for Go.
package uule

import (
	"fmt"
)

// CreateUULE creates a UULE string from the specified city, region and country.
func CreateUULE(city, region, country string) (string, error) {
	location := fmt.Sprintf("%s,%s,%s", city, region, country)

	uuleString, err := encodeUULEString(location)
	if err != nil {
		return "", err
	}

	return uuleString, nil
}

// CreateUULEFromString creates a UULE string from the specified location string.
func CreateUULEFromString(location string) (string, error) {
	uuleString, err := encodeUULEString(location)
	if err != nil {
		return "", err
	}

	return uuleString, nil
}
