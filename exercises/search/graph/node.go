package graph

type Node struct {
	Val string
	Edges []*Edge
	Visited bool
}

func (n *Node) Add(e *Edge) {
	n.Edges = append(n.Edges, e)
}

func (n *Node) String() string {
	return n.Val
}

func NewNode(val string) *Node {
	return &Node{Val: val, Visited: false}
}
