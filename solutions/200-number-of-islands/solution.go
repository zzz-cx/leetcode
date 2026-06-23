package main

import "fmt"

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

func main() {
	grid := [][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}
	got := numIslands(grid)
	status := "PASS"
	if got != 1 {
		status = "FAIL"
	}
	fmt.Printf("%s | islands=%d\n", status, got)
}
