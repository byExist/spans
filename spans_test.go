package spans_test

import (
	"fmt"
	"testing"

	"github.com/byExist/spans"
)

func TestTo(t *testing.T) {
	s := spans.To(5)
	if s.Start() != 0 || s.Stop() != 5 || s.Step() != 1 {
		t.Errorf("To(5) = %+v, want start=0, stop=5, step=1", s)
	}
}

func TestRange(t *testing.T) {
	s := spans.Range(2, 6)
	if s.Start() != 2 || s.Stop() != 6 || s.Step() != 1 {
		t.Errorf("Range(2, 6) = %+v, want start=2, stop=6, step=1", s)
	}
}

func TestStride(t *testing.T) {
	s := spans.Stride(3, 9, 2)
	if s.Start() != 3 || s.Stop() != 9 || s.Step() != 2 {
		t.Errorf("Stride(3, 9, 2) = %+v, want start=3, stop=9, step=2", s)
	}

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Stride with step 0 did not panic")
		}
	}()
	_ = spans.Stride(1, 10, 0)
}

func TestClone(t *testing.T) {
	s := spans.Stride(0, 5, 1)
	cloned := spans.Clone(s)
	if cloned.Start() != s.Start() || cloned.Stop() != s.Stop() || cloned.Step() != s.Step() {
		t.Errorf("Clone() = %+v, want same as original", cloned)
	}
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
		var got []int
		for v := range spans.Values(test.s) {
			got = append(got, v)
		}
		if len(got) != len(test.expected) {
			t.Errorf("Values(%+v) = %v, expected %v", test.s, got, test.expected)
			continue
		}
		for i := range got {
			if got[i] != test.expected[i] {
				t.Errorf("Values(%+v)[%d] = %d, expected %d", test.s, i, got[i], test.expected[i])
			}
		}
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
		if got != test.expected {
			t.Errorf("Len(%+v) = %d, expected %d", test.s, got, test.expected)
		}
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
		if got != test.expected {
			t.Errorf("Contains(%+v, %d) = %v, expected %v", test.s, test.elem, got, test.expected)
		}
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
		if ok != test.ok {
			t.Errorf("Find(%+v, %d) ok = %v, want %v", test.s, test.elem, ok, test.ok)
			continue
		}
		if ok && got != test.expected {
			t.Errorf("Find(%+v, %d) = %d, expected %d", test.s, test.elem, got, test.expected)
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
		if ok != test.ok {
			t.Errorf("At(%+v, %d) ok = %v, want %v", s, test.index, ok, test.ok)
			continue
		}
		if ok && got != test.expected {
			t.Errorf("At(%+v, %d) = %d, want %d", s, test.index, got, test.expected)
		}
	}
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
