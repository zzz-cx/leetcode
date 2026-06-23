# 岛屿数量

> LeetCode 200 · [number-of-islands](https://leetcode.cn/problems/number-of-islands/)

## 题目

给定由 '1'（陆地）和 '0'（水）组成的二维网格，计算岛屿数量。

## 题解思路与解析

- 见下方代码实现与注释。

## 解答

### Golang

```go
func numIslands(grid [][]byte) int {
	m, n := len(grid), len(grid[0])
	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}
	count := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if visited[i][j] || grid[i][j] == '0' {
				continue
			}
			dfs(grid, visited, i, j)
			count++
		}
	}
	return count
}
func dfs(grid [][]byte, visited [][]bool, i, j int) {
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) || visited[i][j] || grid[i][j] == '0' {
		return
	}
	visited[i][j] = true
	dfs(grid, visited, i+1, j)
	dfs(grid, visited, i-1, j)
	dfs(grid, visited, i, j+1)
	dfs(grid, visited, i, j-1)
}
```

### Python

```python
from typing import List


class Solution:
    def num_islands(self, grid: List[List[str]]) -> int:
        m, n = len(grid), len(grid[0])
        visited = [[False] * n for _ in range(m)]
        count = 0
        for i in range(m):
            for j in range(n):
                if visited[i][j] or grid[i][j] == "0":
                    continue
                self._dfs(grid, visited, i, j)
                count += 1
        return count

    def _dfs(self, grid: List[List[str]], visited: List[List[bool]], i: int, j: int) -> None:
        if (
            i < 0
            or i >= len(grid)
            or j < 0
            or j >= len(grid[0])
            or visited[i][j]
            or grid[i][j] == "0"
        ):
            return
        visited[i][j] = True
        self._dfs(grid, visited, i + 1, j)
        self._dfs(grid, visited, i - 1, j)
        self._dfs(grid, visited, i, j + 1)
        self._dfs(grid, visited, i, j - 1)
```
