package main

import (
	"reflect"
	"testing"
)

func TestBinary_selectionSort(t *testing.T) {
	array := []int{13, 3, 2, 6, 6, 8, 23, 11, 20, 34, 25}

	sortedArray := SelectionSort(array)

	if !reflect.DeepEqual(sortedArray, []int{2, 3, 6, 6, 8, 11, 13, 20, 23, 25, 34}) {
		t.Errorf("Expected [2 3, 6, 6, 8, 11, 13, 20, 23, 25, 34], got %d", sortedArray)
	}
}

func TestSelectionSort_sorted_array(t *testing.T) {
	array := []int{1, 2, 3, 4}

	sortedArray := SelectionSort(array)

	if !reflect.DeepEqual(sortedArray, []int{1, 2, 3, 4}) {
		t.Errorf("Expected [1, 2, 3, 4], got %d", sortedArray)
	}
}

func TestSelectionSort_with_smallest_as_last(t *testing.T) {
	array := []int{5, 2, 3, 4, 1}

	sortedArray := SelectionSort(array)

	if !reflect.DeepEqual(sortedArray, []int{1, 2, 3, 4, 5}) {
		t.Errorf("Expected [1, 2, 3, 4, 5], got %d", sortedArray)
	}
}
