# spans [![GoDoc](https://pkg.go.dev/badge/github.com/byExist/spans.svg)](https://pkg.go.dev/github.com/byExist/spans) [![Go Report Card](https://goreportcard.com/badge/github.com/byExist/spans)](https://goreportcard.com/report/github.com/byExist/spans)

A minimal, Python-like span/range utility for integers in Go.

The `spans` package provides an efficient and allocation-free way to represent ranges of integers. It supports slicing, indexing, containment, and iteration in O(1) time where possible, and is especially useful for Unicode filtering, numeric range matching, and logic replacement for manual loops.

---

## ✨ Features

- ✅ Define spans with start, stop, and step
- ✅ Allocation-free representation
- ✅ Constant-time: `Len`, `Contains`, `At`, `Find`
- ✅ Interoperable with `iter.Seq`
- ❌ Immutable (no dynamic insertion/removal)
- ❌ Integer-only (no float/string ranges)

---

## 🧱 Example

```go
package main

import (
	"fmt"
	"github.com/byExist/spans"
)

func main() {
	hangul := spans.Range(0xAC00, 0xD7A3 + 1)

	ch := '강'
	fmt.Printf("'%c' (%U) is Hangul? %v\n", ch, ch, spans.Contains(hangul, int(ch)))

	ch = 'A'
	fmt.Printf("'%c' (%U) is Hangul? %v\n", ch, ch, spans.Contains(hangul, int(ch)))
}
```

```go
// Using Stride to generate even numbers from 0 to 10
evens := spans.Stride(0, 10, 2)
for _, v := range spans.Values(evens) {
    fmt.Println(v)
}
```

---

## 📚 Use When

- You need efficient range checking or indexing
- You want a zero-allocation alternative to slices
- You work with Unicode, event codes, or numeric classes

---

## 🚫 Avoid If

- You need dynamic range mutation
- You want to hold non-integer values
- You need interval trees or overlapping span logic

---

## 📊 Performance

Benchmarked on Apple M1 Pro (darwin/arm64):

| Operation   | Time (ns/op) | Allocations |
|-------------|--------------|-------------|
| Contains    | 2.13         | 0           |
| Len         | 2.12         | 0           |
| At          | 2.23         | 0           |
| Find        | 2.52         | 0           |

---

## 🔍 API

| Function | Description |
|----------|-------------|
| `To(stop)` | span from 0 to stop |
| `Range(start, stop)` | span from start to stop |
| `Stride(start, stop, step)` | span from start to stop with step |
| `Len(span)` | number of elements |
| `Contains(span, elem)` | test for membership |
| `At(span, index)` | value at index |
| `Find(span, value)` | index of value |
| `Values(span)` | iterator over span |

## ⚠️ Limitations

- `step` must not be 0 (will panic)
- `start`, `stop`, and `step` must define a finite sequence

---

## 🪪 License

MIT License. See [LICENSE](LICENSE).
