//weighted quick-union

package main

import (
	"fmt"
)

type QuickUnionWeighted struct {
	elements   []int
	numObjetcs []int
}

func main() {

	quickUnionWeighted := QuickUnionWeighted{
		elements:   []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		numObjetcs: []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
	}

	quickUnionWeighted.union(4, 3)
	quickUnionWeighted.union(3, 8)
	quickUnionWeighted.union(6, 5)
	quickUnionWeighted.union(9, 4)
	quickUnionWeighted.union(2, 1)

	fmt.Println(quickUnionWeighted.connected(8, 9))
	fmt.Println(quickUnionWeighted.connected(5, 4))

	quickUnionWeighted.union(5, 0)
	quickUnionWeighted.union(7, 2)
	quickUnionWeighted.union(6, 1)
	quickUnionWeighted.union(7, 3)

	fmt.Println(quickUnionWeighted.connected(5, 4))
	fmt.Println(quickUnionWeighted.connected(5, 0))
}

func (u QuickUnionWeighted) getRoot(p int) int {

	var rootP = u.elements[p]

	for rootP != u.elements[rootP] {
		rootP = u.elements[rootP]
	}

	return rootP
}

func (u QuickUnionWeighted) getRootWithPathCompression(p int) int {

	var rootP = u.elements[p]

	for rootP != u.elements[rootP] {
		u.elements[rootP] = u.elements[u.elements[rootP]]
		rootP = u.elements[rootP]
	}

	return rootP
}

func (u QuickUnionWeighted) connected(p int, q int) bool {
	return u.getRoot(p) == u.getRoot(q)
}

func (u QuickUnionWeighted) union(p int, q int) {

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
