# 最小路径和

> LeetCode 64 · [minimum-path-sum](https://leetcode.cn/problems/minimum-path-sum/)

## 题目

给定 m×n 网格 `grid`，每个格子有非负整数，从左上角到右下角每次只能向下或向右，求路径数字和的最小值。

## 题解思路与解析

- 见下方代码实现与注释。

## 解答

### Golang

```go
func minPathSum(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	dp[0][0] = grid[0][0]
	for i := 1; i < m; i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}
	for j := 1; j < n; j++ {
		dp[0][j] = dp[0][j-1] + grid[0][j]
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}
	return dp[m-1][n-1]
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
```

### Python

```python
from typing import List


class Solution:
    def min_path_sum(self, grid: List[List[int]]) -> int:
        m, n = len(grid), len(grid[0])
        dp = [[0] * n for _ in range(m)]
        dp[0][0] = grid[0][0]
        for i in range(1, m):
            dp[i][0] = dp[i - 1][0] + grid[i][0]
        for j in range(1, n):
            dp[0][j] = dp[0][j - 1] + grid[0][j]
        for i in range(1, m):
            for j in range(1, n):
                dp[i][j] = min(dp[i - 1][j], dp[i][j - 1]) + grid[i][j]
        return dp[m - 1][n - 1]
```
