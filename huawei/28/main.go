package main

import "fmt"

// 给定两个字符串s1 和s2 和正整数k，其中 s1 长度为n1，s2长度为 n2
// 在s2中选一个子串，满足：
// 1：该子串长度为n1+k
// 2：该子串中包含s1中全部字母
// 3：该子串每个字母出现次数不小于s1中对应的字母
// 我们称s2以长度k冗余覆盖s1,
// 给定s1,s2,k,
// 求最左侧的S2以长度k冗余覆盖s1的子串的首个元素的下标，如果没有返回-1。
func main() {
	var (
		s1, s2 string
		k      int
	)
	fmt.Scan(&s1, &s2, &k)

	n1 := len(s1)
	n2 := len(s2)
	if n1+k > n2 {
		fmt.Println(-1)
		return
	}

	m1 := make(map[rune]int)
	for _, c := range s1 {
		m1[c]++
	}

	m := make(map[rune]int)
	for i := 0; i < n1+k; i++ {
		m[rune(s2[i])]++
	}

	for i := 0; i < n2-n1-k+1; i++ {
		if isSubstringCovered(m1, m) {
			fmt.Println(i)
			return
		}

		m[rune(s2[i])]--
		m[rune(s2[i+n1+k])]++
	}

	fmt.Println(-1)
}

// isSubstringCovered checks if the substring in s2 starting at index 0 and has a length n1+k
// satisfies the conditions of being a redundant cover of s1
func isSubstringCovered(s1, s2 map[rune]int) bool {
	for c, count := range s1 {
		if count > s2[c] {
			return false
		}
	}
	return true
}
