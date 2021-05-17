//quick-union-lazy

package main

import (
	"fmt"
)

type QuickUnionLazy struct {
	elements []int
}

func main() {

	quickUnionLazy := QuickUnionLazy{elements: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}}

	quickUnionLazy.union(4, 3)
	quickUnionLazy.union(3, 8)
	quickUnionLazy.union(6, 5)
	quickUnionLazy.union(9, 4)
	quickUnionLazy.union(2, 1)
	fmt.Println(quickUnionLazy.connected(8, 9))
	fmt.Println(quickUnionLazy.connected(5, 4))
	quickUnionLazy.union(5, 0)
	quickUnionLazy.union(7, 2)
	quickUnionLazy.union(6, 1)
	quickUnionLazy.union(7, 3)
	fmt.Println(quickUnionLazy.connected(5, 0))
}

func (u QuickUnionLazy) getRoot(p int) int {

	var rootP = u.elements[p]

	for rootP != u.elements[rootP] {
		rootP = u.elements[rootP]
	}

	return rootP
}

func (u QuickUnionLazy) connected(p int, q int) bool {
	return u.getRoot(p) == u.getRoot(q)
}

func (u QuickUnionLazy) union(p int, q int) {
	u.elements[u.getRoot(p)] = u.getRoot(q)
}
