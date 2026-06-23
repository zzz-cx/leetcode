# 旋转图像

> LeetCode 48 · [rotate-image](https://leetcode.cn/problems/rotate-image/)

## 题目

给定 n×n 矩阵，原地将其顺时针旋转 90 度。

## 题解思路与解析

- 见下方代码实现与注释。

## 解答

### Golang

```go
func rotate(matrix [][]int) {
	n := len(matrix)
	for i := 0; i < n/2; i++ {
		for j := i; j < n-i-1; j++ {
			matrix[i][j], matrix[j][n-i-1], matrix[n-i-1][n-j-1], matrix[n-j-1][i] = matrix[n-j-1][i], matrix[i][j], matrix[j][n-i-1], matrix[n-i-1][n-j-1]
		}
	}
}
```

### Python

```python
from typing import List


class Solution:
    def rotate(self, matrix: List[List[int]]) -> None:
        n = len(matrix)
        for i in range(n // 2):
            for j in range(i, n - i - 1):
                (
                    matrix[i][j],
                    matrix[j][n - i - 1],
                    matrix[n - i - 1][n - j - 1],
                    matrix[n - j - 1][i],
                ) = (
                    matrix[n - j - 1][i],
                    matrix[i][j],
                    matrix[j][n - i - 1],
                    matrix[n - i - 1][n - j - 1],
                )
```
