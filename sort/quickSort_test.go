package main

import (
	"reflect"
	"testing"
)

func TestQuickSort_zero_elements(t *testing.T) {
	array := []int{}

	sortedArray := QuickSort(array)

	if !reflect.DeepEqual(sortedArray, []int{}) {
		t.Errorf("Expected [], got %d", sortedArray)
	}
}

func TestQuickSort_one_element(t *testing.T) {
	array := []int{2}

	sortedArray := QuickSort(array)

	if !reflect.DeepEqual(sortedArray, []int{2}) {
		t.Errorf("Expected [2], got %d", sortedArray)
	}
}

func TestQuickSort_two_elements(t *testing.T) {
	array := []int{5, 1}

	sortedArray := QuickSort(array)

	if !reflect.DeepEqual(sortedArray, []int{1, 5}) {
		t.Errorf("Expected [1, 5], got %d", sortedArray)
	}
}

func TestQuickSort_three_elements(t *testing.T) {
	array := []int{5, 1, 4}

	sortedArray := QuickSort(array)

	if !reflect.DeepEqual(sortedArray, []int{1, 4, 5}) {
		t.Errorf("Expected [1, 4, 5], got %d", sortedArray)
	}
}

func TestQuickSort(t *testing.T) {
	array := []int{13, 3, 2, 6, 6, 8, 23, 11, 20, 34, 25}

	sortedArray := QuickSort(array)

	if !reflect.DeepEqual(sortedArray, []int{2, 3, 6, 6, 8, 11, 13, 20, 23, 25, 34}) {
		t.Errorf("Expected [2 3, 6, 6, 8, 11, 13, 20, 23, 25, 34], got %d", sortedArray)
	}
}

func TestQuickSort_test1(t *testing.T) {
	array := []int{13, 2, 6, 8, 23, 20, 34, 25}

	sortedArray := QuickSort(array)

	if !reflect.DeepEqual(sortedArray, []int{2, 6, 8, 13, 20, 23, 25, 34}) {
		t.Errorf("Expected [2, 6, 8, 13, 20, 23, 25, 34], got %d", sortedArray)
	}
}
