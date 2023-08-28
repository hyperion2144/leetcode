package main

import (
	"fmt"
)

// 给定一个mxn的矩阵，由若干个字符x和o租成，x表示该处已被占据，o表示该处空闲
// 请找到最大的单入口空闲区域
func main() {
	var m, n int
	fmt.Scan(&m, &n)

	matrix := make([][]string, m)
	for i := 0; i < m; i++ {
		matrix[i] = make([]string, n)
		for j := 0; j < n; j++ {
			fmt.Scan(&matrix[i][j])
		}
	}

	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}

	regions := make([][3]int, 0)
	direction := [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if visited[i][j] || matrix[i][j] == "X" {
				continue
			}

			if matrix[i][j] == "O" {
				queue := [][2]int{{i, j}}
				regionSize := 0
				numberOfEntrances := 0
				entrancePoints := make([][2]int, 0)

				for len(queue) > 0 {
					from := queue[0]
					queue = queue[1:]

					if visited[from[0]][from[1]] {
						continue
					}

					visited[from[0]][from[1]] = true
					regionSize++
					if from[0] == 0 || from[0] == m-1 || from[1] == 0 || from[1] == n-1 {
						numberOfEntrances++
						entrancePoints = append(entrancePoints, [2]int{from[0], from[1]})
					}

					for _, d := range direction {
						dstx, dsty := from[0]+d[0], from[1]+d[1]
						if dstx >= 0 && dstx < m && dsty >= 0 && dsty < n && matrix[dstx][dsty] == "O" && !visited[dstx][dsty] {
							queue = append(queue, [2]int{dstx, dsty})
						}
					}
				}

				if numberOfEntrances == 1 {
					if len(regions) == 0 {
						regions = append(regions, [3]int{regionSize, entrancePoints[0][0], entrancePoints[0][1]})
					} else {
						if regionSize > regions[0][0] {
							regions = [][3]int{{regionSize, entrancePoints[0][0], entrancePoints[0][1]}}
						} else if regionSize == regions[0][0] {
							regions = append(regions, [3]int{regionSize, entrancePoints[0][0], entrancePoints[0][1]})
						}
					}
				}
			}
		}
	}

	if len(regions) == 0 {
		fmt.Println("NULL")
	}
	if len(regions) == 1 {
		fmt.Println(regions[0])
	}
	if len(regions) > 1 {
		fmt.Println(regions[0][0])
	}
}
