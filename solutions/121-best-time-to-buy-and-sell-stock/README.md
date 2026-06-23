# 买卖股票的最佳时机

> LeetCode 121 · [best-time-to-buy-and-sell-stock](https://leetcode.cn/problems/best-time-to-buy-and-sell-stock/)

## 题目

给定股票价格数组 `prices`，最多完成一笔交易，设计算法计算最大利润。

## 题解思路与解析

- 见下方代码实现与注释。

## 解答

### Golang

```go
func maxProfit(prices []int) int {
	minPrice := prices[0]
	maxProfit := 0
	for _, price := range prices {
		if price < minPrice {
			minPrice = price
		} else if price-minPrice > maxProfit {
			maxProfit = price - minPrice
		}
	}
	return maxProfit
}
```

### Python

```python
from typing import List


class Solution:
    def max_profit(self, prices: List[int]) -> int:
        min_price = prices[0]
        max_profit = 0
        for price in prices:
            if price < min_price:
                min_price = price
            elif price - min_price > max_profit:
                max_profit = price - min_price
        return max_profit
```
