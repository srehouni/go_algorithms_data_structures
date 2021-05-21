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

	array2 := []int{13, 3, 2, 6, 6, 8, 23, 11, 20, 34, 25}
	fmt.Println(MergeSort(array2))
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

func MergeSort(array []int) []int {
	if len(array) < 2 {
		return array
	}

	if len(array) == 2 {
		if array[0] > array[1] {
			return []int{array[1], array[0]}
		} else {
			return array
		}
	}

	mid := (len(array) - 1) / 2
	array1 := MergeSort(array[:mid+1])
	array2 := MergeSort(array[mid+1:])

	return merge(array1, array2)
}

func merge(array1 []int, array2 []int) []int {
	sortedArray := []int{}
	i := 0
	j := 0

	for i < len(array1) || j < len(array2) {
		if i < len(array1) && j < len(array2) {
			if array1[i] < array2[j] {
				sortedArray = append(sortedArray, array1[i])
				i++
			} else {
				sortedArray = append(sortedArray, array2[j])
				j++
			}
		} else {
			if i < len(array1) {
				sortedArray = append(sortedArray, array1[i])
				i++
			} else {
				sortedArray = append(sortedArray, array2[j])
				j++
			}
		}
	}
	return sortedArray
}
