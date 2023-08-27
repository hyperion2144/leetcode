package main

// 895. 最大频率栈
// https://leetcode-cn.com/problems/maximum-frequency-stack/
type FreqStack struct {
	stack map[int][]int
	count map[int]int
	max   int
}

func Constructor() FreqStack {
	return FreqStack{
		stack: make(map[int][]int),
		count: make(map[int]int),
	}
}

func (f *FreqStack) Push(val int) {
	f.count[val]++
	f.stack[f.count[val]] = append(f.stack[f.count[val]], val)
	f.max = max(f.max, f.count[val])
}

func (f *FreqStack) Pop() int {
	val := f.stack[f.max][len(f.stack[f.max])-1]
	f.stack[f.max] = f.stack[f.max][:len(f.stack[f.max])-1]
	f.count[val]--
	if len(f.stack[f.max]) == 0 {
		f.max--
	}
	return val
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
