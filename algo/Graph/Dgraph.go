package graph

import queue "github.com/kldgleb/algorithms/01_dataTypes/02_queues"

// DirectedGraph
type DGraph struct {
	bag         map[int][]int
	vertexExist map[int]bool
	edgeTo      map[int]int
	marked      map[int]bool
	queue       queue.Queue
	E           int
	V           int
}

func NewDGraph() *DGraph {
	return &DGraph{
		bag:         make(map[int][]int, 0),
		vertexExist: make(map[int]bool),
		edgeTo:      make(map[int]int, 0),
		marked:      make(map[int]bool),
		queue:       queue.NewArrayQueue(),
		E:           0,
		V:           0,
	}
}

func (g *DGraph) AddEdge(v, w int) {
	if g.EdgeExist(v, w) {
		return
	}
	g.E++

	if !g.vertexExist[w] {
		g.V++
	}
	if !g.vertexExist[v] {
		g.V++
	}
	g.bag[v] = append(g.bag[v], w)
}

func (g *DGraph) EdgeExist(v, w int) bool {
	if !g.vertexExist[v] {
		return false
	}
	for _, ver := range g.bag[v] {
		if ver == w {
			return true
		}
	}
	return false
}

func (g *DGraph) DepthFirstSearch(s int) map[int]int {
	g.edgeTo = make(map[int]int, 0)
	g.marked = make(map[int]bool)
	g.dfs(s)
	return g.edgeTo
}

func (g *DGraph) BreadthFirstSearch(s int) map[int]int {
	g.edgeTo = make(map[int]int, 0)
	g.marked = make(map[int]bool)
	g.bfs(s)
	return g.edgeTo
}

func (g *DGraph) dfs(v int) {
	g.marked[v] = true
	for _, ver := range g.bag[v] {
		if !g.marked[ver] {
			g.dfs(ver)
			g.edgeTo[ver] = v
		}
	}
}

func (g *DGraph) bfs(v int) {
	if !g.marked[v] {
		g.queue.Enqueue(v)
	}
	g.marked[v] = true
	for !g.queue.IsEmpty() {
		dequedVertex, err := g.queue.Dequeue()
		if err != nil {
			panic(err)
		}
		for _, ver := range g.bag[dequedVertex] {
			if !g.marked[ver] {
				g.queue.Enqueue(ver)
				g.marked[ver] = true
				g.edgeTo[ver] = dequedVertex
			}
		}
	}
}
