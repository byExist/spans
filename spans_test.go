package spans_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/byExist/spans"
	"github.com/stretchr/testify/assert"
)

func TestSpanString(t *testing.T) {
	s := spans.Stride(1, 5, 1)
	assert.Equal(t, "Span(1, 5, 1)", s.String())
}

func TestSpanMarshalUnmarshalJSON(t *testing.T) {
	original := spans.Stride(1, 5, 1)
	data, err := json.Marshal(original)
	assert.NoError(t, err)

	var decoded spans.Span
	assert.NoError(t, json.Unmarshal(data, &decoded))

	assert.Equal(t, original, decoded)
}

func TestTo(t *testing.T) {
	s := spans.To(5)
	expected := spans.Stride(0, 5, 1)
	assert.Equal(t, expected, s)
}

func TestRange(t *testing.T) {
	s := spans.Range(2, 6)
	expected := spans.Stride(2, 6, 1)
	assert.Equal(t, expected, s)
}

func TestStride(t *testing.T) {
	s := spans.Stride(3, 9, 2)
	expected := spans.Stride(3, 9, 2)
	assert.Equal(t, expected, s)

	assert.Panics(t, func() {
		_ = spans.Stride(1, 10, 0)
	})
}

func TestClone(t *testing.T) {
	s := spans.Stride(0, 5, 1)
	cloned := spans.Clone(s)
	assert.Equal(t, cloned, s)
}

func TestValues(t *testing.T) {
	tests := []struct {
		s        spans.Span
		expected []int
	}{
		{spans.Stride(0, 5, 1), []int{0, 1, 2, 3, 4}},
		{spans.Stride(5, 0, -1), []int{5, 4, 3, 2, 1}},
		{spans.Stride(0, 0, 1), []int{}},
		{spans.Stride(0, 10, 2), []int{0, 2, 4, 6, 8}},
	}

	for _, test := range tests {
		got := make([]int, 0)
		for v := range spans.Values(test.s) {
			got = append(got, v)
		}
		assert.Equal(t, test.expected, got)
	}
}

func TestLen(t *testing.T) {
	tests := []struct {
		s        spans.Span
		expected int
	}{
		{spans.Stride(0, 5, 1), 5},
		{spans.Stride(5, 0, -1), 5},
		{spans.Stride(0, 0, 1), 0},
		{spans.Stride(0, 10, 2), 5},
	}

	for _, test := range tests {
		got := spans.Len(test.s)
		assert.Equal(t, test.expected, got)
	}
}

func TestContains(t *testing.T) {
	tests := []struct {
		s        spans.Span
		elem     int
		expected bool
	}{
		{spans.Stride(0, 5, 1), 3, true},
		{spans.Stride(0, 5, 1), 5, false},
		{spans.Stride(5, 0, -1), 2, true},
		{spans.Stride(5, 0, -1), -1, false},
		{spans.Stride(0, 0, 1), 0, false},
		{spans.Stride(0, 10, 2), 4, true},
		{spans.Stride(0, 10, 2), 5, false},
	}

	for _, test := range tests {
		got := spans.Contains(test.s, test.elem)
		assert.Equal(t, test.expected, got)
	}
}

func TestFind(t *testing.T) {
	tests := []struct {
		s        spans.Span
		elem     int
		expected int
		ok       bool
	}{
		{spans.Stride(0, 5, 1), 3, 3, true},
		{spans.Stride(0, 5, 1), 5, 0, false},
		{spans.Stride(5, 0, -1), 2, 3, true},
		{spans.Stride(5, 0, -1), -1, 0, false},
		{spans.Stride(0, 10, 2), 4, 2, true},
		{spans.Stride(0, 10, 2), 5, 0, false},
	}

	for _, test := range tests {
		got, ok := spans.Find(test.s, test.elem)
		assert.Equal(t, test.ok, ok)
		if ok {
			assert.Equal(t, test.expected, got)
		}
	}
}

