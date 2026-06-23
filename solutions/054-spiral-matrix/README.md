# 螺旋矩阵

> LeetCode 54 · [spiral-matrix](https://leetcode.cn/problems/spiral-matrix/)

## 题目

给定 m×n 矩阵，按顺时针螺旋顺序返回矩阵中的所有元素。

## 题解思路与解析

- 初始位置是矩阵的左上角，初始方向是向右，当路径超出界限或者进入之前访问过的位置时，顺时针旋转，进入下一个方向。

## 解答

### Golang

```go
func spiralOrder(matrix [][]int) []int {
	//初始位置是矩阵的左上角，初始方向是向右，当路径超出界限或者进入之前访问过的位置时，顺时针旋转，进入下一个方向。
	m, n := len(matrix), len(matrix[0])
	out := make([]int, 0, m*n)
	direction := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}
	row, col := 0, 0
	dir := 0
	for i := 0; i < m*n; i++ {
		out = append(out, matrix[row][col])
		visited[row][col] = true
		nextRow, nextCol := row+direction[dir][0], col+direction[dir][1]
		if nextRow < 0 || nextRow >= m || nextCol < 0 || nextCol >= n || visited[nextRow][nextCol] {
			dir = (dir + 1) % 4
			nextRow, nextCol = row+direction[dir][0], col+direction[dir][1]
		}
		row, col = nextRow, nextCol
	}
	return out
}
```

### Python

```python
# 初始位置是矩阵的左上角，初始方向是向右，当路径超出界限或者进入之前访问过的位置时，顺时针旋转，进入下一个方向。
from typing import List


class Solution:
    def spiral_order(self, matrix: List[List[int]]) -> List[int]:
        m, n = len(matrix), len(matrix[0])
        out: List[int] = []
        direction = [(0, 1), (1, 0), (0, -1), (-1, 0)]
        visited = [[False] * n for _ in range(m)]
        row, col, dir_idx = 0, 0, 0
        for _ in range(m * n):
            out.append(matrix[row][col])
            visited[row][col] = True
            next_row, next_col = row + direction[dir_idx][0], col + direction[dir_idx][1]
            if (
                next_row < 0
                or next_row >= m
                or next_col < 0
                or next_col >= n
                or visited[next_row][next_col]
            ):
                dir_idx = (dir_idx + 1) % 4
                next_row, next_col = row + direction[dir_idx][0], col + direction[dir_idx][1]
            row, col = next_row, next_col
        return out
```
