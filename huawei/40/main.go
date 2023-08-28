package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// 记账本上记录了若干条多国货币金额，需要转换成人民币分，汇总后输出
// 每行记录一条金额，金额带有货币单位，格式为数字+单位，可能是单独元，或者单独分，或者元与分的组合
// 要求将这些货币全部换算成人民币分后进行汇总，汇总结果仅保留整数，小数部分舍弃
// 元和分的换算关系都是1:100，如下:
// 1CNY=100fen（1元=100分）
// 1HKD=100cents（1港元=100港分）
// 1JPY=100sens（1日元=100仙）
// 1EUR=100eurocents（1欧元=100欧分）
// 1GBP=100pence（1英镑=100便士）
// 汇率如下表:
// | CNY | JPY  | HKD | EUR | GBP |
// | 100 | 1825 | 123 | 14  | 12  |
// 即100CNY=1825JPY=123HKD=14EUR=12GBP
func main() {
	var n int
	fmt.Scan(&n)

	reader := bufio.NewReader(os.Stdin)

	var sum float64
	for i := 0; i < n; i++ {

		num := 0.0
		unit := ""

		digits := make([]byte, 0)
		strs := make([]byte, 0)
		for {
			b, _ := reader.ReadByte()
			if b == '\n' {
				unit = string(strs)
				break
			}

			if b >= '0' && b <= '9' {
				if len(strs) > 0 {
					strs = strs[:0]
				}

				digits = append(digits, b)
				continue
			}
			if ('a' <= b && b <= 'z') || ('A' <= b && b <= 'Z') {
				if len(digits) > 0 {
					temp, _ := strconv.ParseFloat(string(digits), 64)
					digits = digits[:0]

					if num > 0 {
						num = num*100 + temp
					} else {
						num = temp
					}
				}

				strs = append(strs, b)
				continue
			}
		}

		switch unit {
		case "CNY":
			sum += num * 100
		case "fen":
			sum += num
		case "JPY":
			sum += num * (100.0 / 1825) * 100
		case "sen":
			sum += num * (100.0 / 1825)
		case "HKD":
			sum += num * (100.0 / 123) * 100
		case "cents":
			sum += num * (100.0 / 123)
		case "EUR":
			sum += num * (100.0 / 14) * 100
		case "eurocents":
			sum += num * (100.0 / 14)
		case "GBP":
			sum += num * (100.0 / 12) * 100
		case "pence":
			sum += num * (100.0 / 12)
		}
	}

	fmt.Println(int(sum))
}
