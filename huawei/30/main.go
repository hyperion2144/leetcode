package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Message struct {
	Price     int
	Number    int
	UnitPrice float64
}

// 某云短信厂商，为庆祝国庆，推出充值优惠活动。
// 现在给出客户预算，和优惠售价序列，求最多可获得的短信总条数。
func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var budget int
	fmt.Scan(&budget)

	var priceList []Message
	scanner.Scan()
	for price, digit := range strings.Split(scanner.Text(), " ") {
		number, _ := strconv.Atoi(digit)
		priceList = append(priceList, Message{Price: price + 1, Number: number, UnitPrice: float64(price+1) / float64(number)})
	}

	sort.Slice(priceList, func(i, j int) bool {
		return int(priceList[i].UnitPrice) < int(priceList[j].UnitPrice)
	})

	var sum int
	for i := len(priceList) - 1; i >= 0; i-- {
		message := priceList[i]
		if message.Price > budget {
			continue
		}

		budget -= message.Price
		sum += message.Number
	}

	fmt.Println(sum)
}
