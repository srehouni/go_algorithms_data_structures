package main

import (
	"fmt"
)

func main() {
	array := []int{13, 2, 6, 8, 23, 20, 34, 25}
	fmt.Println(SelectionSort(array))
}

func SelectionSort(array []int) []int {
	var sortedArray = []int{}
	slice := array
	for len(slice) > 0 {
		index := indexOfSmaller(slice)
		sortedArray = append(sortedArray, slice[index])
		slice = append(slice[:index], slice[index+1:]...)
	}

	return sortedArray
}

func indexOfSmaller(array []int) int {
	var smallest = array[0]
	var smallestIndex = 0

	for i := 0; i < len(array); i++ {
		if smallest > array[i] {
			smallest = array[i]
			smallestIndex = i
		}
	}

	return smallestIndex
}
