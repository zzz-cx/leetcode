# 矩阵置零

> LeetCode 73 · [set-matrix-zeroes](https://leetcode.cn/problems/set-matrix-zeroes/)

## 题目

给定 m×n 矩阵，若某元素为 0，则将其所在行和列的所有元素都设为 0。

## 题解思路与解析

- 给定一个 m x n 的矩阵，如果一个元素为 0 ，则将其所在行和列的所有元素都设为 0 。请使用 原地 算法。
- 先遍历一下数组，查看0的位置，对应的行和列标记一下，最后将对应行和列置零

## 解答

### Golang

```go
//给定一个 m x n 的矩阵，如果一个元素为 0 ，则将其所在行和列的所有元素都设为 0 。请使用 原地 算法。
//先遍历一下数组，查看0的位置，对应的行和列标记一下，最后将对应行和列置零

func setZeroes(matrix [][]int) {
	m, n := len(matrix), len(matrix[0])
	row, col := make([]bool, m), make([]bool, n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				row[i] = true
				col[j] = true
			}
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if row[i] || col[j] {
				matrix[i][j] = 0
			}
		}
	}
}
```

### Python

```python
# 给定一个 m x n 的矩阵，如果一个元素为 0 ，则将其所在行和列的所有元素都设为 0 。请使用原地算法。
# 先遍历一下数组，查看0的位置，对应的行和列标记一下，最后将对应行和列置零
from typing import List


class Solution:
    def set_zeroes(self, matrix: List[List[int]]) -> None:
        m, n = len(matrix), len(matrix[0])
        row = [False] * m
        col = [False] * n
        for i in range(m):
            for j in range(n):
                if matrix[i][j] == 0:
                    row[i] = True
                    col[j] = True
        for i in range(m):
            for j in range(n):
                if row[i] or col[j]:
                    matrix[i][j] = 0
```
