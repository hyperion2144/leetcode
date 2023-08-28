package main

import (
	"fmt"
	"sort"
)

// 小明有n块木板，第i(1<=i<=n）块木板长度力ai。
// 小明买了一块长度为 m 的木料，这块木料可以切割成任意块，拼接到已有的木板上，用来加长木板。
// 小明想让最短的木板尽量长。
// 请问小明加长木板后，最短木板的长度可以为多少？
func main() {
	var n, m int
	fmt.Scan(&n, &m)

	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	sort.Ints(a)

	var ans int
	for i := 0; i < m; i++ {
		a[0] += 1
		sort.Ints(a)
		ans = a[0]
	}

	fmt.Println(ans)
}
