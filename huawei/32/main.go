package main

import (
	"fmt"
	"sort"
)

// 张三要去外地出差，需要做核酸，需要在指定时间点前做完核酸，请帮他找到满足条件的核酸检测点。
// 给出一组核酸检测点的距离和每个核酸检测点当前的人数
// 给出张三要去做核酸的出发时间，出发时间是10的倍数，同时给出张三做核酸的最晚结束时间
// 题目中给出的距离是整数，单位是公里，时间1分钟为一基本单位
// 去找核酸点时，有如下的限制:
// - 去往核酸点的路上，每公里距离花费时间10分钟，费用是10元
// - 核酸点每检测一个人的时间花费是一分钟
// - 每个核酸点工作时间都是8点钟到20点钟中间不休息，核酸点准时工作，早到晚到都不检测
// - 核酸检测就结果可立刻知道
// 在张三去某个核酸点的路上花费的时间内，此核酸点的人数是动态变化的，变化规则是:
// - 在非核酸检测时间内，没有人排队
// - 8点到10点每分钟增加3人
// - 12点到14点每分钟增加10人
// 要求将所有满足条件的核酸检测点按照优选规则排序列出:
// - 花费时间少的核酸检测点拍在前面
// - 花费时间一样，花费费用最少的核酸检测点排在前面
// - 时间和花费一样，则ID值最小的拍在前面
func main() {
	var currentHour, currentMinute, targetHour, targetMinute int
	fmt.Scan(&currentHour, &currentMinute, &targetHour, &targetMinute)

	var n int
	fmt.Scan(&n)

	points := make([]struct {
		ID       int
		Distance int
		Wait     int
	}, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&points[i].ID, &points[i].Distance, &points[i].Wait)
	}

	var answer []struct {
		ID   int
		cost int
		Time int
	}
	for _, point := range points {
		cost := 10 * point.Distance
		startTime := currentHour*60 + currentMinute
		arrivalTime := startTime + point.Distance*10
		wait := point.Wait

		for i := startTime; i <= arrivalTime; i++ {
			if 480 < i && i < 600 {
				wait += 3
			}
			if 720 < i && i < 840 {
				wait += 10
			}
			if 480 < i && i < 1200 {
				wait = max(0, wait-1)
			}
		}
		finishTime := arrivalTime + wait
		if finishTime <= targetHour*60+targetMinute {
			answer = append(answer, struct {
				ID   int
				cost int
				Time int
			}{point.ID, cost, finishTime - startTime})
		}
	}

	sort.Slice(answer, func(i, j int) bool {
		if answer[i].Time == answer[j].Time {
			if answer[i].cost == answer[j].cost {
				return answer[i].ID < answer[j].ID
			}
			return answer[i].cost < answer[j].cost
		}
		return answer[i].Time < answer[j].Time
	})

	fmt.Println(len(answer))
	for _, ans := range answer {
		fmt.Printf("%d %d %d\n", ans.ID, ans.Time, ans.cost)
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
