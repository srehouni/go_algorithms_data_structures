package main

import "fmt"

func main() {
	array := []int{1, 2, 5, 6, 7, 19, 20, 32, 38, 45, 56, 123}

	fmt.Println(binarySearch(array, 19))
	fmt.Println(binarySearchRecursive(array, 19, 0, len(array)-1))
}

func binarySearch(array []int, num int) int {
	first := 0
	last := len(array) - 1

	for first <= last {
		mid := (first + last) / 2

		if array[mid] == num {
			return mid
		} else if array[mid] > num {
			last = mid - 1
		} else {
			first = mid + 1
		}
	}

	return -1
}

func binarySearchRecursive(array []int, num int, first int, last int) int {
	mid := (first + last) / 2

	if first > last {
		return -1
	}

	if array[mid] == num {
		return mid
	}

	if array[mid] > num {
		last = mid - 1
	} else {
		first = mid + 1
	}

	return binarySearchRecursive(array, num, first, last)
}
