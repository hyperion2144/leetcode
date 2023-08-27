package main

import (
	"fmt"
	"sort"
)

type Printer struct {
	Queue [][2]int
	Sort  bool
}

func (p *Printer) Insert(no, priority int) {
	p.Queue = append(p.Queue, [2]int{no, priority})
	p.Sort = false
}

func (p *Printer) Out() int {
	if len(p.Queue) == 0 {
		return -1
	}

	if !p.Sort {
		sort.Slice(p.Queue, func(i, j int) bool {
			if p.Queue[i][1] == p.Queue[j][1] {
				return p.Queue[i][0] < p.Queue[j][0]
			}
			return p.Queue[i][1] > p.Queue[j][1]
		})

		p.Sort = true
	}

	v := p.Queue[0]
	p.Queue = p.Queue[1:]
	return v[0]
}

// 有5台打印机打印文件，每台打印机有自己的待打印队列。因为打印的文件内容有轻重缓急之分，所以队列中的文件有1~10不同的代先级，其中数字越大优先级越高。
// 打印机会从自己的待打印队列中选择优先级最高的文件来打印。
// 如果存在两个优先级一样的文件，则选择最早进入队列的那个文件。
// 现在请你来模拟汶5台打印机的打印过程。
func main() {
	printers := []*Printer{
		{Queue: [][2]int{}},
		{Queue: [][2]int{}},
		{Queue: [][2]int{}},
		{Queue: [][2]int{}},
		{Queue: [][2]int{}},
	}

	var n int
	fmt.Scan(&n)

	var no int
	for i := 0; i < n; i++ {
		var (
			command string
			p, val  int
		)
		fmt.Scan(&command)

		switch command {
		case "IN":
			fmt.Scan(&p, &val)
			no++

			printers[p-1].Insert(no, val)
		case "OUT":
			fmt.Scan(&p)

			out := printers[p-1].Out()
			if out == -1 {
				fmt.Println("NULL")
			} else {
				fmt.Println(out)
			}
		}
	}
}
