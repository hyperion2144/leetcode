
//leetcode submit region begin(Prohibit modification and deletion)
func singleNumber(nums []int) []int {
	// 使用异或运算，计算出x1 XOR x2的结果
	var x int
	for _, v := range nums {
		x ^= v
	}
	// 计算x中最低位的1 所在的位置l
	l := x & -x
	// 将nums分成两部分，分别计算出x1和x2
	var x1, x2 int
	for _, v := range nums {
		if v&l > 0 {
			x1 ^= v
		} else {
			x2 ^= v
		}
	}
	return []int{x1, x2}
}
//leetcode submit region end(Prohibit modification and deletion)