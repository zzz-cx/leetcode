package main

import "fmt"

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

func main() {
	got := spiralOrder([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}})
	fmt.Printf("PASS | %v\n", got)
}
