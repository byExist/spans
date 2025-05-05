# spans [![GoDoc](https://pkg.go.dev/badge/github.com/byExist/spans.svg)](https://pkg.go.dev/github.com/byExist/spans) [![Go Report Card](https://goreportcard.com/badge/github.com/byExist/spans)](https://goreportcard.com/report/github.com/byExist/spans) 

## What is "spans"?

spans is a lightweight Go package that provides a convenient and Python-like interface for working with integer ranges (spans). It supports customizable start, stop, and step values, and includes a suite of utility functions for iteration, slicing, indexing, and more. Most operations, such as length checks, indexing, and containment tests, are constant time (O(1)), while iteration scales linearly with span size (O(n)).

## Features

- Define spans with start, stop, and step
- Iterate over spans using iter.Seq
- Get the length of a span
- Check if an element is contained in a span
- Get an element by index
- Find the index of an element

## Installation

To install spans, use the following command:

```bash
go get github.com/byExist/spans
```

## Quick Start

```go
package main

import (
	"fmt"
	"github.com/byExist/spans"
)

func main() {
	s := spans.Stride(0, 10, 2)

	// Iterate over the span
	for v := range spans.Values(s) {
		fmt.Println(v)
	}

	// Get the length of the span
	fmt.Println("Length:", spans.Len(s))

	// Check if a value is in the span
	fmt.Println("Contains 4?", spans.Contains(s, 4))

	// Get a value at an index
	val, _ := spans.At(s, 2)
	fmt.Println("Value at index 2:", val)
}
```

## Usage

The spans package allows you to easily create and manipulate integer ranges with customizable start, stop, and step values. You can iterate over spans, check for element containment, retrieve elements by index, and find element indices. This makes spans a versatile tool for range-based operations in Go.

## API Overview

### Constructors

| Function | Description | Time Complexity |
|----------|-------------|------------------|
| `To(stop int)` | Creates a Span from 0 to stop (exclusive), step 1 | O(1) |
| `Range(start, stop int)` | Creates a Span from start to stop (exclusive), step 1 | O(1) |
| `Stride(start, stop, step int)` | Creates a Span from start to stop with given step | O(1) |

### Methods

| Method | Description | Time Complexity |
|--------|-------------|------------------|
| `Start(s Span)` | Returns the start value of the Span | O(1) |
| `Stop(s Span)` | Returns the stop value of the Span | O(1) |
| `Step(s Span)` | Returns the step value of the Span | O(1) |

### Utilities

| Function | Description | Time Complexity |
|----------|-------------|------------------|
| `Clone(s Span)` | Returns a copy of the Span | O(1) |
| `Values(s Span)` | Returns an iterator over the Span | O(n) |
| `Len(s Span)` | Returns the number of elements in the Span | O(1) |
| `Contains(s Span, elem int)` | Checks if a value is in the Span | O(1) |
| `Find(s Span, elem int)` | Finds the index of a value in the Span | O(1) |
| `At(s Span, index int)` | Returns the value at the given index | O(1) |

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
