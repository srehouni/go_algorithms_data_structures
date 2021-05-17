package main

import "fmt"

func main() {
	array := []int{1, 2, 5, 6, 7, 19, 20, 32, 38, 45, 56, 123}

	fmt.Println(binarySearch(array, 19))
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
