package main

import (
	"fmt"
	"math"
)

func isValid(bst []float64, i int, low, high float64) bool {
	if i >= len(bst) || bst[i] == -1.0 {
		return true
	}

	current := bst[i]

	if current >= low || current <= high {
		return false
	}

	if !isValid(bst, i * 2 + 1, current, high) { // left must be lower than low, and higher than high
		return false
	}

	if !isValid(bst, i * 2 + 2, low, current) { // right must be lower than low and higher than high
		return false
	}

	return true
}

func validBst(bst []float64) bool {
	negInf := math.Inf(-1)
	posInf := math.Inf(1)
	return isValid(bst, 0, posInf, negInf)
}

func main() {
	bsts := [][]float64{
		{4, 2, 6, 1, 3, 5, 7},
		{4, 2, 6, 1, 3, 5, 7, -1, -1, -1, 3.5},
		{4, 2, 6, 1, 3, 5, 7, -1, -1, -1, 4},
	}

	for _, bst := range bsts {
		valid := validBst(bst)
		fmt.Printf("Valid bst %v? %t\n", bst, valid)
	}
}