func TestAt(t *testing.T) {
	s := spans.Stride(10, 20, 2)
	tests := []struct {
		index    int
		expected int
		ok       bool
	}{
		{0, 10, true},
		{1, 12, true},
		{4, 18, true},
		{5, 0, false},  // out of range
		{-1, 0, false}, // negative index
	}

	for _, test := range tests {
		got, ok := spans.At(s, test.index)
		assert.Equal(t, test.ok, ok)
		if ok {
			assert.Equal(t, test.expected, got)
		}
	}
}

func ExampleSpan_String() {
	s := spans.Stride(2, 8, 2)
	fmt.Println(s.String())
	// Output: Span(2, 8, 2)
}

func ExampleSpan_MarshalJSON() {
	s := spans.Stride(0, 6, 2)
	data, _ := json.Marshal(s)
	fmt.Println(string(data))
	// Output: [0,6,2]
}

func ExampleSpan_UnmarshalJSON() {
	var s spans.Span
	_ = json.Unmarshal([]byte(`[1,5,1]`), &s)
	fmt.Println(s.String())
	// Output: Span(1, 5, 1)
}

func ExampleTo() {
	s := spans.To(3)
	for v := range spans.Values(s) {
		fmt.Print(v, " ")
	}
	// Output: 0 1 2
}

func ExampleRange() {
	s := spans.Range(1, 4)
	for v := range spans.Values(s) {
		fmt.Print(v, " ")
	}
	// Output: 1 2 3
}

func ExampleStride() {
	s := spans.Stride(0, 6, 2)
	for v := range spans.Values(s) {
		fmt.Print(v, " ")
	}
	// Output: 0 2 4
}

func ExampleSpan_Start() {
	fmt.Println(spans.Stride(3, 10, 2).Start())
	// Output: 3
}

func ExampleSpan_Stop() {
	fmt.Println(spans.Stride(3, 10, 2).Stop())
	// Output: 10
}

func ExampleSpan_Step() {
	fmt.Println(spans.Stride(3, 10, 2).Step())
	// Output: 2
}

func ExampleClone() {
	s := spans.Stride(1, 5, 1)
	cloned := spans.Clone(s)
	for v := range spans.Values(cloned) {
		fmt.Print(v, " ")
	}
	// Output: 1 2 3 4
}

func ExampleValues() {
	s := spans.Stride(1, 6, 2)
	for v := range spans.Values(s) {
		fmt.Print(v, " ")
	}
	// Output: 1 3 5
}

func ExampleLen() {
	fmt.Println(spans.Len(spans.Stride(0, 5, 2)))
	fmt.Println(spans.Len(spans.Stride(5, 0, -2)))
	// Output:
	// 3
	// 3
}

func ExampleContains() {
	s := spans.Stride(0, 6, 2)
	fmt.Println(spans.Contains(s, 4))
	fmt.Println(spans.Contains(s, 5))
	// Output:
	// true
	// false
}

func ExampleFind() {
	s := spans.Stride(0, 10, 2)
	i, ok := spans.Find(s, 6)
	fmt.Println(i)
	fmt.Println(ok)
	_, ok = spans.Find(s, 5)
	fmt.Println(ok)
	// Output:
	// 3
	// true
	// false
}

func ExampleAt() {
	s := spans.Stride(10, 20, 2)
	v, ok := spans.At(s, 2)
	fmt.Println(v)
	fmt.Println(ok)
	_, ok = spans.At(s, 10)
	fmt.Println(ok)
	// Output:
	// 14
	// true
	// false
}

func BenchmarkContains(b *testing.B) {
	s := spans.Stride(0, 1000, 2)
	for i := 0; i < b.N; i++ {
		spans.Contains(s, 500)
	}
}

func BenchmarkLen(b *testing.B) {
	s := spans.Stride(0, 1000, 2)
	for i := 0; i < b.N; i++ {
		spans.Len(s)
	}
}

func BenchmarkAt(b *testing.B) {
	s := spans.Stride(0, 1000, 2)
	for i := 0; i < b.N; i++ {
		spans.At(s, 500)
	}
}

func BenchmarkFind(b *testing.B) {
	s := spans.Stride(0, 1000, 2)
	for i := 0; i < b.N; i++ {
		spans.Find(s, 500)
	}
}
