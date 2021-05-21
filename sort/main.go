package main

import (
	"fmt"
	"math/rand"
)

func main() {
	array := []int{13, 2, 6, 8, 23, 20, 34, 25}
	fmt.Println(SelectionSort(array))
	counter(10)
	fmt.Println(factorial(10))
	fmt.Println(sum([]int{1, 2, 3, 4, 5}, 0))
	fmt.Println(countItems([]int{1, 2, 3, 4, 5}))
	fmt.Println(maxItem([]int{9, 2, 3, 4, 5}, 1))
	fmt.Println(maxItem2([]int{9, 2, 3, 4, 5}))
}

func counter(num int) {
	fmt.Println(num)
	if num > 0 {
		counter(num - 1)
	}
}

func factorial(num int) int {
	if num <= 1 {
		return 1
	} else {
		return num * factorial(num-1)
	}
}

func sum(array []int, index int) int {
	if index == len(array)-1 {
		return array[index]
	}
	return array[index] + sum(array, index+1)
}

func countItems(array []int) int {
	if len(array) == 1 {
		return 1
	}
	return 1 + countItems(array[1:])
}

func maxItem(array []int, max int) int {
	if max < array[0] {
		max = array[0]
	}

	if len(array) == 1 {
		return max
	}

	return maxItem(array[1:], max)
}

func maxItem2(array []int) int {
	if len(array) == 2 {
		if array[0] > array[1] {
			return array[0]
		} else {
			return array[1]
		}
	}

	max := maxItem2(array[1:])

	if max > array[0] {
		return max
	} else {
		return array[0]
	}
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
