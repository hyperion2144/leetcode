package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 有M(1<=M<=10）个端口組
// 每个端口組是长度为 N(1<=N<=100)的整数数组，如果端口组间存在2个及以上的端口相同，则认为这
// 2个端口组互相关联，可以合并
// 第一行输入端口组个数M，再输入M行，每行逗号分隔，代表端口组。输出合并后的端口组，用二维数组表示
func main() {
	var m int
	fmt.Scan(&m)

	scanner := bufio.NewScanner(os.Stdin)

	portGroups := make([]map[int]struct{}, m)
	for i := 0; i < m; i++ {
		scanner.Scan()
		ports := strings.Split(scanner.Text(), ",")
		portGroups[i] = make(map[int]struct{})
		for j := 0; j < len(ports); j++ {
			port, _ := strconv.Atoi(ports[j])
			portGroups[i][port] = struct{}{}
		}
	}

	fmt.Println(merge(portGroups))
}

func merge(portGroups []map[int]struct{}) [][]int {
	// Iterate over each port group in the slice
	for i := 0; i < len(portGroups); i++ {
		for j := i + 1; j < len(portGroups); j++ {
			group1 := portGroups[i]
			group2 := portGroups[j]
			// Merge the two groups into a single group
			repeat := 0
			for group := range group1 {
				if _, ok := group2[group]; ok {
					repeat++
				}
			}
			if repeat > 1 {
				for port := range group2 {
					group1[port] = struct{}{}
				}

				portGroups = append(portGroups[:j], portGroups[j+1:]...)
			}
		}
	}

	res := make([][]int, len(portGroups))
	for i := 0; i < len(portGroups); i++ {
		res[i] = make([]int, 0)
		for port := range portGroups[i] {
			res[i] = append(res[i], port)
		}
	}

	return res
}
