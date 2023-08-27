package main

func shortestPathBinaryMatrix(grid [][]int) int {
	// check if the grid is empty or has only one cell
	if len(grid) == 0 || len(grid[0]) == 0 {
		return -1
	}

	// check if the start and end cells are blocked
	if grid[0][0] == 1 || grid[len(grid)-1][len(grid[0])-1] == 1 {
		return -1
	}

	// create a queue to store the cells to be visited
	queue := [][]int{{0, 0}}

	// set the distance of the start cell to 1
	grid[0][0] = 1

	// define the directions for moving to adjacent cells
	directions := [][]int{
		{-1, -1},
		{-1, 0},
		{-1, 1},
		{0, -1},
		{0, 1},
		{1, -1},
		{1, 0},
		{1, 1},
	}

	// perform breadth-first search to find the shortest path
	for len(queue) > 0 {
		// dequeue a cell from the queue
		cell := queue[0]
		queue = queue[1:]

		// check if the current cell is the end cell
		if cell[0] == len(grid)-1 && cell[1] == len(grid[0])-1 {
			return grid[cell[0]][cell[1]]
		}

		// explore the neighboring cells
		for _, dir := range directions {
			// calculate the coordinates of the neighboring cell
			newRow := cell[0] + dir[0]
			newCol := cell[1] + dir[1]

			// check if the neighboring cell is within the grid boundaries
			if newRow >= 0 && newRow < len(grid) && newCol >= 0 && newCol < len(grid[0]) {
				// check if the neighboring cell is not blocked
				if grid[newRow][newCol] == 0 {
					// enqueue the neighboring cell
					queue = append(queue, []int{newRow, newCol})

					// update the distance of the neighboring cell
					grid[newRow][newCol] = grid[cell[0]][cell[1]] + 1
				}
			}
		}
	}

	return -1
}
