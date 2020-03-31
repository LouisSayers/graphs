package bfs

import (
	"fmt"
	"graphs/exercises/search/graph"
	"graphs/exercises/search/util"
)

type BFS struct {
	nodes []*graph.Node
	edges []*graph.Edge
}

func (bfs *BFS) Search(s *graph.Node) {
	util.ResetVisited(bfs.nodes)
	s.Visited = true
	queue := []*graph.Node{s}

	for len(queue) != 0 {
		n := queue[0]
		queue = queue[1:]

		for _, e := range n.Edges {
			n2 := e.N2
			if !n2.Visited {
				n2.Visited = true
				fmt.Printf("Visited (%v, %v)\n", n, n2)
				queue = append(queue, n2)
			}
		}
	}
}

func NewBFS(nodes []*graph.Node, edges []*graph.Edge) *BFS {
	return &BFS{nodes, edges}
}
