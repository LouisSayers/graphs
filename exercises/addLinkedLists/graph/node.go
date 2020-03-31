package graph

type LinkedNode struct {
	Val int
	next *LinkedNode
}

func (n *LinkedNode) Next() *LinkedNode {
	return n.next
}

func (n *LinkedNode) SetNext(otherN *LinkedNode) {
	n.next = otherN
}

func (n *LinkedNode) String() string {
	return string(n.Val)
}

func NewNode(val int) *LinkedNode {
	return &LinkedNode{Val: val}
}
