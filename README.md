# ShinyID

[![Go Reference](https://pkg.go.dev/badge/github.com/itpey/shinyid.svg)](https://pkg.go.dev/github.com/itpey/shinyid)
[![license](https://img.shields.io/github/license/itpey/shinyid)](https://github.com/itpey/shinyid/blob/main/LICENSE)

## About
ShinyID is a high-performance Go package inspired by the Instagram shortcode system. It allows you to encode and decode unique identifiers (IDs) into a human-readable and URL-safe string format called 'shiny'. This package is designed for scenarios where speed and efficiency are crucial, making it ideal for applications that need to handle large volumes of encoded IDs.

## Table of Contents
- [Features](#features)
- [Installation](#installation)
- [Usage](#usage)
  - [Encoding (ID to Shiny)](#encoding-id-to-shiny)
  - [Decoding (Shiny to ID)](#decoding-shiny-to-id)
- [Example](#example) 
- [Benchmark Results](#benchmark-results)
- [Running Tests](#running-tests)
- [Running Benchmarks](#running-benchmarks)
- [License](#license)
- [Author](#author)

## Features

- **Efficient Encoding :** ShinyID offers a highly efficient method for converting numeric IDs into shiny strings.
- **Lightning-Fast Decoding :** Decoding shiny strings into their original numeric IDs is optimized for speed and performance.
- **URL-Safe :** Shiny strings are designed to be URL-safe, making them suitable for web applications.
- **Inspired by Instagram :** The package draws inspiration from Instagram's shortcode system, ensuring a familiar and intuitive approach to representing IDs.

## Installation
To integrate ShinyID into your Go project, you can easily install it using `go get`:

```bash
go get github.com/itpey/shinyid
```

## Usage

### Encoding (ID to Shiny)

You can quickly convert a numeric ID to its shiny representation using the ToShiny function:

```go
shiny := shinyid.ToShiny(9375)
```

### Decoding (Shiny to ID)

To decode a shiny string back into its original numeric ID, use the ToId function:

```go
id, err := shinyid.ToId("CSf")
if err != nil {
    // Handle error
}
```

## Example
Here's a simple example showcasing the use of ShinyID:

```go
package main

import (
	"fmt"
	"github.com/itpey/shinyid"
)

func main() {
	// Encoding an ID to a shiny string
	shiny := shinyid.ToShiny(12345)
	fmt.Println("Shiny:", shiny)

	// Decoding a shiny string back to an ID
	id, err := shinyid.ToId(shiny)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("ID:", id)
	}
}
```

## Benchmark Results

### Benchmark results on an Intel Core i7-6700HQ CPU @ 2.60GHz:

- **BenchmarkShinyIDEncoding :** ShinyID encoding took approximately **15.96** nanoseconds per operation.
- **BenchmarkShinyIDDecoding :** ShinyID decoding took approximately **28.04** nanoseconds per operation.
- **BenchmarkBase64Encoding :** base64 encoding took approximately **105.2** nanoseconds per operation.
- **BenchmarkBase64Decoding :** base64 decoding took approximately **84.01** nanoseconds per operation.

Based on these benchmark results, ShinyID significantly outperforms base64 encoding and decoding, especially in encoding operations. This performance advantage makes ShinyID an excellent choice for applications that require fast encoding and decoding of unique identifiers.

## Running Tests

To run tests for ShinyID, use the following command:

```bash
go test github.com/itpey/shinyid
```
## Running Benchmarks

To run benchmarks for ShinyID, use the following command:

```bash
go test -bench=. github.com/itpey/shinyid
```
## License
This package is distributed under the Apache License, Version 2.0. See the [LICENSE](https://github.com/itpey/shinyid/blob/main/LICENSE) file for more details.

## Author
ShinyID was created by [itpey](https://github.com/itpey).
\
Experience the power and speed of ShinyID for your Go applications!
