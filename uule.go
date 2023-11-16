package uule

import (
	"encoding/base64"
	"fmt"
)

type UULEParams struct {
	City          string
	Region        string
	Country       string
	CanonicalName string
}

func CreateUULE(city string, region string, country string) *UULEParams {
	return &UULEParams{
		City:          city,
		Region:        region,
		Country:       country,
		CanonicalName: "",
	}
}

func (uule *UULEParams) Encode() string {
	location := ""
	if uule.CanonicalName != "" {
		location = uule.CanonicalName
	} else {
		location = uule.City + "," + uule.Region + "," + uule.Country
	}

	encodedLocation := base64.StdEncoding.EncodeToString([]byte(location))

	lengthPrefix, exists := LengthSecret[len(location)]
	if !exists {
		fmt.Printf("No length prefix found for length %d\n", len(location))
		return ""
	}

	return Identifier + lengthPrefix + encodedLocation
}
