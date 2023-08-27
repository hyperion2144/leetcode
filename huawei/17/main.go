package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 某公司研发了一款高性能AI处理路。每台物理设备具备8颗AI处理器，编号分别为0、1、2、3、4、5、6、7。
// 编号0-3的处理品处于同一个链路中，编号4-7的处理品处于另外一个链路中，不同链路中的处理器不能通信。
// 现给定服务品可用的处理品编号数组array，以及任务申请的处理品数是num，找出符合下列亲和性调度原则的芯片组合。
// 如果不存在符合要求的组合，则返回空列表。
// 亲和性调度原则：
// -如果申请处理器个数为1，则选择同一链路，剩余可用的处理器数量为1个的最佳，其次是剩余3个的为次佳，然后是剩余2个，最后是剩余4个。
// -如果申请处理器个数为2，则选择同一链路剩余可用的处理器数量2个的为最佳，其次是剩余4个，退后是剩余3个。
// -如果申请处理器个数为4，则必须选择同一链路剩余可用的处理器数为4个。
// 如果申请处理器个数为8，则申请节点所有8个处理器。
// 提示：
// 任务申请的处理器数最只能是1、2、4、 8。
// 编号0-3的处理路处于一个链路，编号4-7的处理器处于另外一个链路。
// 处理器编号唯一，且不存在相同编号处理品。
func main() {
	cpuGroup1 := make([]int, 0)
	cpuGroup2 := make([]int, 0)

	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	for _, cpu := range strings.Split(strings.TrimRight(line, "\n"), " ") {
		cpuNo, _ := strconv.Atoi(cpu)
		if cpuNo < 4 {
			cpuGroup1 = append(cpuGroup1, cpuNo)
		} else {
			cpuGroup2 = append(cpuGroup2, cpuNo)
		}
	}

	var num int
	fmt.Scan(&num)

	answer := make([][]int, 0)
	switch num {
	case 1:
		for _, l := range []int{1, 3, 2, 4} {
			find := false
			for _, g := range [][]int{cpuGroup1, cpuGroup2} {
				if len(g) == l {
					find = true
					for _, c := range g {
						answer = append(answer, []int{c})
					}
				}
			}
			if find {
				break
			}
		}
	case 2:
		for _, l := range []int{2, 4, 3} {
			find := false
			for _, g := range [][]int{cpuGroup1, cpuGroup2} {
				if len(g) == l {
					find = true
					for i := 0; i < l; i++ {
						for j := i; j < l; j++ {
							answer = append(answer, []int{g[i], g[j]})
						}
					}
				}
			}
			if find {
				break
			}
		}
	case 4:
		if len(cpuGroup1) == 4 {
			answer = append(answer, cpuGroup1)
		}
		if len(cpuGroup2) == 4 {
			answer = append(answer, cpuGroup2)
		}
	case 8:
		if len(cpuGroup1)+len(cpuGroup2) == 8 {
			answer = append(answer, append(cpuGroup1, cpuGroup2...))
		}
	}

	fmt.Println(answer)
}
