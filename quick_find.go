//quick-find

package main

import (
	"fmt"
)

type QuickFind struct {
	elements []int
}

func main() {

	quickFind := QuickFind{elements: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}}

	quickFind.union(4, 3)
	quickFind.union(3, 8)
	quickFind.union(6, 5)
	quickFind.union(9, 4)
	quickFind.union(2, 1)
	quickFind.union(8, 9)
	quickFind.union(5, 0)
	quickFind.union(7, 2)
	quickFind.union(6, 1)

	fmt.Println(quickFind.connected(6, 7))
	fmt.Println(quickFind.connected(0, 3))
	fmt.Println(quickFind.connected(4, 8))
}

func (u QuickFind) union(p int, q int) {

	if u.elements[p] == u.elements[q] {
		return
	}

	oldValue := u.elements[p]

	for index, element := range u.elements {
		if element == oldValue {
			u.elements[index] = u.elements[q]
		}
	}
}

func (u QuickFind) connected(p int, q int) bool {
	return u.elements[p] == u.elements[q]
}
