package graph

type Edge struct {
	N1 *Node
	N2 *Node
}

func NewEdge(a, b *Node) *Edge {
	e := &Edge{a, b}
	a.Add(e)
	b.Add(e)
	return e
}
