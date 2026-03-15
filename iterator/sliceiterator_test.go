package iterator

import (
	"reflect"
	"testing"
)

func TestSliceIterator(t *testing.T) {
	numbers := []int{}
	iter := NewSliceIterator(numbers)
	if iter.HasNext() {
		t.Errorf("Empty iterator HasNext return true")
	}

	if iter.HasPrev() {
		t.Errorf("Empty iterator HasPrev return true")
	}
	numbers = []int{1}
	expected := []int{1}
	result := make([]int, 0)
	iter = NewSliceIterator(numbers)
	for iter.HasNext() {
		val, _ := iter.Next()
		result = append(result, val)
	}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Incorrect iteration, got %v, expected %v", result, expected)
	}
	prevResult := make([]int, 0)
	for iter.HasPrev() {
		val, _ := iter.Prev()
		prevResult = append(prevResult, val)
	}
	if !reflect.DeepEqual(prevResult, expected) {
		t.Errorf("Incorrect iteration, got %v, expected %v", result, expected)
	}

	numbers = []int{1, 2, 3}
	expected = []int{1, 2, 3}
	result = make([]int, 0)
	iter = NewSliceIterator(numbers)
	for iter.HasNext() {
		val, _ := iter.Next()
		result = append(result, val)
	}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Incorrect iteration, got %v, expected %v", result, expected)
	}

	prevExpected := []int{3, 2, 1}
	prevResult = make([]int, 0)
	for iter.HasPrev() {
		val, _ := iter.Prev()
		prevResult = append(prevResult, val)
	}
	if !reflect.DeepEqual(prevResult, prevExpected) {
		t.Errorf("Incorrect iteration, got %v, expected %v", result, expected)
	}
}

func TestReverse(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}
	expected := []int{1, 2, 3, 2, 1}
	result := make([]int, 0)
	iter := NewSliceIterator(numbers)
	idx := 0
	for iter.HasNext() {
		val, _ := iter.Next()
		result = append(result, val)
		if idx == 2 {
			iter.Reverse()
		}
		idx++
	}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Incorrect iteration, got %v, expected %v", result, expected)
	}

	iter.Reset()
	result = make([]int, 0)
	for iter.HasPrev() {
		val, _ := iter.Prev()
		result = append(result, val)
	}
	expected = []int{1, 2, 3, 4, 5}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Incorrect iteration, got %v, expected %v", result, expected)
	}

	iter.Reverse()
	result = make([]int, 0)
	for iter.HasPrev() {
		val, _ := iter.Prev()
		result = append(result, val)
	}
	expected = []int{5, 4, 3, 2, 1}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Incorrect iteration, got %v, expected %v", result, expected)
	}

	numbers = []int{0}
	iter = NewSliceIterator(numbers)
	iter.Reverse()
	if iter.HasNext() {
		t.Errorf("Reversed iterator should not have next elements")
	}
	if !iter.HasPrev() {
		t.Errorf("Reversed iterator should have prev elements")
	}
}
