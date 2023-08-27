package main

// Given the number of nodes in an undirected graph and all the edges,
// where each node can be colored red or black, but with the constraint
// that two nodes connected by an edge cannot be colored the same,
// calculate the number of different coloring schemes for the graph.
func caseNumbers(n int, edges [][]int) int {
	// Create an adjacency list representation of the graph
	adj := make([][]int, n+1)
	for _, edge := range edges {
		u, v := edge[0], edge[1]
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	// Create a color array to store the color of each node
	colors := make([]int, n+1)

	// Call the recursive helper function to calculate the number of coloring schemes
	return helper(1, adj, colors)
}

// Recursive helper function to calculate the number of coloring schemes
func helper(node int, adj [][]int, colors []int) int {
	// Base case: If all nodes have been colored, return 1 (valid coloring scheme)
	if node == len(colors) {
		return 1
	}

	// Initialize the count of coloring schemes
	count := 0

	// Try coloring the current node with red color (color 1)
	colors[node] = 1
	valid := true

	// Check if the current coloring is valid (i.e., no adjacent nodes have the same color)
	for _, neighbor := range adj[node] {
		if colors[neighbor] == colors[node] {
			valid = false
			break
		}
	}

	// If the current coloring is valid, recursively call the helper function for the next node
	if valid {
		count += helper(node+1, adj, colors)
	}

	// Try coloring the current node with black color (color 2)
	colors[node] = 2
	valid = true

	// Check if the current coloring is valid (i.e., no adjacent nodes have the same color)
	for _, neighbor := range adj[node] {
		if colors[neighbor] == colors[node] {
			valid = false
			break
		}
	}

	// If the current coloring is valid, recursively call the helper function for the next node
	if valid {
		count += helper(node+1, adj, colors)
	}

	// Reset the color of the current node
	colors[node] = 0

	// Return the final count of coloring schemes
	return count
}
