package main

import "fmt"

// 羊、狼、农夫都在岸边，当羊的数量小于狼的数量时，狼就会攻击羊，农夫则会损失羊。
// 农夫有一艘容量固定的船，能够承载固定数量的动物。
// 要求求出不损失羊的情况下将全部羊和狼运到对岸需要的最小次数。
// 只计算农夫去对岸的次数，回程时农夫不会运送羊和狼。
// 备注：农夫在或者农夫离开后羊的数量大于狼的数量时狼不会攻击羊。
// 农夫自身不占用船的容量.
func main() {
	var sheep, wolf, cap int
	fmt.Scan(&sheep, &wolf, &cap)

	var sheepi, wolfi, count int
	for wolf+sheep > 0 {
		count++
		isCan := true

		for x := cap; x > 0; x-- {
			if wolf+sheep == x {
				sheepi += sheep
				sheep = 0
				wolfi += wolf
				wolf = 0
				break
			}
			// 全部运送羊后，若羊数量大于狼数量
			if sheep-x > wolf {
				sheepi += sheep - x
				sheep -= x
				break
			}
			// 全部运送狼后，若羊数量大于狼数量
			if wolfi+cap < sheepi {
				wolfi += wolf - x
				wolf -= x
				break
			}

			if sheepi == 0 && x-1 > 0 {
				wolfi += x - 1
				wolf -= x - 1
				break
			}

			isCan = false
		}

		if !isCan {
			count = 0
			break
		}
	}

	fmt.Println(count)
}
