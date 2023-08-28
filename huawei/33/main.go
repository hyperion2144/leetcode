package main

import "fmt"

// 给定a-z, 26个英文字母小写字符组成的字符串A和B
// 其中A可能存在重复字母，B不会存在重复字母，现从字符串A中按规则挑选一些字母可以组成字符串B
// 挑选规则如下:
// 1. 同一个位置的字母只能挑选一次
// 2. 被挑选字母的相对先后顺序不能改变
func main() {
	var a, b string
	fmt.Scan(&a, &b)

	wait := make([]int, 0)
	var count int
	for i := 0; i < len(a); i++ {
		use := false
		for j, w := range wait {
			if a[i] == b[w] {
				if w+1 == len(b) {
					count++
					use = true
					wait = append(wait[:j], wait[j+1:]...)
					break
				}
				wait[j] = w + 1
				break
			}
		}
		if !use && a[i] == b[0] {
			wait = append(wait, 1)
		}
	}

	fmt.Println(count)
}
