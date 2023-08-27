package main

import "fmt"

// 给应两个数组a，b，若a[i]=b[j] 则称[i,j]为一个二元组，求在给定的两个数组中，二元组的个数
func main() {
	var m, n int
	fmt.Scan(&m)
	a := make(map[int]int)
	for i := 0; i < m; i++ {
		var v int
		fmt.Scan(&v)

		a[v]++
	}

	fmt.Scan(&n)
	var count int
	for i := 0; i < n; i++ {
		var v int
		fmt.Scan(&v)

		count += a[v]
	}

	fmt.Println(count)
}
