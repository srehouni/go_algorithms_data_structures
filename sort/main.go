package main

import (
	"fmt"
	"math/rand"
)

func main() {
	array := []int{13, 2, 6, 8, 23, 20, 34, 25}
	fmt.Println(SelectionSort(array))

	array1 := []int{13, 3, 2, 6, 6, 8, 23, 11, 20, 34, 25}
	fmt.Println(QuickSort(array1))
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

func QuickSort(array []int) []int {
	if len(array) < 2 {
		return array
	}

	smallerItems := []int{}
	biggerItems := []int{}

	pivot := rand.Intn(len(array))

	for i := 0; i < len(array); i++ {
		if i != pivot {
			if array[i] < array[pivot] {
				smallerItems = append(smallerItems, array[i])
			} else {
				biggerItems = append(biggerItems, array[i])
			}
		}
	}

	return append(append(QuickSort(smallerItems), array[pivot]), QuickSort(biggerItems)...)
}
