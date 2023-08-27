package main

import (
	"fmt"
	"sort"
)

// 在星球争霸篮球赛对抗赛中，最大的宇宙战队希望每个人都能拿到 MVP，MVP的条件是单场最高分得分获得者。
// 可以并列所以宇宙战队决定在比寒中尽可能让更多队员上场，并且让所有得分的选手得分都相同
// 然而比襄过程中的每1分钟的得分都只能由某一个人包搅。输出有得分的队员都是MVP时，最少的MVP得分。
func main() {
	var minutes int
	fmt.Scan(&minutes)

	score := make([]int, minutes)
	sum := 0
	for i := 0; i < minutes; i++ {
		fmt.Scan(&score[i])
		sum += score[i]
	}

	sort.Ints(score)

	answer := score[len(score)-1]
	for i := 0; i < len(score); i++ {
		if i == len(score)-1 {
			fmt.Println(sum)
			return
		}
		if sum%answer != 0 {
			answer += score[i]
			continue
		}

		sums := make([]int, 1<<minutes)
		dp := make([]bool, 1<<minutes)
		dp[0] = true
		for j := 0; j < 1<<minutes; j++ {
			if !dp[j] {
				continue
			}

			for k := 0; k < minutes; k++ {
				if sums[j]+score[k] > answer {
					break
				}

				if j&(1<<k) == 0 {
					next := j | (1 << k)
					if !dp[next] {
						dp[next] = true
						sums[next] = (sums[j] + score[k]) % answer
					}
				}
			}
		}

		if dp[(1<<minutes)-1] {
			fmt.Println(answer)
			return
		}
	}
}
