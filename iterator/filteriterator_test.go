package iterator

import (
	"reflect"
	"testing"
)

func TestEmptyFilterIterator(t *testing.T) {
	n := []int{}
	sliceIter := NewSliceIterator(n)
	filterIter := NewFilterIterator(sliceIter, func(val int) bool { return val%2 == 0 })
	if filterIter.HasNext() {
		t.Errorf("Empty iterator HasNext return true")
	}
}

func TestFilterIterator(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5, 6}
	sliceIter := NewSliceIterator(numbers)
	filterIter := NewFilterIterator(sliceIter, func(val int) bool { return val%2 == 0 })
	result := []int{}
	for filterIter.HasNext() {
		val, _ := filterIter.Next()
		result = append(result, val)
	}
	expected := []int{1, 3, 5}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Incorrect iteration, got %v, expected %v", result, expected)
	}

	result = []int{}
	for filterIter.HasPrev() {
		val, _ := filterIter.Prev()
		result = append(result, val)
	}
	expected = []int{5, 3, 1}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Incorrect iteration, got %v, expected %v", result, expected)
	}

	filterIter.Reset()
	result = []int{}
	for filterIter.HasNext() {
		val, _ := filterIter.Next()
		result = append(result, val)
	}
	expected = []int{1, 3, 5}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Incorrect iteration, got %v, expected %v", result, expected)
	}

	numbers = []int{1, 2, 3, 4, 5}
	sliceIter = NewSliceIterator(numbers)
	filterIter = NewFilterIterator(sliceIter, func(val int) bool { return val%2 == 0 })
	result = []int{}
	for filterIter.HasNext() {
		val, _ := filterIter.Next()
		result = append(result, val)
	}
	expected = []int{1, 3, 5}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Incorrect iteration, got %v, expected %v", result, expected)
	}

	result = []int{}
	for filterIter.HasPrev() {
		val, _ := filterIter.Prev()
		result = append(result, val)
	}
	expected = []int{5, 3, 1}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Incorrect iteration, got %v, expected %v", result, expected)
	}

	filterIter.Reset()
	result = []int{}
	for filterIter.HasNext() {
		val, _ := filterIter.Next()
		result = append(result, val)
	}
	expected = []int{1, 3, 5}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Incorrect iteration, got %v, expected %v", result, expected)
	}
}

func TestFilterIteratorReverse(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5, 6}
	sliceIter := NewSliceIterator(numbers)
	filterIter := NewFilterIterator(sliceIter, func(val int) bool { return val%2 == 0 })
	result := []int{}
	for filterIter.HasNext() {
		_, _ = filterIter.Next()
	}
	filterIter.Reverse()
	for filterIter.HasNext() {
		val, _ := filterIter.Next()
		result = append(result, val)
	}
	expected := []int{5, 3, 1}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Incorrect iteration, got %v, expected %v", result, expected)
	}
}
