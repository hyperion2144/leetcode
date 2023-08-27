package main

import "fmt"

// findCircleNum uses union-find algorithm to find the number of connected components in a graph.
func findCircleNum(isConnected [][]int) int {
	n := len(isConnected)

	// Create a parent array to store the parent of each node
	parent := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
	}

	// Helper function to find the parent of a node
	var find func(int) int
	find = func(x int) int {
		if parent[x] != x {
			parent[x] = find(parent[x])
		}
		return parent[x]
	}

	// Helper function to merge two nodes into the same set
	merge := func(x, y int) {
		parent[find(x)] = find(y)
	}

	// Iterate through the matrix and merge connected nodes
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if isConnected[i][j] == 1 {
				merge(i, j)
			}
		}
	}

	// Count the number of unique parent nodes
	count := 0
	for i := 0; i < n; i++ {
		if parent[i] == i {
			count++
		}
	}

	return count
}

func main() {
	fmt.Printf("%d\n", findCircleNum([][]int{
		{1, 1, 1},
		{1, 1, 0},
		{1, 0, 1},
	}))
}
