package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 在一个狭小的路口，每秒只能通过一辆车，假设车辆的颜色只有3种，找出 N 秒内经过的最多颜色的车辆数量。
// 三种颜色编号为0,1,2
func colorNums(cars []int, n int) int {
	carCounts := make([]int, 3)
	for i := 0; i < n; i++ {
		carCounts[cars[i]]++
	}

	count := max(max(carCounts[0], carCounts[1]), carCounts[2])
	for i := n; i < len(cars); i++ {
		carCounts[cars[i]]++
		carCounts[cars[i-n]]--
		count = max(count, max(max(carCounts[0], carCounts[1]), carCounts[2]))
	}

	return count
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	var cars []int

	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	for _, color := range strings.Split(strings.TrimRight(line, "\n"), " ") {
		car, _ := strconv.Atoi(color)
		cars = append(cars, car)
	}
	line, _ = reader.ReadString('\n')
	n, _ := strconv.Atoi(strings.TrimRight(line, "\n"))

	fmt.Println(colorNums(cars, int(n)))
}
