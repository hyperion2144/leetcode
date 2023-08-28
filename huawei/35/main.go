package main

import "fmt"

// 有一批箱子（形式为字符串，设为str）
// 要求将这批箱子按从上到下以之字形的顺序摆放在宽度为n的空地，请输出箱子的拜访位置
func main() {
	var (
		str string
		n   int
	)
	fmt.Scanf("%s %d", &str, &n)

	answer := make([][]byte, n)
	for i := 0; i < n; i++ {
		answer[i] = make([]byte, len(str))
	}

	position, up := 0, 0
	for i := 0; i < len(str); i++ {
		answer[position] = append(answer[position], str[i])
		if up == 1 {
			position--
			if position == -1 {
				position = 0
				up = 0
			}
		} else {
			position++
			if position == n {
				position = n - 1
				up = 1
			}
		}
	}

	for i := 0; i < n; i++ {
		fmt.Println(string(answer[i]))
	}
}
