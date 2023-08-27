package main

import (
	"fmt"
	"strings"
)

// 数字0、1、2、3.4、5、6、7、8、9分别关联a~z 26个英文字母。
// 0关联 "a”,"b","c"
// 1关睽 “d”,"e","f"
// 2关联 "g","h","i”
// 3关眹 "j","k","l”
// 4关联 "m","n","o”
// 5关联 "p","q","r"
// 6关联 "s","t"
// 7关联 "u","v"
// 8关朕 "w","x"
// 9关联 "y","z”
// 例如7关联 u","v'，8关联"x”，w'，输入一个字符串例如"78"，和一个屏蔽字符串"ux",那么"78”可以組成多个字符串例如：“ux”，“uw","vx”，“vw”，过滤这些完全包含屏蔽字符串的每一个字符的字符串，然后输出剩下的字符串，
func main() {
	relation := map[int][]string{
		0: {"a", "b", "c"},
		1: {"d", "e", "f"},
		2: {"g", "h", "i"},
		3: {"j", "k", "l"},
		4: {"m", "n", "o"},
		5: {"p", "q", "r"},
		6: {"s", "t"},
		7: {"u", "v"},
		8: {"w", "x"},
		9: {"y", "z"},
	}

	var numberStr, maskStr string
	fmt.Scan(&numberStr, &maskStr)

	var result []string
	var generator func(string, string)
	generator = func(s1, s2 string) {
		if len(s2) == 0 {
			if !strings.Contains(s1, maskStr) {
				result = append(result, s1)
			}
			return
		}

		digit := s2[0] - '0'
		letters := relation[int(digit)]
		for _, letter := range letters {
			generator(s1+letter, s2[1:])
		}
	}

	generator("", numberStr)
	fmt.Println(strings.Join(result, " "))
}
