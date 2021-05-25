package main

import (
	"testing"
)

func TestBreadthFirstSearch(t *testing.T) {
	graph := make(map[string][]string)
	graph["you"] = []string{"bob", "claire", "alice"}
	graph["claire"] = []string{"thom", "jonny"}

	if !BreadthFirstSearch(graph, "thom") {
		t.Errorf("Expected true, got false")
	}
}
