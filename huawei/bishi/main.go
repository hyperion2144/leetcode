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

	// 使用长度为3的数组的列表存储不同任务类型的种类数量、该种类的任务数量和当前冷却情况
	tasks := make([][3]int, 0)
	for _, taskNum := range taskNums {
		tasks = append(tasks, [3]int{len(taskNums), taskNum, -n - 1})
	}

	// 对任务数量进行排序,数量多的排在前面
	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i][1] > tasks[j][1]
	})

	var count int
	// 循环执行任务
	for i, k := 0, 0; len(tasks) > 0; k++ {
		count++
		task := tasks[i]

		if n+1-(k-task[2]) > 0 {
			continue
		}

		tasks[i][1]--
		tasks[i][2] = k
		tasks[i][0] = len(tasks)

		if tasks[i][1] == 0 {
			tasks = append(tasks[:i], tasks[i+1:]...)
			if len(tasks) > 0 {
				i = i % len(tasks)
			}
		} else {
			i = (i + 1) % len(tasks)
		}

	}

	fmt.Println(count)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
