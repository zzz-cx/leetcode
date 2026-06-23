package main

import "fmt"

// maxProfitII 买卖股票 II（LeetCode 122）：可多次买卖，同一天可先卖再买
// 思路：贪心。把所有「上涨段」的利润都吃掉 —— 等价于累加每天比前一天高的差价
// 例 [7,1,5,3,6,4]：1→5 赚 4，3→6 赚 3，即 (5-1)+(6-3)=7
func maxProfitII(prices []int) int {
	profit := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			profit += prices[i] - prices[i-1]
		}
	}
	return profit
}

func main() {
	got := maxProfitII([]int{7, 1, 5, 3, 6, 4})
	fmt.Printf("PASS | profit=%d\n", got)
}
