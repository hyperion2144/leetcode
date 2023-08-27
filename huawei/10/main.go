package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// 在一行中输入一个字符串数组，如果其中一个字符串的所有以索引0开头的子串在数组中都有，那么这个字符串就是潜在密码，在所有潜在密码中最长的是真正的密码，如果有多个长度相同的真正的密码，那么取字典序Q最大的为唯一的真正的密码，求唯—的真正的密码。
func main() {
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')

	strMaps := make(map[string]struct{})
	for _, str := range strings.Split(strings.TrimRight(line, "\n"), " ") {
		strMaps[str] = struct{}{}
	}

	var answer string
	for str := range strMaps {
		contains := true
		for i := 1; i <= len(str); i++ {
			subStr := str[:i]
			if _, ok := strMaps[subStr]; !ok {
				contains = false
			}
		}

		if contains && len(str) > len(answer) {
			answer = str
		} else if contains && len(str) == len(answer) && str > answer {
			answer = str
		}
	}

	fmt.Println(answer)
}
