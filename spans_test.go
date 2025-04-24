package spans

import (
	"fmt"
	"testing"
)

func TestTo(t *testing.T) {
	s := To(5).(span)
	if s.start != 0 || s.stop != 5 || s.step != 1 {
		t.Errorf("To(5) = %+v, want start=0, stop=5, step=1", s)
	}
}

func TestRange(t *testing.T) {
	s := Range(2, 6).(span)
	if s.start != 2 || s.stop != 6 || s.step != 1 {
		t.Errorf("Range(2, 6) = %+v, want start=2, stop=6, step=1", s)
	}
}

func TestStride(t *testing.T) {
	s := Stride(3, 9, 2).(span)
	if s.start != 3 || s.stop != 9 || s.step != 2 {
		t.Errorf("Stride(3, 9, 2) = %+v, want start=3, stop=9, step=2", s)
	}

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Stride with step 0 did not panic")
		}
	}()
	_ = Stride(1, 10, 0)
}

func TestValues(t *testing.T) {
	tests := []struct {
		s        span
		expected []int
	}{
		{span{0, 5, 1}, []int{0, 1, 2, 3, 4}},
		{span{5, 0, -1}, []int{5, 4, 3, 2, 1}},
		{span{0, 0, 1}, []int{}},
		{span{0, 10, 2}, []int{0, 2, 4, 6, 8}},
	}

	for _, test := range tests {
		var got []int
		for v := range Values(test.s) {
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
		s        span
		expected int
	}{
		{span{0, 5, 1}, 5},
		{span{5, 0, -1}, 5},
		{span{0, 0, 1}, 0},
		{span{0, 10, 2}, 5},
	}

	for _, test := range tests {
		got := Len(test.s)
		if got != test.expected {
			t.Errorf("Len(%+v) = %d, expected %d", test.s, got, test.expected)
		}
	}
}

func TestContains(t *testing.T) {
	tests := []struct {
		s        span
		elem     int
		expected bool
	}{
		{span{0, 5, 1}, 3, true},
		{span{0, 5, 1}, 5, false},
		{span{5, 0, -1}, 2, true},
		{span{5, 0, -1}, -1, false},
		{span{0, 0, 1}, 0, false},
		{span{0, 10, 2}, 4, true},
		{span{0, 10, 2}, 5, false},
	}

	for _, test := range tests {
		got := Contains(test.s, test.elem)
		if got != test.expected {
			t.Errorf("Contains(%+v, %d) = %v, expected %v", test.s, test.elem, got, test.expected)
		}
	}
}

func TestFind(t *testing.T) {
	tests := []struct {
		s        span
		elem     int
		expected int
		ok       bool
	}{
		{span{0, 5, 1}, 3, 3, true},
		{span{0, 5, 1}, 5, 0, false},
		{span{5, 0, -1}, 2, 3, true},
		{span{5, 0, -1}, -1, 0, false},
		{span{0, 10, 2}, 4, 2, true},
		{span{0, 10, 2}, 5, 0, false},
	}

	for _, test := range tests {
		got, err := Find(test.s, test.elem)
		if (err == nil) != test.ok {
			t.Errorf("Find(%+v, %d) error = %v, want ok=%v", test.s, test.elem, err, test.ok)
			continue
		}
		if err == nil && got != test.expected {
			t.Errorf("Find(%+v, %d) = %d, expected %d", test.s, test.elem, got, test.expected)
		}
	}
}

func TestAt(t *testing.T) {
	s := span{start: 10, stop: 20, step: 2}
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
		got, err := At(s, test.index)
		if (err == nil) != test.ok {
			t.Errorf("At(%+v, %d) error = %v, want ok=%v", s, test.index, err, test.ok)
			continue
		}
		if err == nil && got != test.expected {
			t.Errorf("At(%+v, %d) = %d, want %d", s, test.index, got, test.expected)
		}
	}
}

func TestSlice(t *testing.T) {
	s := span{start: 0, stop: 10, step: 2} // [0,2,4,6,8]
	tests := []struct {
		from, to int
		want     span
		ok       bool
	}{
		{0, 3, span{0, 6, 2}, true},
		{1, 4, span{2, 8, 2}, true},
		{0, 5, span{0, 10, 2}, true},
		{3, 3, span{6, 6, 2}, true},
		{4, 2, span{}, false},  // invalid: from > to
		{-1, 2, span{}, false}, // invalid: from < 0
		{0, 6, span{}, false},  // invalid: to > len
	}

	for _, test := range tests {
		got, err := Slice(s, test.from, test.to)
		if (err == nil) != test.ok {
			t.Errorf("Slice(%+v, %d, %d) error = %v, want ok=%v", s, test.from, test.to, err, test.ok)
			continue
		}
		if err == nil && got != test.want {
			t.Errorf("Slice(%+v, %d, %d) = %+v, want %+v", s, test.from, test.to, got, test.want)
		}
	}
}

// ExampleTo demonstrates usage of the To function.
func ExampleTo() {
	s := To(5)
	for v := range Values(s) {
		fmt.Print(v, " ")
	}
	// Output: 0 1 2 3 4
}

// ExampleRange demonstrates usage of the Range function.
func ExampleRange() {
	s := Range(3, 6)
	for v := range Values(s) {
		fmt.Print(v, " ")
	}
	// Output: 3 4 5
}

// ExampleStride demonstrates usage of the Stride function.
func ExampleStride() {
	s := Stride(2, 10, 3)
	for v := range Values(s) {
		fmt.Print(v, " ")
	}
	// Output: 2 5 8
}

// ExampleValues demonstrates usage of the Values function.
func ExampleValues() {
	s := Stride(0, 10, 2)
	for v := range Values(s) {
		fmt.Print(v, " ")
	}
	// Output: 0 2 4 6 8
}

// ExampleLen demonstrates usage of the Len function.
func ExampleLen() {
	fmt.Println(Len(To(5)))
	fmt.Println(Len(Stride(5, 0, -1)))
	// Output:
	// 5
	// 5
}

// ExampleContains demonstrates usage of the Contains function.
func ExampleContains() {
	s := Stride(0, 10, 2)
	fmt.Println(Contains(s, 4))
	fmt.Println(Contains(s, 5))
	// Output:
	// true
	// false
}

// ExampleFind demonstrates usage of the Find function.
func ExampleFind() {
	s := Stride(0, 10, 2)
	i, _ := Find(s, 6)
	fmt.Println(i)
	_, err := Find(s, 5)
	fmt.Println(err != nil)
	// Output:
	// 3
	// true
}

// ExampleAt demonstrates usage of the At function.
func ExampleAt() {
	s := Stride(10, 20, 2)
	v, _ := At(s, 2)
	fmt.Println(v)
	_, err := At(s, 10)
	fmt.Println(err != nil)
	// Output:
	// 14
	// true
}

// ExampleSlice demonstrates usage of the Slice function.
func ExampleSlice() {
	s := To(10) // 0 1 2 3 4 5 6 7 8 9
	sub, _ := Slice(s, 2, 5)
	for v := range Values(sub) {
		fmt.Print(v, " ")
	}
	// Output: 2 3 4
}
