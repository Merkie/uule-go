# uule-go

[![GoDoc](https://pkg.go.dev/badge/github.com/merkie/uule-go.svg)](https://pkg.go.dev/github.com/merkie/uule-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/merkie/uule-go)](https://goreportcard.com/report/github.com/merkie/uule-go)
![License](https://img.shields.io/badge/license-MIT-green)

A simple and efficient UULE (URL-encoded Unicode Location Element) generator library for Go. This library allows you to create UULE parameters for use in location-specific Google searches.

## Installation

```bash
go get -u "github.com/merkie/uule-go@latest"
```

## Usage

```go
package main

import (
	"fmt"

	"github.com/merkie/uule-go"
)

func main() {
	// Create a UULE string for New York, New York, United States using location parameters
	nyUule, _ := uule.CreateUULE("New York", "New York", "United States")
	fmt.Println(nyUule) // w+CAIQICIfTmV3IFlvcmssTmV3IFlvcmssVW5pdGVkIFN0YXRlcw==

	// Create a UULE string for Austin, Texas, United States using location string
	austinUule, _ := uule.CreateUULEFromString("Austin,Texas,United States")
	fmt.Println(austinUule) // w+CAIQICIaQXVzdGluLFRleGFzLFVuaXRlZCBTdGF0ZXM=

	// Decode UULE strings
	decodedNyUule, _ := uule.DecodeUULEString("w+CAIQICIfTmV3IFlvcmssTmV3IFlvcmssVW5pdGVkIFN0YXRlcw==")
	fmt.Println(decodedNyUule) // New York,New York,United States

	decodedAustinUule, _ := uule.DecodeUULEString("w+CAIQICIaQXVzdGluLFRleGFzLFVuaXRlZCBTdGF0ZXM=")
	fmt.Println(decodedAustinUule) // Austin,Texas,United States
}
```

## Contributing

Contributions to the UULE Generator are welcome! Please refer to the project's issues page on GitHub for planned improvements and feature requests.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

Special thanks to [petrpatek](https://github.com/petrpatek) for sharing his work on [create-uule](https://github.com/petrpatek/create-uule).

---

## Legal Disclaimer

This project is an independent, community-driven effort and is not officially affiliated with, endorsed by, or in any way connected to Google or any of its subsidiaries or affiliates. The name Google and any related trademarks are the property of their respective owners and are used here for identification purposes only. This project is developed under the MIT License and is not associated with any official offerings from these entities.
