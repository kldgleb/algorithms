package main

import (
	"fmt"

	graph "github.com/kldgleb/algorithms/algo/Graph"
)

func main() {
	g := graph.NewUdGraph()
	g.AddEdge(0, 1)
	g.AddEdge(0, 2)
	g.AddEdge(0, 6)
	g.AddEdge(0, 5)
	g.AddEdge(6, 4)
	g.AddEdge(4, 3)
	g.AddEdge(4, 5)
	g.AddEdge(5, 3)
	g.AddEdge(5, 4)

	g.AddEdge(7, 8)

	g.AddEdge(10, 9)
	g.AddEdge(9, 11)
	g.AddEdge(9, 12)
	g.AddEdge(11, 12)

	fmt.Println(g.BreadthFirstSearch(0))
}
