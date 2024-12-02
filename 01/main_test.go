package main

import (
	"reflect"
	"testing"
)

func TestSmallestNumber(t *testing.T) {
	numbers := []int{42, 32, 10, 5, 11, 9}
	result := SmallestNumber(numbers)
	expected := 5
	if result != expected {
		t.Errorf("SmallestNumber(%d) = %d, want %d", numbers, result, expected)
	}
}
func TestSortArray(t *testing.T) {
	numbers := []int{42, 32, 10, 5, 11, 9}
	result := SortArray(numbers)
	expected := []int{5, 9, 10, 11, 32, 42}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("SortArray(%d) = %d, want %d", numbers, result, expected)
	}
}
