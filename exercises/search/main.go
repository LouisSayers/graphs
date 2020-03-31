package main

import (
	"graphs/exercises/search/bfs"
	"graphs/exercises/search/graph"
)

func main() {
	a := graph.NewNode("a")
	b := graph.NewNode("b")
	c := graph.NewNode("c")
	d := graph.NewNode("d")
	e := graph.NewNode("e")
	e1 := graph.NewEdge(a, b)
	e2 := graph.NewEdge(a, c)
	e3 := graph.NewEdge(b, d)
	e4 := graph.NewEdge(c, d)
	e5 := graph.NewEdge(d, e)

	nodes := []*graph.Node{a, b, c, d, e}
	edges := []*graph.Edge{e1, e2, e3, e4, e5}

  bfs := bfs.NewBFS(nodes, edges)
  bfs.Search(a)
}
