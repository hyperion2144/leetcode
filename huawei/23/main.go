package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// 有一组区间[a0,b0],[a1,b1],…（a,b表示起点，終点)，区间有可能重叠、相邻，重叠或相邻则可以合井力更大的区间
// 给定一组连接器[x1，x2，x3，…](x表示连接器的最大可连接长度，即x>=gap），可用于将分离的区间连接起来，但两个分离区间之间只能使用1个连接器；
// 请编程实现使用连接器后，最少的区间数结果。
// 区间数<10000， a,b均 <=10000
// 连接器数量<10000；x <= 10000
func main() {
	intervals := make([][2]int, 0)
	connector := make([]int, 0)

	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	for _, in := range strings.Split(strings.TrimRight(line, "\n"), "]") {
		if in == "" {
			continue
		}
		var a, b int
		fmt.Sscanf(strings.TrimLeftFunc(in, func(r rune) bool {
			return r == '[' || r == ','
		}), "%d,%d", &a, &b)
		intervals = append(intervals, [2]int{a, b})
	}

	line, _ = reader.ReadString('\n')
	for _, x := range strings.Split(strings.TrimFunc(line, func(r rune) bool {
		return r == '[' || r == ']' || r == '\n'
	}), ",") {
		c, _ := strconv.Atoi(x)
		connector = append(connector, c)
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	sort.Ints(connector)

	temp := [][2]int{intervals[0]}
	distance := []int{}
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] <= temp[len(temp)-1][1] && intervals[i][1] > temp[len(temp)-1][1] {
			temp[len(temp)-1][1] = intervals[i][1]
		} else if intervals[i][0] > temp[len(temp)-1][1] {
			gap := intervals[i][0] - temp[len(temp)-1][1]
			distance = append(distance, gap)
			temp = append(temp, intervals[i])
		}
	}

	sort.Ints(distance)

	var count int
	for i, j := 0, 0; i < len(connector) && j < len(distance); i++ {
		if connector[i] >= distance[j] {
			j++
			count++
		}
	}

	fmt.Println(len(temp) - count)
}
