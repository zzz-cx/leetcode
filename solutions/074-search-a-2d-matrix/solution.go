package main

import "fmt"

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

func main() {
	got := searchMatrix([][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 3)
	fmt.Printf("PASS | %v\n", got)
}
