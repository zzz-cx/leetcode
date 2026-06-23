# 零钱兑换

> LeetCode 322 · [coin-change](https://leetcode.cn/problems/coin-change/)

## 题目

给定不同面额硬币 `coins` 和总金额 `amount`，计算凑成总金额所需的最少硬币数。

## 题解思路与解析

- 见下方代码实现与注释。

## 解答

### Golang

```go
func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		dp[i] = amount + 1
		for _, coin := range coins {
			if coin <= i {
				dp[i] = min(dp[i], dp[i-coin]+1)
			}
		}
	}
	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}
```

### Python

```python
from typing import List


class Solution:
    def coin_change(self, coins: List[int], amount: int) -> int:
        dp = [amount + 1] * (amount + 1)
        dp[0] = 0
        for i in range(1, amount + 1):
            for coin in coins:
                if coin <= i:
                    dp[i] = min(dp[i], dp[i - coin] + 1)
        if dp[amount] > amount:
            return -1
        return dp[amount]
```
