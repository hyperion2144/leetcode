package main

import (
	"fmt"
	"sort"
)

// 给定坐标轴上的一组线段，线段的起点和終点均为整数并且长度不小手1，请你从中找到最少数量的线段，这些线段可以覆盖佳所有线段。
func main() {
	var n int
	fmt.Scan(&n)

	lines := make([][2]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d,%d\n", &lines[i][0], &lines[i][1])
	}

	sort.Slice(lines, func(i, j int) bool {
		if lines[i][0] == lines[j][0] {
			return lines[i][1] > lines[j][1]
		}
		return lines[i][0] < lines[j][0]
	})

	end, count := -1, 0
	for i, line := range lines {
		if line[0] >= end {
			count++
			end = line[1]
			continue
		}
		if line[1] <= end {
			continue
		}

		end = line[1]
		if i+1 < n && lines[i+1][0] >= end {
			count++
		}
		if i+1 == n {
			count++
		}
	}

	fmt.Println(count)
}
