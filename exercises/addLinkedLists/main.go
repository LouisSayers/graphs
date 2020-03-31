package main

import (
	"fmt"
	"graphs/exercises/addLinkedLists/graph"
	"strconv"
)

func add(n1, n2 *graph.LinkedNode) (int64, error) {
	nextN1 := n1
	nextN2 := n2
  carry := 0
  result := ""

	for nextN1 != nil {
		sum := nextN1.Val + nextN2.Val + carry
		carry = sum / 10
		remainder := sum % 10
		result = fmt.Sprintf("%d%v", remainder, result)
		nextN1, nextN2 = nextN1.Next(), nextN2.Next()
	}

	result = fmt.Sprintf("%d%v", carry, result)

	return strconv.ParseInt(result, 10, 64)
}

func main() {
	//   293
	// + 487
	n1 := graph.NewNode(3)
	n2 := graph.NewNode(9)
	n3 := graph.NewNode(2)
	n1.SetNext(n2)
	n2.SetNext(n3)

	n4 := graph.NewNode(7)
	n5 := graph.NewNode(8)
	n6 := graph.NewNode(4)
	n4.SetNext(n5)
	n5.SetNext(n6)

	result, err := add(n1, n4)
	if err != nil {
		fmt.Printf("Received error: %v\n", err)
	} else {
		fmt.Printf("Result: %d\n", result)
	}
}
