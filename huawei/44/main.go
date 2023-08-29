package main

import "fmt"

// 快递业务范围有N个站点，A站点和B站点可以中转快递，则认为A-B站点可达，如果A-B可达，B-C可达，则A-C可达。
// 现在给N个站点编号0、1、...n-1,用s[i][j]表示i-j是否可达，s[i][j]=1表示i-j可达,s[i][j]=0表示i-j不可达。
// 现用二维数组给定N个站点的可达关系，请计算至少选择从几个主站出发，才能可达所有站点（覆盖所有站点业）
// 说明: s[i][j]和s[j][i]取值相同
func main() {
	var n int
	fmt.Scan(&n)

	s := make([][]int, n)
	for i := 0; i < n; i++ {
		s[i] = make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Scan(&s[i][j])
		}
	}

	visited := make([]bool, n)
	count := 0
	for i := 0; i < n; i++ {
		if visited[i] {
			continue
		}

		queue := []int{}
		flag := false
		for j := 0; j < n; j++ {
			if s[i][j] == 1 && !visited[j] {
				queue = append(queue, j)
				flag = true
			}
		}

		for len(queue) > 0 {
			from := queue[0]
			queue = queue[1:]
			visited[from] = true
		}

		visited[i] = true

		if flag {
			count++
		}
	}

	fmt.Println(count)
}
