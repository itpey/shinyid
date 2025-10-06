# ShinyID

[![Go Reference](https://pkg.go.dev/badge/github.com/itpey/shinyid.svg)](https://pkg.go.dev/github.com/itpey/shinyid)
[![license](https://img.shields.io/github/license/itpey/shinyid)](https://github.com/itpey/shinyid/blob/main/LICENSE)

## About

ShinyID is a high-performance Go package inspired by the Instagram shortcode system. It allows you to encode and decode unique identifiers (IDs) into a human-readable and URL-safe string format called 'shiny'. This package is designed for scenarios where speed and efficiency are crucial, making it ideal for applications that need to handle large volumes of encoded IDs.

## Features

- **Efficient Encoding :** ShinyID offers a highly efficient method for converting numeric IDs into shiny strings.
- **Lightning-Fast Decoding :** Decoding shiny strings into their original numeric IDs is optimized for speed and performance.
- **URL-Safe :** Shiny strings are designed to be URL-safe, making them suitable for web applications.
- **Inspired by Instagram :** The package draws inspiration from Instagram's shortcode system, ensuring a familiar and intuitive approach to representing IDs.

## Installation

To integrate ShinyID into your Go project, you can easily install it using `go get`:

```bash
go get -u github.com/itpey/shinyid
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

|                  Benchmark |  Iterations | ns/op | B/op | allocs/op |
| -------------------------: | ----------: | ----: | ---: | --------: |
| BenchmarkShinyIDEncoding-8 |  44,686,579 | 26.32 |    0 |         0 |
| BenchmarkShinyIDDecoding-8 | 100,000,000 | 11.65 |    0 |         0 |
|  BenchmarkBase64Encoding-8 |  12,006,302 | 94.64 |   32 |         1 |
|  BenchmarkBase64Decoding-8 |  10,615,081 | 111.0 |   24 |         1 |

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

## Feedback and Contributions

If you encounter any issues or have suggestions for improvement, please [open an issue](https://github.com/itpey/shinyid/issues) on GitHub.

We welcome contributions! Fork the repository, make your changes, and submit a pull request.

## Support

If you enjoy using ShinyID, please consider giving it a star! Your support helps others discover the project and encourages further development.

## License

ShinyID is distributed under the Apache License, Version 2.0. See the [LICENSE](https://github.com/itpey/shinyid/blob/main/LICENSE) file for more details.

## Author

ShinyID was created by [itpey](https://github.com/itpey).
