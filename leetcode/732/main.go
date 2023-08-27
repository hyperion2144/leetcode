package main

type MyCalendarThree map[int]struct {
	num, lazy int
}

func Constructor() MyCalendarThree {
	return make(MyCalendarThree)
}

func (t MyCalendarThree) update(start, end, l, r, idx int) {
	// 不在区间内
	if r < start || end < l {
		return
	}

	// 比区间大, 或者与区间有交集
	if start <= l && r <= end {
		p := t[idx]
		p.num++
		p.lazy++
		t[idx] = p
	} else {
		mid := (l + r) / 2
		t.update(start, end, l, mid, idx*2)
		t.update(start, end, mid+1, r, idx*2+1)
		p := t[idx]
		p.num = p.lazy + max(t[idx*2].num, t[idx*2+1].num)
		t[idx] = p
	}
}

func (t MyCalendarThree) Book(startTime int, endTime int) int {
	t.update(startTime, endTime-1, 0, 1e9, 1)
	return t[1].num
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	tree := Constructor()
	tree.Book(10, 20)
	tree.Book(50, 60)
	tree.Book(10, 40)
	tree.Book(5, 15)
	tree.Book(5, 10)
	tree.Book(25, 55)
}
