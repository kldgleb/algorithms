package graph

type Grapher interface {
	AddEdge(v, w int)
	DepthFirstSearch(s int) map[int]int
	BreadthFirstSearch(s int) map[int]int
	EdgeExist(v, w int) bool
}
