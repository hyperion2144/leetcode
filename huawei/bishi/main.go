package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// 使用map存放不同任务类型的任务数量，便于统计
	taskNums := make(map[int]int)
	// 冷却时间
	n := 0

	// 处理输入
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	for _, digit := range strings.Split(scanner.Text(), ",") {
		if digit == "" {
			continue
		}
		taskKind, _ := strconv.Atoi(digit)
		taskNums[taskKind]++
	}
	scanner.Scan()
	n, _ = strconv.Atoi(scanner.Text())

	if len(taskNums) == 0 {
		fmt.Println(0)
		return
	}

	// 使用数组存储不同任务类型的任务数量
	tasks := make([]int, 0)
	sum := 0
	for _, taskNum := range taskNums {
		tasks = append(tasks, taskNum)
		sum += taskNum
	}

	// 对任务数量进行排序,数量多的排在前面
	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i] > tasks[j]
	})

	// 计算最大任务数量
	maxCount := tasks[0]
	equalCount := 0
	for i := 1; i < len(tasks); i++ {
		if tasks[i] == maxCount {
			equalCount++
		} else {
			break
		}
	}

	fmt.Println(max((maxCount-1)*(n+1)+1+equalCount, sum))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
