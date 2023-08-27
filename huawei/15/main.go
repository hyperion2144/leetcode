package main

import "fmt"

// 给一块n*m的地块，相当于n*m的二维数组，每个元素的值表示这个小地块的发电量；求在这块地上建立正方形的边长为c的发电站，发电量满足目标电量k的地块数量
// 前缀和求解
func main() {
	var n, m, c, k int
	fmt.Scan(&n, &m, &c, &k)

	ground := make([][]int, n)
	for i := 0; i < n; i++ {
		ground[i] = make([]int, m)
		for j := 0; j < m; j++ {
			fmt.Scan(&ground[i][j])
		}
	}

	// 生成前缀和
	s := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		s[i] = make([]int, m+1)
	}

	for i := 1; i < n+1; i++ {
		for j := 1; j < m+1; j++ {
			s[i][j] = s[i-1][j] + s[i][j-1] - s[i-1][j-1] + ground[i-1][j-1]
		}
	}

	var kind int
	for i := c; i < n+1; i++ {
		for j := c; j < m+1; j++ {
			if s[i][j]-s[i-c][j]-s[i][j-c]+s[i-c][j-c] >= k {
				kind++
			}
		}
	}

	fmt.Println(kind)
}
