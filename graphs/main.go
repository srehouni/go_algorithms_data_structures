package main

import (
	"container/list"
	"fmt"
)

func main() {
	graph := make(map[string][]string)
	graph["you"] = []string{"bob", "claire", "alice"}
	graph["claire"] = []string{"thom", "jonny"}

	fmt.Println(BreadthFirstSearch(graph, "thom"))
}

func BreadthFirstSearch(graph map[string][]string, name string) bool {
	queue := list.New()
	searched := []string{}

	if name == "you" {
		return true
	}

	searched = append(searched, "you")
	addElementsToQueue(queue, graph["you"], searched)

	for queue.Len() > 0 {
		node := queue.Front()
		queue.Remove(node)
		if node.Value.(string) == name {
			return true
		} else {
			addElementsToQueue(queue, graph[node.Value.(string)], searched)
		}
	}

	return false
}

func addElementsToQueue(queue *list.List, elements []string, searched []string) {
	for _, s := range elements {
		if !contains(searched, s) {
			queue.PushBack(s)
		}
	}
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
