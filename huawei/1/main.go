package main

import "fmt"

// 公司创新实验至正在研究如何最小化资源成本，退大化资源利用率，请你设计算法帮他们解决一个任务
// 混部问题：有taskNun项任务，每个任务有开始时间 (startTime），结束时间 (endTime），并行度
//
//	(parallelism)三个属性，并行度是指这个任务运行时将会占用的服务器数量，一个服务器在每个时刻可以被任意任务使用但最多被一个任务占用，任务运行完成立即释放（结束时刻不占用）。任务混部问题是指给定一批任务，让这批任务由同一批服务器承载运行，请你计算完成这批任务混部最少需要多少服务器，从而退大最大化控制资源成本。
//
// 第一行输入为taskNum，表示有taskNum项任务
// 接下来taskNum行，每行三个整数，表示每个任务的开始时间
// (startTime ），结束时间 (endTime ），并行度 (parallelism)
// 输出描述：
// 一个整数，表示最少需要的服务器数量

type MyServers map[int]int

func (s MyServers) update(start, end, l, r, idx, parallelism int) {
	// 不在区间内
	if r < start || end < l {
		return
	}

	// 比区间大, 或者与区间有交集
	if start <= l && r <= end {
		s[idx] += parallelism
		return
	} else {
		mid := (l + r) / 2
		s.update(start, end, l, mid, idx*2, parallelism)
		s.update(start, end, mid+1, r, idx*2+1, parallelism)
		s[idx] = max(s[idx*2], s[idx*2+1])
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func serverNumber(tasks [][3]int) int {
	s := make(MyServers)
	for _, task := range tasks {
		s.update(task[0], task[1], 0, 50001, 1, task[2])
	}
	return s[1]
}

func main() {
	var taskNum int
	fmt.Scan(&taskNum)

	tasks := make([][3]int, taskNum)
	for i := 0; i < taskNum; i++ {
		var startTime, endTime, parallelism int
		fmt.Scanf("%d %d %d", &startTime, &endTime, &parallelism)
		tasks[i] = [3]int{startTime, endTime, parallelism}
	}

	fmt.Println(serverNumber(tasks))
}
