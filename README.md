# spans

[![GoDoc](https://pkg.go.dev/badge/github.com/byExist/spans.svg)](https://pkg.go.dev/github.com/byExist/spans)
[![Go Report Card](https://goreportcard.com/badge/github.com/byExist/spans)](https://goreportcard.com/report/github.com/byExist/spans)

spans is a lightweight Go package that provides a convenient and Python-like interface for working with integer ranges (spans). It supports customizable start, stop, and step values, and includes a suite of utility functions for iteration, slicing, indexing, and more.

## Features

- Define spans with start, stop, and step
- Iterate over spans using iter.Seq
- Get the length of a span
- Check if an element is contained in a span
- Get an element by index
- Find the index of an element
- Slice a span to create sub-spans

## Installation

To install spans, use the following command:
```bash
go get github.com/byExist/spans
```

## Usage

```go
package main

import (
	"fmt"
	"spans"
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

	// Slice the span
	sub, _ := spans.Slice(s, 1, 3)
	for v := range spans.Values(sub) {
		fmt.Println("Sub-span value:", v)
	}
}
```

## API Overview

### Constructors

- To(stop int) Span
- Range(start, stop int) Span
- Stride(start, stop, step int) Span

### Utilities

- Values(s Span) iter.Seq[int]
- Len(s Span) int
- Contains(s Span, elem int) bool
- Find(s Span, elem int) (int, error)
- At(s Span, index int) (int, error)
- Slice(s Span, from, to int) (Span, error)

## License

MIT License