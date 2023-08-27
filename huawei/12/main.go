package main

import "fmt"

// 输入一个长度为4的倍数的字符串，字符串中仅包含WASD四个字母。
// 将这个字符串中的连续子串用同等长度的仅包含WASD的字符串替换Q，如果替换后整个字符串中
// WASD四个字母出现的频数相同，那么我们称替换后的字符串是‘完美走位。
// 求子串的最小长度。
// 如果输入字符串已经平衡则输出0。
func main() {
	var s string
	fmt.Scan(&s)

	countMaps := make(map[rune]int)
	for _, c := range s {
		countMaps[c]++
	}

	balanceNumber := len(s) / 4
	changeChars := make(map[rune]int)
	changeCharNumber := 0

	allEqul := true
	for c, count := range countMaps {
		if count > balanceNumber {
			changeChars[c] = count - balanceNumber
			changeCharNumber += count - balanceNumber
			allEqul = false
		}
	}

	if allEqul {
		fmt.Println(0)
		return
	}

	answer := len(s)
	left, right := 0, changeCharNumber-1

	for i := 0; i < changeCharNumber; i++ {
		char := rune(s[i])
		if _, ok := changeChars[char]; ok {
			changeChars[char]--
		}
	}

	for left <= right && right < len(s) {
		t := true
		for _, count := range changeChars {
			if count > 0 {
				t = false
				break
			}
		}

		if t {
			answer = min(answer, right-left+1)
			left++
			char := rune(s[left-1])
			if _, ok := changeChars[char]; ok {
				changeChars[char]++
			}
		} else {
			right++
			if right < len(s) {
				char := rune(s[right])
				if _, ok := changeChars[char]; ok {
					changeChars[char]--
				}
			}
		}
	}

	fmt.Println(answer)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
