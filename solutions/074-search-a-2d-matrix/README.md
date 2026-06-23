# 搜索二维矩阵

> LeetCode 74 · [search-a-2d-matrix](https://leetcode.cn/problems/search-a-2d-matrix/)

## 题目

编写高效算法判断 m×n 矩阵中是否存在目标值，矩阵每行升序且每行首元素大于上一行尾元素。

## 题解思路与解析

- Z字型查找，从右上角开始查找，如果当前值大于目标值，则向左移动，如果当前值小于目标值，则向下移动

## 解答

### Golang

```go
func searchMatrix(matrix [][]int, target int) bool {
	//Z字型查找，从右上角开始查找，如果当前值大于目标值，则向左移动，如果当前值小于目标值，则向下移动
	m, n := len(matrix), len(matrix[0])
	row, col := 0, n-1
	for row < m && col >= 0 {
		if matrix[row][col] == target {
			return true
		} else if matrix[row][col] > target {
			col--
		} else {
			row++
		}
	}
	return false
}
```

### Python

```python
from typing import List


class Solution:
    def search_matrix(self, matrix: List[List[int]], target: int) -> bool:
        m, n = len(matrix), len(matrix[0])
        row, col = 0, n - 1
        while row < m and col >= 0:
            if matrix[row][col] == target:
                return True
            elif matrix[row][col] > target:
                col -= 1
            else:
                row += 1
        return False
```
