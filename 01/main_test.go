package main

import (
	"reflect"
	"testing"
)

func TestDiff(t *testing.T) {
	expected := 3
	result := Diff(5, 2)
	if result != expected {
		t.Errorf("Diff(%d,%d) = %d want %d", 5, 2, result, expected)
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
