package main

import (
	"testing"
)

func TestBinary_search_found(t *testing.T) {
	array := []int{1, 3, 5, 6, 7, 9, 10, 11, 23, 34, 43, 234, 345, 2323}
	num := 23

	index := binarySearch(array, num)

	if index != 8 {
		t.Errorf("Expected 8, got %d", index)
	}
}

func TestBinary_search_recursive_found(t *testing.T) {
	array := []int{1, 3, 5, 6, 7, 9, 10, 11, 23, 34, 43, 234, 345, 2323}
	num := 23

	index := binarySearchRecursive(array, num, 0, len(array)-1)

	if index != 8 {
		t.Errorf("Expected 8, got %d", index)
	}
}

func TestBinary_search_not_found(t *testing.T) {
	array := []int{1, 3, 5, 6, 7, 9, 10, 11, 23, 34, 43, 234, 345, 2323}
	num := 24

	index := binarySearch(array, num)

	if index != -1 {
		t.Errorf("Expected -1, got %d", index)
	}
}

func TestBinary_search_recursive_not_found(t *testing.T) {
	array := []int{1, 3, 5, 6, 7, 9, 10, 11, 23, 34, 43, 234, 345, 2323}
	num := 24

	index := binarySearchRecursive(array, num, 0, len(array)-1)

	if index != -1 {
		t.Errorf("Expected -1, got %d", index)
	}
}

func TestBinary_search_empty_array(t *testing.T) {
	array := []int{}
	num := 24

	index := binarySearch(array, num)

	if index != -1 {
		t.Errorf("Expected -1, got %d", index)
	}
}

func TestBinary_search_recursive_empty_array(t *testing.T) {
	array := []int{}
	num := 24

	index := binarySearchRecursive(array, num, 0, len(array)-1)

	if index != -1 {
		t.Errorf("Expected -1, got %d", index)
	}
}

func TestBinary_search_one_item_found_array(t *testing.T) {
	array := []int{2}
	num := 2

	index := binarySearch(array, num)

	if index != 0 {
		t.Errorf("Expected 0, got %d", index)
	}
}

func TestBinary_search_recursive_one_item_found_array(t *testing.T) {
	array := []int{2}
	num := 2

	index := binarySearchRecursive(array, num, 0, len(array)-1)

	if index != 0 {
		t.Errorf("Expected 0, got %d", index)
	}
}

func TestBinary_search_one_item_not_found_array(t *testing.T) {
	array := []int{2}
	num := 3

	index := binarySearch(array, num)

	if index != -1 {
		t.Errorf("Expected -1, got %d", index)
	}
}

func TestBinary_search_recursive_one_item_not_found_array(t *testing.T) {
	array := []int{2}
	num := 3

	index := binarySearchRecursive(array, num, 0, len(array)-1)

	if index != -1 {
		t.Errorf("Expected -1, got %d", index)
	}
}
