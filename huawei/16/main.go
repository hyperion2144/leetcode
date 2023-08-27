package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 日志采集是运维系统的的核心组件。日志是按行生成，每行记做一条，由来集系统分批上报。
// 如果上报太频繁，会对服务端造成压力;如果上报太晚，会降低用户的体验,如果一次上报的条数太多，会导致超时失败。为此，项目组设计了如下的上报策路
// 1、每成功上报一条日志，奖励1分
// 2、每条日志每延迟上报1秒，扣1分
// 3、积累日志达到100条，必须立即上报
// 给出日志序列，根据该规则，计算首次上报能获得的晨多积分数
func main() {
	var logSequence []int
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	for _, log := range strings.Split(strings.TrimRight(line, "\n"), " ") {
		seq, _ := strconv.Atoi(log)
		logSequence = append(logSequence, seq)
	}

	var score, totalScore, delScore int
	for _, seq := range logSequence {
		delScore += totalScore
		totalScore += seq

		if totalScore >= 100 {
			if 100-delScore > score {
				score = 100 + delScore
			}
			break
		}

		if totalScore-delScore > score {
			score = totalScore - delScore
		}

	}

	fmt.Println(score)
}
