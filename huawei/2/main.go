package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

// 给定一个字符串s，s包括以空格分隔的若干个单词，请对s进行如下处理后输出：
// 1、单词内部调整：对每个单词字母重新按字典序排序
// 2、单词间顺序调整：
// 1） 统计每个单词出现的次数，并按次数降序排列
// 2）次数相同，按单词长度升序排列
// 3）次数和单词长度均相同，按字典升序排列
// 请输出处理后的字符串，每个单词以一个空格分隔。
func main() {
	scanner := bufio.NewReader(os.Stdin)
	input, _, _ := scanner.ReadLine()

	words := make(map[string]int)
	var slow int
	for fast := 1; fast < len(input); fast++ {
		if input[fast] != ' ' {
			continue
		}

		words[string(input[slow:fast])]++
		slow = fast + 1
	}
	words[string(input[slow:])]++

	countWords := make(map[int][]string, len(words))
	counts := make([]int, 0)
	for word, count := range words {
		countWords[count] = append(countWords[count], word)
	}
	for count := range countWords {
		counts = append(counts, count)
	}

	sort.Slice(counts, func(i, j int) bool {
		return counts[i] > counts[j]
	})

	ans := &strings.Builder{}
	for _, count := range counts {
		words := countWords[count]

		sort.Slice(words, func(i, j int) bool {
			if len(words[i]) != len(words[j]) {
				return len(words[i]) < len(words[j])
			}
			return words[i] < words[j]
		})

		for _, word := range words {
			word := []rune(word)
			sort.Slice(word, func(i, j int) bool {
				return word[i] < word[j]
			})

			for i := 0; i < count; i++ {
				ans.WriteString(string(word))
				ans.WriteByte(' ')
			}
		}
	}
	answer := strings.TrimRight(ans.String(), " ")
	fmt.Println(answer)
}
