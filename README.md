# uule-go

[![GoDoc](https://pkg.go.dev/badge/github.com/merkie/uule-go.svg)](https://pkg.go.dev/github.com/merkie/uule-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/merkie/uule-go)](https://goreportcard.com/report/github.com/merkie/uule-go)
![License](https://img.shields.io/badge/license-MIT-green)

A simple and efficient UULE (URL-encoded Unicode Location Element) generator library for Go. This library allows you to create UULE parameters for use in location-specific Google searches.

## Features

- **Easy UULE Generation**: Quickly generate UULE parameters by specifying city, state, and country.
- **Lightweight and Fast**: Designed with performance in mind, the library has a minimal footprint, ensuring it doesn't bloat your application.

## Installation

To install the UULE Generator library, use the following command:

```bash
go get -u "github.com/merkie/uule-go"
```

## Usage

Here's a quick example to get you started:

```go
package main

import (
	"fmt"

	"github.com/merkie/uule-go"
)

func main() {
	// Generate a UULE by passing a city, state, and country
	uuleEncoded, err := uule.CreateUULE("New York", "New York", "United States")
	if err != nil {
		panic(err)
	}
	fmt.Println(uuleEncoded) // Output -> w+CAIQICIfTmV3IFlvcmssTmV3IFlvcmssVW5pdGVkIFN0YXRlcw==

	// You can also use uule.CreateUULEFromString("New York,New York,United States")
}
```

## Contributing

Contributions to the UULE Generator are welcome! Please refer to the project's issues page on GitHub for planned improvements and feature requests.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

Special thanks to [petrpatek](https://github.com/petrpatek) for sharing his work on [create-uule](https://github.com/petrpatek/create-uule).

---

_Note: This project is not affiliated with Google LLC._
