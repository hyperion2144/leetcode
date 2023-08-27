package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Queue struct {
	data   []int
	isSort bool
}

func NewQueue() *Queue {
	return &Queue{
		data:   make([]int, 0),
		isSort: true,
	}
}

func (q *Queue) InsertFront(value int) {
	q.data = append([]int{value}, q.data...)
}

func (q *Queue) InsertLast(value int) {
	q.data = append(q.data, value)
}

func (q *Queue) DeleteFront() {
	q.data = q.data[1:]
}

// 给定一个队列，但是这个队列比较特殊，可以从头部添加数据，也可以从尾部添加数据，但是只能从头部删除数据。输入一个数字n，会依次添加数字1~n（也就是添加n次）。
// 但是在添加数据的过程中，也会删除数据，要求删除必须按照1~n按照顺序进行删除，所以在删除时，可以根据需要调整队列中数字的顺序以满足删除条件。
// 计算最少需要调整多少次才能满足删除条件。
func main() {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	n, _ := strconv.Atoi(strings.TrimRight(input, "\n"))

	count := 0
	queue := NewQueue()
	for i := 0; i < 2*n; i++ {
		input, _ := reader.ReadString('\n')
		items := strings.Split(strings.TrimRight(input, "\n"), " ")

		switch items[0] {
		case "head":
			val, _ := strconv.Atoi(items[2])
			queue.InsertFront(val)
			if len(queue.data) > 1 {
				queue.isSort = false
			}
		case "tail":
			val, _ := strconv.Atoi(items[2])
			queue.InsertLast(val)
		case "remove":
			if !queue.isSort {
				sort.Ints(queue.data)
				queue.isSort = true
				count++
			}
			queue.DeleteFront()
		}
	}

	fmt.Println(count)
}
