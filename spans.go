package spans

import (
	"errors"
	"fmt"
	"iter"
)

type (
	// Span represents a range of integers with a start, stop, and step.
	Span struct {
		start int
		stop  int
		step  int
	}
)

// Start returns the starting value of the span.
func (i Span) Start() int {
	return i.start
}

// Stop returns the stopping value of the span.
func (i Span) Stop() int {
	return i.stop
}

// Step returns the step size of the span.
func (i Span) Step() int {
	return i.step
}

// To returns a Span starting at 0 and ending before the given stop value, with a step of 1.
func To(stop int) Span {
	return Span{
		start: 0,
		stop:  stop,
		step:  1,
	}
}

// Range returns a Span starting at the given start and ending before stop, with a step of 1.
func Range(start, stop int) Span {
	return Span{
		start: start,
		stop:  stop,
		step:  1,
	}
}

// Stride returns a Span with the specified start, stop, and step values.
// Panics if the step is 0.
func Stride(start, stop, step int) Span {
	if step == 0 {
		panic("step cannot be zero")
	}
	return Span{
		start: start,
		stop:  stop,
		step:  step,
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// Values returns an iterator that yields values in the span.
func Values(s Span) iter.Seq[int] {
	start, stop, step := s.Start(), s.Stop(), s.Step()
	return func(yield func(int) bool) {
		for i := start; (step > 0 && i < stop) || (step < 0 && i > stop); i += step {
			if !yield(i) {
				break
			}
		}
	}
}

// Len returns the number of elements in the span.
func Len(s Span) int {
	start, stop, step := s.Start(), s.Stop(), s.Step()
	if (step > 0 && start >= stop) || (step < 0 && start <= stop) {
		return 0
	}
	diff := stop - start
	if step < 0 {
		diff = start - stop
	}
	return (diff + abs(step) - 1) / abs(step)
}

// Contains returns true if the given element is contained within the span.
func Contains(s Span, elem int) bool {
	start, stop, step := s.Start(), s.Stop(), s.Step()
	if (step > 0 && (elem < start || elem >= stop)) ||
		(step < 0 && (elem > start || elem <= stop)) {
		return false
	}
	diff := elem - start
	if step < 0 {
		diff = start - elem
	}
	return diff%abs(step) == 0
}

// Find returns the index of the element in the span.
// Returns an error if the element is not in the span.
func Find(s Span, elem int) (int, error) {
	if !Contains(s, elem) {
		return 0, errors.New("element not found in span")
	}
	start, step := s.Start(), s.Step()
	if step > 0 {
		return (elem - start) / step, nil
	}
	return (start - elem) / abs(step), nil
}

// At returns the element at the given index in the span.
// Returns an error if the index is out of bounds.
func At(s Span, index int) (int, error) {
	l := Len(s)
	if index < 0 || index >= l {
		return 0, fmt.Errorf("index %d out of bounds [0, %d)", index, l)
	}
	return s.Start() + index*s.Step(), nil
}
