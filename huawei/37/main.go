package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// 输入单行英文句子，里面包含英文字母，空格以及.?三种字符，请将句子内每个单词进行倒序，并输出倒序后的语句
func main() {
	reader := bufio.NewReader(os.Stdin)

	answer := &strings.Builder{}
	words := make([]byte, 0)
	for b, _ := reader.ReadByte(); b != '\n'; b, _ = reader.ReadByte() {
		if b == '.' || b == ' ' || b == '?' {
			reverseArray(words)
			answer.Write(words)
			answer.WriteByte(b)

			words = words[:0]
			continue
		}

		words = append(words, b)
	}

	if len(words) > 0 {
		reverseArray(words)
		answer.Write(words)
	}

	fmt.Println(answer.String())
}

func reverseArray(nums []byte) {
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		tmp := nums[i]
		nums[i] = nums[j]
		nums[j] = tmp
	}
}
