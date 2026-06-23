# 买卖股票的最佳时机 II

> LeetCode 122 · [best-time-to-buy-and-sell-stock-ii](https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-ii/)

## 题目

给定股票价格数组，可多次买卖，设计算法计算最大利润。

## 题解思路与解析

- maxProfitII 买卖股票 II（LeetCode 122）：可多次买卖，同一天可先卖再买
- 思路：贪心。把所有「上涨段」的利润都吃掉 —— 等价于累加每天比前一天高的差价
- 例 [7,1,5,3,6,4]：1→5 赚 4，3→6 赚 3，即 (5-1)+(6-3)=7

## 解答

### Golang

```go
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
```

### Python

```python
# max_profit_ii 买卖股票 II（LeetCode 122）：可多次买卖，同一天可先卖再买
# 思路：贪心。把所有「上涨段」的利润都吃掉 —— 等价于累加每天比前一天高的差价
from typing import List


class Solution:
    def max_profit_ii(self, prices: List[int]) -> int:
        profit = 0
        for i in range(1, len(prices)):
            if prices[i] > prices[i - 1]:
                profit += prices[i] - prices[i - 1]
        return profit
```
