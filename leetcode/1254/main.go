package main

import "fmt"

func closedIsland(grid [][]int) int {
	directions := [][]int{
		{-1, 0},
		{0, -1},
		{0, 1},
		{1, 0},
	}

	row := len(grid)
	col := len(grid[0])

	var ans int
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if grid[i][j] == 1 {
				continue
			}

			grid[i][j] = 1

			isCloed := true
			queue := [][2]int{{i, j}}
			for len(queue) > 0 {
				from := queue[0]
				queue = queue[1:]

				if from[0] == 0 || from[0] == row-1 || from[1] == 0 || from[1] == col-1 {
					isCloed = false
				}

				for _, dir := range directions {
					to := [2]int{from[0] + dir[0], from[1] + dir[1]}
					if 0 <= to[0] && to[0] < row && 0 <= to[1] && to[1] < col && grid[to[0]][to[1]] == 0 {
						grid[to[0]][to[1]] = 1
						queue = append(queue, to)
					}
				}
			}

			if isCloed {
				ans++
			}

		}
	}
	return ans
}

func main() {
	fmt.Println(closedIsland([][]int{
		{1, 1, 1, 1, 1, 1, 1, 0},
		{1, 0, 0, 0, 0, 1, 1, 0},
		{1, 0, 1, 0, 1, 1, 1, 0},
		{1, 0, 0, 0, 0, 1, 0, 1},
		{1, 1, 1, 1, 1, 1, 1, 0},
	}))
}
