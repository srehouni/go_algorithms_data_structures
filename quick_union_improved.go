//quick-union-lazy

package main

import (
	"fmt"
)

type QuickUnionImproved struct {
	elements   []int
	numObjetcs []int
}

func main() {

	quickUnionImproved := QuickUnionImproved{
		elements:   []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		numObjetcs: []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
	}

	quickUnionImproved.union(4, 3)
	quickUnionImproved.union(3, 8)
	quickUnionImproved.union(6, 5)
	quickUnionImproved.union(9, 4)
	quickUnionImproved.union(2, 1)

	fmt.Println(quickUnionImproved.connected(8, 9))
	fmt.Println(quickUnionImproved.connected(5, 4))

	quickUnionImproved.union(5, 0)
	quickUnionImproved.union(7, 2)
	quickUnionImproved.union(6, 1)
	quickUnionImproved.union(7, 3)

	fmt.Println(quickUnionImproved.connected(5, 4))
	fmt.Println(quickUnionImproved.connected(5, 0))
}

func (u QuickUnionImproved) getRoot(p int) int {

	var rootP = u.elements[p]

	for rootP != u.elements[rootP] {
		rootP = u.elements[rootP]
	}

	return rootP
}

func (u QuickUnionImproved) connected(p int, q int) bool {
	return u.getRoot(p) == u.getRoot(q)
}

func (u QuickUnionImproved) union(p int, q int) {

	var rootP = u.getRoot(p)
	var rootQ = u.getRoot(q)

	if rootP == rootQ {
		return
	}

	if u.numObjetcs[rootP] > u.numObjetcs[rootQ] {
		u.elements[rootQ] = rootP
		u.numObjetcs[rootP] += u.numObjetcs[rootQ]
	} else {
		u.elements[rootP] = rootQ
		u.numObjetcs[rootQ] += u.numObjetcs[rootP]
	}
}
