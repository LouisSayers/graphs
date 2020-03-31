package util

import "graphs/exercises/search/graph"

func ResetVisited(nodes []*graph.Node) {
	for _, n := range nodes {
		n.Visited = false
	}
}
